package main

import (
	log "github.com/sirupsen/logrus"
	"k8s-vault-webhook/registry"
	"k8s.io/client-go/kubernetes"
)

type secretManagerConfig struct {
	vault
	aws
	// explicitSecrets bool // only get secrets that match the prefix `secret:`
}

type mutatingWebhook struct {
	k8sClient kubernetes.Interface
	registry  registry.ImageRegistry
	logger    log.FieldLogger
}
