<div align="center">
    <img src="./static/k8s-vault-webhook-logo.svg" height="160" width="120">
</div>

<p align="center">
  <a href="https://dev.azure.com/opstreedevops/DevOps/_build?definitionId=4">
    <img src="https://dev.azure.com/opstreedevops/DevOps/_apis/build/status/k8s-vault-webhook/k8s-vault-webhook?branchName=master" alt="Azure Pipelines">
  </a>
  <a href="https://goreportcard.com/report/github.com/OT-CONTAINER-KIT/k8s-vault-webhook">
    <img src="https://goreportcard.com/badge/github.com/OT-CONTAINER-KIT/k8s-vault-webhook" alt="GoReportCard">
  </a>
  <a href="http://golang.org">
    <img src="https://img.shields.io/github/go-mod/go-version/OT-CONTAINER-KIT/k8s-vault-webhook" alt="GitHub go.mod Go version (subdirectory of monorepo)">
  </a>
  <a href="http://golang.org">
    <img src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg" alt="made-with-Go">
  </a>
  <a href="https://quay.io/repository/opstree/k8s-vault-webhook">
    <img src="https://img.shields.io/badge/container-ready-green" alt="Docker">
  </a>
  <a href="https://github.com/OT-CONTAINER-KIT/k8s-vault-webhook/master/LICENSE">
    <img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg" alt="License">
  </a>
</p>

k8s-vault-webhook is a Kubernetes admission webhook which listen for the events related to Kubernetes resources for injecting secret directly from secret manager to pod, secret, and configmap.
The motive of creating this project is to provide a dynamic secret injection to containers/pods running inside Kubernetes from different secret managers for enhanced security.

Documentation is available here:- https://ot-container-kit.github.io/k8s-vault-webhook/

The secret managers which are currently supported:-

- **[Hashicorp Vault](https://www.vaultproject.io/)**
- **[AWS Secret Manager](https://aws.amazon.com/secrets-manager/)**
- **[Azure Key Vault](https://azure.microsoft.com/en-in/services/key-vault/)**
- **[GCP Secret Manager](https://cloud.google.com/secret-manager)**

### Supported Features

- Authentication to Hashicorp vault using Kubernetes service-account
- RBAC implementation of vault using different policies of vault and association of policy with service-account
- Inject secret directly to pods/containers running inside Kubernetes
- Inject secret directly to pods/containers from AWS Secret Manager
- Authentication with AWS Secret Manager with access key and iam role
- Fetch secrets from Azure Key Vault and inject them in pods/containers
- Pod AD identity and Service principal based authentication in Azure
- Authentication with AWS Secret Manager with access key and iam role
- Authenticate and authorize using GCP service-account and annotations
- Secret injection in pods/containers from GCP Secret Manager
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

### Quickstart

For setting up a quickstart environment for demo, you can start quickstart from [here](https://ot-container-kit.github.io/k8s-vault-webhook/)

### Development

If you like to contribute to this project, you are more than welcome. Please see our [DEVELOPMENT.md](./DEVELOPMENT.md) for details.

### Release History

Please see our [CHANGELOG.md](./CHANGELOG.md) for details.

### Contact

If you have any suggestion or query. Contact us at

[opensource@opstree.com](mailto:opensource@opstree.com)

