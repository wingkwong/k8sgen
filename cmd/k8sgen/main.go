package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wingkwong/k8sgen/internal/pkg/cli"
)

const k8sgenLogo = `
 _    ___                       
| | _( _ ) ___  __ _  ___ _ __  
| |/ / _ \/ __|/ _  |/ _ | |_ \ 
|   | (_) \__ | (_| |  __| | | |
|_|\_\___/|___/\__, |\___|_| |_|
	 	|___/`

func main() {
	cmd := buildRootCmd()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func buildRootCmd() *cobra.Command {
	fmt.Println(k8sgenLogo)
	cmd := &cobra.Command{
		Use:              "k8sgen",
		Short:            "Generating Kubernetes Resource Configurations in an Interactive CLI",
		Example:          `k8sgen`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {},
		SilenceUsage:     true,
	}

	cmd.AddCommand(cli.BuildJumpStartCmd())

	return cmd
}
