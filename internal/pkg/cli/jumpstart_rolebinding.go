package cli

import (
	"fmt"
)

func (o *askOpts) AskRoleBindingCmdOpts() error {
	if err := o.Ask("RoleName"); err != nil {
		return err
	}

	if err := o.Ask("ClusterRoleName"); err != nil {
		return err
	}

	if err := o.Ask("User"); err != nil {
		return err
	}

	if err := o.Ask("Group"); err != nil {
		return err
	}

	if err := o.AskOutputInfo(); err != nil {
		return err
	}

	return nil
}

func (o *askOpts) ExecuteRoleBindingCmd() error {
	var cmd string

	// # Create a RoleBinding for user1, user2, and group1 using the admin ClusterRole
	// kubectl create rolebinding admin --clusterrole=admin --user=user1 --user=user2 --group=group1

	cmd = fmt.Sprintf("kubectl create rolebinding %s ", o.RoleName)

	if o.ClusterRoleName != "" {
		cmd = cmd + fmt.Sprintf("--clusterrole=%s ", o.ClusterRoleName)
	}

	if o.User != "" {
		cmd = cmd + fmt.Sprintf("--user=%s ", o.User)
	}

	if o.Group != "" {
		cmd = cmd + fmt.Sprintf("--group=%s ", o.Group)
	}

	cmd = cmd + fmt.Sprintf("--output=%s --dry-run=true > %s", o.OutputFormat, o.OutputPath)

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %s", cmd, err.Error())
	}

	return nil
}

func (o *askOpts) ExecuteJumpStartRoleBindingCmd() error {
	if err := o.AskRoleBindingCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteRoleBindingCmd(); err != nil {
		return err
	}

	return nil
}
