package cli

import (
	"fmt"
)

func (o *jumpStartOpts) AskPodDisruptionBudgetCmdOpts() error {
	// TODO:
	return nil
}

func (o *jumpStartOpts) ExecutePodDisruptionBudgetCmd() error {
	// TODO:
	cmd := ""

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *jumpStartOpts) ExecuteJumpStartPodDisruptionBudgetCmd() error {
	if err := o.AskPodDisruptionBudgetCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecutePodDisruptionBudgetCmd(); err != nil {
		return err
	}

	return nil
}
