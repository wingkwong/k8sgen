package cli

import (
	"fmt"
)

func (o *jumpStartOpts) AskClusterRoleCmdOpts() error {
	// TODO:
	return nil
}

func (o *jumpStartOpts) ExecuteClusterRoleCmd() error {
	// TODO:
	cmd := ""

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *jumpStartOpts) ExecuteJumpStartClusterRoleCmd() error {
	if err := o.AskClusterRoleCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteClusterRoleCmd(); err != nil {
		return err
	}

	return nil
}