# Development Guide

## Pre-requisites

**Access to Kubernetes cluster**

First of all, you will need access to a Kubernetes cluster. The easiest way to start is minikube.

- [Virtualbox](https://www.virtualbox.org/wiki/Downloads) - hypervisor to run a Kubernetes cluster
- [Minikube](https://kubernetes.io/docs/setup/minikube/) - for Kubernetes cluster creation on local machine
- [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) - to interact with Kubernetes cluster

**Tools to build K8s Vault Webhook**

Apart from kubernetes cluster, there are some tools which are needed to build and test k8s-vault-webhook.

- [Git](https://git-scm.com/downloads)
- [Go](https://golang.org/dl/)
- [Docker](https://docs.docker.com/install/)
- [Make](https://www.gnu.org/software/make/manual/make.html)

## Build

**Build Localy**

To achieve this, execute this command:-

```shell
make build-code
```

**Build Image**

k8s-vault-webhook gets packaged as a container image for running on Kubernetes cluster. These instructions will guide you to build image.

```shell
make build-image
```

## Testing
If you want to play it on Kubernetes. You can use minikube.

```shell
#Start minikube
$ minikube start --vm-driver virtualbox
...
😄  minikube v1.0.1 on linux (amd64)
🤹  Downloading Kubernetes v1.14.1 images in the background ...
🔥  Creating kvm2 VM (CPUs=2, Memory=2048MB, Disk=20000MB) ...
📶  "minikube" IP address is 192.168.39.240
🐳  Configuring Docker as the container runtime ...
🐳  Version of container runtime is 18.06.3-ce
⌛  Waiting for image downloads to complete ...
✨  Preparing Kubernetes environment ...
🚜  Pulling images required by Kubernetes v1.14.1 ...
🚀  Launching Kubernetes v1.14.1 using kube😄  minikube v1.0.1 on linux (amd64)
🤹  Downloading Kubernetes v1.14.1 images in the background ...
🔥  Creating kvm2 VM (CPUs=2, Memory=2048MB, Disk=20000MB) ...
📶  "minikube" IP address is 192.168.39.240
🐳  Configuring Docker as the container runtime ...
🐳  Version of container runtime is 18.06.3-ce
⌛  Waiting for image downloads to complete ...
✨  Preparing Kubernetes environment ...
🚜  Pulling images required by Kubernetes v1.14.1 ...
🚀  Launching Kubernetes v1.14.1 using kubeadm ... 
⌛  Waiting for pods: apiserver proxy etcd scheduler controller dns
🔑  Configuring cluster permissions ...
🤔  Verifying component health .....
💗  kubectl is now configured to use "minikube"
🏄  Done! Thank you for using minikube!adm ... 
⌛  Waiting for pods: apiserver proxy etcd scheduler controller dns
🔑  Configuring cluster permissions ...
🤔  Verifying component health .....
💗  kubectl is now configured to use "minikube"
🏄  Done! Thank you for using minikube!
```

```shell
#Deploy the image on minikube
$ helm upgrade k8s-vault-webhook \
    ot-helm/k8s-vault-webhook --namespace vault --install
...
Release "k8s-vault-webhook" has been upgraded. Happy Helming!
NAME: k8s-vault-webhook
LAST DEPLOYED: Mon May  3 19:55:03 2021
NAMESPACE: vault
STATUS: deployed
REVISION: 2
TEST SUITE: None
```

**Running Test Cases**

```shell
make test
```

