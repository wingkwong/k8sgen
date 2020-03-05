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
	// 	Options:
	//       --dry-run=false: If true, only print the object that would be sent, without sending it.
	//       --image=[]: Image name to run.
	//   -o, --output='': Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-file.

	cmd := fmt.Sprintf("kubectl create deployment %s --image=%s --output=%s --dry-run=true > %s", o.deploymentName, o.imageName, o.outputFormat, o.outputPath)

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
