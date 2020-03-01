# k8sgen

k8sgen is an utility which is designed to guide users to build their Kubernetes resources in an interactive CLI. 

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Usage

``k8sgen jumpstart`` - jumpstart a resource file

This command utilises ``kubectl`` to create a jumpstart version of a resource file.

```bash
k8sgen jumpstart

? What kind of object you want to create? [Use arrows to move, type to filter]
  ClusterRole
  ClusterRoleBinding
  Configmap
> Deployment
  Job
  Namespace
  PodDisruptionBudget
  PriorityClass
  Quota
  Role
  RoleBinding
  Secret
  Service
  ServiceAccount

? What deployment you want to name? my-deployment

? What image you want to name to run? busybox

? Please select an output format yaml
> json
  yaml

? What directory you want to save? /home/wingkwong/deployment.json
```

/home/wingkwong/deployment.json
```json
{
    "apiVersion": "apps/v1",
    "kind": "Deployment",
    "metadata": {
        "creationTimestamp": "2020-03-01T08:18:23Z",
        "generation": 1,
        "labels": {
            "app": "my-deployment"
        },
        "name": "my-deployment",
        "namespace": "default",
        "resourceVersion": "59974",
        "selfLink": "/apis/apps/v1/namespaces/default/deployments/my-deployment",
        "uid": "aa654409-14a3-4c85-8ea6-1a28ad94e49c"
    },
    "spec": {
        "progressDeadlineSeconds": 600,
        "replicas": 1,
        "revisionHistoryLimit": 10,
        "selector": {
            "matchLabels": {
                "app": "my-deployment"
            }
        },
        "strategy": {
            "rollingUpdate": {
                "maxSurge": "25%",
                "maxUnavailable": "25%"
            },
            "type": "RollingUpdate"
        },
        "template": {
            "metadata": {
                "creationTimestamp": null,
                "labels": {
                    "app": "my-deployment"
                }
            },
            "spec": {
                "containers": [
                    {
                        "image": "busybox",
                        "imagePullPolicy": "Always",
                        "name": "busybox",
                        "resources": {},
                        "terminationMessagePath": "/dev/termination-log",
                        "terminationMessagePolicy": "File"
                    }
                ],
                "dnsPolicy": "ClusterFirst",
                "restartPolicy": "Always",
                "schedulerName": "default-scheduler",
                "securityContext": {},
                "terminationGracePeriodSeconds": 30
            }
        }
    },
    "status": {}
}
```

``k8sgen create`` - create a resource file from scratch 

TBC