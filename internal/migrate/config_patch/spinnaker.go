package config_patch

import (
	"strings"

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
			str += GetSpinnakerExtensibilityPlugins(spinnakerReference.Extensibility.Plugins)
		} else {
			str += `
		  plugins: {}`
		}

		if nil != spinnakerReference.Extensibility.Repositories /*&& "" != spinnakerReference.Extensibility.Repositories.Name*/ {
			str += `
		  repositories:
		    ` + spinnakerReference.Extensibility.Repositories.Name + `:` +
				helpers.PrintFmtStr(`id: `, spinnakerReference.Extensibility.Repositories.Id, 7, true) +
				helpers.PrintFmtStr(`url: `, spinnakerReference.Extensibility.Repositories.Url, 7, true)
		} else {
			str += `
		  repositories: {}`
		}
	}

	return str
}

func GetSpinnakerExtensibilityPlugins(pluginsReference *spinnaker.Test) string {
	str := ""

	if nil != pluginsReference.Plugins {
		for key, plugin := range pluginsReference.Plugins {
			str += `
		  plugins:
		    ` + key + `:` +
				helpers.PrintFmtStr(`id: `, plugin.Id, 7, true) +
				helpers.PrintFmtBool(`enabled: `, plugin.Enabled, 7, true) +
				helpers.PrintFmtStr(`version: `, plugin.Version, 7, true) +
				GetSpinnakerPluginConfig(plugin.Config) +
				GetSpinnakerPluginExtensions(plugin.Extensions)
		}
	} else {
		str += `
		  plugins: {}`
	}

	return str
}

func GetSpinnakerPluginConfig(pluginConfigRef *spinnaker.Config) string {
	str := ""

	if nil != pluginConfigRef {
		str += `	//TODO Missing proto
		      config: ` /*
		   enabled: ` + strconv.FormatBool(plugin.Enabled) + `
		   version: ` + plugin.Version*/
	} else {
		str += `
		    config: {}`
	}

	return str
}

func GetSpinnakerPluginExtensions(pluginExtensionRef []*spinnaker.Extensions) string {
	str := ""

	if nil != pluginExtensionRef {
		for _, plugin := range pluginExtensionRef {
			str += `
		    ` + plugin.Name + `:` +
				helpers.PrintFmtStr(`id: `, plugin.Id, 7, true) +
				helpers.PrintFmtBool(`enabled: `, plugin.Enabled, 7, true) +
				strings.Replace(GetSpinnakerPluginConfig(plugin.Config), "\t", "     ", -1)
		}
	} else {
		str += `
		    extensions: {}`
	}

	return str
}
