## Development

### Prerequisites

##### Access to a Kubernetes Cluster

First of all, you will need access to a Kubernetes cluster. The easiest way to start is minikube.

- [Virtualbox](https://www.virtualbox.org/wiki/Downloads) - hypervisor to run a kubernetes cluster
- [Minikube](https://kubernetes.io/docs/setup/minikube/) - for kubernetes cluster creation on local machine
- [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) - to interact with kubernetes cluster

##### Tools to build and test k8s-vault-webhook

Apart from kubernetes cluster, there are some tools which are needed to build and test k8s-vault-webhook.

- [Git](https://git-scm.com/downloads)
- [Go](https://golang.org/dl/)
- [Docker](https://docs.docker.com/install/)
- [Make](https://www.gnu.org/software/make/manual/make.html)

## Build Locally

To achieve this, execute this command:-

```shell
make build-code
```

## Build Image

k8s-vault-webhook gets packaged as a container image for running on Kubernetes cluster. These instructions will guide you to build image.

```shell
make build-image
```

## Testing
If you want to play it on Kubernetes. You can use minikube.

```shell
#Start minikube
minikube start --vm-driver virtualbox
```

```shell
#Deploy the image on minikube
$ helm upgrade k8s-vault-webhook ot-helm/k8s-vault-webhook --namespace vault --install
```

## Run Tests
```shell
make test
```
