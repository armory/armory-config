package config_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/spinnaker"
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
		    ` + spinnakerReference.Extensibility.Repositories.Name + `:
		      id:` + spinnakerReference.Extensibility.Repositories.Id + `
		      url:` + spinnakerReference.Extensibility.Repositories.Url
		} else {
			str += `
		  repositories: {}`
		}

		str = strings.Replace(str, "\t", "    ", -1)
	}

	return str
}

func GetSpinnakerExtensibilityPlugins(pluginsReference *spinnaker.Test) string {
	str := ""

	if nil != pluginsReference.Plugins {
		for _, plugin := range pluginsReference.Plugins {
			str += `
		  plugins:
		    PLUGIN` + /*TODO plugin.name + */ `
		      id: ` + plugin.Id + `
		      enabled: ` + strconv.FormatBool(plugin.Enabled) + `
		      version: ` + plugin.Version +
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
		    ` + plugin.Name + `:
		      id: ` + plugin.Id + `
		      enabled: ` + strconv.FormatBool(plugin.Enabled) +
				strings.Replace(GetSpinnakerPluginConfig(plugin.Config), "\t", "     ", -1)
		}
	} else {
		str += `
		    extensions: {}`
	}

	return str
}
