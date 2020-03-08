package cli

import (
	"fmt"
)

func (o *askOpts) AskDeploymentCmdOpts() error {
	if err := o.Ask("DeploymentName"); err != nil {
		return err
	}

	if err := o.Ask("Image"); err != nil {
		return err
	}

	if err := o.AskOutputInfo(); err != nil {
		return err
	}

	return nil
}

func (o *askOpts) ExecuteDeploymentCmd() error {
	// Example:

	// # Create a new deployment named my-dep that runs the busybox image.
	// kubectl create deployment my-dep --image=busybox

	cmd := fmt.Sprintf("kubectl create deployment %s --image=%s --output=%s --dry-run=true > %s", o.DeploymentName, o.ImageName, o.OutputFormat, o.OutputPath)

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command: \n %w", err)
	}

	return nil
}

func (o *askOpts) ExecuteJumpStartDeploymentCmd() error {
	if err := o.AskDeploymentCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteDeploymentCmd(); err != nil {
		return err
	}

	return nil
}
