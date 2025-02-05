package aws_lb

import (
	"eksdemo/pkg/application"
	"eksdemo/pkg/cmd"
	"eksdemo/pkg/installer"
	"eksdemo/pkg/resource"
	"eksdemo/pkg/resource/irsa"
	"eksdemo/pkg/template"
)

// Docs:    https://kubernetes-sigs.github.io/aws-load-balancer-controller/
// GitHub:  https://github.com/kubernetes-sigs/aws-load-balancer-controller
// Helm:    https://github.com/aws/eks-charts/tree/master/stable/aws-load-balancer-controller
// Repo:    602401143452.dkr.ecr.us-west-2.amazonaws.com/amazon/aws-load-balancer-controller
// Version: Latest is v2.4.3 (as of 08/14/22)

func NewApp() *application.Application {
	app := &application.Application{
		Command: cmd.Command{
			Name:        "aws-lb-controller",
			Description: "AWS Load Balancer Controller",
			Aliases:     []string{"aws-lb", "awslb"},
		},

		Dependencies: []*resource.Resource{
			irsa.NewResourceWithOptions(&irsa.IrsaOptions{
				CommonOptions: resource.CommonOptions{
					Name: "aws-lb-controller-irsa",
				},
				PolicyType: irsa.WellKnown,
				Policy:     []string{"awsLoadBalancerController"},
			}),
		},

		Options: &application.ApplicationOptions{
			Namespace:      "awslb",
			ServiceAccount: "aws-load-balancer-controller",
			DefaultVersion: &application.LatestPrevious{
				LatestChart:   "1.4.4",
				Latest:        "v2.4.3",
				PreviousChart: "1.4.2",
				Previous:      "v2.4.2",
			},
		},

		Installer: &installer.HelmInstaller{
			ChartName:     "aws-load-balancer-controller",
			ReleaseName:   "aws-lb-controller",
			RepositoryURL: "https://aws.github.io/eks-charts",
			ValuesTemplate: &template.TextTemplate{
				Template: valuesTemplate,
			},
		},
	}
	return app
}

const valuesTemplate = `---
replicaCount: 1
image:
  tag: {{ .Version }}
fullnameOverride: aws-load-balancer-controller
clusterName: {{ .ClusterName }}
serviceAccount:
  annotations:
    {{ .IrsaAnnotation }}
  name: {{ .ServiceAccount }}
region: {{ .Region }}
vpcId: {{ .Cluster.ResourcesVpcConfig.VpcId }}
`
