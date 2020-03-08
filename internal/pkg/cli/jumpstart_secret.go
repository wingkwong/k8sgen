package cli

import (
	"fmt"
)

func (o *askOpts) AskSecretCmdOpts() error {
	if err := o.Ask("SecretCmdName"); err != nil {
		return err
	}

	switch o.SecretCmdName {
	case dockerRegistryCmdName:
		if err := o.Ask("SecretName"); err != nil {
			return err
		}
		if err := o.Ask("DockerServerName"); err != nil {
			return err
		}
		if err := o.Ask("DockerUserName"); err != nil {
			return err
		}
		if err := o.Ask("DockerPassword"); err != nil {
			return err
		}
		if err := o.Ask("DockerEmail"); err != nil {
			return err
		}
	case genericCmdName:
		if err := o.Ask("SecretName"); err != nil {
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

	case tlsCmdName:
		if err := o.Ask("SecretName"); err != nil {
			return err
		}

		if err := o.Ask("CertPath"); err != nil {
			return err
		}

		if err := o.Ask("KeyPath"); err != nil {
			return err
		}
	default:
		return fmt.Errorf("No available Secret option: %s", o.SecretCmdName)
	}

	if err := o.Ask("AppendHash"); err != nil {
		return err
	}

	if err := o.AskOutputInfo(); err != nil {
		return err
	}

	return nil
}

func (o *askOpts) ExecuteSecretCmd() error {
	var cmd string

	switch o.SecretCmdName {
	case dockerRegistryCmdName:
		cmd = fmt.Sprintf("kubectl create secret docker-registry %s --docker-server=%s --docker-username=%s --docker-password=%s --docker-email=%s --output=%s--dry-run=true > %s", o.SecretName, o.DockerServer, o.DockerUserName, o.DockerUserPassword, o.DockerEmail, o.OutputFormat, o.OutputPath)
	case genericCmdName:
		cmd = fmt.Sprintf("kubectl create secret generic %s ", o.SecretName)

		for i := 0; i < len(o.FromFile); i++ {
			cmd = cmd + fmt.Sprintf("--from-file=%s ", o.FromFile[i])
		}

		for i := 0; i < len(o.FromLiteral); i++ {
			cmd = cmd + fmt.Sprintf("--from-literal=%s ", o.FromLiteral[i])
		}

		if len(o.FromLiteral) == 0 && len(o.FromFile) == 0 && o.FromEnvFile != "" {
			cmd = cmd + fmt.Sprintf("--from-env-file=%s ", o.FromEnvFile)
		}

		cmd = cmd + fmt.Sprintf("--output=%s > %s", o.OutputFormat, o.OutputPath)
	case tlsCmdName:
		cmd = fmt.Sprintf("kubectl create secret tls %s --cert=%s --key=%s --output=%s --dry-run=true > %s", o.SecretName, o.CertPath, o.KeyPath, o.OutputFormat, o.OutputPath)

	default:
		return fmt.Errorf("No execution available for Secret: %s", o.SecretCmdName)
	}

	if o.AppendHash {
		cmd = cmd + fmt.Sprintf("--append-hash=%t ", o.AppendHash)
	}

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *askOpts) ExecuteJumpStartSecretCmd() error {
	if err := o.AskSecretCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteSecretCmd(); err != nil {
		return err
	}

	return nil
}
