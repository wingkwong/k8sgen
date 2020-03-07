package cli

import (
	"fmt"
)

func (o *askOpts) AskPriorityClassCmdOpts() error {
	if err := o.Ask("PriorityClassName"); err != nil {
		return err
	}

	if err := o.Ask("Value"); err != nil {
		return err
	}

	if err := o.Ask("Description"); err != nil {
		return err
	}

	if err := o.Ask("GlobalDefault"); err != nil {
		return err
	}
	if err := o.Ask("PreemptionPolicy"); err != nil {
		return err
	}

	if err := o.AskOutputInfo(); err != nil {
		return err
	}

	return nil
}

func (o *askOpts) ExecutePriorityClassCmd() error {
	var cmd string

	// Example

	// 	# Create a priorityclass named high-priority
	//   kubectl create priorityclass high-priority --value=1000 --description="high priority"

	//   # Create a priorityclass named default-priority that considered as the global default priority
	//   kubectl create priorityclass default-priority --value=1000 --global-default=true --description="default priority"

	//   # Create a priorityclass named high-priority that can not preempt pods with lower priority
	//   kubectl create priorityclass high-priority --value=1000 --description="high priority" --preemption-policy="Never"

	cmd = fmt.Sprintf("kubectl create priorityclass %s ", o.PriorityClassName)

	if o.Value != 0 {
		cmd = cmd + fmt.Sprintf("--value=%d ", o.Value)
	}

	if o.Description != "" {
		cmd = cmd + fmt.Sprintf("--description=%s ", o.Description)
	}

	if o.GlobalDefault {
		cmd = cmd + fmt.Sprintf("--global-default=%t ", o.GlobalDefault)
	}

	if o.PreemptionPolicy != "" {
		cmd = cmd + fmt.Sprintf("--preemption-policy=%s ", o.PreemptionPolicy)
	}

	cmd = cmd + fmt.Sprintf("--output=%s > %s", o.OutputFormat, o.OutputPath)

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *askOpts) ExecuteJumpStartPriorityClassCmd() error {
	if err := o.AskPriorityClassCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecutePriorityClassCmd(); err != nil {
		return err
	}

	return nil
}
