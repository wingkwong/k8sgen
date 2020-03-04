package cli

func (o *askOpts) AskDeploymentSpecOpts() error {
	if err := o.AskDeploymentName(); err != nil {
		return err
	}

	if err := o.AskImageName(); err != nil {
		return err
	}

	if err := o.AskOutputInfo(); err != nil {
		return err
	}

	return nil
}

func (o *askOpts) ExecuteDeploymentSpec() error {
	// TODO:
	return nil
}

func (o *askOpts) ExecuteCreateDeploymentSpec() error {
	if err := o.AskDeploymentSpecOpts(); err != nil {
		return err
	}

	if err := o.ExecuteDeploymentSpec(); err != nil {
		return err
	}

	return nil
}
