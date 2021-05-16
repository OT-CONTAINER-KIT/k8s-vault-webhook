# Annotations

This section of document contains the information about Kubernetes annotation which will be used for Webhook mutation.
Based on these annotations, the secrets will be mutated.

The annotations which are currently supported:-

- **[Hashicorp Vault](https://www.vaultproject.io/)**
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

## AWS Annotations

|**Name**|**Description**|**Required**|**Default**|
|--------|---------------|------------|-----------|
|`aws.secret.manager/enabled`| Enable the AWS secret manager | - | false |
|`aws.secret.manager/region`| AWS secret manager region | no | us-east-1 |
|`aws.secret.manager/role-arn`| AWS IAM Role to access the secret | no | |
|`aws.secret.manager/secret-name`| Name of the AWS secret | no | |
|`aws.secret.manager/previous-version`| If the secret is rotated, set to "true" | no | |

## Azure Annotations

|**Name**|**Description**|**Required**|**Default**|
|--------|---------------|------------|-----------|
|`azure.secret.manager/enabled`| Enable the Azure Key Vault | - | false |
|`azure.secret.manager/vault-name`| Name of the Azure Key Vault in which secrets are held | no | test-secret |

## GCP Annotations

|**Name**|**Description**|**Required**|**Default**|
|--------|---------------|------------|-----------|
|`gcp.opstree.secret.manager/enabled`| enable the GCP secret manager | - | false |
|`gcp.opstree.secret.manager/project-id` | GCP Project ID | Yes | - |
|`gcp.opstree.secret.manager/gcp-service-account-key-secret-name` | GCP IAM service account secret name (file name **must be** `service-account.json`) | No | Google Default Application Credentials |
|`gcp.opstree.secret.manager/secret-name` | secret name | Yes | - |
|`gcp.opstree.secret.manager/secret-version` | specify the secret version as string | No | Latest |
