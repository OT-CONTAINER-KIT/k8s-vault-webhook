package main

import (
	"k8s-vault-webhook/registry"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
)

type secretManagerConfig struct {
	vault
	explicitSecrets bool // only get secrets that match the prefix `secret:`
}

type mutatingWebhook struct {
	k8sClient kubernetes.Interface
	registry  registry.ImageRegistry
	logger    log.FieldLogger
}
