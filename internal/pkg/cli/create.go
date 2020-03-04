package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

func (o *askOpts) AskCreateSpec() error {
	if err := o.Ask("Kind"); err != nil {
		return err
	}

	return nil
}

func (o *askOpts) ExecuteCreateSpec() error {
	switch o.KindName {
	case deploymentName:
		if err := o.ExecuteCreateDeploymentSpec(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("No execution available for kind: %s", o.KindName)
	}

	return nil
}

func BuildCreateCmd() *cobra.Command {
	vars := askVars{
		GlobalOpts: NewGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create your Kubernetes resources",
		Example: `
  $ k8sgen create`,
		RunE: runCmdE(func(cmd *cobra.Command, args []string) error {
			opts, err := newAskOpts(vars)
			if err != nil {
				return err
			}
			if err := opts.AskCreateSpec(); err != nil {
				return err
			}
			if err := opts.ExecuteCreateSpec(); err != nil {
				return err
			}
			return nil
		}),
	}

	return cmd
}
