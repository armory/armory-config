// Package validate supports validating a *config.Hal and reporting any
// validation failures.
package validate

import (
	"errors"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations"
)

type validationFailure struct {
	msg string
}

type halConfigValidator func(*deploymentConfigurations.HalFile) []validationFailure

func getValidators() []halConfigValidator {
	return []halConfigValidator{
		validateKindsAndOmitKinds,
		// validateCurrentDeploymentExists,
	}
}

// HalConfig validates the supplied *config.Hal, returning any errors encountered.
func HalConfig(h *deploymentConfigurations.HalFile) error {
	messages := getValidationMessages(h, getValidators())
	if len(messages) > 0 {
		msg := strings.Join(messages, "\n")
		return errors.New(msg)
	}
	return nil
}

func getValidationMessages(h *deploymentConfigurations.HalFile, fa []halConfigValidator) []string {
	var messages []string
	for _, f := range fa {
		rs := f(h)
		for _, r := range rs {
			messages = append(messages, r.msg)
		}
	}
	return messages
}

func validateKindsAndOmitKinds(h *deploymentConfigurations.HalFile) []validationFailure {
	var messages []validationFailure
	// for _, a := range h.GetProviders().GetKubernetes().GetAccounts() {
	// 	if !(len(a.GetKinds()) == 0) && !(len(a.GetOmitKinds()) == 0) {
	// 		messages = append(messages, fatalResult("Cannot specify both kinds and omitKinds."))
	// 	}
	// }
	return messages
}

func validateCurrentDeploymentExists(h *deploymentConfigurations.HalFile) []validationFailure {
	var messages []validationFailure
	ok := false

	for _, a := range h.GetDeploymentConfiguration() {
		if h.Name == a.Name {
			ok = true
		}
	}

	if !ok {
		messages = append(messages, fatalResult("Cannot find current deployment"))
	}

	return messages
}

func fatalResult(msg string) validationFailure {
	return validationFailure{
		msg: msg,
	}
}
