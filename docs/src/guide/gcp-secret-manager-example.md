# GCP Secret Manager

Let's try to create a deployment to inject secrets directly from GCP Secret Manager. For example, purpose we are taking mysql as deployment and then we will try to set mysql root password using k8s-vault-webhook.

We can use our [example](https://github.com/OT-CONTAINER-KIT/k8s-vault-webhook/tree/master/example) folder.

The environment variables will get substitute automatically, we just have to provide some custom annotations.

```yaml
  template:
    metadata:
      labels:
        app: k8s-gcp-mysql
        tier: mysql
      annotations:
        gcp.opstree.secret.manager/enabled: "true"
        gcp.opstree.secret.manager/project-id: "graceful-flag-209120"
        gcp.opstree.secret.manager/secret-name: "test-secret"
        gcp.opstree.secret.manager/secret-version: "3"
        gcp.opstree.secret.manager/gcp-service-account-key-secret-name: "gcp-sa"
    spec:
      containers:
      - image: opstree/mysql:latest
        name: mysql
        ports:
        - containerPort: 3306
          name: mysql
```

Let's try to apply the deployment manifest.

```shell
$ kubectl apply -f example/gcp-mysql-example.yaml
...
deployment.apps/k8s-gcp-mysql configured
```

Verify the mysql pods are running or not by using `kubectl` command line.

```shell
$ kubectl get pods
...
NAME                                       READY   STATUS    RESTARTS   AGE
k8s-gcp-mysql-7b45bbc486-8w55w             1/1     Running   0          16h
```

Now let's try to get inside the `mysql` pod and see if the GCP Secret Manager's password is working fine or not.

```shell
$ kubectl exec -it k8s-gcp-mysql-7b45bbc486-8w55w \
    -- mysql -u root -pgcppassword -e "show databases;"
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
$ kubectl exec -it k8s-gcp-mysql-7b45bbc486-8w55w \
    -- env | grep ROOT
...
No output
```
