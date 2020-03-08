package cli

import (
	"fmt"
)

func (o *askOpts) AskConfigMapCmdOpts() error {
	if err := o.Ask("ConfigMapName"); err != nil {
		return err
	}

	if err := o.Ask("AppendHash"); err != nil {
		return err
	}

	if err := o.AskWithIterator("FromFile"); err != nil {
		return err
	}

	if err := o.AskWithIterator("FromLiteral"); err != nil {
		return err
	}

	// from-env-file cannot be combined with from-file or from-literal
	if len(o.FromLiteral) == 0 && len(o.FromFile) == 0 {
		if err := o.Ask("FromEnvFile"); err != nil {
			return err
		}
	}

	if err := o.AskOutputInfo(); err != nil {
		return err
	}

	return nil
}

func (o *askOpts) ExecuteConfigMapCmd() error {
	var cmd string

	// Example:

	// # Create a new configmap named my-config based on folder bar
	// kubectl create configmap my-config --from-file=path/to/bar

	// # Create a new configmap named my-config with specified keys instead of file basenames on disk
	// kubectl create configmap my-config --from-file=key1=/path/to/bar/file1.txt --from-file=key2=/path/to/bar/file2.txt

	// # Create a new configmap named my-config with key1=config1 and key2=config2
	// kubectl create configmap my-config --from-literal=key1=config1 --from-literal=key2=config2

	// # Create a new configmap named my-config from the key=value pairs in the file
	// kubectl create configmap my-config --from-file=path/to/bar

	// # Create a new configmap named my-config from an env file
	// kubectl create configmap my-config --from-env-file=path/to/bar.env

	cmd = fmt.Sprintf("kubectl create configmap %s ", o.ConfigMapName)

	for i := 0; i < len(o.FromFile); i++ {
		cmd = cmd + fmt.Sprintf("--from-file=%s ", o.FromFile[i])
	}

	for i := 0; i < len(o.FromLiteral); i++ {
		cmd = cmd + fmt.Sprintf("--from-literal=%s ", o.FromLiteral[i])
	}

	if len(o.FromLiteral) == 0 && len(o.FromFile) == 0 && o.FromEnvFile != "" {
		cmd = cmd + fmt.Sprintf("--from-env-file=%s ", o.FromEnvFile)
	}

	if o.AppendHash {
		cmd = cmd + fmt.Sprintf("--append-hash=%t ", o.AppendHash)
	}

	cmd = cmd + fmt.Sprintf("--output=%s > %s", o.OutputFormat, o.OutputPath)

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *askOpts) ExecuteJumpStartConfigMapCmd() error {
	if err := o.AskConfigMapCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteConfigMapCmd(); err != nil {
		return err
	}

	return nil
}
