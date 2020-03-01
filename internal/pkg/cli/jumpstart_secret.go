package cli

import (
	"fmt"
)

func (o *jumpStartOpts) AskSecretCmdOpts() error {
	// TODO:
	return nil
}

func (o *jumpStartOpts) ExecuteSecretCmd() error {
	// TODO:
	cmd := ""

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *jumpStartOpts) ExecuteJumpStartSecretCmd() error {
	if err := o.AskSecretCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteSecretCmd(); err != nil {
		return err
	}

	return nil
}