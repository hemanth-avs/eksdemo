package velero

import (
	"eksdemo/pkg/application"
	"eksdemo/pkg/cmd"
	"eksdemo/pkg/installer"
	"eksdemo/pkg/resource"
	"eksdemo/pkg/resource/irsa"
	"eksdemo/pkg/resource/s3_bucket"
	"eksdemo/pkg/template"
)

// Docs:    https://velero.io/docs/
// GitHub:  https://github.com/vmware-tanzu/velero
// Helm:    https://github.com/vmware-tanzu/helm-charts/tree/main/charts/velero
// Repo:    velero/velero
// Version: Latest is chart 2.30.1, app v1.9.0 (as of 07/15/22)

func NewApp() *application.Application {
	options, flags := newOptions()

	app := &application.Application{
		Command: cmd.Command{
			Name:        "velero",
			Description: "Backup and migrate Kubernetes applications",
		},

		Dependencies: []*resource.Resource{
			irsa.NewResourceWithOptions(&irsa.IrsaOptions{
				CommonOptions: resource.CommonOptions{
					Name: "velero-irsa",
				},
				PolicyType: irsa.PolicyDocument,
				PolicyDocTemplate: &template.TextTemplate{
					Template: policyDocument,
				},
			}),
			s3_bucket.NewResourceWithOptions(options.BucketOptions),
		},

		Installer: &installer.HelmInstaller{
			ChartName:     "velero",
			ReleaseName:   "velero",
			RepositoryURL: "https://vmware-tanzu.github.io/helm-charts",
			ValuesTemplate: &template.TextTemplate{
				Template: valuesTemplate,
			},
		},
	}
	app.Options = options
	app.Flags = flags

	return app
}

const policyDocument = `
Version: '2012-10-17'
Statement:
- Effect: Allow
  Action:
  - ec2:DescribeVolumes
  - ec2:DescribeSnapshots
  - ec2:CreateTags
  - ec2:CreateVolume
  - ec2:CreateSnapshot
  - ec2:DeleteSnapshot
  Resource: "*"
- Effect: Allow
  Action:
  - s3:GetObject
  - s3:DeleteObject
  - s3:PutObject
  - s3:AbortMultipartUpload
  - s3:ListMultipartUploadParts
  Resource: arn:aws:s3:::eksdemo-{{ .Account }}-velero/*
- Effect: Allow
  Action: s3:ListBucket
  Resource: arn:aws:s3:::eksdemo-{{ .Account }}-velero
`

const valuesTemplate = `---
image:
  tag: {{ .Version }}
initContainers:
- name: velero-plugin-for-aws
  image: velero/velero-plugin-for-aws:{{ .PluginVersion }}
  volumeMounts:
  - mountPath: /target
    name: plugins
configuration:
  provider: aws
  backupStorageLocation:
    bucket: eksdemo-{{ .Account }}-velero
  volumeSnapshotLocation:
    config:
      region: {{ .Region }}
serviceAccount:
  server:
    name: {{ .ServiceAccount }}
    annotations:
      {{ .IrsaAnnotation }}
credentials:
  useSecret: false
`
