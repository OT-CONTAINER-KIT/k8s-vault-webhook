package main

const (
	// VaultTLSMountPath path where to mount the vault TLS secret
	VaultTLSMountPath = "/etc/tls/"

	// VaultTLSVolumeName name of the volume for the vault TLS certs and keys
	VaultTLSVolumeName = "vault-tls"
)
