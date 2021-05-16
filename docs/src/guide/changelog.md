### v4.0
**May 16, 2021**

**:tada: [Features Added]**

- Added GCP secret Manager support
- Added CI pipeline using Azure DevOps
- Authenticate and authorize using GCP service-account and annotations
- Secret injection in pods/containers from GCP Secret Manager

### v3.0
**May 9, 2021**

**:tada: [Features Added]**

- Added Azure Key Vault support
- Fetch secrets from Azure Key Vault and inject them in pods/containers
- Pod AD identity and Service principal based authentication in Azure

### v2.0
**May 8, 2021**

**:tada: [Features Added]**

- Added AWS Secret Manager support
- Inject secret directly to pods/containers from AWS Secret Manager
- Authentication with AWS Secret Manager with access key and iam role 

### v1.0
**April 11, 2021**

**:tada: [Features Added]**

- Authentication to Hashicorp vault using Kubernetes service-account
- RBAC implementation of vault using different policies of vault and association of policy with service-account
- Inject secret directly to pods/containers running inside Kubernetes
- Support regex to inject all secrets from a certain path of Vault
- Inject secrets directly to the process of container, i.e. after the injection you cannot read secrets from the environment variable
