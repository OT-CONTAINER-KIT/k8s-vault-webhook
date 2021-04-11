## Vault Setup

Create a K8s namespace.

```shell
$ kubectl create namespace vault
```

To access the Vault Helm chart, add the Hashicorp Helm repository.

```shell
$ helm repo add hashicorp https://helm.releases.hashicorp.com
"hashicorp" has been added to your repositories
```

```shell
$ helm install vault hashicorp/vault --namespace vault --version 0.9.1
```

Verify the pods

```shell
$ kubectl get pods -n vault
```

