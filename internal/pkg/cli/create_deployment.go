package cli

import (
	"fmt"
)

func (o *createOpts) AskDeploymentCmdOpts() error {
	if err := o.AskDeploymentName(); err != nil {
		return err
	}

	if err := o.AskImageName(); err != nil {
		return err
	}

	if err := o.AskOutputFormat(); err != nil {
		return err
	}

	if err := o.AskOutputPath(); err != nil {
		return err
	}

	return nil
}

func (o *createOpts) AskDeploymentName() error {
	// TODO: Add regex to check deployment name: [a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*

	deploymentName, err := o.prompt.Get(inputDeploymentNamePrompt, "", nil /*no validation*/)

	if err != nil {
		return fmt.Errorf("Prompt for deployment name: %w", err)
	}
	
	o.deploymentName = deploymentName

	return nil
}

func (o *createOpts) AskImageName() error {
	imageName, err := o.prompt.Get(inputImageNamePrompt, "", nil /*no validation*/)
	if err != nil {
		return fmt.Errorf("Prompt for image name: %w", err)
	}
	o.imageName = imageName

	return nil
}

func (o *createOpts) AskOutputFormat() error {
	formats := getOutputFormats()
	outputFormat, err := o.prompt.SelectOne(inputOutputFormatPrompt, "", formats)
	if err != nil {
		return fmt.Errorf("Prompt for output format: %w", err)
	}
	o.outputFormat = outputFormat

	return nil
}

func (o *createOpts) AskOutputPath() error {
	outputPath, err := o.prompt.Get(inputOutputPathPrompt, "", nil)
	if err != nil {
		return fmt.Errorf("Prompt for output path: %w", err)
	}

	o.outputPath = outputPath
	return nil
}

func (o *createOpts) ExecuteDeploymentCmd() error {
	// TODO:
	return nil
}

func (o *createOpts)  ExecuteCreateDeploymentCmd() error{
	if err := o.AskDeploymentCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteDeploymentCmd(); err != nil {
		return err
	}

	return nil
}