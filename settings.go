package main

const (
	// VolumeMountGoogleCloudKeyPath  the path where the gcp service acount credentials would be mount to
	VolumeMountGoogleCloudKeyPath = "/var/run/secret/cloud.google.com"

	// VolumeMountGoogleCloudKeyName the name of the volume for the gcp service account
	VolumeMountGoogleCloudKeyName = "google-cloud-key"

	// GCPServiceAccountCredentialsFileName the name of the generated credentials file for gcp
	GCPServiceAccountCredentialsFileName = "service-account.json"

	// VaultTLSMountPath path where to mount the vault TLS secret
	VaultTLSMountPath = "/etc/tls/"

	// VaultTLSVolumeName name of the volume for the vault TLS certs and keys
	VaultTLSVolumeName = "vault-tls"
)
