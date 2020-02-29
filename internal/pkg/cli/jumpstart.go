package cli

import (
	"github.com/k8sgen/internal/pkg/cli/jumpstart"
	"github.com/spf13/cobra"
)

func BuildJumpStartCmd() *cobra.Command {
	vars := jumpStartVars{
		GlobalOpts: NewGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:   "jumpstart",
		Short: "Jumpstart your Kubenetes resources",
		Example: `
  $ k8sgen jumpstart`,
		RunE: runCmdE(func(cmd *cobra.Command, args []string) error {
			opts, err := newJumpStartOpts(vars)
			if err != nil {
				return err
			}

			if err := opts.Ask(); err != nil {
				return err
			}
			if err := opts.Execute(); err != nil {
				return err
			}
			return nil
		}),
	}

	return cmd
}
