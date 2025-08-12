{{/*
Expand the name of the chart.
*/}}
{{- define "base.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "base.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Metadata 
*/}}
{{- define "base.metadata" -}}
{{ include "base.selectorLabels" . }}
managed_by: helm
chart: {{ include "base.chart" . }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "base.labels" -}}
{{ include "base.selectorLabels" . }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "base.selectorLabels" -}}
{{ include "base.governanceLabels" . }}
{{- end }}

{{/*
Governance Labels
*/}}
{{- define "base.governanceLabels" -}}
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