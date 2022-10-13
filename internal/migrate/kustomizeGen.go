// Package migrate is to create and insert the data inside kustomize files
package migrate

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/internal/fileio"
)

type Kustomize struct {
	Spin_flavor string
	Output_dir  string
}

func (KustomizeData Kustomize) CreateKustomization() error {
	str := `resources:
	- SpinnakerService.yml

	# Apply the patches top down order
	patchesStrategicMerge:
	- config-patch.yml
	- profiles-patch.yml
	- files-patch.yml
	- service-settings-patch.yml
`

	str = strings.Replace(str, "\t", "", -1)

	err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/kustomization.yml", str)
	if err != nil {
		return err
	}

	header := KustomizeData.GetHeaders()

	err = KustomizeData.CreateSpinnakerService(header)
	if err != nil {
		return err
	}

	err = KustomizeData.CreateConfigPatch(header)
	if err != nil {
		return err
	}

	err = KustomizeData.CreateProfilesPatch(header)
	if err != nil {
		return err
	}

	err = KustomizeData.CreateFilesPatch(header)
	if err != nil {
		return err
	}

	err = KustomizeData.CreateServiceSettingsPatch(header)
	if err != nil {
		return err
	}

	return nil
}

func (KustomizeData Kustomize) GetHeaders() string {
	apiVersion := ""

	if "OSS" == KustomizeData.Spin_flavor {
		apiVersion = "apiVersion: spinnaker.io/v1alpha2"
	} else {
		apiVersion = "apiVersion: spinnaker.armory.io/v1alpha2"
	}

	header := apiVersion + `
kind: SpinnakerService
metadata:
	name: spinnaker
spec:`

	return header
}

func (KustomizeData Kustomize) CreateSpinnakerService(header string) error {
	SpinnakerServiceStr := KustomizeData.GetSpinnakerService(header)

	err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/SpinnakerService.yml", SpinnakerServiceStr)
	if err != nil {
		return err
	}
	return nil
}

func (KustomizeData Kustomize) CreateConfigPatch(header string) error {
	ConfigPatchStr := KustomizeData.GetConfigPatch(header)

	err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/config-patch.yml", ConfigPatchStr)
	if err != nil {
		return err
	}
	return nil
}

func (KustomizeData Kustomize) CreateProfilesPatch(header string) error {
	ProfilesPatchStr := KustomizeData.GetProfilesPatch(header)

	err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/profiles-patch.yml", ProfilesPatchStr)
	if err != nil {
		return err
	}
	return nil
}

func (KustomizeData Kustomize) CreateFilesPatch(header string) error {
	FilesPatchStr := KustomizeData.GetFilesPatch(header)

	err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/files-patch.yml", FilesPatchStr)
	if err != nil {
		return err
	}
	return nil
}

func (KustomizeData Kustomize) CreateServiceSettingsPatch(header string) error {
	ServiceSettingsPatchStr := KustomizeData.GetServiceSettingsPatch(header)

	err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/service-settings-patch.yml", ServiceSettingsPatchStr)
	if err != nil {
		return err
	}
	return nil
}

func (KustomizeData Kustomize) GetSpinnakerService(header string) string {

	str := header + `
	spinnakerConfig:
		config:
			version: 0.0.0   # the version of Spinnaker to be deployed, see config-patch.yml
			persistentStorage:
			persistentStoreType: s3
			s3:
				bucket: mybucket
				rootFolder: front50
		profiles:
			clouddriver:
			artifacts:
				http:
				enabled: true
				accounts: []
			deck:
			settings-local.js: |
				window.spinnakerSettings.feature.kustomizeEnabled = true;
				window.spinnakerSettings.feature.artifactsRewrite = true;
			echo: {}    # is the contents of ~/.hal/default/profiles/echo.yml
			fiat: {}    # is the contents of ~/.hal/default/profiles/fiat.yml
			front50: {} # is the contents of ~/.hal/default/profiles/front50.yml
			gate: {}    # is the contents of ~/.hal/default/profiles/gate.yml
			igor: {}    # is the contents of ~/.hal/default/profiles/igor.yml
			kayenta: {} # is the contents of ~/.hal/default/profiles/kayenta.yml
			orca: {}    # is the contents of ~/.hal/default/profiles/orca.yml
			rosco: {}   # is the contents of ~/.hal/default/profiles/rosco.yml

		service-settings:
			clouddriver: {}
			deck: {}
			echo: {}
			fiat: {}
			front50: {}
			gate: {}
			igor: {}
			kayenta: {}
			orca: {}
			rosco: {}
		files: {}

		expose:
		type: service
		service:
			type: LoadBalancer
			overrides: {}
`

	str = strings.Replace(str, "\t", "  ", -1)

	return str
}

func (KustomizeData Kustomize) GetConfigPatch(header string) string {

	str := header + `
	# spec.spinnakerConfig - This section is how to specify configuration spinnaker
	spinnakerConfig:
		# spec.spinnakerConfig.config - This section contains the contents of a deployment found in a halconfig .deploymentConfigurations[0]
		config:
			version: 1.17.1  # the version of Spinnaker to be deployed
			persistentStorage:
				persistentStoreType: s3
				s3:
					bucket: mybucket
					rootFolder: front50
`

	str = strings.Replace(str, "\t", "  ", -1)

	return str
}

func (KustomizeData Kustomize) GetProfilesPatch(header string) string {

	str := header + `
	# spec.spinnakerConfig - This section is how to specify configuration spinnaker
	spinnakerConfig:
		# spec.spinnakerConfig.profiles - This section contains the YAML of each service's profile
		profiles:
			clouddriver: {} # is the contents of ~/.hal/default/profiles/clouddriver.yml
			# deck has a special key "settings-local.js" for the contents of settings-local.js
			deck:
				# settings-local.js - contents of ~/.hal/default/profiles/settings-local.js
				# Use the | YAML symbol to indicate a block-style multiline string
				settings-local.js: |
					window.spinnakerSettings.feature.kustomizeEnabled = true;
					window.spinnakerSettings.feature.artifactsRewrite = true;
			echo: {}    # is the contents of ~/.hal/default/profiles/echo.yml
			fiat: {}    # is the contents of ~/.hal/default/profiles/fiat.yml
			front50: {} # is the contents of ~/.hal/default/profiles/front50.yml
			gate: {}    # is the contents of ~/.hal/default/profiles/gate.yml
			igor: {}    # is the contents of ~/.hal/default/profiles/igor.yml
			kayenta: {} # is the contents of ~/.hal/default/profiles/kayenta.yml
			orca: {}    # is the contents of ~/.hal/default/profiles/orca.yml
			rosco: {}   # is the contents of ~/.hal/default/profiles/rosco.yml
`

	str = strings.Replace(str, "\t", "  ", -1)

	return str
}

func (KustomizeData Kustomize) GetFilesPatch(header string) string {

	str := header + `
	# spec.spinnakerConfig - This section is how to specify configuration spinnaker
	spinnakerConfig:
		# spec.spinnakerConfig.files - This section allows you to include any other raw string files not handle above.
		# The KEY is the filepath and filename of where it should be placed
		#	- Files here will be placed into ~/.hal/default/ on halyard
		#	- __ is used in place of / for the path separator
		# The VALUE is the contents of the file.
		#	- Use the | YAML symbol to indicate a block-style multiline string
		#	- We currently only support string files
		#	- NOTE: Kubernetes has a manifest size limitation of 1MB
		files:
#			profiles__rosco__packer__example-packer-config.json: |
#				{
#					"packerSetting": "someValue"
#				}
#			profiles__rosco__packer__my_custom_script.sh: |
#				#!/bin/bash -e
#				echo "hello world!"
`

	str = strings.Replace(str, "\t", "  ", -1)

	return str
}

func (KustomizeData Kustomize) GetServiceSettingsPatch(header string) string {

	str := header + `
	# spec.spinnakerConfig - This section is how to specify configuration spinnaker
	spinnakerConfig:
		# spec.spinnakerConfig.service-settings - This section contains the YAML of the service's service-setting
		# see https://www.spinnaker.io/reference/halyard/custom/#tweakable-service-settings for available settings
		service-settings:
			clouddriver: {}
			deck: {}
			echo: {}
			fiat: {}
			front50: {}
			gate: {}
			igor: {}
			kayenta: {}
			orca: {}
			rosco: {}
`

	str = strings.Replace(str, "\t", "  ", -1)

	return str
}
