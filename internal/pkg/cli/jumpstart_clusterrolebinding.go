package cli

import (
	"fmt"
)

func (o *askOpts) AskClusterRoleBindingCmdOpts() error {
	// TODO:
	return nil
}

func (o *askOpts) ExecuteClusterRoleBindingCmd() error {
	// TODO:
	cmd := ""

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *askOpts) ExecuteJumpStartClusterRoleBindingCmd() error {
	if err := o.AskClusterRoleBindingCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteClusterRoleBindingCmd(); err != nil {
		return err
	}

	return nil
}
