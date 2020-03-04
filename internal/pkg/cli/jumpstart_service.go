package cli

import (
	"fmt"
)

func (o *askOpts) AskServiceCmdOpts() error {
	// TODO:
	return nil
}

func (o *askOpts) ExecuteServiceCmd() error {
	// TODO:
	cmd := ""

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *askOpts) ExecuteJumpStartServiceCmd() error {
	if err := o.AskServiceCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteServiceCmd(); err != nil {
		return err
	}

	return nil
}
