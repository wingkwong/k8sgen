package jumpstart

import (
	"github.com/k8sgen/internal/pkg/cli"
)

type DploymentCmdOpts struct {
	deploymentName string
	imageName string
	outputFormat string
	outputPath string
}

type jumpStartVars struct {
	*GlobalOpts
	*DploymentCmdOpts
	KindName          string
}

type jumpStartOpts struct {
	jumpStartVars
}

func newJumpStartOpts(vars jumpStartVars) (*jumpStartOpts, error) {
	return &deleteAppOpts{
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
		case clusterRoleName:
			if err := ExecuteClusterRoleCmd(&o); err != nil {
				return err
			}
		default:
			return fmt.Errorf("No execution available for kind: %s", o.KindName)
	}

	return nil
}