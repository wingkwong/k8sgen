package cli

import (
	"fmt"
)

func (o *askOpts) AskJobCmdOpts() error {
	if err := o.Ask("JobName"); err != nil {
		return err
	}

	if err := o.Ask("Image"); err != nil {
		return err
	}

	if err := o.Ask("Command"); err != nil {
		return err
	}

	if err := o.Ask("FromResource"); err != nil {
		return err
	}

	if err := o.AskOutputInfo(); err != nil {
		return err
	}

	return nil
}

func (o *askOpts) ExecuteJobCmd() error {
	var cmd string

	// Example:

	// # Create a job
	// kubectl create job my-job --image=busybox

	// # Create a job with command
	// kubectl create job my-job --image=busybox -- date

	// # Create a job from a CronJob named "a-cronjob"
	// kubectl create job test-job --from=cronjob/a-cronjob

	cmd = fmt.Sprintf("kubectl create job %s --image=%s ", o.JobName, o.ImageName)

	if o.Command != "" {
		cmd = cmd + fmt.Sprintf("-- %s ", o.Command)
	}

	if o.FromResource != "" {
		cmd = cmd + fmt.Sprintf("--from=%s ", o.FromResource)
	}

	cmd = cmd + fmt.Sprintf("--output=%s --dry-run=true > %s ", o.OutputFormat, o.OutputPath)

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *askOpts) ExecuteJumpStartJobCmd() error {
	if err := o.AskJobCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteJobCmd(); err != nil {
		return err
	}

	return nil
}
