package resource

import (
	"eksdemo/pkg/cmd"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/spf13/cobra"
)

type Options interface {
	AddCreateFlags(*cobra.Command, cmd.Flags) cmd.Flags
	AddDeleteFlags(*cobra.Command, cmd.Flags) cmd.Flags
	AddGetFlags(*cobra.Command, cmd.Flags) cmd.Flags
	AddUpdateFlags(*cobra.Command, cmd.Flags) cmd.Flags
	Common() *CommonOptions
	PostCreate() error
	PreCreate() error
	PreDelete() error
	SetName(string)
	Validate(args []string) error
}

type CommonOptions struct {
	Name                string
	ArgumentOptional    bool
	ClusterFlagDisabled bool
	ClusterFlagOptional bool
	DeleteByIdFlag      bool
	KubeContext         string
	NamespaceFlag       bool

	Account           string
	Cluster           *eks.Cluster
	ClusterName       string
	DryRun            bool
	Id                string
	KubernetesVersion string
	Namespace         string
	Partition         string
	Region            string
	ServiceAccount    string
}

type Action string

const Create Action = "create"
const Delete Action = "delete"
const Get Action = "get"
const Update Action = "update"

func (o *CommonOptions) AddCreateFlags(cobraCmd *cobra.Command, flags cmd.Flags) cmd.Flags {
	flags = append(flags, o.NewDryRunFlag())

	if !o.ClusterFlagDisabled {
		flags = append(flags, o.NewClusterFlag(Create, true))
	}

	if o.NamespaceFlag {
		flags = append(flags, o.NewNamespaceFlag(Create))
	}

	for _, f := range flags {
		f.AddFlagToCommand(cobraCmd)
	}

	return flags
}

func (o *CommonOptions) AddDeleteFlags(cobraCmd *cobra.Command, flags cmd.Flags) cmd.Flags {
	if !o.ClusterFlagDisabled && !o.ClusterFlagOptional {
		flags = append(flags, o.NewClusterFlag(Delete, true))
	}

	if o.DeleteByIdFlag {
		flags = append(flags, o.NewIdFlag())
	}

	if o.NamespaceFlag {
		flags = append(flags, o.NewNamespaceFlag(Delete))
	}

	for _, f := range flags {
		f.AddFlagToCommand(cobraCmd)
	}

	return flags
}

func (o *CommonOptions) AddGetFlags(cobraCmd *cobra.Command, flags cmd.Flags) cmd.Flags {
	if o.ClusterFlagOptional {
		flags = append(flags, o.NewClusterFlag(Get, false))
	} else if !o.ClusterFlagDisabled {
		flags = append(flags, o.NewClusterFlag(Get, true))
	}

	for _, f := range flags {
		f.AddFlagToCommand(cobraCmd)
	}

	return flags
}

func (o *CommonOptions) AddUpdateFlags(cobraCmd *cobra.Command, flags cmd.Flags) cmd.Flags {
	if !o.ClusterFlagDisabled {
		flags = append(flags, o.NewClusterFlag(Update, true))
	}

	if o.NamespaceFlag {
		flags = append(flags, o.NewNamespaceFlag(Update))
	}

	for _, f := range flags {
		f.AddFlagToCommand(cobraCmd)
	}

	return flags
}

func (o *CommonOptions) Common() *CommonOptions {
	return o
}

func (o *CommonOptions) PostCreate() error {
	return nil
}

func (o *CommonOptions) PreCreate() error {
	return nil
}

func (o *CommonOptions) PreDelete() error {
	return nil
}

func (o *CommonOptions) SetName(name string) {
	o.Name = name
}

func (o *CommonOptions) Validate(args []string) error {
	return nil
}
