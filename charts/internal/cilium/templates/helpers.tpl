{{- define "cilium.selfhosted-controlplane-k8s-service-env" -}}
- name: KUBERNETES_SERVICE_HOST
  value: localhost
- name: KUBERNETES_SERVICE_PORT
  value: "443"
{{- end -}}

