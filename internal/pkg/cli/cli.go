package cli

import (
	"os"
	"github.com/k8sgen/third_party/term/prompt"
	"github.com/spf13/cobra"
)

type GlobalOpts struct {
	prompt prompter
}

type DeploymentCmdOpts struct {
	deploymentName string
	imageName string
	outputFormat string
	outputPath string
}

func NewGlobalOpts() *GlobalOpts {
	return &GlobalOpts{
		prompt: prompt.New(),
	}
}

func runCmdE(f func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 && args[0] == "help" {
			_ = cmd.Help()
			os.Exit(0)
		}
		return f(cmd, args)
	}
}
