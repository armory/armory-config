package structs

import "github.com/austinthao5/golang_proto_test/config/deploymentConfigurations"

type Kustomize struct {
	Spin_flavor                string
	Output_dir                 string
	Halyard                    *deploymentConfigurations.HalFile
	ProfilesConfigFiles        map[string]string
	ServiceSettingsConfigFiles map[string]string
	CurrentDeploymentPos       int
	PatchSizing                bool
	Header                     string
}
