package main

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

type azure struct {
	config struct {
		enabled           bool
		azureKeyVaultName string
	}
}

func (azure *azure) mutateContainer(container corev1.Container) corev1.Container {
	container = azure.setArgs(container)
	return container
}

func (azure *azure) setArgs(c corev1.Container) corev1.Container {
	args := []string{"azure"}
	args = append(args, fmt.Sprintf("--azure-vault-name=%s", azure.config.azureKeyVaultName))

	args = append(args, "--")
	c.Args = append(args, c.Args...)
	return c
}
