package profile_patch

import (
	"regexp"
	"strings"

	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetProfiles(KustomizeData structs.Kustomize) string {

	return GetClouddriverProfile(KustomizeData) +
		GetDeckProfile(KustomizeData) +
		GetEchoProfile(KustomizeData) +
		GetFiatProfile(KustomizeData) +
		GetFront50Profile(KustomizeData) +
		GetGateProfile(KustomizeData) +
		GetIgorProfile(KustomizeData) +
		GetKayentaProfile(KustomizeData) +
		GetOrcaProfile(KustomizeData) +
		GetRoscoProfile(KustomizeData)
}

func GetClouddriverProfile(KustomizeData structs.Kustomize) string {
	str := `
# === Clouddriver Config ===
		clouddriver: {} # is the contents of ~/.hal/default/profiles/clouddriver.yml`

	if value, isMapContainsKey := KustomizeData.ProfilesConfigFiles["clouddriver-local.yml"]; isMapContainsKey {
		str = `
# === Clouddriver Config ===
		clouddriver:
			` + value
	}

	return ProfilesFormatStr(str)
}

func GetDeckProfile(KustomizeData structs.Kustomize) string {
	str := `
# === Deck Config ===
		# deck has a special key "settings-local.js" for the contents of settings-local.js
		deck:
			# settings-local.js - contents of ~/.hal/default/profiles/settings-local.js
			# Use the | YAML symbol to indicate a block-style multiline string
			settings-local.js: |
				window.spinnakerSettings.feature.kustomizeEnabled = true;
				window.spinnakerSettings.feature.artifactsRewrite = true;`

	if value, isMapContainsKey := KustomizeData.ProfilesConfigFiles["settings-local.js"]; isMapContainsKey {
		str = `
# === Deck Config ===
		deck:
			settings-local.js: |
				` + value
	}

	return DeckProfilesFormatStr(str)
}

func GetEchoProfile(KustomizeData structs.Kustomize) string {
	str := `
# === Echo Config ===
		echo: {}    # is the contents of ~/.hal/default/profiles/echo.yml`

	if value, isMapContainsKey := KustomizeData.ProfilesConfigFiles["echo-local.yml"]; isMapContainsKey {
		str = `
# === Echo Config ===
		echo:
			` + value
	}

	return ProfilesFormatStr(str)
}

func GetFiatProfile(KustomizeData structs.Kustomize) string {
	str := `
# === Fiat Config ===
		fiat: {}    # is the contents of ~/.hal/default/profiles/fiat.yml`

	if value, isMapContainsKey := KustomizeData.ProfilesConfigFiles["fiat-local.yml"]; isMapContainsKey {
		str = `
# === Fiat Config ===
		fiat:
			` + value
	}

	return ProfilesFormatStr(str)
}

func GetFront50Profile(KustomizeData structs.Kustomize) string {
	str := `
# === Front50 Config ===
		front50: {} # is the contents of ~/.hal/default/profiles/front50.yml`

	if value, isMapContainsKey := KustomizeData.ProfilesConfigFiles["front50-local.yml"]; isMapContainsKey {
		str = `
# === Front50 Config ===
		front50:
			` + value
	}

	return ProfilesFormatStr(str)
}

func GetGateProfile(KustomizeData structs.Kustomize) string {
	str := `
# === Gate Config ===
		gate: {}    # is the contents of ~/.hal/default/profiles/gate.yml`

	if value, isMapContainsKey := KustomizeData.ProfilesConfigFiles["gate-local.yml"]; isMapContainsKey {
		str = `
# === Gate Config ===
		gate:
			` + value
	}

	return ProfilesFormatStr(str)
}

func GetIgorProfile(KustomizeData structs.Kustomize) string {
	str := `
# === Igor Config ===
		igor: {}    # is the contents of ~/.hal/default/profiles/igor.yml`

	if value, isMapContainsKey := KustomizeData.ProfilesConfigFiles["igor-local.yml"]; isMapContainsKey {
		str = `
# === Igor Config ===
		igor:
			` + value
	}

	return ProfilesFormatStr(str)
}

func GetKayentaProfile(KustomizeData structs.Kustomize) string {
	str := `
# === Kayenta Config ===
		kayenta: {} # is the contents of ~/.hal/default/profiles/kayenta.yml`

	if value, isMapContainsKey := KustomizeData.ProfilesConfigFiles["kayenta-local.yml"]; isMapContainsKey {
		str = `
# === Kayenta Config ===
		kayenta:
			` + value
	}

	return ProfilesFormatStr(str)
}

func GetOrcaProfile(KustomizeData structs.Kustomize) string {
	str := `
# === Orca Config ===
		orca: {}    # is the contents of ~/.hal/default/profiles/orca.yml`

	if value, isMapContainsKey := KustomizeData.ProfilesConfigFiles["orca-local.yml"]; isMapContainsKey {
		str = `
# === Orca Config ===
		orca:
			` + value
	}

	return ProfilesFormatStr(str)
}

func GetRoscoProfile(KustomizeData structs.Kustomize) string {
	str := `
# === Rosco Config ===
		rosco: {}   # is the contents of ~/.hal/default/profiles/rosco.yml`

	if value, isMapContainsKey := KustomizeData.ProfilesConfigFiles["rosco-local.yml"]; isMapContainsKey {
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

// Function that edit the file to have the correct spacing for Deck
func DeckProfilesFormatStr(str string) string {
	//Every newline and 2 spaces we add 10 spaces
	reg := regexp.MustCompile(`(\n)  `)
	str = reg.ReplaceAllString(str, `$1              `)
	//Every newline and } change
	reg = regexp.MustCompile(`(\n)(}.*)`)
	str = reg.ReplaceAllString(str, `$1            $2`)

	//Change tabs to 4 spaces
	str = strings.Replace(str, "\t\t", "      ", -1)
	str = strings.Replace(str, "\t", "  ", -1)

	//Every newline and a word we add 8 spaces
	reg = regexp.MustCompile(`(\n)([a-zA-Z]+)`)
	str = reg.ReplaceAllString(str, `$1            $2`)

	return str
}
