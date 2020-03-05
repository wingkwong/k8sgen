package cli

import (
	"fmt"
)

func (o *jumpStartOpts) AskRoleCmdOpts() error {
	// TODO:
	return nil
}

func (o *jumpStartOpts) ExecuteRoleCmd() error {
	// TODO:
	cmd := ""

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %s", cmd, err.Error())
	}

	return nil
}

func (o *jumpStartOpts) ExecuteJumpStartRoleCmd() error {
	if err := o.AskRoleCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteRoleCmd(); err != nil {
		return err
	}

	return nil
}
