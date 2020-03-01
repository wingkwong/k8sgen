package cli

import (
	"fmt"
)

func (o *jumpStartOpts) AskPriorityClassCmdOpts() error {
	// TODO:
	return nil
}

func (o *jumpStartOpts) ExecutePriorityClassCmd() error {
	// TODO:
	cmd := ""

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *jumpStartOpts) ExecuteJumpStartPriorityClassCmd() error {
	if err := o.AskPriorityClassCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecutePriorityClassCmd(); err != nil {
		return err
	}

	return nil
}