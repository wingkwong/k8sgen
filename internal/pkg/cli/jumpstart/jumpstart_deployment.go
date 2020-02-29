package jumpstart

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
	o.outputPath = outputPath

	return nil
}

func (o *jumpStartOpts) ExecuteDeploymentCmdOpts() error {
	//TODO:

	return nil
}

func ExecuteClusterRoleCmd(o *jumpStartOpts) error {
	if err := o.AskDeploymentCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteDeploymentCmdOpts(); err != nil {
		return err
	}

	return nil
}