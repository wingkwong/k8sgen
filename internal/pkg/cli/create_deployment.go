package cli

import (
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"encoding/json"
)

func int32Ptr(i int32) *int32 { return &i }

func (o *askOpts) AskDeploymentSpecOpts() error {

	if err := o.Ask("RequireObjectMeta"); err != nil {
		return err
	}

	if o.requireObjectMeta {
		// Name
		if err := o.Ask("Kind"); err != nil {
			return err
		}

		// Namespace
		if err := o.Ask("Namespace"); err != nil {
			return err
		}
	}

	if err := o.Ask("RequireDeploymentSpec"); err != nil {
		return err
	}

	if o.requireDeploymentSpec {
		// Replicas
		// Selector
		// Template
		// Strategy
		// MinReadySeconds
		// RevisionHistoryLimit
		// Paused
		// ProgressDeadlineSeconds

	}

	if err := o.Ask("RequireDeploymentStatus"); err != nil {
		return err
	}

	if o.requireDeploymentStatus {
		// ObservedGeneration
		// Replicas
		// UpdatedReplicas
		// ReadyReplicas
		// AvailableReplicas
		// UnavailableReplicas
		// Conditions
		// CollisionCount
	}

	if err := o.Ask("DeploymentName"); err != nil {
		return err
	}

	if err := o.Ask("Image"); err != nil {
		return err
	}

	if err := o.AskOutputInfo(); err != nil {
		return err
	}

	return nil
}

func (o *askOpts) ExecuteDeploymentSpec() error {

	// Deployment > DeploymentSpec > PodTemplateSpec > PodSpec
	podSpec := apiv1.PodSpec{
		Containers: []apiv1.Container{
			{
				Name:  "web",
				Image: "nginx:1.12",
				Ports: []apiv1.ContainerPort{
					{
						Name:          "http",
						Protocol:      apiv1.ProtocolTCP,
						ContainerPort: 80,
					},
				},
			},
		},
	}

	// Deployment > DeploymentSpec > LabelSelector
	labelSelector := &metav1.LabelSelector{
		MatchLabels: map[string]string{
			"app": "demo",
		},
	}

	// Deployment > DeploymentSpec > PodTemplateSpec
	podTemplateSpec := apiv1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app": "demo",
			},
		},
		Spec: podSpec,
	}

	// Deployment > ObjectMeta
	objectMeta := metav1.ObjectMeta{
		Name: "demo-deployment",
	}

	// Deployment > DeploymentSpec
	spec := appsv1.DeploymentSpec{
		Replicas: int32Ptr(2),
		Selector: labelSelector,
		Template: podTemplateSpec,
	}

	// Deployment
	deployment := &appsv1.Deployment{
		ObjectMeta: objectMeta,
		Spec:       spec,
	}

	b, err := json.Marshal(deployment)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(b))

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
