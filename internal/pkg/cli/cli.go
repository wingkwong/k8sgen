package cli

import (
	"github.com/spf13/cobra"
	"github.com/wingkwong/k8sgen/third_party/term/prompt"
	"os"
)

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
