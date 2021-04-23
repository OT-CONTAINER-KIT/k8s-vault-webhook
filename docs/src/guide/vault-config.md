# Vault Configuration

For Vault configuration, we are going to use vault-cli and vault-ui both. Download and install vault-cli on your local system. It can be downloaded from [here](https://www.vaultproject.io/downloads).

First, forward the port of vault to your local system.

```shell
$ kubectl port-forward vault-0 8200:8200 -n vault
```

Login into the vault using the root token which we got using init command.

```shell
$ vault login
  Token (will be hidden):
```

Enable the Kubernetes auth backend in the vault cluster

```shell
$ vault auth enable kubernetes
Success! Enabled kubernetes auth method at: kubernetes/
```

Create a file with name `vault-reviewer.yaml` which will have the service-account and cluster role access information.

```yaml
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: vault-reviewer
  namespace: vault
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: role-tokenreview-binding
  namespace: vault
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:auth-delegator
subjects:
- kind: ServiceAccount
  name: vault-reviewer
  namespace: vault
```

```shell
$ kubectl apply -f vault-reviewer.yaml
```

Configure Vault with the vault-reviewer token and Kubernetes CA to fetch secrets.

```shell
$ VAULT_SA_TOKEN_NAME=$(kubectl get sa vault-reviewer -n vault -o jsonpath="{.secrets[*]['name']}")

$ SA_JWT_TOKEN=$(kubectl get secret -n vault "$VAULT_SA_TOKEN_NAME" -o jsonpath="{.data.token}" | base64 --decode; echo)

$ SA_CA_CRT=$(kubectl get secret -n vault "$VAULT_SA_TOKEN_NAME" -o jsonpath="{.data['ca\.crt']}" | base64 --decode; echo)
```

```shell
$ vault write auth/kubernetes/config token_reviewer_jwt="$SA_JWT_TOKEN" kubernetes_host=https://kubernetes.default kubernetes_ca_cert="$SA_CA_CRT"
Success! Data written to: auth/kubernetes/config
```

Create a policy in vault for Kubernetes to read the secrets.

```hcl
path "secret/*" {
  capabilities = ["read", "list"]
}
```

```shell
$ vault policy write k8s_policy policy.hcl
```

Create a service-account which can be associated with the application pod to fetch the secrets.

```shell
$ kubectl create sa tester
```

Associate the role to service-account.

```shell
$ vault write auth/kubernetes/role/k8s_role \
  bound_service_account_names=tester \
  bound_service_account_namespaces=default \
  policies=k8s_policy \
  ttl=1h
```

Let’s try to put some secret inside Vault.

```shell
$ vault kv put secret/mysql MYSQL_ROOT_PASSWORD=password
```

![](./images/vault-ui.png)
