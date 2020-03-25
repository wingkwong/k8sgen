package cli

import (
	"fmt"
)

func (o *askOpts) AskClusterRoleCmdOpts() error {
	if err := o.Ask("ClusterRoleName"); err != nil {
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

	if err := o.Ask("AggregationRule"); err != nil {
		return err
	}

	if err := o.AskOutputInfo(); err != nil {
		return err
	}

	return nil
}

func (o *askOpts) ExecuteClusterRoleCmd() error {
	var cmd string

	// Example:

	// # Create a ClusterRole named "pod-reader" that allows user to perform "get", "watch" and "list" on pods
	// kubectl create clusterrole pod-reader --verb=get,list,watch --resource=pods

	// # Create a ClusterRole named "pod-reader" with ResourceName specified
	// kubectl create clusterrole pod-reader --verb=get --resource=pods --resource-name=readablepod --resource-name=anotherpod

	// # Create a ClusterRole named "foo" with API Group specified
	// kubectl create clusterrole foo --verb=get,list,watch --resource=rs.extensions

	// # Create a ClusterRole named "foo" with SubResource specified
	// kubectl create clusterrole foo --verb=get,list,watch --resource=pods,pods/status

	// # Create a ClusterRole name "foo" with NonResourceURL specified
	// kubectl create clusterrole "foo" --verb=get --non-resource-url=/logs/*

	// # Create a ClusterRole name "monitoring" with AggregationRule specified
	// kubectl create clusterrole monitoring --aggregation-rule="rbac.example.com/aggregate-to-monitoring=true"

	cmd = fmt.Sprintf("kubectl create clusterrole %s ", o.RoleName)

	if o.Resource != "" {
		cmd = cmd + fmt.Sprintf("--resource=%s ", o.Resource)
	}

	if o.ResourceName != "" {
		cmd = cmd + fmt.Sprintf("--resource-name=%s ", o.ResourceName)
	}

	if o.Verb != "" {
		cmd = cmd + fmt.Sprintf("--verb=%s ", o.Verb)
	}

	if o.AggregationRule != "" {
		cmd = cmd + fmt.Sprintf("--aggregation-rule=%s ", o.AggregationRule)
	}

	cmd = cmd + fmt.Sprintf("--output=%s --dry-run=true > %s", o.OutputFormat, o.OutputPath)

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *askOpts) ExecuteJumpStartClusterRoleCmd() error {
	if err := o.AskClusterRoleCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteClusterRoleCmd(); err != nil {
		return err
	}

	return nil
}
