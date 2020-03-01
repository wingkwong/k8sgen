package cli

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

type DploymentCmdOpts struct {
	deploymentName string
	imageName string
	outputFormat string
	outputPath string
}

type jumpStartVars struct {
	*GlobalOpts
	DploymentCmdOpts
	KindName          string
}

type jumpStartOpts struct {
	jumpStartVars
}

func newJumpStartOpts(vars jumpStartVars) (*jumpStartOpts, error) {
	return &jumpStartOpts{
		jumpStartVars: vars,
	}, nil
}

func (o *jumpStartOpts) Ask() error {
	if err := o.askKindName(); err != nil {
		return err
	}

	return nil
}

func (o *jumpStartOpts) askKindName() error {
	if o.KindName != "" {
		return nil
	}

	names := getKindNames()

	if len(names) == 0 {
		return errors.New("No object is found")
	}

	selectedKindName, err := o.prompt.SelectOne("What kind of object you want to create?", "", names)
	if err != nil {
		return fmt.Errorf("Select kind name: %w", err)
	}
	o.KindName = selectedKindName
	return nil
}

func (o *jumpStartOpts) Execute() error {
	switch o.KindName {
		case deploymentName:
			if err := o.ExecuteJumpStartDeploymentCmd(); err != nil {
				return err
			}
		default:
			return fmt.Errorf("No execution available for kind: %s", o.KindName)
	}

	return nil
}

func BuildJumpStartCmd() *cobra.Command {
	vars := jumpStartVars{
		GlobalOpts: NewGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:   "jumpstart",
		Short: "Jumpstart your Kubernetes Resources",
		Example: `
  $ k8sgen jumpstart`,
		RunE: runCmdE(func(cmd *cobra.Command, args []string) error {
			opts, err := newJumpStartOpts(vars)
			if err != nil {
				return err
			}

			if err := opts.Ask(); err != nil {
				return err
			}
			if err := opts.Execute(); err != nil {
				return err
			}
			return nil
		}),
	}

	return cmd
}
