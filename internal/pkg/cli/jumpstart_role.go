package cli

import (
	"fmt"
)

func (o *askOpts) AskRoleCmdOpts() error {
	if err := o.Ask("RoleName"); err != nil {
		return err
	}

	if err := o.Ask("Resource"); err != nil {
		return err
	}

	if err := o.Ask("ResourceName"); err != nil {
		return err
	}

	if err := o.Ask("Verb"); err != nil {
		return err
	}

	if err := o.AskOutputInfo(); err != nil {
		return err
	}

	return nil
}

func (o *askOpts) ExecuteRoleCmd() error {
	var cmd string

	// Example:

	// # Create a Role named "pod-reader" that allows user to perform "get", "watch" and "list" on pods
	// kubectl create role pod-reader --verb=get --verb=list --verb=watch --resource=pods

	// # Create a Role named "pod-reader" with ResourceName specified
	// kubectl create role pod-reader --verb=get --resource=pods --resource-name=readablepod --resource-name=anotherpod

	// # Create a Role named "foo" with API Group specified
	// kubectl create role foo --verb=get,list,watch --resource=rs.extensions

	// # Create a Role named "foo" with SubResource specified
	// kubectl create role foo --verb=get,list,watch --resource=pods,pods/status

	cmd = fmt.Sprintf("kubectl create role %s ", o.RoleName)

	if o.Resource != "" {
		cmd = cmd + fmt.Sprintf("--resource=%s ", o.Resource)
	}

	if o.ResourceName != "" {
		cmd = cmd + fmt.Sprintf("--resource-name=%s ", o.ResourceName)
	}

	if o.Verb != "" {
		cmd = cmd + fmt.Sprintf("--verb=%s ", o.Verb)
	}

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	cmd = cmd + fmt.Sprintf("--output=%s > %s", o.OutputFormat, o.OutputPath)

	return nil
}

func (o *askOpts) ExecuteJumpStartRoleCmd() error {
	if err := o.AskRoleCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteRoleCmd(); err != nil {
		return err
	}

	return nil
}
