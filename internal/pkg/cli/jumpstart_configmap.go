package cli

import (
	"fmt"
)

func (o *askOpts) AskConfigMapCmdOpts() error {
	// TODO:
	return nil
}

func (o *askOpts) ExecuteConfigMapCmd() error {
	// TODO:
	cmd := ""

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *askOpts) ExecuteJumpStartConfigMapCmd() error {
	if err := o.AskConfigMapCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteConfigMapCmd(); err != nil {
		return err
	}

	return nil
}
