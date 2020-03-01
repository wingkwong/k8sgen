package cli

import (
	"fmt"
)

func (o *jumpStartOpts) AskJobCmdOpts() error {
	// TODO:
	return nil
}

func (o *jumpStartOpts) ExecuteJobCmd() error {
	// TODO:
	cmd := ""

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *jumpStartOpts) ExecuteJumpStartJobCmd() error {
	if err := o.AskJobCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteJobCmd(); err != nil {
		return err
	}

	return nil
}