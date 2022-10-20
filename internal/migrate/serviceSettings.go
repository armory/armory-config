package migrate

import "strings"

func GetServiceSettings(KustomizeData Kustomize) string {

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

func GetClouddriverServiceSettings(KustomizeData Kustomize) string {
	str := `
	clouddriver: {}`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Clouddriver*/

	str = strings.Replace(str, "\t", "      ", -1)
	return str
}

func GetDeckServiceSettings(KustomizeData Kustomize) string {
	str := `
	deck: {}`

	str = strings.Replace(str, "\t", "      ", -1)
	return str
}

func GetEchoServiceSettings(KustomizeData Kustomize) string {
	str := `
	echo: {}`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Echo*/

	str = strings.Replace(str, "\t", "      ", -1)
	return str
}

func GetFiatServiceSettings(KustomizeData Kustomize) string {
	str := `
	fiat: {}`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Fiat*/

	str = strings.Replace(str, "\t", "      ", -1)
	return str
}

func GetFront50ServiceSettings(KustomizeData Kustomize) string {
	str := `
	front50: {}`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Front50*/

	str = strings.Replace(str, "\t", "      ", -1)
	return str
}

func GetGateServiceSettings(KustomizeData Kustomize) string {
	str := `
	gate: {}`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Gate*/

	str = strings.Replace(str, "\t", "      ", -1)
	return str
}

func GetIgorServiceSettings(KustomizeData Kustomize) string {
	str := `
	igor: {}`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Igor*/

	str = strings.Replace(str, "\t", "      ", -1)
	return str
}

func GetKayentaServiceSettings(KustomizeData Kustomize) string {
	str := `
	kayenta: {}`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Kayenta*/

	str = strings.Replace(str, "\t", "      ", -1)
	return str
}

func GetOrcaServiceSettings(KustomizeData Kustomize) string {
	str := `
	orca: {}`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Orca*/

	str = strings.Replace(str, "\t", "      ", -1)
	return str
}

func GetRoscoServiceSettings(KustomizeData Kustomize) string {
	str := `
	rosco: {}`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Rosco*/

	str = strings.Replace(str, "\t", "      ", -1)
	return str
}
