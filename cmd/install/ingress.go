package install

import (
	"eksdemo/pkg/application"
	"eksdemo/pkg/application/ingress/emissary"
	"eksdemo/pkg/application/ingress/nginx"

	"github.com/spf13/cobra"
)

var ingressControllers []func() *application.Application

func NewInstallIngressCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ingress",
		Short: "Ingress Controllers",
	}

	// Don't show flag errors for `install ingress` without a subcommand
	cmd.DisableFlagParsing = true

	for _, a := range ingressControllers {
		cmd.AddCommand(a().NewInstallCmd())
	}

	return cmd
}

func NewUninstallIngressCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ingress",
		Short: "Ingress Controllers",
	}

	// Don't show flag errors for `uninstall ingress` without a subcommand
	cmd.DisableFlagParsing = true

	for _, a := range ingressControllers {
		cmd.AddCommand(a().NewUninstallCmd())
	}

	return cmd
}

func init() {
	ingressControllers = []func() *application.Application{
		emissary.NewApp,
		nginx.NewApp,
	}
}
