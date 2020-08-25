{{- define "prometheus.tls-config.kube-cert-auth" -}}
ca_file: /etc/prometheus/seed/ca.crt
cert_file: /etc/prometheus/seed/prometheus.crt
key_file: /etc/prometheus/seed/prometheus.key
{{- end -}}

{{- define "prometheus.tls-config.kube-cert-auth-insecure" -}}
insecure_skip_verify: true
cert_file: /etc/prometheus/seed/prometheus.crt
key_file: /etc/prometheus/seed/prometheus.key
{{- end -}}


{{- define "prometheus.keep-metrics.metric-relabel-config" -}}
- source_labels: [ __name__ ]
  regex: ^({{ . | join "|" }})$
  action: keep
{{- end -}}
