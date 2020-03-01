# k8sgen

k8sgen is an utility which is designed to guide users to build their Kubernetes resources in an interactive CLI. 

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Usage

``k8sgen create`` - create a resource file from scratch 

TBC

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

? What deployment you want to name? my_deployment

? What image you want to name to run? busybox

? Please select an output format yaml
> json
  yaml

? What directory you want to save? /home/wingkwong/deployment.json
```