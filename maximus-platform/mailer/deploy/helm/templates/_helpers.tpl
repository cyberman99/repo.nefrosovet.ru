{{- define "mailer.name" -}}
{{- .Chart.Name | lower -}}
{{- end -}}

{{- define "mailer.fullname" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | lower | replace "." "-"}}
{{- end -}}

{{- define "mailer.release" -}}
{{- printf "%s-%s" .Chart.Version .Chart.AppVersion | lower | replace "." "-"}}
{{- end -}}

{{- define "mailer.version" -}}
{{- .Chart.Version | lower -}}
{{- end -}}
