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
๐  minikube v1.0.1 on linux (amd64)
๐คน  Downloading Kubernetes v1.14.1 images in the background ...
๐ฅ  Creating kvm2 VM (CPUs=2, Memory=2048MB, Disk=20000MB) ...
๐ถ  "minikube" IP address is 192.168.39.240
๐ณ  Configuring Docker as the container runtime ...
๐ณ  Version of container runtime is 18.06.3-ce
โ  Waiting for image downloads to complete ...
โจ  Preparing Kubernetes environment ...
๐  Pulling images required by Kubernetes v1.14.1 ...
๐  Launching Kubernetes v1.14.1 using kube๐  minikube v1.0.1 on linux (amd64)
๐คน  Downloading Kubernetes v1.14.1 images in the background ...
๐ฅ  Creating kvm2 VM (CPUs=2, Memory=2048MB, Disk=20000MB) ...
๐ถ  "minikube" IP address is 192.168.39.240
๐ณ  Configuring Docker as the container runtime ...
๐ณ  Version of container runtime is 18.06.3-ce
โ  Waiting for image downloads to complete ...
โจ  Preparing Kubernetes environment ...
๐  Pulling images required by Kubernetes v1.14.1 ...
๐  Launching Kubernetes v1.14.1 using kubeadm ... 
โ  Waiting for pods: apiserver proxy etcd scheduler controller dns
๐  Configuring cluster permissions ...
๐ค  Verifying component health .....
๐  kubectl is now configured to use "minikube"
๐  Done! Thank you for using minikube!adm ... 
โ  Waiting for pods: apiserver proxy etcd scheduler controller dns
๐  Configuring cluster permissions ...
๐ค  Verifying component health .....
๐  kubectl is now configured to use "minikube"
๐  Done! Thank you for using minikube!
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

