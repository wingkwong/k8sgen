package cli

import (
	"fmt"
)

func (o *jumpStartOpts) AskServiceAccountCmdOpts() error {
	// TODO:
	return nil
}

func (o *jumpStartOpts) ExecuteServiceAccountCmd() error {
	// TODO:
	cmd := ""

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *jumpStartOpts) ExecuteJumpStartServiceAccountCmd() error {
	if err := o.AskServiceAccountCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteServiceAccountCmd(); err != nil {
		return err
	}

	return nil
}