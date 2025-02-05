package ec2_instance

import (
	"eksdemo/pkg/aws"
	"eksdemo/pkg/resource"
	"fmt"

	"github.com/spf13/cobra"
)

type Manager struct {
	DryRun bool
	Getter
}

func (m *Manager) Create(options resource.Options) error {
	return fmt.Errorf("feature not supported")
}

func (m *Manager) Delete(options resource.Options) (err error) {
	instanceId := options.Common().Name

	ec2, err := m.GetInstanceById(instanceId)
	if err != nil {
		return err
	}

	if aws.StringValue(ec2.State.Name) == "terminated" {
		return fmt.Errorf("ec2-instance %q already terminated", instanceId)
	}

	if err := aws.EC2TerminateInstances(instanceId); err != nil {
		return err
	}
	fmt.Println("EC2 Instance terminating...")

	return nil
}

func (m *Manager) SetDryRun() {
	m.DryRun = true
}

func (m *Manager) Update(options resource.Options, cmd *cobra.Command) error {
	return fmt.Errorf("feature not supported")
}
