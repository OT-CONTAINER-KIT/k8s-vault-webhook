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

From a new terminal execute:-

```shell
kubectl port-forward vault-0 8200:8200 -n vault
```

**Unseal your vault**

## Vault webhook environment setup

Make sure you are logged in to vault using the root token

```shell
$ vault login
  Token (will be hidden):
```

Enable the Kubernetes auth backend

```shell
$ vault auth enable kubernetes
Success! Enabled kubernetes auth method at: kubernetes/
```

Create a file with name `vault-reviewer.yaml`

```yaml
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
kubectl apply -f vault-reviewer.yaml
```

Configure Vault with the vault-reviewer token and Kubernetes CA:

```shell
VAULT_SA_TOKEN_NAME=$(kubectl get sa vault-reviewer -n vault -o jsonpath="{.secrets[*]['name']}")

SA_JWT_TOKEN=$(kubectl get secret -n vault "$VAULT_SA_TOKEN_NAME" -o jsonpath="{.data.token}" | base64 --decode; echo)

SA_CA_CRT=$(kubectl get secret -n vault "$VAULT_SA_TOKEN_NAME" -o jsonpath="{.data['ca\.crt']}" | base64 --decode; echo)

$ vault write auth/kubernetes/config token_reviewer_jwt="$SA_JWT_TOKEN" kubernetes_host=https://kubernetes.default kubernetes_ca_cert="$SA_CA_CRT"
Success! Data written to: auth/kubernetes/config
```

Create a policy in vault:

```hcl
path "secret/*" {
  capabilities = ["read", "list"]
}
```

```shell
$ vault policy write test_policy test-policy.hcl
```

Create a kubernetes serviceaccount:

```shell
$ kubectl create sa tester
```

Create a role inside vault:

```shell
$ vault write auth/kubernetes/role/tester \
  bound_service_account_names=tester \
  bound_service_account_namespaces=default \
  policies=test_policy \
  ttl=1h
```

## Vault webhook setup

```shell
$ helm repo add ot-helm https://github.com/OT-CONTAINER-KIT/helm-charts

$ helm upgrade k8s-vault-webhook ot-helm/k8s-vault-webhook --namespace vault --install
```

```shell
$ kubectl get pods -n vault
```

## Testing

Testing can be done via using MySQL

```shell
$ kubectl apply -f example/mysql-deployment.yaml
```

```shell
$ kubectl get pods
```
