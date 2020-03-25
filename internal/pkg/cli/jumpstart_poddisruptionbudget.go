package cli

import (
	"fmt"
)

func (o *askOpts) AskPodDisruptionBudgetCmdOpts() error {
	if err := o.Ask("PodDisruptionBudgetName"); err != nil {
		return err
	}

	if err := o.Ask("MaxUnavailable"); err != nil {
		return err
	}

	if err := o.Ask("MinAvailable"); err != nil {
		return err
	}

	if err := o.Ask("Selector"); err != nil {
		return err
	}

	if err := o.AskOutputInfo(); err != nil {
		return err
	}
	return nil
}

func (o *askOpts) ExecutePodDisruptionBudgetCmd() error {
	var cmd string

	// # Create a pod disruption budget named my-pdb that will select all pods with the app=rails label
	// # and require at least one of them being available at any point in time.
	// kubectl create poddisruptionbudget my-pdb --selector=app=rails --min-available=1

	// # Create a pod disruption budget named my-pdb that will select all pods with the app=nginx label
	// # and require at least half of the pods selected to be available at any point in time.
	// kubectl create pdb my-pdb --selector=app=nginx --min-available=50%

	cmd = fmt.Sprintf("kubectl create poddisruptionbudget %s ", o.PodDisruptionBudgetName)

	if o.MaxUnavailable != "" {
		cmd = cmd + fmt.Sprintf("--max-unavailable=%s ", o.MaxUnavailable)
	}

	if o.MinAvailable != "" {
		cmd = cmd + fmt.Sprintf("--min-available=%s ", o.MinAvailable)
	}

	if o.Selector != "" {
		cmd = cmd + fmt.Sprintf("--selector=%s ", o.Selector)
	}

	cmd = cmd + fmt.Sprintf("--output=%s --dry-run=true > %s", o.OutputFormat, o.OutputPath)

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
