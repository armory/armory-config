package config_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/features"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetFeatures(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Features {
		str = GetFeaturesData(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Features)
	}
	return str
}

func GetFeaturesData(featuresReference *features.Features) string {
	str := ""

	str = `
		auth: ` + strconv.FormatBool(featuresReference.Auth) + `
		fiat: ` + strconv.FormatBool(featuresReference.Fiat) + `
		chaos: ` + strconv.FormatBool(featuresReference.Chaos) + `
		entityTags: ` + strconv.FormatBool(featuresReference.EntityTags) + `
		pipelineTemplates: ` + strconv.FormatBool(featuresReference.PipelineTemplates) + `
		artifacts: ` + strconv.FormatBool(featuresReference.Artifacts) + `
		appengineContainerImageUrlDeployments: ` + strconv.FormatBool(featuresReference.AppengineContainerImageUrlDeployments) + `
		managedPipelineTemplatesV2UI: ` + strconv.FormatBool(featuresReference.ManagedPipelineTemplatesV2UI) + `
		gremlin: ` + strconv.FormatBool(featuresReference.Gremlin)

	str = strings.Replace(str, "\t", "    ", -1)

	return str
}
