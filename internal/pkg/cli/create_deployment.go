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

	if err := o.AskRequireObjectMeta(); err != nil {
		return err
	}

	if o.requireObjectMeta {
		// Name
		if err := o.AskDeploymentName(); err != nil {
			return err
		}

		// Namespace
		if err := o.AskNamespace(); err != nil {
			return err
		}
	}

	if err := o.AskRequireDeploymentSpec(); err != nil {
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

	if err := o.AskRequireDeploymentStatus(); err != nil {
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
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
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
				},
			},
		},
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
