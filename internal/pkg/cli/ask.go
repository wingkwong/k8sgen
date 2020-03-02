package cli

import (
	"fmt"
)

const (
	inputDeploymentNamePrompt = "What deployment you want to name?"
	inputImageNamePrompt      = "What image you want to name to run?"
	inputOutputFormatPrompt   = "Please select an output format"
	inputOutputPathPrompt     = "What directory you want to save?"
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
