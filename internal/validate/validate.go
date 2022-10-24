// Package validate supports validating a *config.Hal and reporting any
// validation failures.
package validate

import (
	"errors"
	"fmt"
	"strings"

	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

type validationFailure struct {
	msg string
}

type validator func(*structs.Kustomize) []validationFailure

func getValidators() []validator {
	return []validator{
		validateKindsAndOmitKinds,
		// validateCurrentDeploymentExists,
		// validateProfilesExists,
		validateProfilesFileHasData,
	}
}

// KustomizeConfig validates the supplied *structs.Kustomize, returning any errors encountered.
func KustomizeConfig(KustomizeData *structs.Kustomize) error {

	fmt.Println("Running Validations")

	messages := getValidationMessages(KustomizeData, getValidators())
	if len(messages) > 0 {
		msg := strings.Join(messages, "\n")
		return errors.New(msg)
	}
	return nil
}

func getValidationMessages(KustomizeData *structs.Kustomize, fa []validator) []string {
	var messages []string
	for _, f := range fa {
		rs := f(KustomizeData)
		for _, r := range rs {
			messages = append(messages, r.msg)
		}
	}
	return messages
}

func validateKindsAndOmitKinds(KustomizeData *structs.Kustomize) []validationFailure {
	var messages []validationFailure
	// for _, a := range h.GetProviders().GetKubernetes().GetAccounts() {
	// 	if !(len(a.GetKinds()) == 0) && !(len(a.GetOmitKinds()) == 0) {
	// 		messages = append(messages, fatalResult("Cannot specify both kinds and omitKinds."))
	// 	}
	// }
	return messages
}

func validateCurrentDeploymentExists(KustomizeData *structs.Kustomize) []validationFailure {
	var messages []validationFailure
	ok := false

	for _, a := range KustomizeData.Halyard.GetDeploymentConfiguration() {
		if KustomizeData.Halyard.Name == a.Name {
			ok = true
		}
	}

	if !ok {
		messages = append(messages, fatalResult("Cannot find current deployment"))
	}

	return messages
}

func validateProfilesFileHasData(KustomizeData *structs.Kustomize) []validationFailure {
	var messages []validationFailure

	for key, value := range KustomizeData.ProfilesConfigFiles {
		if _, isMapContainsKey := KustomizeData.ProfilesConfigFiles[key]; isMapContainsKey {
			if "" != value {
				messages = append(messages, fatalResult(key+" is empty"))
			}
		}
	}

	return messages
}

func fatalResult(msg string) validationFailure {
	return validationFailure{
		msg: msg,
	}
}
