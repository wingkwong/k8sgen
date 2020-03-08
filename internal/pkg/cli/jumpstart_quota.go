package cli

import (
	"fmt"
)

func (o *askOpts) AskQuotaCmdOpts() error {
	if err := o.Ask("QuotaName"); err != nil {
		return err
	}

	if err := o.Ask("Hard"); err != nil {
		return err
	}

	if err := o.Ask("Scopes"); err != nil {
		return err
	}

	if err := o.AskOutputInfo(); err != nil {
		return err
	}

	return nil
}

func (o *askOpts) ExecuteQuotaCmd() error {
	var cmd string

	// Example:

	// # Create a new resourcequota named my-quota
	// kubectl create quota my-quota --hard=cpu=1,memory=1G,pods=2,services=3,replicationcontrollers=2,resourcequotas=1,secrets=5,persistentvolumeclaims=10

	// # Create a new resourcequota named best-effort
	// kubectl create quota best-effort --hard=pods=100 --scopes=BestEffort

	cmd = fmt.Sprintf("kubectl create quota %s ", o.QuotaName)

	if o.Hard != "" {
		cmd = cmd + fmt.Sprintf("--hard=%s ", o.Hard)
	}

	if o.Scopes != "" {
		cmd = cmd + fmt.Sprintf("--scopes=%s ", o.Scopes)
	}

	cmd = cmd + fmt.Sprintf("--output=%s --dry-run=true > %s", o.OutputFormat, o.OutputPath)

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *askOpts) ExecuteJumpStartQuotaCmd() error {
	if err := o.AskQuotaCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteQuotaCmd(); err != nil {
		return err
	}

	return nil
}
