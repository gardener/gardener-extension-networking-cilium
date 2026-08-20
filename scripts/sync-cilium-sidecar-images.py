#!/usr/bin/env python3
"""sync-cilium-sidecar-images.py <cilium-version>

Fetches the authoritative cilium/cilium values.yaml at the given Cilium version and
updates imagevector/images.yaml with the exact cilium-envoy and certgen tags it pins.
Also updates the c-ares and v8 scanning-hint versions from the Envoy build graph
and advances the cilium/proxy minor-version comment.

Intended to be called from Renovate postUpgradeTasks after a cilium-agent bump.
Requires: python3 (stdlib only)
"""

import re
import sys
import urllib.request


def fetch_text(url: str) -> str:
    with urllib.request.urlopen(url, timeout=30) as resp:
        return resp.read().decode()


def extract_image_tag(lines: list[str], top_key: str) -> str:
    in_top = in_image = False
    for line in lines:
        if line == f'{top_key}:':
            in_top = True
            in_image = False
        elif in_top and line == '  image:':
            in_image = True
        elif in_top and in_image and line.startswith('    tag:'):
            return line.split('tag:', 1)[1].strip().strip('"').strip("'")
        elif in_top and line and not line[0].isspace():
            in_top = in_image = False
    raise ValueError(f'tag not found for {top_key!r}')


def extract_bzl_version(bzl: str, project_name: str) -> str:
    """Extract version = "..." for a project identified by project_name in repository_locations.bzl.

    Uses project_name (e.g. "c-ares", "V8") rather than the dict key, which can
    be renamed across Envoy versions (e.g. com_github_c_ares_c_ares → com_github_cares_cares).
    """
    # Find a dict block containing project_name = "<name>" and extract its version field.
    # Strategy: find the project_name anchor, then search backwards to the opening `= dict(`
    # and forwards to the closing `),` to extract the version within that block.
    m = re.search(
        r'=\s*dict\([^)]*?project_name\s*=\s*"' + re.escape(project_name) + r'".*?version\s*=\s*"([^"]+)"',
        bzl,
        re.DOTALL,
    )
    if not m:
        raise ValueError(f'version not found for project_name {project_name!r}')
    return m.group(1)


def update_images_yaml(
    target: str,
    envoy_tag: str,
    certgen_tag: str,
    envoy_ver: str,
    cilium_ver: str,
    cares_ver: str,
    v8_ver: str,
) -> None:
    with open(target) as f:
        content = f.read()

    # 1. Replace cilium-envoy tag (unique format: v<maj>.<min>.<pat>-<10-digit-ts>-<sha>)
    content = re.sub(
        r'(tag: )v\d+\.\d+\.\d+-\d{10}-[0-9a-f]+',
        r'\g<1>' + envoy_tag,
        content,
    )

    # 2. Replace certgen tag (simple semver, scoped to the certgen block)
    def replace_certgen_tag(m: re.Match) -> str:
        block = m.group(0)
        return re.sub(r'(    tag: )v\d+\.\d+\.\d+(\s)', r'\g<1>' + certgen_tag + r'\2', block)

    content = re.sub(
        r'- name: certgen\b.*?(?=\n  - name:|\Z)',
        replace_certgen_tag,
        content,
        flags=re.DOTALL,
    )

    # 3. Update Envoy doc URL comments (both cilium-agent and cilium-envoy blocks)
    content = re.sub(
        r'(envoy/v)\d+\.\d+\.\d+(/intro)',
        r'\g<1>' + envoy_ver + r'\g<2>',
        content,
    )

    # 4. Update cilium/proxy minor-version comment (both blocks)
    #    Format: # https://github.com/cilium/proxy: v1.19.x -> v1.36.x
    cilium_minor = re.match(r'v(\d+\.\d+)', cilium_ver).group(1)   # e.g. "1.19"
    envoy_minor  = re.match(r'(\d+\.\d+)',  envoy_ver).group(1)    # e.g. "1.36"
    content = re.sub(
        r'(# https://github\.com/cilium/proxy: v)\d+\.\d+(\.x -> v)\d+\.\d+(\.x)',
        rf'\g<1>{cilium_minor}\g<2>{envoy_minor}\g<3>',
        content,
    )

    # 5. Update c-ares scanning-hint version (cilium-envoy block)
    content = re.sub(
        r"(- name: 'c-ares'\s*\n\s*version: ')[^']+(')",
        r'\g<1>' + cares_ver + r'\g<2>',
        content,
    )

    # 6. Update v8 scanning-hint version (cilium-agent block)
    content = re.sub(
        r"(- name: 'v8'\s*\n\s*version: ')[^']+(')",
        r'\g<1>' + v8_ver + r'\g<2>',
        content,
    )

    # 7. Update or insert a sync-info comment before the cilium-envoy entry
    sync_comment = (
        f'  # Synced from cilium/cilium {cilium_ver}:'
        f' cilium-envoy={envoy_tag}, certgen={certgen_tag}\n'
    )
    if '  # Synced from cilium/cilium' in content:
        content = re.sub(r'  # Synced from cilium/cilium[^\n]*\n', sync_comment, content)
    else:
        content = content.replace(
            '  - name: cilium-envoy\n',
            sync_comment + '  - name: cilium-envoy\n',
            1,
        )

    with open(target, 'w') as f:
        f.write(content)

    print(f'Updated {target}: cilium-envoy={envoy_tag}, certgen={certgen_tag},'
          f' c-ares={cares_ver}, v8={v8_ver}')


def main() -> None:
    if len(sys.argv) != 2:
        print(f'Usage: {sys.argv[0]} <cilium-version>', file=sys.stderr)
        sys.exit(1)

    cilium_ver = sys.argv[1]
    values_url = (
        f'https://raw.githubusercontent.com/cilium/cilium/{cilium_ver}'
        f'/install/kubernetes/cilium/values.yaml'
    )
    target = 'imagevector/images.yaml'

    print(f'Fetching {values_url}...')
    lines = fetch_text(values_url).splitlines()

    envoy_tag   = extract_image_tag(lines, 'envoy')
    certgen_tag = extract_image_tag(lines, 'certgen')

    m = re.match(r'v(\d+\.\d+\.\d+)', envoy_tag)
    if not m:
        print(f'Could not extract semver from envoy tag: {envoy_tag!r}', file=sys.stderr)
        sys.exit(1)
    envoy_ver = m.group(1)  # e.g. "1.36.9"

    bzl_url = (
        f'https://raw.githubusercontent.com/envoyproxy/envoy/v{envoy_ver}'
        f'/bazel/repository_locations.bzl'
    )
    print(f'Fetching {bzl_url}...')
    bzl = fetch_text(bzl_url)

    cares_ver = extract_bzl_version(bzl, 'c-ares')
    v8_ver    = extract_bzl_version(bzl, 'V8')

    print(f'cilium-envoy: {envoy_tag}')
    print(f'certgen:      {certgen_tag}')
    print(f'envoy ver:    {envoy_ver}')
    print(f'c-ares:       {cares_ver}')
    print(f'v8:           {v8_ver}')

    update_images_yaml(target, envoy_tag, certgen_tag, envoy_ver, cilium_ver,
                       cares_ver, v8_ver)


if __name__ == '__main__':
    main()
