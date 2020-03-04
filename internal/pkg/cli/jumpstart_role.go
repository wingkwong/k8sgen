package cli

import (
	"fmt"
)

func (o *askOpts) AskRoleCmdOpts() error {
	// TODO:
	return nil
}

func (o *askOpts) ExecuteRoleCmd() error {
	// TODO:
	cmd := ""

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *askOpts) ExecuteJumpStartRoleCmd() error {
	if err := o.AskRoleCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteRoleCmd(); err != nil {
		return err
	}

	return nil
}
