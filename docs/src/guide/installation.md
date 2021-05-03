# Installation

K8s Vault Webhook is a golang based binary that is packaged as a Docker image and needs to be deployed inside the Kubernetes cluster. It cannot be deployed outside the Kubernetes cluster because of Kubernetes architecture dependency. It uses the concept of Admission Webhook and if you want to explore the concept of admission webhook please refer to [official documentation](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/).

k8s-vault-webhook can easily get installed by using [Helm](https://helm.sh/) . We just simple need to add the repository of our [helm charts](https://github.com/OT-CONTAINER-KIT/helm-charts). It requires `=>3.0.0` version of helm.

First we need to add the repository for helm.

```shell
$ helm repo add ot-helm https://ot-container-kit.github.io/helm-charts
...
"ot-helm" has been added to your repositories
```

Once the repo is added, we can install the k8s-vault-webhook.

```shell
$ helm upgrade k8s-vault-webhook ot-helm/k8s-vault-webhook \
    --namespace vault --install
...
Release "k8s-vault-webhook" has been upgraded. Happy Helming!
NAME: k8s-vault-webhook
LAST DEPLOYED: Mon May  3 19:55:03 2021
NAMESPACE: vault
STATUS: deployed
REVISION: 2
TEST SUITE: None
```

Verify the deployment is successful by listing pods of k8s-vault-webhook.

```shell
$ kubectl get pods -n vault
...
NAME                                 READY   STATUS    RESTARTS   AGE
k8s-vault-webhook-6cf98dbd6f-5gjvv   1/1     Running   0          61s
k8s-vault-webhook-6cf98dbd6f-tjdv5   1/1     Running   0          48s
```
