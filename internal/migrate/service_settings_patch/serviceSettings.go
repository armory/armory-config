package service_settings_patch

import (
	"regexp"
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
# === Clouddriver Config ===
			clouddriver: {} # is the contents of ~/.hal/default/service-settings/clouddriver.yml`

	if value, isMapContainsKey := KustomizeData.ServiceSettingsConfigFiles["clouddriver.yml"]; isMapContainsKey {
		str = `
# === Clouddriver Config ===
		clouddriver:
			` + value
	}

	return ProfilesFormatStr(str)
}

func GetDeckServiceSettings(KustomizeData structs.Kustomize) string {
	str := `
# === Deck Config ===
			deck: {} # is the contents of ~/.hal/default/service-settings/deck.yml`

	if value, isMapContainsKey := KustomizeData.ServiceSettingsConfigFiles["deck.yml"]; isMapContainsKey {
		str = `
# === Deck Config ===
		deck:
			` + value
	}

	return ProfilesFormatStr(str)
}

func GetEchoServiceSettings(KustomizeData structs.Kustomize) string {
	str := `
# === Echo Config ===
			echo: {} # is the contents of ~/.hal/default/service-settings/echo.yml`

	if value, isMapContainsKey := KustomizeData.ServiceSettingsConfigFiles["echo.yml"]; isMapContainsKey {
		str = `
# === Echo Config ===
		echo:
			` + value
	}

	return ProfilesFormatStr(str)
}

func GetFiatServiceSettings(KustomizeData structs.Kustomize) string {
	str := `
# === Fiat Config ===
			fiat: {} # is the contents of ~/.hal/default/service-settings/fiat.yml`

	if value, isMapContainsKey := KustomizeData.ServiceSettingsConfigFiles["fiat.yml"]; isMapContainsKey {
		str = `
# === Fiat Config ===
		fiat:
			` + value
	}

	return ProfilesFormatStr(str)
}

func GetFront50ServiceSettings(KustomizeData structs.Kustomize) string {
	str := `
# === Front50 Config ===
			front50: {} # is the contents of ~/.hal/default/service-settings/front50.yml`

	if value, isMapContainsKey := KustomizeData.ServiceSettingsConfigFiles["front50.yml"]; isMapContainsKey {
		str = `
# === Front50 Config ===
		front50:
			` + value
	}

	return ProfilesFormatStr(str)
}

func GetGateServiceSettings(KustomizeData structs.Kustomize) string {
	str := `
# === Gate Config ===
			gate: {} # is the contents of ~/.hal/default/service-settings/gate.yml`

	if value, isMapContainsKey := KustomizeData.ServiceSettingsConfigFiles["gate.yml"]; isMapContainsKey {
		str = `
# === Gate Config ===
		gate:
			` + value
	}

	return ProfilesFormatStr(str)
}

func GetIgorServiceSettings(KustomizeData structs.Kustomize) string {
	str := `
# === Igor Config ===
			igor: {} # is the contents of ~/.hal/default/service-settings/igor.yml`

	if value, isMapContainsKey := KustomizeData.ServiceSettingsConfigFiles["igor.yml"]; isMapContainsKey {
		str = `
# === Igor Config ===
		igor:
			` + value
	}

	return ProfilesFormatStr(str)
}

func GetKayentaServiceSettings(KustomizeData structs.Kustomize) string {
	str := `
# === Kayenta Config ===
			kayenta: {} # is the contents of ~/.hal/default/service-settings/kayenta.yml`

	if value, isMapContainsKey := KustomizeData.ServiceSettingsConfigFiles["kayenta.yml"]; isMapContainsKey {
		str = `
# === Kayenta Config ===
		kayenta:
			` + value
	}

	return ProfilesFormatStr(str)
}

func GetOrcaServiceSettings(KustomizeData structs.Kustomize) string {
	str := `
# === Orca Config ===
			orca: {} # is the contents of ~/.hal/default/service-settings/orca.yml`

	if value, isMapContainsKey := KustomizeData.ServiceSettingsConfigFiles["orca.yml"]; isMapContainsKey {
		str = `
# === Orca Config ===
		orca:
			` + value
	}

	return ProfilesFormatStr(str)
}

func GetRoscoServiceSettings(KustomizeData structs.Kustomize) string {
	str := `
# === Rosco Config ===
			rosco: {} # is the contents of ~/.hal/default/service-settings/rosco.yml`

	if value, isMapContainsKey := KustomizeData.ServiceSettingsConfigFiles["rosco.yml"]; isMapContainsKey {
		str = `
# === Rosco Config ===
		rosco:
			` + value
	}

	return ProfilesFormatStr(str)
}

// Function that edit the file to have the correct spacing
func ProfilesFormatStr(str string) string {
	//Every newline and 2 spaces we add 10 spaces
	reg := regexp.MustCompile(`(\n)  `)
	str = reg.ReplaceAllString(str, `$1          `)

	//Every newline and a word we add 8 spaces
	reg = regexp.MustCompile(`(\n)([a-zA-Z]+)`)
	str = reg.ReplaceAllString(str, `$1        $2`)

	//Change tabs to 4 spaces
	str = strings.Replace(str, "\t\t", "      ", -1)
	str = strings.Replace(str, "\t", "  ", -1)

	return str
}
