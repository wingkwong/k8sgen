package cli

import (
	"fmt"
)

func (o *askOpts) AskPodDisruptionBudgetCmdOpts() error {
	// TODO:
	return nil
}

func (o *askOpts) ExecutePodDisruptionBudgetCmd() error {
	// TODO:
	cmd := ""

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *askOpts) ExecuteJumpStartPodDisruptionBudgetCmd() error {
	if err := o.AskPodDisruptionBudgetCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecutePodDisruptionBudgetCmd(); err != nil {
		return err
	}

	return nil
}
