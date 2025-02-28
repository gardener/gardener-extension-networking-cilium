{{/*
Convert a map to a comma-separated string: key1=value1,key2=value2
*/}}
{{- define "mapToString" -}}
{{- $list := list -}}
{{- range $k, $v := . -}}
{{- $list = append $list (printf "%s=%s" $k $v) -}}
{{- end -}}
{{ join "," $list }}
{{- end -}}
