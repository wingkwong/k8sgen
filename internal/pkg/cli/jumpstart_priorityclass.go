package cli

import (
	"fmt"
)

func (o *askOpts) AskPriorityClassCmdOpts() error {
	// TODO:
	return nil
}

func (o *askOpts) ExecutePriorityClassCmd() error {
	// TODO:
	cmd := ""

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *askOpts) ExecuteJumpStartPriorityClassCmd() error {
	if err := o.AskPriorityClassCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecutePriorityClassCmd(); err != nil {
		return err
	}

	return nil
}
