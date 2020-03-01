package cli

import (
	"fmt"
)

const (
	inputDeploymentNamePrompt = "What deployment you want to name?"
	inputImageNamePrompt = "What image you want to name to run?"
	inputOutputFormatPrompt = "Please select an output format"
	inputOutputPathPrompt = "What directory you want to save?"
)

func (o *jumpStartOpts) AskDeploymentCmdOpts() error {
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

func (o *jumpStartOpts) AskDeploymentName() error {
	deploymentName, err := o.prompt.Get(inputDeploymentNamePrompt, "", nil /*no validation*/)

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

func (o *jumpStartOpts) ExecuteDeploymentCmd() error {
// 	Options:
//       --allow-missing-template-keys=true: If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.
//       --dry-run=false: If true, only print the object that would be sent, without sending it.
//       --generator='': The name of the API generator to use.
//       --image=[]: Image name to run.
//   -o, --output='': Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-file.
//       --save-config=false: If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.
//       --template='': Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].
//       --validate=true: If true, use a schema to validate the input before sending it

	cmd := fmt.Sprintf("kubectl create deployment %s --image=%s --output=%s > %s", o.deploymentName, o.imageName, o.outputFormat, o.outputPath)

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *jumpStartOpts) ExecuteJumpStartDeploymentCmd() error {
	if err := o.AskDeploymentCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteDeploymentCmd(); err != nil {
		return err
	}

	return nil
}