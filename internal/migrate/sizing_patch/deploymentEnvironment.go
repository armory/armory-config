package sizing_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/deploymentEnv"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetDeploymentEnvironment(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment.CustomSizing {
		str = GetCustomSizing(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment.CustomSizing)
	}

	return str
}

func GetCustomSizing(sizingReference *deploymentEnv.CustomSizing) string {

	str := `
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
		getServiceSizing(sizingReference.SpinClouddriverCaching, "spin-clouddriver-caching") +
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
		getServiceSizing(sizingReference.SpinFront50, "spin-front50")

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}

func getServiceSizing(serviceSizingReference *deploymentEnv.ServiceSizing, name string) string {
	str := ""

	if nil != serviceSizingReference {
		str = `
		  ` + name + `:` + `
		    replicas: ` + strconv.FormatInt(int64(serviceSizingReference.Replicas), 10) +
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
		      cpu: ` + strconv.FormatInt(int64(sizingReference.Cpu), 10) + `
		      memory: ` + sizingReference.Memory
	} else {
		str = `
		  ` + name + `: {}`
	}

	return str
}
