// Package validate supports validating a *config.Hal and reporting any
// validation failures.
package validate

import (
	"errors"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentconfigurations/providers"
)

type validationFailure struct {
	msg string
}

type halConfigValidator func(*providers.AppEngine) []validationFailure

func getValidators() []halConfigValidator {
	return []halConfigValidator{
		validateKindsAndOmitKinds,
	}
}

// HalConfig validates the supplied *config.Hal, returning any errors encountered.
// func HalConfig(h *config.Hal) error {
func HalConfig(h *providers.AppEngine) error {
	messages := getValidationMessages(h, getValidators())
	if len(messages) > 0 {
		msg := strings.Join(messages, "\n")
		return errors.New(msg)
	}
	return nil
}

// func getValidationMessages(h *config.Hal, fa []halConfigValidator) []string {
func getValidationMessages(h *providers.AppEngine, fa []halConfigValidator) []string {
	var messages []string
	for _, f := range fa {
		rs := f(h)
		for _, r := range rs {
			messages = append(messages, r.msg)
		}
	}
	return messages
}

// func validateKindsAndOmitKinds(h *config.Hal) []validationFailure {
func validateKindsAndOmitKinds(h *providers.AppEngine) []validationFailure {
	var messages []validationFailure
	// for _, a := range h.GetProviders().GetKubernetes().GetAccounts() {
	// 	if !(len(a.GetKinds()) == 0) && !(len(a.GetOmitKinds()) == 0) {
	// 		messages = append(messages, fatalResult("Cannot specify both kinds and omitKinds."))
	// 	}
	// }
	return messages
}

func fatalResult(msg string) validationFailure {
	return validationFailure{
		msg: msg,
	}
}
