package log_stream

import (
	"eksdemo/pkg/cmd"
	"eksdemo/pkg/resource"
)

func NewResource() *resource.Resource {
	res := &resource.Resource{
		Command: cmd.Command{
			Name:        "log-stream",
			Description: "CloudWatch Log Stream",
			Aliases:     []string{"log-streams", "logstream", "ls"},
			Args:        []string{"NAME_PREFIX"},
		},

		Getter: &Getter{},
	}

	res.Options, res.GetFlags = newOptions()

	return res
}
