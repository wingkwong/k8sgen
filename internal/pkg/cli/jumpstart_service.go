package cli

import (
	"fmt"
)

func (o *askOpts) AskServiceCmdOpts() error {
	if err := o.Ask("ServiceCmdName"); err != nil {
		return err
	}

	switch o.ServiceCmdName {
	case clusterIPCmdName:
		if err := o.Ask("ClusterIPName"); err != nil {
			return err
		}

		if err := o.Ask("TCP"); err != nil {
			return err
		}

		if err := o.Ask("ClusterIP"); err != nil {
			return err
		}
	case externalNameCmdName:
		if err := o.Ask("ExternalName"); err != nil {
			return err
		}

		if err := o.Ask("ExternalServiceName"); err != nil {
			return err
		}

	case loadbalancerCmdName:
		if err := o.Ask("LoadbalancerName"); err != nil {
			return err
		}

		if err := o.Ask("TCP"); err != nil {
			return err
		}
	case nodePortCmdName:
		if err := o.Ask("NodePortName"); err != nil {
			return err
		}

		if err := o.Ask("TCP"); err != nil {
			return err
		}
	default:
		return fmt.Errorf("No available Sevice option: %s", o.ServiceCmdName)
	}

	if err := o.AskOutputInfo(); err != nil {
		return err
	}

	return nil
}

func (o *askOpts) ExecuteServiceCmd() error {
	var cmd string

	switch o.ServiceCmdName {
	case clusterIPCmdName:
		// # Create a new ClusterIP service named my-cs
		// kubectl create service clusterip my-cs --tcp=5678:8080

		// # Create a new ClusterIP service named my-cs (in headless mode)
		// kubectl create service clusterip my-cs --clusterip="None"

		cmd = fmt.Sprintf("kubectl create service %s ", o.ClusterIPName)

		if o.TCP != "" {
			cmd = cmd + fmt.Sprintf("--tcp=%s ", o.TCP)
		}

		if o.ClusterIP != "" {
			cmd = cmd + fmt.Sprintf("--clusterip=%s ", o.ClusterIP)
		}

	case externalNameCmdName:
		// # Create a new ExternalName service named my-ns
		// kubectl create service externalname my-ns --external-name bar.com
		cmd = fmt.Sprintf("kubectl create service %s ", o.ExternalName)

		if o.ExternalServiceName != "" {
			cmd = cmd + fmt.Sprintf("--external-name=%s ", o.ExternalServiceName)
		}
	case loadbalancerCmdName:
		// 	# Create a new LoadBalancer service named my-lbs
		// kubectl create service loadbalancer my-lbs --tcp=5678:8080
		cmd = fmt.Sprintf("kubectl create service %s ", o.LoadbalancerName)

		if o.TCP != "" {
			cmd = cmd + fmt.Sprintf("--tcp=%s ", o.TCP)
		}

	case nodePortCmdName:
		// # Create a new NodePort service named my-ns
		// kubectl create service nodeport my-ns --tcp=5678:8080
		cmd = fmt.Sprintf("kubectl create service %s ", o.NodePortName)

		if o.TCP != "" {
			cmd = cmd + fmt.Sprintf("--tcp=%s ", o.TCP)
		}
	default:
		return fmt.Errorf("No available Sevice option: %s", o.ServiceCmdName)
	}

	cmd = cmd + fmt.Sprintf("--output=%s --dry-run=true > %s ", o.OutputFormat, o.OutputPath)

	if err := ExecCmd(cmd); err != nil {
		return fmt.Errorf("Failed To execute command `%s` \n %w", cmd, err)
	}

	return nil
}

func (o *askOpts) ExecuteJumpStartServiceCmd() error {
	if err := o.AskServiceCmdOpts(); err != nil {
		return err
	}

	if err := o.ExecuteServiceCmd(); err != nil {
		return err
	}

	return nil
}
