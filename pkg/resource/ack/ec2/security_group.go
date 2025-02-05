package ec2

import (
	"eksdemo/pkg/cmd"
	"eksdemo/pkg/kubernetes"
	"eksdemo/pkg/resource"
	"eksdemo/pkg/template"
)

type SecurityGroupOptions struct {
	resource.CommonOptions
	Description string
}

func NewSecurityGroupResource() *resource.Resource {
	res := &resource.Resource{
		Command: cmd.Command{
			Name:        "security-group",
			Description: "EC2 Security Group",
			Aliases:     []string{"security-groups", "sg"},
			Args:        []string{"NAME"},
		},

		Manager: &kubernetes.ResourceManager{
			Template: &template.TextTemplate{
				Template: securityGroupYamlTemplate,
			},
		},
	}

	options := &SecurityGroupOptions{
		CommonOptions: resource.CommonOptions{
			Name:          "ack-ec2-security-group",
			Namespace:     "default",
			NamespaceFlag: true,
		},
	}

	flags := cmd.Flags{
		&cmd.StringFlag{
			CommandFlag: cmd.CommandFlag{
				Name:        "desc",
				Description: "description",
				Required:    true,
			},
			Option: &options.Description,
		},
	}

	res.Options = options
	res.CreateFlags = flags

	return res
}

const securityGroupYamlTemplate = `
---
apiVersion: ec2.services.k8s.aws/v1alpha1
kind: SecurityGroup
metadata:
  name: {{ .Name }}
  namespace: {{ .Namespace }}
spec:
  description: {{ .Description }}
  name: {{ .Name }}
  tagSpecifications:
  - resourceType: security-group
    tags:
    - key: Name
      value: {{ .Name }}
  vpcID: {{ .Cluster.ResourcesVpcConfig.VpcId }}
...
`
