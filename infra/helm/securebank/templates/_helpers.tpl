{{/*
Common labels
*/}}
{{- define "securebank.labels" -}}
app.kubernetes.io/managed-by: Helm
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/version: {{ .Chart.AppVersion }}
environment: {{ .Values.global.environment }}
{{- end }}

{{/*
Image helper
*/}}
{{- define "securebank.image" -}}
{{ .Values.global.imageRegistry }}/{{ .image }}:{{ .tag }}
{{- end }}
