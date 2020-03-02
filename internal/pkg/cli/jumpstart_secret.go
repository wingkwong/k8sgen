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

		if err := o.AskSecretGenericOpts(); err != nil {
			return err
		}

		if err := o.AskSecretName(); err != nil {
			return err
		}

		switch o.secretCmdName {
		// Create a new secret with keys for each file in folder
		case genericOpt1:
			if err := o.AskFromFilePath(); err != nil {
				return err
			}
		// Create a new secret with specified keys instead of names on disk
		case genericOpt2:
			// TODO: allow multiple --from-file
			if err := o.AskFromFilePath(); err != nil {
				return err
			}
		// Create a new secret with keys
		case genericOpt3:
			// TODO: allow multiple --from-literal
			if err := o.AskFromLiteral(); err != nil {
				return err
			}
		// Create a new secret using a combination of a file and a literal
		case genericOpt4:
			if err := o.AskFromFilePath(); err != nil {
				return err
			}

			if err := o.AskFromLiteral(); err != nil {
				return err
			}
		// Create a new secret from an env file
		case genericOpt5:
			if err := o.AskFromEnv(); err != nil {
				return err
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
		// 	# Create a new secret named my-secret with keys for each file in folder bar
		// 	kubectl create secret generic my-secret --from-file=path/to/bar

		// 	# Create a new secret named my-secret with specified keys instead of names on disk
		// 	kubectl create secret generic my-secret --from-file=ssh-privatekey=path/to/id_rsa
		//   --from-file=ssh-publickey=path/to/id_rsa.pub

		// 	# Create a new secret named my-secret with key1=supersecret and key2=topsecret
		// 	kubectl create secret generic my-secret --from-literal=key1=supersecret --from-literal=key2=topsecret

		// 	# Create a new secret named my-secret using a combination of a file and a literal
		// 	kubectl create secret generic my-secret --from-file=ssh-privatekey=path/to/id_rsa --from-literal=passphrase=topsecret

		// 	# Create a new secret named my-secret from an env file
		// 	kubectl create secret generic my-secret --from-env-file=path/to/bar.env

		switch o.secretCmdName {
		// Create a new secret with keys for each file in folder
		case genericOpt1:
			//TODO
			return nil
		// Create a new secret with specified keys instead of names on disk
		case genericOpt2:
			//TODO
			return nil

		// Create a new secret with keys
		case genericOpt3:
			//TODO
			return nil

		// Create a new secret using a combination of a file and a literal
		case genericOpt4:
			//TODO
			return nil

		// Create a new secret from an env file
		case genericOpt5:
			//TODO
			return nil
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
