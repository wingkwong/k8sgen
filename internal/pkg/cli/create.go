package cli

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

type createVars struct {
	*GlobalOpts
	DeploymentCmdOpts
	KindName string
}

type createOpts struct {
	createVars
}

func newCreateOpts(vars createVars) (*createOpts, error) {
	return &createOpts{
		createVars: vars,
	}, nil
}

func (o *createOpts) Ask() error {
	if err := o.askKindName(); err != nil {
		return err
	}

	return nil
}

func (o *createOpts) askKindName() error {
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

func (o *createOpts) Execute() error {
	switch o.KindName {
	case deploymentName:
		if err := o.ExecuteCreateDeploymentCmd(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("No execution available for kind: %s", o.KindName)
	}

	return nil
}

func BuildCreateCmd() *cobra.Command {
	vars := createVars{
		GlobalOpts: NewGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create your Kubenetes resources",
		Example: `
  $ k8sgen create`,
		RunE: runCmdE(func(cmd *cobra.Command, args []string) error {
			opts, err := newCreateOpts(vars)
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
