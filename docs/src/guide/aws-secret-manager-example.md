# AWS Secret Manager

Let's try to create a deployment to inject secrets directly from AWS Secret Manager. For example, purpose we are taking mysql as deployment and then we will try to set mysql root password using k8s-vault-webhook.

We can use our [example](https://github.com/OT-CONTAINER-KIT/k8s-vault-webhook/tree/master/example) folder.

The environment variables will get substitute automatically, we just have to provide some custom annotations.

```yaml
  template:
    metadata:
      labels:
        app: k8s-aws-mysql
        tier: mysql
      annotations:
        aws.opstree.secret.manager/enabled: "true"
        aws.opstree.secret.manager/region: "us-west-2"
# Use this role-arn if cluster is configured in AWS
        # aws.opstree.secret.manager/role-arn: "arn:aws:iam::999:role/secretManager"
        aws.opstree.secret.manager/secret-name: "test-secret"
    spec:
      containers:
      - image: opstree/mysql:latest
        name: mysql
# If running outside AWS
        env:
        - name: AWS_ACCESS_KEY_ID
          valueFrom:
            secretKeyRef:
              name: aws-secret
              key: AWS_ACCESS_KEY_ID
        - name: AWS_SECRET_ACCESS_KEY
          valueFrom:
            secretKeyRef:
              name: aws-secret
              key: AWS_SECRET_ACCESS_KEY
```

Let's try to apply the deployment manifest.

```shell
$ kubectl apply -f example/aws-mysql-example.yaml
...
deployment.apps/k8s-aws-mysql configured
```

Verify the mysql pods are running or not by using `kubectl` command line.

```shell
$ kubectl get pods
...
NAME                                       READY   STATUS    RESTARTS   AGE
k8s-aws-mysql-5fcb986486-npjql             1/1     Running   0          16h
```

Now let's try to get inside the `mysql` pod and see if the AWS Secret Manager's password is working fine or not.

```shell
$ kubectl exec -it k8s-aws-mysql-5fcb986486-npjql \
    -- mysql -u root -pawspassword -e "show databases;"
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
$ kubectl exec -it k8s-aws-mysql-5fcb986486-npjql \
    -- env | grep ROOT
...
No output
```
