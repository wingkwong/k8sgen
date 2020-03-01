package cli

import (
	"fmt"
)

func (o *jumpStartOpts) AskClusterRoleBindingCmdOpts() error {
	// TODO:
	return nil
}

func (o *jumpStartOpts) ExecuteClusterRoleBindingCmd() error {
	// TODO:
	cmd := ""

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *jumpStartOpts) ExecuteJumpStartClusterRoleBindingCmd() error {
	if err := o.AskClusterRoleBindingCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteClusterRoleBindingCmd(); err != nil {
		return err
	}

	return nil
}