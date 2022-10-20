package service_settings_patch

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetServiceSettings(KustomizeData structs.Kustomize) string {

	return GetClouddriverServiceSettings(KustomizeData) +
		GetDeckServiceSettings(KustomizeData) +
		GetEchoServiceSettings(KustomizeData) +
		GetFiatServiceSettings(KustomizeData) +
		GetFront50ServiceSettings(KustomizeData) +
		GetGateServiceSettings(KustomizeData) +
		GetIgorServiceSettings(KustomizeData) +
		GetKayentaServiceSettings(KustomizeData) +
		GetOrcaServiceSettings(KustomizeData) +
		GetRoscoServiceSettings(KustomizeData)
}

func GetClouddriverServiceSettings(KustomizeData structs.Kustomize) string {
	str := `
	clouddriver: {}`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Clouddriver*/

	str = strings.Replace(str, "\t", "      ", -1)
	return str
}

func GetDeckServiceSettings(KustomizeData structs.Kustomize) string {
	str := `
	deck: {}`

	str = strings.Replace(str, "\t", "      ", -1)
	return str
}

func GetEchoServiceSettings(KustomizeData structs.Kustomize) string {
	str := `
	echo: {}`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Echo*/

	str = strings.Replace(str, "\t", "      ", -1)
	return str
}

func GetFiatServiceSettings(KustomizeData structs.Kustomize) string {
	str := `
	fiat: {}`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Fiat*/

	str = strings.Replace(str, "\t", "      ", -1)
	return str
}

func GetFront50ServiceSettings(KustomizeData structs.Kustomize) string {
	str := `
	front50: {}`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Front50*/

	str = strings.Replace(str, "\t", "      ", -1)
	return str
}

func GetGateServiceSettings(KustomizeData structs.Kustomize) string {
	str := `
	gate: {}`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Gate*/

	str = strings.Replace(str, "\t", "      ", -1)
	return str
}

func GetIgorServiceSettings(KustomizeData structs.Kustomize) string {
	str := `
	igor: {}`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Igor*/

	str = strings.Replace(str, "\t", "      ", -1)
	return str
}

func GetKayentaServiceSettings(KustomizeData structs.Kustomize) string {
	str := `
	kayenta: {}`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Kayenta*/

	str = strings.Replace(str, "\t", "      ", -1)
	return str
}

func GetOrcaServiceSettings(KustomizeData structs.Kustomize) string {
	str := `
	orca: {}`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Orca*/

	str = strings.Replace(str, "\t", "      ", -1)
	return str
}

func GetRoscoServiceSettings(KustomizeData structs.Kustomize) string {
	str := `
	rosco: {}`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Rosco*/

	str = strings.Replace(str, "\t", "      ", -1)
	return str
}
