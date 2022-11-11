package config_patch

import (
	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/spinnaker"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetSpinnaker(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Spinnaker {
		str = GetSpinnakerExtensibility(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Spinnaker)
	}
	return str
}

func GetSpinnakerExtensibility(spinnakerReference *spinnaker.Spinnaker) string {
	str := ""

	// TODO AAAA Issue with unMash
	// fmt.Printf("\nHalyard: %#v \n\n", spinnakerReference.Extensibility.Repositories)
	// fmt.Printf("\nHalyard: %#v \n\n", spinnakerReference.Extensibility.Plugins)

	if nil != spinnakerReference.Extensibility {
		str += `
		extensibility:`

		if nil != spinnakerReference.Extensibility.Plugins {
			str += GetSpinnakerExtensibilityPlugins(spinnakerReference.Extensibility)
		} else {
			str += `
		  plugins: {}`
		}

		if nil != spinnakerReference.Extensibility.Repositories /*&& "" != spinnakerReference.Extensibility.Repositories.Name*/ {

			i := 0
			str += `
		  repositories:`
			for key := range spinnakerReference.Extensibility.Repositories {
				str += `
			`
				str += key + `:` +
					helpers.PrintFmtStr(`id: `, spinnakerReference.Extensibility.Repositories[key].Id, 7, true) +
					helpers.PrintFmtStr(`url: `, spinnakerReference.Extensibility.Repositories[key].Url, 7, true)

				i++

			}
			// ` + spinnakerReference.Extensibility.Repositories + `:` +
			// helpers.PrintFmtStr(`id: `, spinnakerReference.Extensibility.Repositories.Id, 7, true) +
			// helpers.PrintFmtStr(`url: `, spinnakerReference.Extensibility.Repositories.Url, 7, true)
		} else {
			str += `
		  repositories: {}`
		}
	}

	return str
}

func GetSpinnakerExtensibilityPlugins(pluginsReference *spinnaker.Extensibility) string {
	str := ""
	if nil != pluginsReference.Plugins {
		for key, plugin := range pluginsReference.Plugins {
			str += `
		  plugins:
		    ` + key + `:` +
				helpers.PrintFmtStr(`id: `, plugin.Id, 7, true) +
				helpers.PrintFmtBool(`enabled: `, plugin.Enabled, 7, true) +
				helpers.PrintFmtStr(`version: `, plugin.Version, 7, true)

			configCount := 0
			if len(plugin.Config) != configCount {
				str += `
			  config:
				`
				for configKey := range plugin.Config {
					str += configKey + `: ` + plugin.Config[configKey]
					if len(plugin.Config) != configCount+1 {
						str += `
				`
					}
					configCount++
				}
			} else {
				str += `
			  config: {}`
			}
			// GetSpinnakerPluginConfig(plugin.Config)
			// GetSpinnakerPluginExtensions(plugin.Extensions)
			extCount := 0
			if len(plugin.Extensions) != extCount {
				str += `
			  extensions:`
				for extKey := range plugin.Extensions {
					str += helpers.PrintFmtStr(``, extKey, 8, true) + `:
				  id: ` + plugin.Extensions[extKey].Id + helpers.PrintFmtBool(`enabled: `, plugin.Extensions[extKey].Enabled, 9, true) + `
				`
					extConfigCount := 0
					if len(plugin.Extensions[extKey].Config) != extConfigCount {
						str += `  config:
					`
						for extConfigKey := range plugin.Extensions[extKey].Config {
							if extConfigCount != 0 {
								str += `	`
							}
							str += extConfigKey + `: ` + plugin.Extensions[extKey].Config[extConfigKey]
							if len(plugin.Extensions[extKey].Config) != extConfigCount+1 {
								str += `
				`
							}
							extConfigCount++
						}

						if len(plugin.Extensions) != extCount+1 {
							str += `
					`
						}
						extCount++
					} else {
						str += `config: {}`
					}

				}
			} else {
				str += `
				extensions: {}`
			}
		}
	} else {
		str += `
		  plugins: {}`
	}

	return str
}

// func GetSpinnakerPluginConfig(pluginConfigRef *spinnaker.Config) string {
// 	str := ""

// 	if nil != pluginConfigRef {
// 		str += `	//TODO Missing proto
// 		      config: ` /*
// 		   enabled: ` + strconv.FormatBool(plugin.Enabled) + `
// 		   version: ` + plugin.Version*/
// 	} else {
// 		str += `
// 		    config: {}`
// 	}

// 	return str
// }

// func GetSpinnakerPluginExtensions(pluginExtensionRef []*spinnaker.Extensions) string {
// 	str := ""

// 	if nil != pluginExtensionRef {
// 		for _, plugin := range pluginExtensionRef {
// 			str += `
// 		    ` + plugin.Name + `:` +
// 				helpers.PrintFmtStr(`id: `, plugin.Id, 7, true) +
// 				helpers.PrintFmtBool(`enabled: `, plugin.Enabled, 7, true)
// 			// strings.Replace(GetSpinnakerPluginConfig(plugin.Config), "\t", "     ", -1)
// 		}
// 	} else {
// 		str += `
// 		    extensions: {}`
// 	}

// 	return str
// }
