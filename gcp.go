package main

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

type gcp struct {
	config struct {
		enabled                     bool
		projectID                   string
		secretName                  string
		secretVersion               string
		serviceAccountKeySecretName string
	}
}

func (gcp *gcp) mutateContainer(container corev1.Container) corev1.Container {
	container = gcp.setArgs(container)

	// Mount google service account key if given
	if gcp.config.serviceAccountKeySecretName != "" {
		container.VolumeMounts = append(container.VolumeMounts, []corev1.VolumeMount{
			{
				Name:      VolumeMountGoogleCloudKeyName,
				MountPath: VolumeMountGoogleCloudKeyPath,
			},
		}...)
	}

	return container
}

func (gcp *gcp) setArgs(c corev1.Container) corev1.Container {
	args := []string{"gcp"}
	args = append(args, fmt.Sprintf("--project-id=%s", gcp.config.projectID))

	if gcp.config.secretName != "" {
		args = append(args, fmt.Sprintf("--secret-name=%s", gcp.config.secretName))
	}

	if gcp.config.secretVersion != "" {
		args = append(args, fmt.Sprintf("--secret-version=%s", gcp.config.secretVersion))
	}

	if gcp.config.secretName != "" {
		args = append(args, fmt.Sprintf("--google-application-credentials=%s", fmt.Sprintf("%s/%s", VolumeMountGoogleCloudKeyPath, GCPServiceAccountCredentialsFileName)))
	}

	args = append(args, "--")
	c.Args = append(args, c.Args...)
	return c
}
