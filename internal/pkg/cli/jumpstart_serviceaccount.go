package cli

import (
	"fmt"
)

func (o *askOpts) AskServiceAccountCmdOpts() error {
	if err := o.Ask("ServiceAccountName"); err != nil {
		return err
	}

	if err := o.AskOutputInfo(); err != nil {
		return err
	}

	return nil
}

func (o *askOpts) ExecuteServiceAccountCmd() error {
	// Example:

	// # Create a new service account named my-service-account
	// kubectl create serviceaccount my-service-account

	cmd := fmt.Sprintf("kubectl create serviceaccount %s --output=%s --dry-run=true > %s", o.ServiceAccountName, o.OutputFormat, o.OutputPath)

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command: \n %w", err)
	}

	return nil
}

func (o *askOpts) ExecuteJumpStartServiceAccountCmd() error {
	if err := o.AskServiceAccountCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteServiceAccountCmd(); err != nil {
		return err
	}

	return nil
}
