package cli

import (
	"fmt"
)

func (o *askOpts) AskClusterRoleBindingCmdOpts() error {
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

func (o *askOpts) ExecuteClusterRoleBindingCmd() error {
	var cmd string

	// # Create a ClusterRoleBinding for user1, user2, and group1 using the cluster-admin ClusterRole
	// kubectl create clusterrolebinding admin --clusterrole=admin --user=user1 --user=user2 --group=group1

	cmd = fmt.Sprintf("kubectl create clusterrolebinding %s ", o.RoleName)

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
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *askOpts) ExecuteJumpStartClusterRoleBindingCmd() error {
	if err := o.AskClusterRoleBindingCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteClusterRoleBindingCmd(); err != nil {
		return err
	}

	return nil
}
