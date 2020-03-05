package cli

import (
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"encoding/json"
)

func int32Ptr(i int32) *int32 { return &i }

func (o *askOpts) ExecuteCreateDeploymentSpec() error {
	var objectMeta metav1.ObjectMeta
	var spec appsv1.DeploymentSpec

	if err := o.Ask("RequireObjectMeta"); err != nil {
		return err
	}

	if o.requireObjectMeta {
		// Name
		if err := o.Ask("ObjectMetaName"); err != nil {
			return err
		}

		// Namespace
		if err := o.Ask("Namespace"); err != nil {
			return err
		}

		// Deployment > ObjectMeta
		objectMeta = metav1.ObjectMeta{
			Name: o.Name,
		}

	}

	if err := o.Ask("RequireDeploymentSpec"); err != nil {
		return err
	}

	if o.requireDeploymentSpec {
		// Replicas
		if err := o.Ask("Replicas"); err != nil {
			return err
		}
		// Selector
		// Template
		// Strategy
		// MinReadySeconds
		if err := o.Ask("MinReadySeconds"); err != nil {
			return err
		}
		// RevisionHistoryLimit
		if err := o.Ask("RevisionHistoryLimit"); err != nil {
			return err
		}
		// Paused
		if err := o.Ask("Paused"); err != nil {
			return err
		}
		// ProgressDeadlineSeconds
		if err := o.Ask("ProgressDeadlineSeconds"); err != nil {
			return err
		}

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

		// Deployment > DeploymentSpec > PodTemplateSpec
		podTemplateSpec := apiv1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					"app": "demo",
				},
			},
			Spec: podSpec,
		}

		// Deployment > DeploymentSpec > LabelSelector
		labelSelector := &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"app": "demo",
			},
		}

		// Deployment > DeploymentSpec
		spec = appsv1.DeploymentSpec{
			// FIXME: ambiguous selector
			// Replicas: int32Ptr(o.Replicas),
			Selector: labelSelector,
			Template: podTemplateSpec,
			// Strategy:                nil,
			MinReadySeconds:         o.MinReadySeconds,
			RevisionHistoryLimit:    o.RevisionHistoryLimit,
			Paused:                  o.Paused,
			ProgressDeadlineSeconds: o.ProgressDeadlineSeconds,
		}

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

	if err := o.AskOutputInfo(); err != nil {
		return err
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

	// TODO: yaml conversion

	return nil
}
