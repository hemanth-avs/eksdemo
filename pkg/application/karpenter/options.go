package karpenter

import (
	"eksdemo/pkg/application"
	"eksdemo/pkg/cmd"
)

type KarpenterOptions struct {
	application.ApplicationOptions
}

func NewOptions() (options *KarpenterOptions, flags cmd.Flags) {
	options = &KarpenterOptions{
		ApplicationOptions: application.ApplicationOptions{
			Namespace:      "karpenter",
			ServiceAccount: "karpenter",
			DefaultVersion: &application.LatestPrevious{
				LatestChart:   "0.13.1",
				Latest:        "v0.13.1",
				PreviousChart: "0.11.1",
				Previous:      "v0.11.1",
			},
		},
	}

	flags = cmd.Flags{}
	return
}
