package cli

import (
	"fmt"
)

func (o *askOpts) AskNamespaceCmdOpts() error {
	if err := o.Ask("NamespaceName"); err != nil {
		return err
	}

	if err := o.AskOutputInfo(); err != nil {
		return err
	}

	return nil
}

func (o *askOpts) ExecuteNamespaceCmd() error {
	var cmd string

	// Example

	// # Create a new namespace named my-namespace
	// kubectl create namespace my-namespace

	cmd = fmt.Sprintf("kubectl create namespace %s ", o.NamespaceName)

	cmd = cmd + fmt.Sprintf("--output=%s --dry-run=true > %s", o.OutputFormat, o.OutputPath)

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *askOpts) ExecuteJumpStartNamespaceCmd() error {
	if err := o.AskNamespaceCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteNamespaceCmd(); err != nil {
		return err
	}

	return nil
}
