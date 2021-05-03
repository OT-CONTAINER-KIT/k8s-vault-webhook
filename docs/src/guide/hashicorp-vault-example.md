# Hashicorp Vault

Let's try to create a deployment to inject secrets directly from Vault. For example, purpose we are taking mysql as deployment and then we will try to set mysql root password using k8s-vault-webhook.

We can use our [example](https://github.com/OT-CONTAINER-KIT/k8s-vault-webhook/tree/master/example) folder.

We have defined environment variable in mysql deployment something like this:-

```yaml
  template:
    metadata:
      labels:
        app: k8s-vault-webhook
        tier: mysql
      annotations:
        vault.opstree.secret.manager/enabled: "true"
        vault.opstree.secret.manager/path: secret/mysql
        vault.opstree.secret.manager/role: tester
        vault.opstree.secret.manager/service: http://vault.vault:8200
    spec:
      serviceAccount: tester
      serviceAccountName: tester
      containers:
      - image: mysql:5.6
        name: mysql
        env:
        - name: MYSQL_ROOT_PASSWORD
          value: vault:MYSQL_ROOT_PASSWORD
        - name: MYSQL_PASSWORD
          value: vault:MYSQL_ROOT_PASSWORD
```

```shell
$ kubectl apply -f example/mysql-deployment.yaml
...
deployment.apps/k8s-vault-webhook-mysql configured
```

Verify the mysql pods are running or not by using `kubectl` command line.

```shell
$ kubectl get pods
...
NAME                                       READY   STATUS    RESTARTS   AGE
k8s-vault-webhook-mysql-6d5df45956-nfmhv   1/1     Running   0          79s
```

Now let's try to get inside the `mysql` pod and see if the Vault password is working fine or not.

```shell
$ kubectl exec -it k8s-vault-webhook-mysql-6d5df45956-nfmhv \
    -- mysql -u root -ppassword -e "show databases;"
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
$ kubectl exec -it k8s-vault-webhook-mysql-6d5df45956-nfmhv \
    -- env | grep MYSQL_ROOT_PASSWORD
...
MYSQL_ROOT_PASSWORD=vault:MYSQL_ROOT_PASSWORD
MYSQL_PASSWORD=vault:MYSQL_ROOT_PASSWORD
```
