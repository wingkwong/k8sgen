# k8sgen

k8sgen is an utility which is designed to guide users to build their Kubernetes resources in an interactive CLI. 

[![Build Status](https://travis-ci.com/wingkwong/k8sgen.svg?branch=master)](https://travis-ci.com/wingkwong/k8sgen) ![Stability:Experimental](https://img.shields.io/badge/stability-experimental-orange) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) [![Go Report Card](https://goreportcard.com/badge/github.com/wingkwong/k8sgen)](https://goreportcard.com/report/github.com/wingkwong/k8sgen)

## Usage

``k8sgen jumpstart`` - jumpstart a resource file

This command utilises ``kubectl`` to create a jumpstart version of a resource file.

```bash
k8sgen jumpstart

 _    ___                       
| | _( _ ) ___  __ _  ___ _ __  
| |/ / _ \/ __|/ _  |/ _ | |_ \ 
|   | (_) \__ | (_| |  __| | | |
|_|\_\___/|___/\__, |\___|_| |_|
                |___/

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
  json
> yaml

? What directory you want to save? /home/wingkwong/deployment.yaml
```

Result: 
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: my-deployment
  name: my-deployment
spec:
  replicas: 1
  selector:
    matchLabels:
      app: my-deployment
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: my-deployment
    spec:
      containers:
      - image: busybox
        name: busybox
        resources: {}
status: {}
```

## Contributing

The k8sgen project adheres to the [CNCF Code of
Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md).

We welcome community contributions and pull requests.

## License 

MIT