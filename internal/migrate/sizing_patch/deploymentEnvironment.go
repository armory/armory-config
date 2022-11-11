package sizing_patch

import (
	"regexp"
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

// func IsInitContainerEmpty(KustomizeData structs.Kustomize) bool {
// 	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment.InitContainers {
// 		return false
// 	}
// 	return true
// }

func GetDeploymentEnvironment(KustomizeData structs.Kustomize) string {
	str := ""

	if !IsDeploymentEnvironmentEmpty(KustomizeData) {
		str = GetDeploymentEnvironmentData(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment) +
			getDeploymentEnvConsul(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment.Consul) +
			getDeploymentEnvVault(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment.Vault) +
			GetDeploymentEnvCustomSizing(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment.CustomSizing) +
			GetDeploymentEnvSidecars(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment.Sidecars) +
			// GetDeploymentEnvInitContainers(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment.InitContainers) +
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

	str += helpers.PrintFmtStr(`size: `, deploymentEnvRef.Size, 4, true) +
		helpers.PrintFmtStr(`type: `, deploymentEnvRef.Type, 4, true) +
		helpers.PrintFmtStr(`accountName: `, deploymentEnvRef.AccountName, 4, true) +
		helpers.PrintFmtStr(`imageVariant: `, deploymentEnvRef.ImageVariant, 4, true) +
		helpers.PrintFmtStr(`bootstrapOnly: `, strconv.FormatBool(deploymentEnvRef.BootstrapOnly), 4, true) +
		helpers.PrintFmtStr(`updateVersions: `, strconv.FormatBool(deploymentEnvRef.UpdateVersions), 4, true) +
		helpers.PrintFmtStr(`location: `, deploymentEnvRef.Location, 4, true)

	return str
}

func getDeploymentEnvConsul(consulReference *deploymentEnv.Consul) string {
	str := ""

	if nil != consulReference {
		str = `
	    consul:` +
			helpers.PrintFmtBool(`enabled: `, consulReference.Enabled, 5, true) +
			helpers.PrintFmtStr(`address: `, consulReference.Address, 5, true)

		str = strings.Replace(str, "\t", "    ", -1)
	}

	return str
}

func getDeploymentEnvVault(vaultReference *deploymentEnv.Vault) string {
	str := ""

	if nil != vaultReference {
		str = `
	    vault:` +
			helpers.PrintFmtBool(`enabled: `, vaultReference.Enabled, 5, true) +
			helpers.PrintFmtStr(`address: `, vaultReference.Address, 5, true)

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
		  ` + name + `:` +
			helpers.PrintFmtInt(`replicas: `, serviceSizingReference.Replicas, 6, true) +
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
		    ` + name + `:` +
			helpers.PrintFmtInt(`cpu: `, sizingReference.Cpu, 7, true) +
			helpers.PrintFmtStr(`memory: `, sizingReference.Memory, 7, true)
	} else {
		str = `
		    ` + name + `: {}`
	}

	return str
}

func GetDeploymentEnvSidecars(sidecarsRef *deploymentEnv.Sidecars) string {
	str := ""

	if nil != sidecarsRef {
		str = `
		sidecars:` +
			getSidecarService(sidecarsRef.SpinKayenta, "spin-kayenta") +
			getSidecarService(sidecarsRef.Clouddriver, "clouddriver") +
			getSidecarService(sidecarsRef.Echo, "echo") +
			getSidecarService(sidecarsRef.Fiat, "fiat") +
			getSidecarService(sidecarsRef.SpinIgor, "spin-igor") +
			getSidecarService(sidecarsRef.SpinGate, "spin-gate") +
			getSidecarService(sidecarsRef.SpinOrca, "spin-orca") //+
		// getSidecarService(sidecarsRef.SpinEcho, "spin-echo") +
		// getSidecarService(sidecarsRef.SpinClouddriver, "spin-clouddriver") +
		// getSidecarService(sidecarsRef.SpinFront50, "spin-front50") +
		// getSidecarService(sidecarsRef.SpinRosco, "spin-rosco") +
		// getSidecarService(sidecarsRef.SpinDeck, "spin-deck") +
		// getSidecarService(sidecarsRef.SpinFiat, "spin-fiat")

	} else {
		str = `
		sidecars: {}`
	}
	str = strings.Replace(str, "\t", "    ", -1)

	return str
}

func getSidecarService(sidecarServicesRef []*deploymentEnv.SideCarService, name string) string {
	str := ""

	if nil != sidecarServicesRef {
		str += `
		  ` + name + `:`
		for _, sidecarService := range sidecarServicesRef {
			str += helpers.PrintFmtStr(`- name: `, sidecarService.Name, 5, true) +
				helpers.PrintFmtStr(`dockerImage: `, sidecarService.DockerImage, 6, true) +
				getDeploymentEnvMap(sidecarService.Env, "env") +
				strings.Replace(getDeploymentEnvArray(sidecarService.Args, "args"), "   ", " ", -1) +
				strings.Replace(getDeploymentEnvArray(sidecarService.Command, "command"), "   ", " ", -1) +
				getSidecarConfigMapVolume(sidecarService.ConfigMapVolumeMounts) +
				strings.Replace(getDeploymentEnvArray(sidecarService.SecretVolumeMounts, "secretVolumeMounts"), "   ", " ", -1) +
				helpers.PrintFmtStr(`mountPath: `, sidecarService.MountPath, 6, true) +
				getSecurityContext(sidecarService.SecurityContext)
		}
	} else {
		str = `
		  ` + name + `: {}`
	}

	return str
}

func getSidecarConfigMapVolume(configMapRef []*deploymentEnv.ConfigMapVolumeMounts) string {
	str := ""

	if nil != configMapRef {
		str = `
			configMapVolumeMounts:`
		for _, configMap := range configMapRef {
			str += helpers.PrintFmtStr(`- configMapName: `, configMap.ConfigMapName, 6, true) +
				helpers.PrintFmtStr(`mountPath: `, configMap.MountPath, 7, true)
		}
	} else {
		str = `
			configMapVolumeMounts: {}`
	}
	str = strings.Replace(str, "\t", "    ", -1)

	return str
}

func getSecurityContext(securityContexRtef *deploymentEnv.SecurityContext) string {
	str := ""

	if nil != securityContexRtef {
		str = `
			securityContext:` +
			helpers.PrintFmtBool(`privileged: `, securityContexRtef.Privileged, 7, true)
	} else {
		str = `
			securityContext: {}`
	}
	str = strings.Replace(str, "\t", "    ", -1)

	return str
}

func GetDeploymentEnvInitContainers(initReference *deploymentEnv.InitContainers) string {
	str := ""

	if nil != initReference.Front50 {
		str += `` +
			getServiceInitContainers(initReference.Front50, "front50")
	}
	if nil != initReference.Clouddriver {
		str += `` +
			getServiceInitContainers(initReference.Clouddriver, "clouddriver")
	}
	if nil != initReference.Deck {
		str += `` +
			getServiceInitContainers(initReference.Deck, "deck")
	}
	if nil != initReference.Gate {
		str += `` +
			getServiceInitContainers(initReference.Gate, "gate")
	}
	if nil != initReference.Igor {
		str += `` +
			getServiceInitContainers(initReference.Igor, "igor")
	}
	if nil != initReference.Orca {
		str += `` +
			getServiceInitContainers(initReference.Orca, "orca")
	}
	if nil != initReference.Deck {
		str += `` +
			getServiceInitContainers(initReference.Rosco, "rosco")
	}
	if nil != initReference.SpinClouddriverRw {
		str += `` +
			getServiceInitContainers(initReference.SpinClouddriverRw, "spin-clouddriver-rw")
	}
	if nil != initReference.SpinOrca {
		str += `` +
			getServiceInitContainers(initReference.SpinOrca, "spin-orca")
	}
	if nil != initReference.SpinKayenta {
		str += `` +
			getServiceInitContainers(initReference.SpinKayenta, "spin-kayenta")
	}
	if nil != initReference.SpinIgor {
		str += `` +
			getServiceInitContainers(initReference.SpinIgor, "spin-igor")
	}
	if nil != initReference.SpinEcho {
		str += `` +
			getServiceInitContainers(initReference.SpinEcho, "spin-echo")
	}
	str = strings.Replace(str, "\t", "    ", -1)

	if "" == str {
		str = "--|||EMPTY|||--"
	}

	return str
}

// func IsInitContainerEmpty(initReference *deploymentEnv.InitContainers) string {
// 	str := ""

// 	if nil != initReference.Front50 {
// 		checkInitContainerEmpty(initReference.Front50, "front50")
// 		checkInitContainerEmpty(initReference.Clouddriver, "clouddriver")
// 		checkInitContainerEmpty(initReference.Orca, "orca")
// 		checkInitContainerEmpty(initReference.Rosco, "rosco")
// 		checkInitContainerEmpty(initReference.Gate, "gate")
// 		checkInitContainerEmpty(initReference.Deck, "deck")
// 		checkInitContainerEmpty(initReference.Igor, "igor")
// 		checkInitContainerEmpty(initReference.SpinOrca, "spin-orca")
// 		checkInitContainerEmpty(initReference.SpinKayenta, "spin-kayenta")
// 		checkInitContainerEmpty(initReference.SpinIgor, "spin-igor")
// 		checkInitContainerEmpty(initReference.SpinEcho, "spin-echo")
// 		checkInitContainerEmpty(initReference.SpinClouddriverRw, "spin-clouddriver-rw")

// 		str = strings.Replace(str, "\t", "    ", -1)
// 	}
// 	return str
// }

// func checkInitContainerEmpty(servicesRef []*deploymentEnv.Service, name string) bool {

// 	return true
// }

func getServiceInitContainers(servicesRef []*deploymentEnv.Service, name string) string {
	str := ""
	if nil != servicesRef {
		str = `
	` + name + `:
	  deployment:
		patchesStrategicMerge:
		  - |
			spec:
			  template:
				spec:
				  initContainers:`
		for _, services := range servicesRef {
			str +=
				helpers.PrintFmtStr(`- name: `, services.Name, 9, true) +
					helpers.PrintFmtStr(`image: `, services.Image, 10, true) +
					strings.Replace(getDeploymentEnvArray(services.Args, "args"), "\t", "      ", -1) +
					strings.Replace(getVolumeMounts(services.VolumeMounts), "\t", "      ", -1)
		}
	} else {
		str = ``
	}

	return str
}

// func getServiceInitContainers(servicesRef []*deploymentEnv.Service, name string) string {
// 	str := ""

// 	if nil != servicesRef {
// 		for _, services := range servicesRef {
// 			str = `
// 			` + name + `:
// 			deployment:
// 			patchesStrategicMerge:
// 				- |
// 		spec:
// 			template:
// 			spec:
// 				initContainers:` +
// 				helpers.PrintFmtStr(`- name: `, services.Name, 7, true) +
// 				helpers.PrintFmtStr(`image: `, services.Image, 8, true) +
// 				getDeploymentEnvArray(services.Args, "args") +
// 				getVolumeMounts(services.VolumeMounts)
// 		}
// 	} else {
// 		str = `
// 		    ` + name + `: {}`
// 	}

// 	return str
// }

func GetDeploymentEnvHostAliases(hostAliasesRef *deploymentEnv.HostAliases) string {
	str := ""

	if nil != hostAliasesRef {
		//TODO Missing proto
		str = `
		hostAliases:` +
			getHostAliasService(hostAliasesRef.SpinIgor, "spin-igor") +
			getHostAliasService(hostAliasesRef.SpinKayenta, "spin-kayenta") +
			getHostAliasService(hostAliasesRef.SpinOrca, "spin-orca") +
			getHostAliasService(hostAliasesRef.SpinEcho, "spin-echo") +
			getHostAliasService(hostAliasesRef.SpinClouddriver, "spin-clouddriver") +
			getHostAliasService(hostAliasesRef.SpinFront50, "spin-front50") +
			getHostAliasService(hostAliasesRef.SpinRosco, "spin-rosco") +
			getHostAliasService(hostAliasesRef.SpinDeck, "spin-deck") +
			getHostAliasService(hostAliasesRef.SpinGate, "spin-gate") +
			getHostAliasService(hostAliasesRef.SpinFiat, "spin-fiat")

	} else {
		str = `
		hostAliases: {}`
	}
	str = strings.Replace(str, "\t", "    ", -1)

	return str
}

func getHostAliasService(hostAliasesRef []*deploymentEnv.HostAliasService, name string) string {
	str := ""

	if nil != hostAliasesRef {
		for _, hostAlias := range hostAliasesRef {
			str = `
		  ` + name + `:` +
				helpers.PrintFmtStr(`- ip: `, hostAlias.Ip, 5, true) +
				strings.Replace(getDeploymentEnvArray(hostAlias.Hostnames, "hostnames"), "   ", " ", -1)
		}
	} else {
		str = `
		  ` + name + `: {}`
	}

	return str
}

func GetDeploymentEnvTolerations(tolerationsRef *deploymentEnv.Tolerations) string {
	str := ""

	if nil != tolerationsRef && "" != tolerationsRef.Key {
		str = `
		tolerations:` +
			helpers.PrintFmtStr(`key: `, tolerationsRef.Key, 5, true) +
			helpers.PrintFmtStr(`operator: `, tolerationsRef.Operator, 5, true) +
			helpers.PrintFmtStr(`value: `, tolerationsRef.Value, 5, true) +
			helpers.PrintFmtStr(`effect: `, tolerationsRef.Effect, 5, true)
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
		gitConfig:` +
			helpers.PrintFmtStr(`upstreamUser: `, gitConfigRef.UpstreamUser, 5, true) +
			helpers.PrintFmtStr(`originUser: `, gitConfigRef.OriginUser, 5, true)
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
		livenessProbeConfig:` +
			helpers.PrintFmtBool(`enabled: `, livenessProbeConfigRef.Enabled, 5, true) +
			helpers.PrintFmtInt(`initialDelaySeconds: `, livenessProbeConfigRef.InitialDelaySeconds, 5, true)
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
		  clouddriver:` +
			helpers.PrintFmtBool(`enabled: `, clouddriverHaRef.Enabled, 6, true) +
			helpers.PrintFmtBool(`disableClouddriverRoDeck: `, clouddriverHaRef.DisableClouddriverRoDeck, 6, true)
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
		  echo:` +
			helpers.PrintFmtBool(`enabled: `, echoHaRef.Enabled, 6, true)
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
						str += helpers.PrintFmtStr(`- key: `, matchExpression.Key, 11, true) +
							helpers.PrintFmtStr(`operator: `, matchExpression.Operator, 12, true) +
							strings.Replace(getDeploymentEnvArray(matchExpression.Values, "values"), "\t", "        ", -1)
					}
				}
			}
			str += helpers.PrintFmtStr(`topologyKey: `, parSide.TopologyKey, 9, true)
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
		        - matchExpressions:` +
					helpers.PrintFmtStr(`- key: `, matchExpression.Key, 10, true) +
					helpers.PrintFmtStr(`operator: `, matchExpression.Operator, 11, true) +
					strings.Replace(getDeploymentEnvArray(matchExpression.Values, "values"), "\t", "      ", -1)
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
			str += helpers.PrintFmtStr(`- name: `, container.Name, 5, true) +
				helpers.PrintFmtStr(`image: `, container.Image, 6, true) +
				strings.Replace(getDeploymentEnvArray(container.Command, "command"), "\t", "  ", -1)

			if nil != container.Env {
				str += `
		    env:`
				for _, containerEnv := range container.Env {
					str += helpers.PrintFmtStr(`- name: `, containerEnv.Name, 7, true) +
						helpers.PrintFmtStr(`value: `, containerEnv.Value, 8, true)
				}
			}

			if nil != container.VolumeMounts {
				str += `
		    volumeMounts:`
				for _, containerVolumeMounts := range container.VolumeMounts {
					str += helpers.PrintFmtStr(`- mountPath: `, containerVolumeMounts.MountPath, 7, true) +
						helpers.PrintFmtStr(`name: `, containerVolumeMounts.Name, 8, true)
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
			str += helpers.PrintFmtStr(`- name: `, volume.Name, 5, true)

			if nil != volume.EmptyDir && "" != volume.EmptyDir.Test {
				str += `
		    emptyDir: ` + volume.EmptyDir.Test
			} else {
				str += `
		    emptyDir: {}`
			}

			if nil != volume.Secret {
				str += `
		    secret:` +
					helpers.PrintFmtStr(`secretName: `, volume.Secret.SecretName, 7, true) +
					helpers.PrintFmtInt64(`defaultMode: `, volume.Secret.DefaultMode, 7, true)
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

func getDeploymentEnvArray(stringArray []string, fieldName string) string {
	str := ""

	if nil != stringArray {
		str += `
		        ` + fieldName + `:`
		for _, stringValue := range stringArray {

			str += `
		          - ` + checkGetMultiline(stringValue, 10)
		}
	} else {
		str += `
		        ` + fieldName + `: []`
	}

	return str
}

func getDeploymentEnvMap(stringMap map[string]string, fieldName string) string {
	str := ""
	str2 := ""

	if nil != stringMap {

		for key, value := range stringMap {
			str2 += helpers.PrintFmtStr(key+": ", value, 7, true)
		}
	}

	if "" != str2 {
		str += `
		    ` + fieldName + `:` +
			str2
	} else {
		str += `
		    ` + fieldName + `: {}`
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

func checkGetMultiline(str string, amountTabs int) string {
	resp := ""

	reg := regexp.MustCompile(`(.*)(\n)`)
	resp = reg.ReplaceAllString(str, `|-$2`+helpers.GetSpaces(amountTabs*2)+`$1$2`+helpers.GetSpaces(amountTabs*2))

	return resp
}
