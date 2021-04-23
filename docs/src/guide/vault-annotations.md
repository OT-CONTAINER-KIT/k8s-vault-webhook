# Vault Annotations

The available annotations for k8s vault webhook are:-

|**Name**|**Description**|**Required**|**Default**|
|--------|---------------|------------|-----------|
|`vault.opstree.secret.manager/enabled`| Enables the vault secret manager | - | false |
|`vault.opstree.secret.manager/service`| Vault cluster address with http prefix | yes | - |
|`vault.opstree.secret.manager/tls-secret`| Vault TLS secret name if vault is configured on TLS | no | - |
|`vault.opstree.secret.manager/role`| Vault role created with Kubernetes serviceaccount | yes | - |
|`vault.opstree.secret.manager/path`| Path of the secret in vault | no | - |
