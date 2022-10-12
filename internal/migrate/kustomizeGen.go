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

	err = KustomizeData.CreateSpinnakerService()
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

func (KustomizeData Kustomize) GetSpinnakerService() string {
	header := KustomizeData.GetHeaders()

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

func (KustomizeData Kustomize) CreateSpinnakerService() error {
	SpinnakerServiceStr := KustomizeData.GetSpinnakerService()

	err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/SpinnakerService.yml", SpinnakerServiceStr)
	if err != nil {
		return err
	}
	return nil
}
