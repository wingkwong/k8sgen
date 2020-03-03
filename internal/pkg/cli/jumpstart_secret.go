package cli

import (
	"fmt"
)

func (o *jumpStartOpts) AskSecretCmdOpts() error {
	if err := o.AskSecretCmdName(); err != nil {
		return err
	}

	switch o.secretCmdName {
	case dockerRegistryCmdName:
		if err := o.AskSecretName(); err != nil {
			return err
		}

		if err := o.AskDockerServerName(); err != nil {
			return err
		}

		if err := o.AskDockerUserName(); err != nil {
			return err
		}

		if err := o.AskDockerUserPassword(); err != nil {
			return err
		}

		if err := o.AskDockerEmail(); err != nil {
			return err
		}

	case genericCmdName:
		if err := o.AskSecretName(); err != nil {
			return err
		}

		if err := o.AskFromFilePath(); err != nil {
			return err
		}

		if err := o.AskFromLiteral(); err != nil {
			return err
		}

		// from-env-file cannot be combined with from-file or from-literal
		if len(o.fromLiteral) == 0 && len(o.fromFile) == 0 {
			if err := o.AskFromEnv(); err != nil {
				return err
			}
		}

	case tlsCmdName:
		if err := o.AskSecretName(); err != nil {
			return err
		}

		if err := o.AskCertPath(); err != nil {
			return err
		}

		if err := o.AskKeyPath(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("No available Secret option: %s", o.secretCmdName)
	}

	if err := o.AskOutputPath(); err != nil {
		return err
	}

	return nil
}

func (o *jumpStartOpts) ExecuteSecretCmd() error {

	switch o.secretCmdName {
	case dockerRegistryCmdName:
		cmd := fmt.Sprintf("kubectl create secret docker-registry %s --docker-server=%s --docker-username=%s --docker-password=%s --docker-email=%s --output=%s--dry-run=true > %s", o.secretName, o.dockerServer, o.dockerUserName, o.dockerUserPassword, o.dockerEmail, o.outputFormat, o.outputPath)

		if err := ExecCmd(cmd); err != nil {
			return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
		}

	case genericCmdName:
		var cmd string
		cmd = fmt.Sprintf("kubectl create secret generic %s ", o.secretName)

		for i := 0; i < o.noOfFromFileIteration; i++ {
			cmd = cmd + fmt.Sprintf("--from-file=%s ", o.fromFile[i])
		}

		for i := 0; i < o.noOfFromLiteralIteration; i++ {
			cmd = cmd + fmt.Sprintf("--from-literal=%s ", o.fromLiteral[i])
		}

		if len(o.fromLiteral) == 0 && len(o.fromFile) == 0 && o.fromEnvFile != "" {
			cmd = cmd + fmt.Sprintf("--from-env-file=%s ", o.fromEnvFile)
		}
	case tlsCmdName:
		cmd := fmt.Sprintf("kubectl create secret tls %s --cert=%s --key=%s --output=%s --dry-run=true > %s", o.secretName, o.certPath, o.keyPath, o.outputFormat, o.outputPath)

		if err := ExecCmd(cmd); err != nil {
			return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
		}

	default:
		return fmt.Errorf("No execution available for Secret: %s", o.secretCmdName)
	}

	return nil
}

func (o *jumpStartOpts) ExecuteJumpStartSecretCmd() error {
	if err := o.AskSecretCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteSecretCmd(); err != nil {
		return err
	}

	return nil
}
