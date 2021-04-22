.. _vault_configuration:

Vault Configuration
===================

For Vault configuration, we are going to use vault-cli and vault-ui both. Download and install vault-cli on your local system. It can be downloaded from `here <https://www.vaultproject.io/downloads>`__ .

First, forward the port of vault to your local system.

.. code:: bash

    $ kubectl port-forward vault-0 8200:8200 -n vault

Login into the vault using the root token which we got using `init` command.

.. code:: bash

    $ vault login
      Token (will be hidden):

Enable the Kubernetes auth backend in the vault cluster

.. code:: bash

    $ vault auth enable kubernetes
    Success! Enabled kubernetes auth method at: kubernetes/

Create a file with name `vault-reviewer.yaml` which will have the service-account and cluster role access information.

.. code:: yaml

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

.. code:: bash

    $ kubectl apply -f vault-reviewer.yaml

Configure Vault with the vault-reviewer token and Kubernetes CA to fetch secrets.

.. code:: bash

    $ VAULT_SA_TOKEN_NAME=$(kubectl get sa vault-reviewer -n vault -o jsonpath="{.secrets[*]['name']}")

    $ SA_JWT_TOKEN=$(kubectl get secret -n vault "$VAULT_SA_TOKEN_NAME" -o jsonpath="{.data.token}" | base64 --decode; echo)

    $ SA_CA_CRT=$(kubectl get secret -n vault "$VAULT_SA_TOKEN_NAME" -o jsonpath="{.data['ca\.crt']}" | base64 --decode; echo)

.. code:: bash

    $ vault write auth/kubernetes/config token_reviewer_jwt="$SA_JWT_TOKEN" kubernetes_host=https://kubernetes.default kubernetes_ca_cert="$SA_CA_CRT"
    Success! Data written to: auth/kubernetes/config

Create a policy in vault for Kubernetes to read the secrets.

.. code:: hcl

    path "secret/*" {
      capabilities = ["read", "list"]
    }

.. code:: bash

    $ vault policy write k8s_policy policy.hcl

Create a service-account which can be associated with the application pod to fetch the secrets.

.. code:: bash

    $ kubectl create sa tester

Associate the role to service-account.

.. code:: bash

    $ vault write auth/kubernetes/role/k8s_role \
      bound_service_account_names=tester \
      bound_service_account_namespaces=default \
      policies=k8s_policy \
      ttl=1h

Let's try to put some secret inside Vault.

.. code:: bash

    $ vault kv put secret/mysql MYSQL_ROOT_PASSWORD=password

.. image:: _static/images/vault-ui.png
    :align: center
    :alt: vault-ui
