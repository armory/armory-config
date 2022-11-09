package config_patch

import (
	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/features"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
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

	str = helpers.PrintFmtBool(`auth: `, featuresReference.Auth, 4, true) +
		helpers.PrintFmtBool(`fiat: `, featuresReference.Auth, 4, true) +
		helpers.PrintFmtBool(`chaos: `, featuresReference.Chaos, 4, true) +
		helpers.PrintFmtBool(`entityTags: `, featuresReference.EntityTags, 4, true) +
		helpers.PrintFmtBool(`pipelineTemplates: `, featuresReference.PipelineTemplates, 4, true) +
		helpers.PrintFmtBool(`artifacts: `, featuresReference.Artifacts, 4, true) +
		helpers.PrintFmtBool(`appengineContainerImageUrlDeployments: `, featuresReference.AppengineContainerImageUrlDeployments, 4, true) +
		helpers.PrintFmtBool(`managedPipelineTemplatesV2UI: `, featuresReference.ManagedPipelineTemplatesV2UI, 4, true) +
		helpers.PrintFmtBool(`gremlin: `, featuresReference.Gremlin, 4, true)

	return str
}
