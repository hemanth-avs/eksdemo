package volume

import (
	"eksdemo/pkg/cmd"
	"eksdemo/pkg/resource"
)

func NewResource() *resource.Resource {
	res := &resource.Resource{
		Command: cmd.Command{
			Name:        "volume",
			Description: "EBS Volume",
			Aliases:     []string{"volumes", "vols", "vol"},
			Args:        []string{"ID"},
		},

		Getter: &Getter{},

		Manager: &Manager{},

		Options: &resource.CommonOptions{
			ClusterFlagDisabled: true,
		},
	}

	return res
}
