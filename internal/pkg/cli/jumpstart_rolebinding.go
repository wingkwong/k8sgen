package cli

import (
	"fmt"
)

func (o *askOpts) AskRoleBindingCmdOpts() error {
	// TODO:
	return nil
}

func (o *askOpts) ExecuteRoleBindingCmd() error {
	// TODO:
	cmd := ""

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *askOpts) ExecuteJumpStartRoleBindingCmd() error {
	if err := o.AskRoleBindingCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteRoleBindingCmd(); err != nil {
		return err
	}

	return nil
}
