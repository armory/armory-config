package sizing_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/deploymentEnv"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func IsDeploymentEnvironmentEmpty(KustomizeData structs.Kustomize) bool {
	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment {
		return false
	}
	return true
}

func GetDeploymentEnvironment(KustomizeData structs.Kustomize) string {
	str := ""

	if !IsDeploymentEnvironmentEmpty(KustomizeData) {
		str = GetDeploymentEnvironmentData(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment) +
			getDeploymentEnvConsul(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment.Consul) +
			getDeploymentEnvVault(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment.Vault) +
			GetDeploymentEnvCustomSizing(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment.CustomSizing) +
			// GetDeploymentEnvSidecar(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment.Sidecar) + //TODO Missing PROTO
			GetDeploymentEnvInitContainers(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment.InitContainers) +
			GetDeploymentEnvHostAliases(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment.HostAliases) +
			GetDeploymentEnvTolerations(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment.Tolerations) +
			GetDeploymentEnvNodeSelectors(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment.NodeSelectors) +
			GetDeploymentEnvGitConfig(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment.GitConfig) +
			GetDeploymentEnvLivenessProbeConfig(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment.LivenessProbeConfig) +
			GetDeploymentEnvHaServices(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment.HaServices) +
			GetDeploymentEnvAffinity(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment.Affinity) +
			GetDeploymentEnvContainers(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment.Containers) +
			GetDeploymentEnvVolumes(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment.Volumes)
	}

	return str
}

func GetDeploymentEnvironmentData(deploymentEnvRef *deploymentEnv.DeploymentEnvironment) string {
	str := ""

	str = `
	size: ` + deploymentEnvRef.Size + `
	type: ` + deploymentEnvRef.Type + `
	accountName: ` + deploymentEnvRef.AccountName + `
	imageVariant: ` + deploymentEnvRef.ImageVariant + `
	bootstrapOnly: ` + strconv.FormatBool(deploymentEnvRef.BootstrapOnly) + `
	updateVersions: ` + strconv.FormatBool(deploymentEnvRef.UpdateVersions) + `
	location: ` + deploymentEnvRef.Location

	str = strings.Replace(str, "\t", "        ", -1)

	return str
}

func getDeploymentEnvConsul(consulReference *deploymentEnv.Consul) string {
	str := ""

	if nil != consulReference {
		str = `
	    consul:
	      enabled: ` + strconv.FormatBool(consulReference.Enabled) + `
	      address: ` + consulReference.Address

		str = strings.Replace(str, "\t", "    ", -1)
	}

	return str
}

func getDeploymentEnvVault(vaultReference *deploymentEnv.Vault) string {
	str := ""

	if nil != vaultReference {
		str = `
	    vault:
	      enabled: ` + strconv.FormatBool(vaultReference.Enabled) + `
	      address: ` + vaultReference.Address

		str = strings.Replace(str, "\t", "    ", -1)
	}

	return str
}

func GetDeploymentEnvCustomSizing(sizingReference *deploymentEnv.CustomSizing) string {
	str := ""

	if nil != sizingReference {
		str = `
	    customSizing:` +
			getServiceSizing(sizingReference.SpinClouddriver, "spin-clouddriver") +
			getServiceSizing(sizingReference.SpinVaultServer, "spin-vault-server") +
			getServiceSizing(sizingReference.Deck, "deck") +
			getServiceSizing(sizingReference.Echo, "echo") +
			getServiceSizing(sizingReference.SpinRosco, "spin-rosco") +
			getServiceSizing(sizingReference.Redis, "redis") +
			getServiceSizing(sizingReference.RedisBootstrap, "redis-bootstrap") +
			getServiceSizing(sizingReference.Clouddriver, "clouddriver") +
			getServiceSizing(sizingReference.Front50, "front50") +
			getServiceSizing(sizingReference.ConsulServer, "consul-server") +
			getServiceSizing(sizingReference.Kayenta, "kayenta") +
			getServiceSizing(sizingReference.ClouddriverRw, "clouddriver-rw") +
			getServiceSizing(sizingReference.SpinClouddriverRoDeck, "spin-clouddriver-ro-deck") +
			getServiceSizing(sizingReference.SpinOrcaBootstrap, "spin-orca-bootstrap") +
			getServiceSizing(sizingReference.ClouddriverRo, "clouddriver-ro") +
			getServiceSizing(sizingReference.SpinIgor, "spin-igor") +
			getServiceSizing(sizingReference.SpinKayenta, "spin-kayenta") +
			getServiceSizing(sizingReference.SpinOrca, "spin-orca") +
			getServiceSizing(sizingReference.VaultServer, "vault-server") +
			getServiceSizing(sizingReference.VaultClient, "vault-client") +
			getServiceSizing(sizingReference.EchoScheduler, "echo-scheduler") +
			getServiceSizing(sizingReference.SpinClouddriverRo, "spin-clouddriver-ro") +
			getServiceSizing(sizingReference.SpinRedis, "spin-redis") +
			getServiceSizing(sizingReference.EchoWorker, "echo-worker") +
			getServiceSizing(sizingReference.OrcaBootstrap, "orca-bootstrap") +
			getServiceSizing(sizingReference.SpinRedisBootstrap, "spin-redis-bootstrap") +
			getServiceSizing(sizingReference.SpinConsulClient, "spin-consul-client") +
			getServiceSizing(sizingReference.SpinMonitoringDaemon, "spin-monitoring-daemon") +
			getServiceSizing(sizingReference.ClouddriverRoDeck, "clouddriver-ro-deck") +
			getServiceSizing(sizingReference.SpinDeck, "spin-deck") +
			getServiceSizing(sizingReference.SpinEcho, "spin-echo") +
			getServiceSizing(sizingReference.ClouddriverBootstrap, "clouddriver-bootstrap") +
			getServiceSizing(sizingReference.SpinEchoScheduler, "spin-echo-scheduler") +
			getServiceSizing(sizingReference.SpinClouddriverCaching, "spin-clouddriver-caching") + //spinClouddriverCaching
			getServiceSizing(sizingReference.SpinClouddriverBootstrap, "spin-clouddriver-bootstrap") +
			getServiceSizing(sizingReference.SpinVaultClient, "spin-vault-client") +
			getServiceSizing(sizingReference.SpinFiat, "spin-fiat") +
			getServiceSizing(sizingReference.SpinEchoWorker, "spin-echo-worker") +
			getServiceSizing(sizingReference.ConsulClient, "consul-client") +
			getServiceSizing(sizingReference.SpinGate, "spin-gate") +
			getServiceSizing(sizingReference.ClouddriverCaching, "clouddriver-caching") +
			getServiceSizing(sizingReference.SpinConsulServer, "spin-consul-server") +
			getServiceSizing(sizingReference.Gate, "gate") +
			getServiceSizing(sizingReference.MonitoringDaemon, "monitoring-daemon") +
			getServiceSizing(sizingReference.SpinClouddriverRw, "spin-clouddriver-rw") +
			getServiceSizing(sizingReference.Fiat, "fiat") +
			getServiceSizing(sizingReference.Igor, "igor") +
			getServiceSizing(sizingReference.Orca, "orca") +
			getServiceSizing(sizingReference.Rosco, "rosco") +
			getServiceSizing(sizingReference.SpinFront50, "spin-front50") +
			getServiceSizing(sizingReference.SpinTerraformer, "spin-terraformer") +
			getServiceSizing(sizingReference.SpinDinghy, "spin-dinghy")

		str = strings.Replace(str, "\t", "    ", -1)
	}
	return str
}

func getServiceSizing(serviceSizingReference *deploymentEnv.ServiceSizing, name string) string {
	str := ""

	if nil != serviceSizingReference {
		str = `
		  ` + name + `:` + `
		    replicas: ` + helpers.IntToString(serviceSizingReference.Replicas) +
			getSizing(serviceSizingReference.Limits, "limits") +
			getSizing(serviceSizingReference.Requests, "requests")
	} else {
		str = `
		  ` + name + `: {}`
	}

	return str
}

func getSizing(sizingReference *deploymentEnv.Sizing, name string) string {
	str := ""

	if nil != sizingReference {
		str = `
		    ` + name + `:` + `
		      cpu: ` + helpers.IntToString(sizingReference.Cpu) + `
		      memory: ` + sizingReference.Memory
	} else {
		str = `
		    ` + name + `: {}`
	}

	return str
}

func GetDeploymentEnvInitContainers(initReference *deploymentEnv.InitContainers) string {
	str := ""

	if nil != initReference {
		str = `
	    initContainers:` +
			getServiceInitContainers(initReference.Front50, "front50") +
			getServiceInitContainers(initReference.Clouddriver, "clouddriver") +
			getServiceInitContainers(initReference.Orca, "orca") +
			getServiceInitContainers(initReference.Rosco, "rosco") +
			getServiceInitContainers(initReference.Gate, "gate") +
			getServiceInitContainers(initReference.Deck, "deck") +
			getServiceInitContainers(initReference.Igor, "igor") //+
			//TODO Check the following Services the "args:"field" has an issue
			// getServiceInitContainers(initReference.SpinOrca, "spin-orca") +
			// getServiceInitContainers(initReference.SpinKayenta, "spin-kayenta") +
			// getServiceInitContainers(initReference.SpinIgor, "spin-igor") +
			// getServiceInitContainers(initReference.SpinEcho, "spin-echo") +
			// getServiceInitContainers(initReference.SpinClouddriverRw, "spin-clouddriver-rw")

		str = strings.Replace(str, "\t", "    ", -1)
	}
	return str
}

func getServiceInitContainers(servicesRef []*deploymentEnv.Service, name string) string {
	str := ""

	if nil != servicesRef {
		for _, services := range servicesRef {
			str = `
		    ` + name + `:` + `
		      - name: ` + services.Name + `
		        image: ` + services.Image +
				getDeploymentEnArray(services.Args, "args") +
				getVolumeMounts(services.VolumeMounts)
		}
	} else {
		str = `
		    ` + name + `: {}`
	}

	return str
}

func GetDeploymentEnvHostAliases(hostAliasesRef *deploymentEnv.HostAliases) string {
	str := ""

	if nil != hostAliasesRef {
		//TODO Missing proto
		str = `
		hostAliases: {}` //TODO remove {}

	} else {
		str = `
		hostAliases: {}`
	}
	str = strings.Replace(str, "\t", "    ", -1)

	return str
}

func GetDeploymentEnvTolerations(tolerationsRef *deploymentEnv.Tolerations) string {
	str := ""

	if nil != tolerationsRef && "" != tolerationsRef.Key {
		str = `
		tolerations:
		  key: ` + tolerationsRef.Key + `
		  operator: ` + tolerationsRef.Operator + `
		  value: ` + tolerationsRef.Value + `
		  effect: ` + tolerationsRef.Effect
	} else {
		str = `
		tolerations: {}`
	}
	str = strings.Replace(str, "\t", "    ", -1)

	return str
}

func GetDeploymentEnvNodeSelectors(nodeSelectorsRef *deploymentEnv.NodeSelectors) string {
	str := ""

	if nil != nodeSelectorsRef {
		//TODO Missing proto
		str = `
		nodeSelectors: {}` //TODO remove {}

	} else {
		str = `
		nodeSelectors: {}`
	}
	str = strings.Replace(str, "\t", "    ", -1)

	return str
}

func GetDeploymentEnvGitConfig(gitConfigRef *deploymentEnv.GitConfig) string {
	str := ""

	if nil != gitConfigRef {
		str = `
		gitConfig:
		  UpstreamUser: ` + gitConfigRef.UpstreamUser + `
		  originUser: ` + gitConfigRef.OriginUser
	} else {
		str = `
		gitConfig: {}`
	}
	str = strings.Replace(str, "\t", "    ", -1)

	return str
}

func GetDeploymentEnvLivenessProbeConfig(livenessProbeConfigRef *deploymentEnv.LivenessProbeConfig) string {
	str := ""

	if nil != livenessProbeConfigRef {
		str = `
		livenessProbeConfig:
		  enabled: ` + strconv.FormatBool(livenessProbeConfigRef.Enabled) + `
		  initialDelaySeconds: ` + helpers.IntToString(livenessProbeConfigRef.InitialDelaySeconds)
	} else {
		str = `
		livenessProbeConfig: {}`
	}
	str = strings.Replace(str, "\t", "    ", -1)

	return str
}

func GetDeploymentEnvHaServices(haServicesRef *deploymentEnv.HaServices) string {
	str := ""

	if nil != haServicesRef {
		str = `
		haServices:` +
			getCloudDriverHa(haServicesRef.Clouddriver) +
			getEchoHa(haServicesRef.Echo)
	} else {
		str = `
		haServices: {}`
	}
	str = strings.Replace(str, "\t", "    ", -1)

	return str
}

func getCloudDriverHa(clouddriverHaRef *deploymentEnv.ClouddriverHA) string {
	str := ""

	if nil != clouddriverHaRef {
		str = `
		  clouddriver:
		    enabled: ` + strconv.FormatBool(clouddriverHaRef.Enabled) + `
		    disableClouddriverRoDeck: ` + strconv.FormatBool(clouddriverHaRef.DisableClouddriverRoDeck)
	} else {
		str = `
		clouddriver: {}`
	}
	str = strings.Replace(str, "\t", "    ", -1)

	return str
}

func getEchoHa(echoHaRef *deploymentEnv.EchoHA) string {
	str := ""

	if nil != echoHaRef {
		str = `
		  echo:
		    enabled: ` + strconv.FormatBool(echoHaRef.Enabled) + `
		    image: ` + echoHaRef.Image +
			strings.Replace(getDeploymentEnArray(echoHaRef.Args, "args"), "  ", " ", -1) +
			strings.Replace(getVolumeMounts(echoHaRef.VolumeMounts), "  ", " ", -1)
	} else {
		str = `
		  echo: {}`
	}
	str = strings.Replace(str, "\t", "    ", -1)

	return str
}

func GetDeploymentEnvAffinity(affinityRef *deploymentEnv.Affinity) string {
	str := ""

	if nil != affinityRef {
		str = `
		affinity:` +
			getServiceServiceAffinity(affinityRef.SpinClouddriver, "spin-clouddriver") +
			getServiceServiceAffinity(affinityRef.SpinClouddriverCaching, "spin-clouddriver-caching") +
			getServiceServiceAffinity(affinityRef.SpinClouddriverRo, "spin-clouddriver-ro") /* TODO Regex error*/ +
			getServiceServiceAffinity(affinityRef.SpinClouddriverRoDeck, "spin-clouddriver-ro-deck") +
			getServiceServiceAffinity(affinityRef.SpinClouddriverRw, "spin-clouddriver-rw") +
			getServiceServiceAffinity(affinityRef.SpinDeck, "spin-deck") +
			getServiceServiceAffinity(affinityRef.SpinEcho, "spin-echo") +
			getServiceServiceAffinity(affinityRef.SpinFiat, "spin-fiat") +
			getServiceServiceAffinity(affinityRef.SpinFront50, "spin-front50") +
			getServiceServiceAffinity(affinityRef.SpinGate, "spin-gate") +
			getServiceServiceAffinity(affinityRef.SpinIgor, "spin-igor") +
			getServiceServiceAffinity(affinityRef.SpinOrca, "spin-orca") +
			getServiceServiceAffinity(affinityRef.SpinRedis, "spin-redis") +
			getServiceServiceAffinity(affinityRef.SpinRosco, "spin-rosco") +
			getServiceServiceAffinity(affinityRef.SpinKayenta, "spin-kayenta") +
			getServiceServiceAffinity(affinityRef.Echo, "echo") +
			getServiceServiceAffinity(affinityRef.Clouddriver, "clouddriver")
	} else {
		str = `
		affinity: {}`
	}
	str = strings.Replace(str, "\t", "    ", -1)

	return str
}

func getServiceServiceAffinity(servicesAffinityRef *deploymentEnv.ServiceAffinity, name string) string {
	str := ""

	if nil != servicesAffinityRef {
		str = `
		` + name + `:`

		if nil != servicesAffinityRef.PodAntiAffinity && nil != servicesAffinityRef.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
			str += `
		  podAntiAffinity:
		    requiredDuringSchedulingIgnoredDuringExecution:` +
				getLabelSelector(servicesAffinityRef.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution)
		} else {
			str += `
		  podAntiAffinity: {}`
		}

		if nil != servicesAffinityRef.NodeAffinity && nil != servicesAffinityRef.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
			str += `
		  nodeAffinity:
		    requiredDuringSchedulingIgnoredDuringExecution:` +
				getNodeSelectorTerms(servicesAffinityRef.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms)
		} else {
			str += `
		  nodeAffinity: {}`
		}
	} else {
		str = `
		` + name + `: {}`
	}

	str = strings.Replace(str, "\t", "     ", -1)

	return str
}

func getLabelSelector(parSidesRef []*deploymentEnv.PAARDSIDE) string {
	str := ""

	if nil != parSidesRef {
		str += `
		      - labelSelector:`
		for _, parSide := range parSidesRef {
			if nil != parSide.LabelSelector && nil != parSide.LabelSelector.MatchExpressions {
				str += `
		          matchExpressions:`
				for _, matchExpression := range parSide.LabelSelector.MatchExpressions {
					if nil != matchExpression {
						str += `
		            - key: ` + matchExpression.Key + `
		              operator: ` + matchExpression.Operator +
							strings.Replace(getDeploymentEnArray(matchExpression.Values, "values"), "\t", "        ", -1)
					}
				}
			}
			str += `
			   topologyKey: ` + parSide.TopologyKey
		}
	} else {
		str += `
		      - labelSelector: []`
	}

	return str
}

func getNodeSelectorTerms(NodeSelectorTermsRef []*deploymentEnv.NodeSelectorTerms) string {
	str := ""

	if nil != NodeSelectorTermsRef {
		for _, NodeSelectorTerm := range NodeSelectorTermsRef {
			str += `
		      nodeSelectorTerms:`
			for _, matchExpression := range NodeSelectorTerm.MatchExpressions {
				str += `
		        - matchExpressions:
		          - key: ` + matchExpression.Key + `
		            operator: ` + matchExpression.Operator +
					strings.Replace(getDeploymentEnArray(matchExpression.Values, "values"), "\t", "      ", -1)
			}
		}
	} else {
		str += `
		      nodeSelectorTerms: []`
	}

	return str
}

func GetDeploymentEnvContainers(containersRef []*deploymentEnv.Containers) string {
	str := ""

	if nil != containersRef {
		str = `
		containers:`
		for _, container := range containersRef {
			str += `
		  - name: ` + container.Name + `
		    image: ` + container.Image +
				strings.Replace(getDeploymentEnArray(container.Command, "command"), "\t", "  ", -1)

			if nil != container.Env {
				str += `
		    env:`
				for _, containerEnv := range container.Env {
					str += `
		      - name: ` + containerEnv.Name + `
		        value: ` + containerEnv.Value
				}
			}

			if nil != container.VolumeMounts {
				str += `
		    volumeMounts:`
				for _, containerVolumeMounts := range container.VolumeMounts {
					str += `
		      - mountPath: ` + containerVolumeMounts.MountPath + `
		        name: ` + containerVolumeMounts.Name
				}
			}
		}
	} else {
		str = `
		containers: {}`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}

func GetDeploymentEnvVolumes(volumesRef []*deploymentEnv.Volumes) string {
	str := ""

	if nil != volumesRef {
		str = `
		volumes:`
		for _, volume := range volumesRef {
			str += `
		  - name: ` + volume.Name

			if nil != volume.EmptyDir && "" != volume.EmptyDir.Test {
				str += `
		    emptyDir: ` + volume.EmptyDir.Test
			} else {
				str += `
		    emptyDir: {}`
			}

			if nil != volume.Secret {
				str += `
		    secret:
		      secretName: ` + volume.Secret.SecretName + `
		      defaultMode: ` + strconv.FormatInt(volume.Secret.DefaultMode, 10)
			} else {
				str += `
		    secret: {}`
			}
		}
	} else {
		str = `
		volumes: {}`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}

func getDeploymentEnArray(stringArray []string, fieldName string) string {
	str := ""

	if nil != stringArray {
		str += `
		        ` + fieldName + `:`
		for _, stringValue := range stringArray {
			str += `
		          - ` + stringValue
		}
	} else {
		str += `
		        ` + fieldName + `: []`
	}

	return str
}

func getVolumeMounts(mounts []*deploymentEnv.VolumeMounts) string {
	str := ""

	if nil != mounts {
		str += `
		        volumeMounts:`
		for _, mount := range mounts {
			str += `
		          - mountPath: ` + mount.MountPath + `
		            name: ` + mount.Name
		}
	} else {
		str += `
		        volumeMounts: []`
	}

	return str
}
