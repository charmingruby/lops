{{/*
Expand the name of the chart.
*/}}
{{- define "api.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "api.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Metadata 
*/}}
{{- define "api.metadata" -}}
{{ include "api.selectorLabels" . }}
managed_by: helm
chart: {{ include "api.chart" . }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "api.labels" -}}
{{ include "api.selectorLabels" . }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "api.selectorLabels" -}}
{{ include "api.governanceLabels" . }}
{{- end }}

{{/*
Governance Labels
*/}}
{{- define "api.governanceLabels" -}}
app_name: {{ .Values.governanceData.appName }}
app_type: {{ .Values.governanceData.appType }}
team: {{ .Values.governanceData.team }}
{{- end }}

{{/*
Configs map resource name
*/}}
{{- define "config.configMapName" -}}
{{- printf "%s-config-map" . -}}
{{- end -}}