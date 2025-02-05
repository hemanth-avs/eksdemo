package resource

import (
	"eksdemo/pkg/cmd"
	"eksdemo/pkg/printer"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type Resource struct {
	cmd.Command
	Options

	CreateFlags cmd.Flags
	DeleteFlags cmd.Flags
	GetFlags    cmd.Flags
	UpdateFlags cmd.Flags

	Getter
	Manager
}

func (r *Resource) Create() error {
	return r.Manager.Create(r.Options)
}

func (r *Resource) Delete() error {
	return r.Manager.Delete(r.Options)
}

func (r *Resource) Update(cmd *cobra.Command) error {
	return r.Manager.Update(r.Options, cmd)
}

func (r *Resource) NewCreateCmd() *cobra.Command {
	use := r.Command.Name
	if len(r.Args) > 0 {
		use += " " + strings.Join(r.Args, " ")
	}

	cmd := &cobra.Command{
		Use:     use,
		Short:   r.Description,
		Long:    "Create " + r.Description,
		Aliases: r.Aliases,
		Args:    cobra.ExactArgs(len(r.Args)),
		Hidden:  r.Hidden,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := r.CreateFlags.ValidateFlags(cmd, args); err != nil {
				return err
			}

			if err := r.Options.Validate(args); err != nil {
				return err
			}

			cmd.SilenceUsage = true
			if len(r.Args) > 0 {
				r.SetName(args[0])
			}

			if r.Common().DryRun {
				r.SetDryRun()
			}

			if err := r.PreCreate(); err != nil {
				return err
			}

			if r.Manager == nil {
				return fmt.Errorf("feature not yet implemented")
			}

			if err := r.Create(); err != nil {
				return err
			}

			return r.PostCreate()
		},
	}
	r.CreateFlags = r.Options.AddCreateFlags(cmd, r.CreateFlags)

	return cmd
}

func (r *Resource) NewDeleteCmd() *cobra.Command {
	var args cobra.PositionalArgs
	use := r.Command.Name

	if len(r.Args) > 0 && r.Common().ArgumentOptional {
		use += " " + "[" + r.Args[0] + "]"
	} else if len(r.Args) > 0 {
		use += " " + strings.Join(r.Args, " ")
	}

	if r.Common().ArgumentOptional {
		args = cobra.RangeArgs(0, len(r.Args))
	} else {
		args = cobra.ExactArgs(len(r.Args))
	}

	cmd := &cobra.Command{
		Use:     use,
		Short:   r.Description,
		Long:    "Delete " + r.Description,
		Aliases: r.Aliases,
		Args:    args,
		Hidden:  r.Hidden,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := r.DeleteFlags.ValidateFlags(cmd, args); err != nil {
				return err
			}

			if r.Common().DeleteByIdFlag && r.Common().Id == "" && len(r.Args) > 0 && len(args) == 0 {
				return fmt.Errorf("must include either %s or --id flag", r.Args[0])
			}
			cmd.SilenceUsage = true

			if len(args) > 0 {
				r.SetName(args[0])
			}

			if err := r.PreDelete(); err != nil {
				return err
			}

			if r.Manager == nil {
				return fmt.Errorf("feature not yet implemented")
			}

			return r.Delete()
		},
	}
	r.DeleteFlags = r.Options.AddDeleteFlags(cmd, r.DeleteFlags)

	return cmd
}

func (r *Resource) NewGetCmd() *cobra.Command {
	var args cobra.PositionalArgs
	var output printer.Output
	use := r.Command.Name

	if len(r.Args) > 0 {
		use += " " + "[" + r.Args[0] + "]"
	}

	if len(r.Args) == 0 {
		args = cobra.NoArgs
	} else {
		args = cobra.RangeArgs(0, 1)
	}

	cobraCmd := &cobra.Command{
		Use:     use,
		Short:   r.Description,
		Long:    "Get " + r.Description,
		Aliases: r.Aliases,
		Args:    args,
		Hidden:  r.Hidden,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := r.GetFlags.ValidateFlags(cmd, args); err != nil {
				return err
			}

			if err := r.Options.Validate(args); err != nil {
				return err
			}
			cmd.SilenceUsage = true

			if r.Getter == nil {
				return fmt.Errorf("feature not yet implemented")
			}

			name := ""
			if len(args) > 0 {
				name = args[0]
			}

			return r.Getter.Get(name, output, r.Options)
		},
	}
	cobraCmd.Flags().VarP(cmd.NewOutputFlag(&output), "output", "o", "output format (json|table|yaml)")

	r.GetFlags = r.Options.AddGetFlags(cobraCmd, r.GetFlags)

	return cobraCmd
}

func (r *Resource) NewUpdateCmd() *cobra.Command {
	use := r.Command.Name
	if len(r.Args) > 0 {
		use += " " + strings.Join(r.Args, " ")
	}

	cobraCmd := &cobra.Command{
		Use:     use,
		Short:   r.Description,
		Long:    "Update " + r.Description,
		Aliases: r.Aliases,
		Args:    cobra.ExactArgs(len(r.Args)),
		Hidden:  r.Hidden,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := r.UpdateFlags.ValidateFlags(cmd, args); err != nil {
				return err
			}
			cmd.SilenceUsage = true

			if len(args) > 0 {
				r.SetName(args[0])
			}

			if r.Manager == nil {
				return fmt.Errorf("feature not yet implemented")
			}

			return r.Update(cmd)
		},
	}

	r.UpdateFlags = r.Options.AddUpdateFlags(cobraCmd, r.UpdateFlags)

	return cobraCmd
}
