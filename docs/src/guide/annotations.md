# Annotations

This section of document contains the information about Kubernetes annotation which will be used for Webhook mutation.
Based on these annotations, the secrets will be mutated.

The annotations which are currently supported:-

- **[Hashicorp Vault](https://www.vaultproject.io/)**

There are some other annotations which are planned to be implemented in future.

- **[AWS Secret Manager](https://aws.amazon.com/secrets-manager/)**
- **[Azure Key Vault](https://azure.microsoft.com/en-in/services/key-vault/)**
- **[GCP Secret Manager](https://cloud.google.com/secret-manager)**

## Vault Annotations

The available annotations for k8s vault webhook are:-

|**Name**|**Description**|**Required**|**Default**|
|--------|---------------|------------|-----------|
|`vault.opstree.secret.manager/enabled`| Enables the vault secret manager | - | false |
|`vault.opstree.secret.manager/service`| Vault cluster address with http prefix | yes | - |
|`vault.opstree.secret.manager/tls-secret`| Vault TLS secret name if vault is configured on TLS | no | - |
|`vault.opstree.secret.manager/role`| Vault role created with Kubernetes serviceaccount | yes | - |
|`vault.opstree.secret.manager/path`| Path of the secret in vault | no | - |
|`vault.opstree.secret.manager/k8s-token-path`| Alternate kubernetes service account token path | no | `/var/run/secrets/kubernetes.io/serviceaccount/token` |
|`vault.opstree.secret.manager/path` | Vault secret path | Yes | - | 
|`vault.opstree.secret.manager/secret-version` | Vault secret version (if using v2 secret engine) | Yes | - | 
|`vault.opstree.secret.manager/use-secret-names-as-keys` | treat secret path ending with / as directory where secret name is the key and a single value in each | No | - |
|`vault.opstree.secret.manager/auth-path`| alternate kubernetes backend auth path | No | `auth/kubernetes/login` |
