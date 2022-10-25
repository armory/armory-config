package config_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetSecurity(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security {
		str = GetUiApiSecurity(KustomizeData) +
			GetAuthnSecurity(KustomizeData) +
			GetAuthzSecurity(KustomizeData)
	}
	return str
}

func GetUiApiSecurity(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.ApiSecurity {
		str = `
		apiSecurity:
		  ssl:
		    enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.ApiSecurity.Ssl.Enabled) + `
		    keyAlias: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.ApiSecurity.Ssl.KeyAlias + `
		    keyStore: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.ApiSecurity.Ssl.KeyStore + `
		    keyStoreType: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.ApiSecurity.Ssl.KeyStoreType + `
		    keyStorePassword: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.ApiSecurity.Ssl.KeyStorePassword + `
		    trustStore: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.ApiSecurity.Ssl.TrustStore + `
		    trustStoreType: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.ApiSecurity.Ssl.TrustStoreType + `
		    trustStorePassword: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.ApiSecurity.Ssl.TrustStorePassword + `
		  overrideBaseUrl: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.ApiSecurity.OverrideBaseUrl
	}

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.UiSecurity {
		str += `
		uiSecurity:
		  ssl:
		    enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.UiSecurity.Ssl.Enabled) + `
		    keyAlias: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.UiSecurity.Ssl.KeyAlias + `
		    keyStore: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.UiSecurity.Ssl.KeyStore + `
		    keyStoreType: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.UiSecurity.Ssl.KeyStoreType + `
		    keyStorePassword: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.UiSecurity.Ssl.KeyStorePassword + `
		    trustStore: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.UiSecurity.Ssl.TrustStore + `
		    trustStoreType: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.UiSecurity.Ssl.TrustStoreType + `
		    trustStorePassword: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.UiSecurity.Ssl.TrustStorePassword + `
		  overrideBaseUrl: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.UiSecurity.OverrideBaseUrl
	}

	str = strings.Replace(str, "\t", "    ", -1)

	return str
}

func GetAuthnSecurity(KustomizeData structs.Kustomize) string {
	str := ""

	if nil == KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn {
		return str
	}

	str = `
		authn:`

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Oauth2 {
		str += `
		  oauth2:
		    enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Oauth2.Enabled) + `
		    client:
		      clientId: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Oauth2.Client.ClientId + `
		      clientSecret: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Oauth2.Client.ClientSecret + `
		      accessTokenUri: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Oauth2.Client.AccessTokenUri + `
		      userAuthorizationUri: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Oauth2.Client.UserAuthorizationUri + `
		      scope: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Oauth2.Client.Scope + `
		    resource:
		      userInfoUri: XXXX` + /*+ KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Oauth2.Resource.UserInfoUri + `
			  userInfoMapping:
			    email: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Oauth2.UserInfoMapping.Email + `
			    firstName: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Oauth2.UserInfoMapping.FirstName + `
			    lastName: "` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Oauth2.UserInfoMapping.LastName + `"
			    username: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Oauth2.UserInfoMapping.UserName +*/`
		    provider: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Oauth2.Provider + `                # One of AZURE, GITHUB, ORACLE, OTHER, GOOGLE`
	}

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Saml {
		str += `
		  saml:
		    enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Saml.Enabled) + `
		    metadataRemote: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Saml.MetadataRemote + `
		    metadataLocal: ` + /*KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Saml.MetadataLocal + */ `
		    issuerId: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Saml.IssuerId + `
		    keyStore: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Saml.KeyStore + `
		    keyStorePassword: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Saml.KeyStorePassword + `
		    keyStoreAliasName: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Saml.KeyStoreAliasName + `
		    serviceAddress: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Saml.ServiceAddress + `
		    userAttributeMapping: {}` //TODO
	}

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Ldap {
		str += `
		  ldap:
		    enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Ldap.Enabled) + `
		    url: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Ldap.Url + `
		    userDnPattern: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Ldap.UserDnPattern + `
		    userSearchFilter: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Ldap.UserSearchFilter + `
		    managerDn: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Ldap.ManagerDn + `
		    managerPassword: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Ldap.ManagerPassword + `
		    groupSearchBase: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Ldap.GroupSearchBase
	}

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.X509 {
		str += `
		  x509:
		    enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.X509.Enabled) + `
		    subjectPrincipalRegex: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.X509.SubjectPrincipalRegex + `
		    roleOid: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.X509.RoleOid
	}

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Iap {
		str += `
		  iap:
		    enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Iap.Enabled) + `
		    jwtHeader: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Iap.JwtHeader + `
		    issuerId: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Iap.IssuerId + `
		    audience: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Iap.Audience + `
		    iapVerifyKeyUrl: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authn.Iap.IapVerifyKeyUrl
	}

	str = strings.Replace(str, "\t", "    ", -1)

	return str
}

func GetAuthzSecurity(KustomizeData structs.Kustomize) string {
	str := ""

	if nil == KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authz.GroupMembership {
		return str
	}

	str = `
		authz:
		  groupMembership:
		    service: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authz.GroupMembership.Service + `
		    enable: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authz.GroupMembership.Enabled)

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authz.GroupMembership.Google {
		str += `
		    google:
		      roleProviderType: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authz.GroupMembership.Google.RoleProviderType + `
		      path: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authz.GroupMembership.Google.Path
	}

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authz.GroupMembership.Github {
		str += `
		    github:
		      roleProviderType: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authz.GroupMembership.Github.RoleProviderType + `
		      path: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authz.GroupMembership.Github.Path
	}
	// TODO
	//     github:
	//       roleProviderType: GITHUB
	//       baseUrl: https://api.github.com
	//       accessToken: XXXX
	//       organization: XXXX

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authz.GroupMembership.File {
		str += `
		    file:
		      roleProviderType: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authz.GroupMembership.File.RoleProviderType + `
		      path: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authz.GroupMembership.File.Path
	}

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authz.GroupMembership.Ldap {
		str += `
		    ldap:
		      roleProviderType: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authz.GroupMembership.Ldap.RoleProviderType + `
		      path: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authz.GroupMembership.Ldap.Path
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
