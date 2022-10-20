package config_patch

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetSecurity(KustomizeData structs.Kustomize) string {

	return GetUiApiSecurity(KustomizeData) +
		GetAuthnSecurity(KustomizeData) +
		GetAuthzSecurity(KustomizeData)
}

func GetUiApiSecurity(KustomizeData structs.Kustomize) string {
	str := `
	apiSecurity:
	  ssl:
	    enabled: false
	    keyAlias: gate
	    keyStore: XXXX
	    keyStoreType: jks
	    keyStorePassword: XXXX
	    trustStore: XXXX
	    trustStoreType: jks
	    trustStorePassword: XXXX
	  overrideBaseUrl: https://api.spinnaker.mycomany.com` /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Security.ApiSecurity.OverrideBaseUrl*/ + `
	uiSecurity:
	  ssl:
	    enabled: false
	  overrideBaseUrl: https://spinnaker.mycomany.com` /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Security.UiSecurity.OverrideBaseUrl*/

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetAuthnSecurity(KustomizeData structs.Kustomize) string {
	str := `
	authn:
	  oauth2:
	    enabled: true
	    client:
	      clientId: XXXX
	        clientSecret: XXXX
	        accessTokenUri: XXXX
	        userAuthorizationUri: XXXX
	        scope: user:email
	    resource:
	      userInfoUri: XXXX
	    userInfoMapping:
	      email: email
	      firstName: name
	      lastName: ""
	      username: login
	    provider: GITHUB                                                   # One of AZURE, GITHUB, ORACLE, OTHER, GOOGLE
	  saml:
	    enabled: true
	    metadataLocal: XXXX
	    issuerId: XXXX
	    keyStore: XXXX
	    keyStorePassword: XXXX
	    keyStoreAliasName: saml
	    serviceAddress: hXXXX
	    userAttributeMapping: {}
	  ldap:
	    enabled: false
	  x509:
	    enabled: true
	    roleOid: 1.2.840.10070.8.1
	  iap:
	    enabled: false
	  enabled: false`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Security.UiSecurity.OverrideBaseUrl*/

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetAuthzSecurity(KustomizeData structs.Kustomize) string {
	str := `
	authz:
	  groupMembership:
	    service: FILE
	    google:
	      roleProviderType: GOOGLE
	    github:
	      roleProviderType: GITHUB
	      baseUrl: https://api.github.com
	      accessToken: XXXX
	      organization: XXXX
	    file:
	      roleProviderType: FILE
	      path: /home/spinnaker/.hal/rbac-generated.yaml
	    ldap:
	      roleProviderType: LDAP
	  enabled: false`
	/*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Security.UiSecurity.OverrideBaseUrl*/

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}
