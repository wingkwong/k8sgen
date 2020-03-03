package cli

import (
	"fmt"
	"strconv"
)

const (
	// text input
	inputDeploymentNamePrompt           = "What deployment you want to name?"
	inputImageNamePrompt                = "What image you want to name to run?"
	inputOutputPathPrompt               = "What directory you want to save?"
	inputSecretNamePrompt               = "What secret you want to name?"
	inputDockerServerNamePrompt         = "What docker server you want to name?"
	inputDockerUserNamePrompt           = "What is the username for Docker registry authentiation?"
	inputDockerUserPasswordPrompt       = "What is the password for Docker registry authentiation?"
	inputDockerEmailPrompt              = "What is the email address for Docker registry?"
	inputFromFilePrompt                 = "Input key names from file: (e.g path/to/bar):"
	inputFromLiteralPrompt              = "Input a key-value pair secret (e.g foo='bar'):"
	inputFromEnvFilePrompt              = "Where is the env file path?"
	inputNoOfFromFileIterationPrompt    = "How many from-file iterations for your input?"
	inputNoOfFromLiteralIterationPrompt = "How many from-literal iterations for your input?"

	// select
	inputOutputFormatPrompt  = "Please select an output format:"
	inputSecretCmdNamePrompt = "Please select the type of secret:"
)

func (o *jumpStartOpts) AskDeploymentName() error {
	deploymentName, err := o.prompt.Get(inputDeploymentNamePrompt, "", validateDeploymentName)

	if err != nil {
		return fmt.Errorf("Prompt for deployment name: %w", err)
	}

	o.deploymentName = deploymentName

	return nil
}

func (o *jumpStartOpts) AskImageName() error {
	imageName, err := o.prompt.Get(inputImageNamePrompt, "", nil /*no validation*/)
	if err != nil {
		return fmt.Errorf("Prompt for image name: %w", err)
	}
	o.imageName = imageName

	return nil
}

func (o *jumpStartOpts) AskOutputFormat() error {
	formats := getOutputFormats()
	outputFormat, err := o.prompt.SelectOne(inputOutputFormatPrompt, "", formats)
	if err != nil {
		return fmt.Errorf("Prompt for output format: %w", err)
	}
	o.outputFormat = outputFormat

	return nil
}

func (o *jumpStartOpts) AskOutputPath() error {
	outputPath, err := o.prompt.Get(inputOutputPathPrompt, "", nil)
	if err != nil {
		return fmt.Errorf("Prompt for output path: %w", err)
	}

	// if err := VerifyDirectory(outputPath); err != nil {
	// 	return fmt.Errorf("Failed to verify directory: %w", err)
	// }

	o.outputPath = outputPath
	return nil
}

func (o *jumpStartOpts) AskSecretCmdName() error {
	secrets := getSecretNames()
	secretCmdName, err := o.prompt.SelectOne(inputSecretCmdNamePrompt, `
docker-registry Create a secret for use with a Docker registry
generic         Create a secret from a local file, directory or literal value
tls             Create a TLS secret`, secrets)

	if err != nil {
		return fmt.Errorf("Prompt for secret cmd name: %w", err)
	}

	o.secretCmdName = secretCmdName

	return nil
}

func (o *jumpStartOpts) AskSecretName() error {
	secretName, err := o.prompt.Get(inputSecretNamePrompt, "", nil /*no validation*/)
	if err != nil {
		return fmt.Errorf("Prompt for secret name: %w", err)
	}
	o.secretName = secretName

	return nil
}

func (o *jumpStartOpts) AskDockerServerName() error {
	dockerServer, err := o.prompt.Get(inputDockerServerNamePrompt, "", nil /*no validation*/)
	if err != nil {
		return fmt.Errorf("Prompt for docker server name: %w", err)
	}
	o.dockerServer = dockerServer

	return nil
}

func (o *jumpStartOpts) AskDockerUserName() error {
	dockerUserName, err := o.prompt.Get(inputDockerUserNamePrompt, "", nil /*no validation*/)
	if err != nil {
		return fmt.Errorf("Prompt for docker username: %w", err)
	}
	o.dockerUserName = dockerUserName

	return nil
}

func (o *jumpStartOpts) AskDockerUserPassword() error {
	dockerUserPassword, err := o.prompt.GetSecret(inputDockerUserPasswordPrompt, "")
	if err != nil {
		return fmt.Errorf("Prompt for docker password: %w", err)
	}
	o.dockerUserPassword = dockerUserPassword

	return nil
}

func (o *jumpStartOpts) AskDockerEmail() error {
	// TODO: email validation
	dockerEmail, err := o.prompt.Get(inputDockerEmailPrompt, "", nil /*no validation*/)
	if err != nil {
		return fmt.Errorf("Prompt for docker email: %w", err)
	}
	o.dockerEmail = dockerEmail

	return nil
}

func (o *jumpStartOpts) AskCertPath() error {
	certPath, err := o.prompt.Get(inputDockerUserNamePrompt, "", nil /*no validation*/)
	if err != nil {
		return fmt.Errorf("Prompt for cert path: %w", err)
	}
	o.certPath = certPath

	return nil
}

func (o *jumpStartOpts) AskKeyPath() error {
	keyPath, err := o.prompt.Get(inputDockerUserNamePrompt, "", nil /*no validation*/)
	if err != nil {
		return fmt.Errorf("Prompt for key path: %w", err)
	}
	o.keyPath = keyPath

	return nil
}

func (o *jumpStartOpts) AskFromFilePath() error {
	if err := o.AskFromFileIteration(); err != nil {
		return err
	}

	for i := 0; i < o.noOfFromFileIteration; i++ {
		fromFile, err := o.prompt.Get(inputFromFilePrompt, "", nil /*no validation*/)
		if err != nil {
			return fmt.Errorf("Prompt for from-file: %w", err)
		}
		o.fromFile = append(o.fromFile, fromFile)
	}

	return nil
}

func (o *jumpStartOpts) AskFromLiteral() error {
	if err := o.AskFromLiteralIteration(); err != nil {
		return err
	}

	for i := 0; i < o.noOfFromLiteralIteration; i++ {
		fromLiteral, err := o.prompt.Get(inputFromLiteralPrompt, "", nil /*no validation*/)
		if err != nil {
			return fmt.Errorf("Prompt for from-literal: %w", err)
		}
		o.fromLiteral = append(o.fromLiteral, fromLiteral)
	}

	return nil
}

func (o *jumpStartOpts) AskFromEnv() error {
	fromEnvFile, err := o.prompt.Get(inputFromEnvFilePrompt, "", nil /*no validation*/)
	if err != nil {
		return fmt.Errorf("Prompt for env: %w", err)
	}
	o.fromEnvFile = fromEnvFile

	return nil
}

func (o *jumpStartOpts) AskOutputInfo() error {
	if err := o.AskOutputFormat(); err != nil {
		return err
	}

	if err := o.AskOutputPath(); err != nil {
		return err
	}

	return nil
}

func (o *jumpStartOpts) AskFromFileIteration() error {
	// TODO: int vaildation
	noOfIterationStr, err := o.prompt.Get(inputNoOfFromFileIterationPrompt, "", nil /*no validation*/)
	if err != nil {
		return fmt.Errorf("Prompt for from-file iteration: %w", err)
	}

	noOfFromFileIteration, err := strconv.Atoi(noOfIterationStr)
	if err != nil {
		return fmt.Errorf("Prompt for from-file iteration: %w", err)
	}
	o.noOfFromFileIteration = noOfFromFileIteration

	return nil
}

func (o *jumpStartOpts) AskFromLiteralIteration() error {
	// TODO: int vaildation
	noOfFromLiteralIterationStr, err := o.prompt.Get(inputNoOfFromLiteralIterationPrompt, "", nil /*no validation*/)
	if err != nil {
		return fmt.Errorf("Prompt for from-literal iteration: %w", err)
	}

	noOfFromLiteralIteration, err := strconv.Atoi(noOfFromLiteralIterationStr)
	o.noOfFromLiteralIteration = noOfFromLiteralIteration

	return nil
}
