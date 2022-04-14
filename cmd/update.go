package cmd

import (
	"eksdemo/pkg/resource/nodegroup"

	"github.com/spf13/cobra"
)

func newCmdUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "update resource(s)",
	}

	// Don't show flag errors for delete without a subcommand
	cmd.DisableFlagParsing = true

	cmd.AddCommand(nodegroup.NewResource().NewUpdateCmd())

	return cmd
}
