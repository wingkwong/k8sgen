package cli

import (
	"fmt"
)

func (o *askOpts) AskNamespaceCmdOpts() error {
	// TODO:
	return nil
}

func (o *askOpts) ExecuteNamespaceCmd() error {
	// TODO:
	cmd := ""

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
