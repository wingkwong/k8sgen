package cli

import (
	"fmt"
)

func (o *jumpStartOpts) AskRoleBindingCmdOpts() error {
	// TODO:
	return nil
}

func (o *jumpStartOpts) ExecuteRoleBindingCmd() error {
	// TODO:
	cmd := ""

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %s", cmd, err.Error())
	}

	return nil
}

func (o *jumpStartOpts) ExecuteJumpStartRoleBindingCmd() error {
	if err := o.AskRoleBindingCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteRoleBindingCmd(); err != nil {
		return err
	}

	return nil
}
