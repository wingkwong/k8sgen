package cli

import (
	"fmt"
)

func (o *jumpStartOpts) AskConfigMapCmdOpts() error {
	// TODO:
	return nil
}

func (o *jumpStartOpts) ExecuteConfigMapCmd() error {
	// TODO:
	cmd := ""

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *jumpStartOpts) ExecuteJumpStartConfigMapCmd() error {
	if err := o.AskConfigMapCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteConfigMapCmd(); err != nil {
		return err
	}

	return nil
}