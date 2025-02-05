package argo_workflows

import (
	"eksdemo/pkg/application"
	"eksdemo/pkg/cmd"
	"eksdemo/pkg/installer"
	"eksdemo/pkg/template"
)

// Docs:    https://argoproj.github.io/argo-workflows/
// GitHub:  https://github.com/argoproj/argo-workflows
// Helm:    https://github.com/argoproj/argo-helm/tree/main/charts/argo-workflows
// Repo:    quay.io/argoproj/argocli, quay.io/argoproj/workflow-controller
// Version: Latest Chart is 0.16.8, Argo Workflows v3.3.8 (as of 07/27/22)

func NewApp() *application.Application {
	app := &application.Application{
		Command: cmd.Command{
			Parent:      "argo",
			Name:        "workflows",
			Description: "Workflow engine for Kubernetes",
		},

		Installer: &installer.HelmInstaller{
			ChartName:     "argo-workflows",
			ReleaseName:   "argo-workflows",
			RepositoryURL: "https://argoproj.github.io/argo-helm",
			ValuesTemplate: &template.TextTemplate{
				Template: valuesTemplate,
			},
		},
	}
	app.Options, app.Flags = newOptions()

	return app
}

const valuesTemplate = `---
images:
  tag: {{ .Version }}
fullnameOverride: argo
server:
  serviceType: {{ .ServiceType }}
{{- if eq .ServiceType "LoadBalancer" }}
  serviceAnnotations:
    {{- .ServiceAnnotations | nindent 4 }}
{{- end }}
  extraArgs:
  - --auth-mode={{ .AuthMode }}
{{- if .IngressHost }}
  ingress:
    enabled: true
    annotations:
      {{- .IngressAnnotations | nindent 6 }}
    ingressClassName: {{ .IngressClass }}
    hosts:
    - {{ .IngressHost }}
    paths:
    - /
    pathType: Prefix
    tls:
    - hosts:
      - {{ .IngressHost }}
    {{- if ne .IngressClass "alb" }}
      secretName: argo-workflows-tls
    {{- end }}
{{- end }}
`
