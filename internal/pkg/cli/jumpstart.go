package cli

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wingkwong/k8sgen/internal/pkg/cli/jumpstart"
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
