package main

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

type aws struct {
	config struct {
		enabled         bool
		region          string
		secretName      string
		previousVersion string
		roleARN         string
	}
}

func (aws *aws) mutateContainer(container corev1.Container) corev1.Container {
	container = aws.setArgs(container)
	return container
}

func (aws *aws) setArgs(c corev1.Container) corev1.Container {
	args := []string{"aws"}
	args = append(args, fmt.Sprintf("--region=%s", aws.config.region))

	if aws.config.secretName != "" {
		args = append(args, fmt.Sprintf("--secret-name=%s", aws.config.secretName))
	}

	if aws.config.roleARN != "" {
		args = append(args, fmt.Sprintf("--role-arn=%s", aws.config.roleARN))
	}

	if aws.config.secretName != "" {
		args = append(args, fmt.Sprintf("--previous-version=%s", aws.config.previousVersion))
	}

	args = append(args, "--")
	c.Args = append(args, c.Args...)
	return c
}
