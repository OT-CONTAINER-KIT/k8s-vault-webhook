### v2.0
##### May 8, 2021

#### :tada: [Features Added]

- Added AWS Secret Manager support
- Inject secret directly to pods/containers from AWS Secret Manager
- Authentication with AWS Secret Manager with access key and iam role 

### v1.0
##### April 11, 2021

#### :tada: [Features Added]

- Authentication to Hashicorp vault using Kubernetes service-account
- RBAC implementation of vault using different policies of vault and association of policy with service-account
- Inject secret directly to pods/containers running inside Kubernetes
- Support regex to inject all secrets from a certain path of Vault
- Inject secrets directly to the process of container, i.e. after the injection you cannot read secrets from the environment variable
