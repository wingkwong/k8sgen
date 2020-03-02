package cli

import (
	"fmt"
)

func (o *jumpStartOpts) AskQuotaCmdOpts() error {
	// TODO:
	return nil
}

func (o *jumpStartOpts) ExecuteQuotaCmd() error {
	// TODO:
	cmd := ""

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *jumpStartOpts) ExecuteJumpStartQuotaCmd() error {
	if err := o.AskQuotaCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteQuotaCmd(); err != nil {
		return err
	}

	return nil
}
