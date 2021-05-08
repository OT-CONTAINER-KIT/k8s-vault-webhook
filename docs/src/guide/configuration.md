# Configuration

There is not alot of configuration changes requires to deploy K8s Vault Webhook. But the configuration can be customized using **Helm**, in that case [values.yaml](https://github.com/OT-CONTAINER-KIT/helm-charts/blob/main/charts/k8s-vault-webhook/values.yaml) can be updated.

## Helm Parameters

| Parameter| Description | Default |
|----------|-------------|---------|
|affinity| affinities to use | `{}` |
|debug| debug logs for webhook | `false` |
|image.pullPolicy| image pull policy | `IfNotPresent`|
|image.repository| image repo that contains the admission server | `quay.io/opstree/k8s-vault-webhook` |
|image.tag| image tag for admission server | `2.0` |
|image.imagePullSecrets| image pull secrets for private repositories | `[]` |
|namespaceSelector| namespace selector to use, will limit webhook scope | `{}` |
|nodeSelector|node selector to use | `{}` |
|podAnnotations|extra annotations to add to pod metadata | `{}` |
|replicaCount|number of replicas | `2` |
|resources|resources to request | `{}` |
|service.externalPort | webhook service external port | `443` |
|service.name |webhook service name | `k8s-vault-webhook` |
|service.type |webhook service type | `ClusterIP` |
|tolerations |tolerations to add  | `[]` |
|rbac.enabled |use rbac | `true` |
|rbac.psp.enabled |use pod security policy | `true` |
|env.VAULT_IMAGE | vault image  | `vault:latest` |
|env.K8S_SECRET_INJECTOR_IMAGE | vault-env image  | `quay.io/opstree/k8s-secret-injector:2.0` |
|volumes |extra volume definitions  | `[]` |
|volumeMounts |extra volume mounts  | `[]` |
| configMapMutation                | enable injecting values from Vault to ConfigMaps                             | `false`                             |
| podDisruptionBudget.enabled      | enable PodDisruptionBudget                                                   | `false`                             |
| podDisruptionBudget.minAvailable | represents the number of Pods that must be available (integer or percentage) | `1`                                 |
| certificate.generate             | should a new CA and TLS certificate be generated for the webhook             | `true`                              |
| certificate.useCertManager       | should request cert-manager for getting a new CA and TLS certificate         | `false`                             |
| certificate.ca.crt               | Base64 encoded CA certificate                                                | ``                                  |
| certificate.server.tls.crt       | Base64 encoded TLS certificate signed by the CA                              | ``                                  |
| certificate.server.tls.key       | Base64 encoded  private key of TLS certificate signed by the CA              | ``                                  |
| apiSideEffectValue               | Webhook sideEffect value                                                     | `NoneOnDryRun`                      |
