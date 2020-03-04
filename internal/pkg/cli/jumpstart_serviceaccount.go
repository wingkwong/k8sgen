package cli

import (
	"fmt"
)

func (o *askOpts) AskServiceAccountCmdOpts() error {
	// TODO:
	return nil
}

func (o *askOpts) ExecuteServiceAccountCmd() error {
	// TODO:
	cmd := ""

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *askOpts) ExecuteJumpStartServiceAccountCmd() error {
	if err := o.AskServiceAccountCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteServiceAccountCmd(); err != nil {
		return err
	}

	return nil
}
