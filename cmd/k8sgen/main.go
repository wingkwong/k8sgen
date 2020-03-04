package main

import (
	"github.com/wingkwong/k8sgen/internal/pkg/cli"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := buildRootCmd()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func buildRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "k8sgen",
		Short:            "Generating Kubernetes Resource Configurations in an Interactive CLI",
		Example:          `k8sgen`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {},
		SilenceUsage:     true,
	}

	cmd.AddCommand(cli.BuildCreateCmd())
	cmd.AddCommand(cli.BuildJumpStartCmd())

	return cmd
}
