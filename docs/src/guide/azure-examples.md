# Azure Key Vault

Let's try to create a deployment to inject secrets directly from Azure Key Vault. For example, purpose we are taking mysql as deployment and then we will try to set mysql root password using k8s-vault-webhook.

We can use our [example](https://github.com/OT-CONTAINER-KIT/k8s-vault-webhook/tree/master/example) folder.

The environment variables will get substitute automatically, we just have to provide some custom annotations.

```yaml
  template:
    metadata:
      labels:
        app: k8s-azure-mysql
        tier: mysql
# Use Azure App Pod Pinding if cluster is configured in Azure
        # aadpodidbinding: POD_IDENTITY_NAME
      annotations:
        azure.opstree.secret.manager/enabled: "true"
        azure.opstree.secret.manager/vault-name: "vault-k8s-secret"
    spec:
      containers:
      - image: opstree/mysql:latest
        name: mysql
# If running outside Azure
        env:
        - name: AZURE_CLIENT_SECRET
          valueFrom:
            secretKeyRef:
              name: azure-secret
              key: AZURE_CLIENT_SECRET
        - name: AZURE_CLIENT_ID
          valueFrom:
            secretKeyRef:
              name: azure-secret
              key: AZURE_CLIENT_ID
        - name: AZURE_TENANT_ID
          valueFrom:
            secretKeyRef:
              name: azure-secret
              key: AZURE_TENANT_ID
```

Let's try to apply the deployment manifest.

```shell
$ kubectl apply -f example/azure-mysql-example.yaml
...
deployment.apps/k8s-azure-mysql configured
```

Verify the mysql pods are running or not by using `kubectl` command line.

```shell
$ kubectl get pods
...
NAME                                       READY   STATUS             RESTARTS   AGE
k8s-azure-mysql-658b99f8dc-k9r58           1/1     Running            0          128m
```

Now let's try to get inside the `mysql` pod and see if the Azure Key Vault's password is working fine or not.

```shell
$ kubectl exec -it k8s-azure-mysql-658b99f8dc-k9r58 \
    -- mysql -u root -pazurepassword -e "show databases;"
...
Warning: Using a password on the command line interface can be insecure.
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
+--------------------+
```

Also, try to check the value in environment variable of MySQL pod.

```shell
$ kubectl exec -it k8s-azure-mysql-658b99f8dc-k9r58 \
    -- env | grep ROOT
...
No output
```
