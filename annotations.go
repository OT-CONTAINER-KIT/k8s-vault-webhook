package main

const (

	// AnnotationVaultEnabled if enabled use vault as the secret manager
	AnnotationVaultEnabled = "vault.opstree.secret.manager/enabled"

	// AnnotationVaultService vault address in the http/https format including the port number
	// for example https://vault.vault.svc:8200
	AnnotationVaultService = "vault.opstree.secret.manager/service"

	// AnnotationVaultAuthPath specifies the mount path to be used for the Kubernetes auto-auth method.
	AnnotationVaultAuthPath = "vault.opstree.secret.manager/auth-path"

	// AnnotationVaultSecretPath the secret path in vault - will auto detect if kv2 is used and auto-append `data` to it
	AnnotationVaultSecretPath = "vault.opstree.secret.manager/path"

	// AnnotationVaultRole specifies the role to be used for the Kubernetes auto-auth method.
	AnnotationVaultRole = "vault.opstree.secret.manager/role"

	// AnnotationVaultTLSSecret is the name of the Kubernetes secret containing
	// client TLS certificates and keys.
	AnnotationVaultTLSSecret = "vault.opstree.secret.manager/tls-secret"

	// AnnotationVaultCACert is the filename of the CA certificate used to verify Vault's
	// CA certificate.
	AnnotationVaultCACert = "vault.opstree.secret.manager/ca-cert"

	// AnnotationVaultK8sTokenPath override the token that will be used for vault authentication
	AnnotationVaultK8sTokenPath = "vault.opstree.secret.manager/k8s-token-path"

	// AnnotationVaultUseSecretNamesAsKeys is used with a path that has a tree under it,
	// will be using the secret names as the keys and ignore the real key in the secret itself
	AnnotationVaultUseSecretNamesAsKeys = "vault.opstree.secret.manager/use-secret-names-as-keys"

	// AnnotationVaultSecretVersion get the specified secret version, default to latest version
	AnnotationVaultSecretVersion = "vault.opstree.secret.manager/secret-version"

	// AnnotationVaultMultiSecretPrefix allow multi secret by order
	// vault.secret.manager/secret-config-1: '{"Path": "secrets/v2/plain/secrets/path/app", "Version": "2", "use-secret-names-as-keys": "true"}'
	AnnotationVaultMultiSecretPrefix = "vault.opstree.secret.manager/secret-config-"
)
