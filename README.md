<div align="left">
    <img src="./static/k8s-vault-webhook-logo.svg" height="120" width="120">
</div>

## K8s Vault Webhook

k8s-vault-webhook is a Kubernetes admission webhook which listen for the events related to Kubernetes resources for injecting secret directly from secret manager to pod, secret, and configmap.
The motive of creating this project is to provide a dynamic secret injection to containers/pods running inside Kubernetes from different secret managers for enhanced security.

The secret managers which are currently supported:-

- **[Hashicorp Vault](https://www.vaultproject.io/)**

There are some secret managers which are planned to be implemented in future.

- **[AWS Secret Manager](https://aws.amazon.com/secrets-manager/)**
- **[Azure Key Vault](https://azure.microsoft.com/en-in/services/key-vault/)**
- **[GCP Secret Manager](https://cloud.google.com/secret-manager)**

### Supported Features

- Authentication to Hashicorp vault using Kubernetes service-account
- RBAC implementation of vault using different policies of vault and association of policy with service-account
- Inject secret directly to pods/containers running inside Kubernetes
- Support regex to inject all secrets from a certain path of Vault
- Inject secrets directly to the process of container, i.e. after the injection you cannot read secrets from the environment variable

### Architecture

<div align="center">
    <img src="./static/k8s-vault-webhook-arc.png">
</div>

### Installation

k8s-vault-webhook can easily get installed by using [Helm](https://helm.sh/). We just simple need to add the repository of our [helm charts](https://github.com/OT-CONTAINER-KIT/helm-charts).

```shell
$ helm repo add ot-helm https://github.com/OT-CONTAINER-KIT/helm-charts

$ helm upgrade k8s-vault-webhook ot-helm/k8s-vault-webhook --namespace <namespace> --install
```

If you want to pass your custom values file while installing the chart, you can find the values file [here](https://github.com/OT-CONTAINER-KIT/helm-charts/blob/main/charts/k8s-vault-webhook/values.yaml)

### Annotations

|**Name**|**Description**|**Required**|**Default**|
|--------|---------------|------------|-----------|
|`vault.opstree.secret.manager/enabled`| Enables the vault secret manager | - | false |
|`vault.opstree.secret.manager/service`| Vault cluster address with http prefix | yes | - |
|`vault.opstree.secret.manager/tls-secret`| Vault TLS secret name if vault is configured on TLS | no | - |
|`vault.opstree.secret.manager/role`| Vault role created with Kubernetes serviceaccount | yes | - |
|`vault.opstree.secret.manager/path`| Path of the secret in vault | no | - |

### Quickstart

For setting up a quickstart environment for demo, you can follow [Quickstart.md](Quickstart.md).

