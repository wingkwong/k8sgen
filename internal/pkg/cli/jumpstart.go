package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

func (o *askOpts) AskJumpStart() error {
	if err := o.askKindName(); err != nil {
		return err
	}

	return nil
}

func (o *askOpts) ExecuteJumpStart() error {
	switch o.KindName {
	case clusterRoleName:
		if err := o.ExecuteJumpStartClusterRoleCmd(); err != nil {
			return err
		}
	case clusterRoleBindingName:
		if err := o.ExecuteJumpStartClusterRoleBindingCmd(); err != nil {
			return err
		}
	case configmapName:
		if err := o.ExecuteJumpStartConfigMapCmd(); err != nil {
			return err
		}
	case deploymentName:
		if err := o.ExecuteJumpStartDeploymentCmd(); err != nil {
			return err
		}
	case jobName:
		if err := o.ExecuteJumpStartJobCmd(); err != nil {
			return err
		}
	case namespaceName:
		if err := o.ExecuteJumpStartNamespaceCmd(); err != nil {
			return err
		}
	case podDisruptionBudgetName:
		if err := o.ExecuteJumpStartPodDisruptionBudgetCmd(); err != nil {
			return err
		}
	case priorityClassName:
		if err := o.ExecuteJumpStartPriorityClassCmd(); err != nil {
			return err
		}
	case quotaName:
		if err := o.ExecuteJumpStartQuotaCmd(); err != nil {
			return err
		}
	case roleName:
		if err := o.ExecuteJumpStartRoleCmd(); err != nil {
			return err
		}
	case roleBindingName:
		if err := o.ExecuteJumpStartRoleBindingCmd(); err != nil {
			return err
		}
	case secretName:
		if err := o.ExecuteJumpStartSecretCmd(); err != nil {
			return err
		}
	case serviceName:
		if err := o.ExecuteJumpStartServiceCmd(); err != nil {
			return err
		}
	case serviceAccountName:
		if err := o.ExecuteJumpStartServiceAccountCmd(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("No execution available for kind: %s", o.KindName)
	}

	return nil
}

func BuildJumpStartCmd() *cobra.Command {
	vars := askVars{
		GlobalOpts: NewGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:   "jumpstart",
		Short: "Jumpstart your Kubernetes Resources",
		Example: `
  $ k8sgen jumpstart`,
		RunE: runCmdE(func(cmd *cobra.Command, args []string) error {
			if err := VerifyCmd("kubectl"); err != nil {
				return fmt.Errorf("kubectl is not installed")
			}

			opts, err := newAskOpts(vars)
			if err != nil {
				return err
			}

			if err := opts.AskJumpStart(); err != nil {
				return err
			}
			if err := opts.ExecuteJumpStart(); err != nil {
				return err
			}
			return nil
		}),
	}

	return cmd
}
