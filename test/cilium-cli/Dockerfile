FROM alpine:3.20
ARG TARGETARCH
ARG CILIUM_CLI_VERSION
RUN apk add --no-cache ca-certificates curl && \
    curl -L https://github.com/cilium/cilium-cli/releases/download/${CILIUM_CLI_VERSION}/cilium-linux-${TARGETARCH}.tar.gz \
    | tar xz -C /usr/local/bin && \
    mv /usr/local/bin/cilium /usr/local/bin/cilium-cli && \
    apk del curl
