package main

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"k8s-vault-webhook/registry"
	"k8s-vault-webhook/version"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	whhttp "github.com/slok/kubewebhook/pkg/http"
	"github.com/slok/kubewebhook/pkg/observability/metrics"
	whcontext "github.com/slok/kubewebhook/pkg/webhook/context"
	"github.com/slok/kubewebhook/pkg/webhook/mutating"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	// "k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	// kubeVer "k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/kubernetes"
	kubernetesConfig "sigs.k8s.io/controller-runtime/pkg/client/config"
)

func (mw *mutatingWebhook) getVolumes(existingVolumes []corev1.Volume, secretManagerConfig secretManagerConfig) []corev1.Volume {
	mw.logger.Debugf("Adding generic volumes to podspec")

	volumes := []corev1.Volume{
		{
			Name: "k8s-secret-injector",
			VolumeSource: corev1.VolumeSource{
				EmptyDir: &corev1.EmptyDirVolumeSource{
					Medium: corev1.StorageMediumMemory,
				},
			},
		},
	}

	if secretManagerConfig.vault.config.tlsSecretName != "" {
		mw.logger.Debugf("Adding Vault TLS Volume to podspec")
		volumes = append(volumes, []corev1.Volume{
			{
				Name: "vault-tls",
				VolumeSource: corev1.VolumeSource{
					Secret: &corev1.SecretVolumeSource{
						SecretName: secretManagerConfig.vault.config.tlsSecretName,
					},
				},
			},
		}...)
	}
	return volumes
}

func getInitContainers(originalContainers []corev1.Container, secretManagerConfig secretManagerConfig, initContainersMutated bool, containersMutated bool) []corev1.Container {
	var containers = []corev1.Container{}

	if initContainersMutated || containersMutated {
		containers = append(containers, corev1.Container{
			Name:            "copy-k8s-secret-injector",
			Image:           viper.GetString("k8s_secret_injector_image"),
			ImagePullPolicy: corev1.PullPolicy(viper.GetString("k8s_secret_injector_image_pull_policy")),
			Command:         []string{"sh", "-c", "cp /usr/local/bin/k8s-secret-injector /k8s-secret/"},
			VolumeMounts: []corev1.VolumeMount{
				{
					Name:      "k8s-secret-injector",
					MountPath: "/k8s-secret/",
				},
			},
			Resources: corev1.ResourceRequirements{
				Limits: corev1.ResourceList{
					corev1.ResourceCPU:    resource.MustParse("50m"),
					corev1.ResourceMemory: resource.MustParse("64Mi"),
				},
			},
		})
	}

	return containers
}

func hasSecretPrefix(value string) bool {
	return strings.HasPrefix(value, "vault:") || strings.HasPrefix(value, ">>secret:") || strings.HasPrefix(value, "secret:")
}

func (mw *mutatingWebhook) getDataFromConfigmap(cmName string, ns string) (map[string]string, error) {
	ctx := context.Background()
	configMap, err := mw.k8sClient.CoreV1().ConfigMaps(ns).Get(ctx, cmName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return configMap.Data, nil
}

func (mw *mutatingWebhook) getDataFromSecret(secretName string, ns string) (map[string][]byte, error) {
	ctx := context.Background()
	secret, err := mw.k8sClient.CoreV1().Secrets(ns).Get(ctx, secretName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return secret.Data, nil
}

func (mw *mutatingWebhook) lookForValueFrom(env corev1.EnvVar, ns string) (*corev1.EnvVar, error) {
	if env.ValueFrom.ConfigMapKeyRef != nil {
		data, err := mw.getDataFromConfigmap(env.ValueFrom.ConfigMapKeyRef.Name, ns)
		if err != nil {
			if apierrors.IsNotFound(err) {
				return nil, nil
			}
			return nil, err
		}
		if hasSecretPrefix(data[env.ValueFrom.ConfigMapKeyRef.Key]) {
			fromCM := corev1.EnvVar{
				Name:  env.Name,
				Value: data[env.ValueFrom.ConfigMapKeyRef.Key],
			}
			return &fromCM, nil
		}
	}
	if env.ValueFrom.SecretKeyRef != nil {
		data, err := mw.getDataFromSecret(env.ValueFrom.SecretKeyRef.Name, ns)
		if err != nil {
			if apierrors.IsNotFound(err) {
				return nil, nil
			}
			return nil, err
		}
		if hasSecretPrefix(string(data[env.ValueFrom.SecretKeyRef.Key])) {
			fromSecret := corev1.EnvVar{
				Name:  env.Name,
				Value: string(data[env.ValueFrom.SecretKeyRef.Key]),
			}
			return &fromSecret, nil
		}
	}
	return nil, nil
}

func (mw *mutatingWebhook) lookForEnvFrom(envFrom []corev1.EnvFromSource, ns string) ([]corev1.EnvVar, error) {
	var envVars []corev1.EnvVar

	for _, ef := range envFrom {
		if ef.ConfigMapRef != nil {
			data, err := mw.getDataFromConfigmap(ef.ConfigMapRef.Name, ns)
			if err != nil {
				if apierrors.IsNotFound(err) || (ef.ConfigMapRef.Optional != nil && *ef.ConfigMapRef.Optional) {
					continue
				} else {
					return envVars, err
				}
			}
			for key, value := range data {
				if hasSecretPrefix(value) {
					envFromCM := corev1.EnvVar{
						Name:  key,
						Value: value,
					}
					envVars = append(envVars, envFromCM)
				}
			}
		}
		if ef.SecretRef != nil {
			data, err := mw.getDataFromSecret(ef.SecretRef.Name, ns)
			if err != nil {
				if apierrors.IsNotFound(err) || (ef.SecretRef.Optional != nil && *ef.SecretRef.Optional) {
					continue
				} else {
					return envVars, err
				}
			}
			for key, value := range data {
				if hasSecretPrefix(string(value)) {
					envFromSec := corev1.EnvVar{
						Name:  key,
						Value: string(value),
					}
					envVars = append(envVars, envFromSec)
				}
			}
		}
	}
	return envVars, nil
}

func (mw *mutatingWebhook) mutateContainers(containers []corev1.Container, podSpec *corev1.PodSpec, secretManagerConfig secretManagerConfig, ns string) (bool, error) {
	mutated := false
	var mutationInProgress bool
	for i, container := range containers {
		var envVars []corev1.EnvVar
		if len(container.EnvFrom) > 0 {
			envFrom, err := mw.lookForEnvFrom(container.EnvFrom, ns) //nolint
			if err != nil {
				return false, err
			}
			envVars = append(envVars, envFrom...) //nolint
		}
		for _, env := range container.Env {
			if hasSecretPrefix(env.Value) {
				envVars = append(envVars, env) //nolint
			}
			if env.ValueFrom != nil {
				valueFrom, err := mw.lookForValueFrom(env, ns)
				if err != nil {
					return false, err
				}
				if valueFrom == nil {
					continue
				}
				envVars = append(envVars, *valueFrom) //nolint
			}
		}

		args := container.Command

		// the container has no explicitly specified command
		if len(args) == 0 {
			mw.logger.Info("No command was given - attempting to get image metadata")
			imageConfig, err := mw.registry.GetImageConfig(mw.k8sClient, ns, &container, podSpec)
			if err != nil {
				return false, err
			}

			args = append(args, imageConfig.Entrypoint...)

			// If no Args are defined we can use the Docker CMD from the image
			// https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#notes
			if len(container.Args) == 0 {
				args = append(args, imageConfig.Cmd...)
			}
		}
		args = append(args, container.Args...)

		container.Command = []string{"/k8s-secret/k8s-secret-injector"}
		container.Args = args

		if secretManagerConfig.azure.config.enabled {
			container = secretManagerConfig.azure.mutateContainer(container)
			mutationInProgress = true
		}

		if secretManagerConfig.aws.config.enabled {
			container = secretManagerConfig.aws.mutateContainer(container)
			mutationInProgress = true
		}

		if secretManagerConfig.vault.config.enabled {
			container = secretManagerConfig.vault.mutateContainer(container)
			mutationInProgress = true
		}

		if !mutationInProgress {
			continue
		}
		mutated = true

		// add the volume mount for k8s-secret-injector
		container.VolumeMounts = append(container.VolumeMounts, []corev1.VolumeMount{
			{
				Name:      "k8s-secret-injector",
				MountPath: "/k8s-secret",
			},
		}...)

		containers[i] = container
	}
	return mutated, nil
}

func (mw *mutatingWebhook) mutatePod(pod *corev1.Pod, secretManagerConfig secretManagerConfig, ns string, dryRun bool) error {
	mw.logger.Debugf("Successfully connected to the API")

	initContainersMutated, err := mw.mutateContainers(pod.Spec.InitContainers, &pod.Spec, secretManagerConfig, ns)
	if err != nil {
		return err
	}

	if initContainersMutated {
		mw.logger.Debugf("Successfully mutated pod init containers")
	} else {
		mw.logger.Debugf("No pod init containers were mutated")
	}

	containersMutated, err := mw.mutateContainers(pod.Spec.Containers, &pod.Spec, secretManagerConfig, ns)
	if err != nil {
		return err
	}

	if containersMutated {
		mw.logger.Debugf("Successfully mutated pod containers")
	} else {
		mw.logger.Debugf("No pod containers were mutated")
	}

	if initContainersMutated || containersMutated {
		pod.Spec.InitContainers = append(getInitContainers(pod.Spec.Containers, secretManagerConfig, initContainersMutated, containersMutated), pod.Spec.InitContainers...)
		mw.logger.Debugf("Successfully appended pod init containers to spec")

		pod.Spec.Volumes = append(pod.Spec.Volumes, mw.getVolumes(pod.Spec.Volumes, secretManagerConfig)...)
		mw.logger.Debugf("Successfully appended pod spec volumes")
	}

	if viper.GetString("k8s_secret_injector_image_pull_secret_name") != "" {
		pod.Spec.ImagePullSecrets = append(pod.Spec.ImagePullSecrets, corev1.LocalObjectReference{Name: viper.GetString("k8s_secret_injector_image_pull_secret_name")})
	}

	return nil
}

// take all the annotations (m), filter the prefix(delemiter) "secret-config-" and sort them alpha-numeric
func filterAndSortMapNumStr(m map[string]string, delimiter string) ([]string, error) {
	keys := make([]string, 0, len(m))

	for k := range m {
		if strings.HasPrefix(k, delimiter) {
			keys = append(keys, k)
		}
	}

	if len(keys) > 0 {
		sort.Slice(keys, func(i, j int) bool {
			numA, _ := strconv.Atoi(strings.Split(keys[i], delimiter)[1])
			numB, _ := strconv.Atoi(strings.SplitAfter(keys[j], delimiter)[1])
			return numA < numB
		})
	}

	return keys, nil
}

func (mw *mutatingWebhook) parseSecretManagerConfig(obj metav1.Object) secretManagerConfig {
	var smCfg secretManagerConfig
	annotations := obj.GetAnnotations()

	smCfg.azure.config.enabled, _ = strconv.ParseBool(annotations[AnnotationAzureKeyVaultEnabled])
	smCfg.azure.config.azureKeyVaultName = annotations[AnnotationAzureKeyVaultName]

	smCfg.aws.config.enabled, _ = strconv.ParseBool(annotations[AnnotationAWSSecretManagerEnabled])
	smCfg.aws.config.region = annotations[AnnotationAWSSecretManagerRegion]
	smCfg.aws.config.roleARN = annotations[AnnotationAWSSecretManagerRoleARN]
	smCfg.aws.config.secretName = annotations[AnnotationAWSSecretManagerSecretName]
	smCfg.aws.config.previousVersion = annotations[AnnotationAWSSecretManagerPreviousVersion]

	smCfg.vault.config.enabled, _ = strconv.ParseBool(annotations[AnnotationVaultEnabled])
	smCfg.vault.config.addr = annotations[AnnotationVaultService]
	smCfg.vault.config.path = annotations[AnnotationVaultSecretPath]
	smCfg.vault.config.role = annotations[AnnotationVaultRole]
	smCfg.vault.config.tlsSecretName = annotations[AnnotationVaultTLSSecret]
	smCfg.vault.config.vaultCACert = annotations[AnnotationVaultCACert]
	smCfg.vault.config.tokenPath = annotations[AnnotationVaultK8sTokenPath]
	smCfg.vault.config.backend = annotations[AnnotationVaultAuthPath]
	smCfg.vault.config.useSecretNamesAsKeys, _ = strconv.ParseBool(annotations[AnnotationVaultUseSecretNamesAsKeys])
	smCfg.vault.config.version = annotations[AnnotationVaultSecretVersion]
	smCfg.vault.config.kubernetesBackend = annotations[AnnotationVaultAuthPath]
	smCfg.vault.config.secretConfigs = []string{}
	keys, err := filterAndSortMapNumStr(annotations, AnnotationVaultMultiSecretPrefix)

	if err != nil {
		mw.logger.Warnf("sorting annotations of %s failed! %+v", AnnotationVaultMultiSecretPrefix, err)
	}

	for _, k := range keys {
		smCfg.vault.config.secretConfigs = append(smCfg.vault.config.secretConfigs, annotations[k])
	}

	return smCfg
}

// SecretsMutator if object is Pod mutate pod specs
// return a stop boolean to stop executing the chain and also an error.
func (mw *mutatingWebhook) SecretsMutator(ctx context.Context, obj metav1.Object) (bool, error) {
	smCfg := mw.parseSecretManagerConfig(obj)
	mw.logger.Debugf("Secret Managers config: %#v", smCfg)

	switch v := obj.(type) {
	case *corev1.Pod:
		if smCfg.azure.config.enabled {
			mw.logger.Infof("Using Azure Key Vault")

			if smCfg.azure.config.azureKeyVaultName == "" {
				return true, fmt.Errorf("Error getting azure key vault - make sure you set the annotation %s on the Pod", AnnotationAzureKeyVaultName)
			}

			return false, mw.mutatePod(v, smCfg, whcontext.GetAdmissionRequest(ctx).Namespace, whcontext.IsAdmissionRequestDryRun(ctx))
		}

		if smCfg.aws.config.enabled {
			mw.logger.Infof("Using AWS Secret Manager")

			if smCfg.aws.config.secretName == "" {
				return true, fmt.Errorf("Error getting aws secret name - make sure you set the annotation %s on the Pod", AnnotationAWSSecretManagerSecretName)
			}

			return false, mw.mutatePod(v, smCfg, whcontext.GetAdmissionRequest(ctx).Namespace, whcontext.IsAdmissionRequestDryRun(ctx))
		}

		if smCfg.vault.config.enabled {
			var err error
			mw.logger.Info("Using Vault Secret Manager")

			if smCfg.vault.config.addr == "" {
				err = fmt.Errorf("Error getting vault service address - make sure you set the annotation %s on the Pod", AnnotationVaultEnabled)
			}

			if smCfg.vault.config.path == "" && len(smCfg.vault.config.secretConfigs) == 0 {
				err = fmt.Errorf("Error getting vault secret path - make sure you either set the annotation %s or use the annotation %s-x where x is the secret number", AnnotationVaultSecretPath, AnnotationVaultMultiSecretPrefix)
			}

			if smCfg.vault.config.role == "" {
				err = fmt.Errorf("Error getting vault role - make sure you set the annotation %s", AnnotationVaultRole)
			}

			if smCfg.vault.config.tlsSecretName != "" && smCfg.vault.config.vaultCACert == "" {
				err = fmt.Errorf("Error getting CA cert filename - make sure you set the annotation %s with the CA cert file name", AnnotationVaultCACert)
			}

			if err != nil {
				return true, err
			}

			return false, mw.mutatePod(v, smCfg, whcontext.GetAdmissionRequest(ctx).Namespace, whcontext.IsAdmissionRequestDryRun(ctx))
		}
		return false, nil
	default:
		return false, nil
	}
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (mw *mutatingWebhook) serveMetrics(addr string) {
	mw.logger.Infof("Telemetry on http://%s", addr)

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		mw.logger.Fatalf("error serving telemetry: %s", err)
	}
}

func handlerFor(config mutating.WebhookConfig, mutator mutating.Mutator, recorder metrics.Recorder, logger logrus.FieldLogger) http.Handler {
	webhook, err := mutating.NewWebhook(config, mutator, nil, recorder, logger)
	if err != nil {
		logger.Fatalf("error creating webhook: %s", err)
	}

	handler, err := whhttp.HandlerFor(webhook)
	if err != nil {
		logger.Fatalf("error creating webhook: %s", err)
	}

	return handler
}

func newK8SClient() (kubernetes.Interface, error) {
	kubeConfig, err := kubernetesConfig.GetConfig()
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(kubeConfig)
}

func init() {
	viper.SetDefault("k8s_secret_injector_image", "quay.io/opstree/k8s-secret-injector:3.0")
	viper.SetDefault("k8s_secret_injector_image_pull_policy", string(corev1.PullIfNotPresent))
	viper.SetDefault("k8s_secret_injector_image_pull_secret_name", "")
	viper.SetDefault("tls_cert_file", "")
	viper.SetDefault("tls_private_key_file", "")
	viper.SetDefault("listen_address", ":8443")
	viper.SetDefault("debug", "false")
	viper.SetDefault("enable_json_log", "false")
	viper.SetDefault("telemetry_listen_address", "")
	viper.AutomaticEnv()
}

func main() {
	var logger logrus.FieldLogger
	{
		log := logrus.New()

		if viper.GetBool("enable_json_log") {
			log.SetFormatter(&logrus.JSONFormatter{})
		}

		if viper.GetBool("debug") {
			log.SetLevel(logrus.DebugLevel)
			log.Debug("Debug mode enabled")
		}

		logger = log.WithField("app", "k8s-secret-injector")
	}
	fmt.Printf("K8s Vault Webhook Version: %s\n", version.GetVersion())
	fmt.Printf("K8s Secret Injector Version: %s\n", viper.GetString("k8s_secret_injector_image"))

	k8sClient, err := newK8SClient()
	if err != nil {
		logger.Fatalf("error creating k8s client: %s", err)
	}

	mutatingWebhook := mutatingWebhook{
		k8sClient: k8sClient,
		registry:  registry.NewRegistry(),
		logger:    logger,
	}

	mutator := mutating.MutatorFunc(mutatingWebhook.SecretsMutator)

	metricsRecorder := metrics.NewPrometheus(prometheus.DefaultRegisterer)

	podHandler := handlerFor(mutating.WebhookConfig{Name: "k8s-secret-injector-pods", Obj: &corev1.Pod{}}, mutator, metricsRecorder, logger)

	mux := http.NewServeMux()
	mux.Handle("/pods", podHandler)
	mux.Handle("/healthz", http.HandlerFunc(healthzHandler))

	telemetryAddress := viper.GetString("telemetry_listen_address")
	listenAddress := viper.GetString("listen_address")
	tlsCertFile := viper.GetString("tls_cert_file")
	tlsPrivateKeyFile := viper.GetString("tls_private_key_file")

	if len(telemetryAddress) > 0 {
		// Serving metrics without TLS on separated address
		go mutatingWebhook.serveMetrics(telemetryAddress)
	} else {
		mux.Handle("/metrics", promhttp.Handler())
	}

	if tlsCertFile == "" && tlsPrivateKeyFile == "" {
		logger.Infof("Listening on http://%s", listenAddress)
		err = http.ListenAndServe(listenAddress, mux)
	} else {
		logger.Infof("Listening on https://%s", listenAddress)
		err = http.ListenAndServeTLS(listenAddress, tlsCertFile, tlsPrivateKeyFile, mux)
	}

	if err != nil {
		log.Fatalf("error serving webhook: %s", err)
	}
}
