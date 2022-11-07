package config_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/security"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetSecurity(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security {
		str = GetUiApiSecurity(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security) +
			GetAuthnSecurity(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security) +
			GetAuthzSecurity(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security)
	}
	return str
}

func GetUiApiSecurity(securityReference *security.Security) string {
	str := ""

	if nil != securityReference.ApiSecurity {
		str = `
		apiSecurity:
		  ssl:
		    enabled: ` + strconv.FormatBool(securityReference.ApiSecurity.Ssl.Enabled) + `
		    keyAlias: ` + securityReference.ApiSecurity.Ssl.KeyAlias + `
		    keyStore: ` + securityReference.ApiSecurity.Ssl.KeyStore + `
		    keyStoreType: ` + securityReference.ApiSecurity.Ssl.KeyStoreType + `
		    keyStorePassword: ` + securityReference.ApiSecurity.Ssl.KeyStorePassword + `
		    trustStore: ` + securityReference.ApiSecurity.Ssl.TrustStore + `
		    trustStoreType: ` + securityReference.ApiSecurity.Ssl.TrustStoreType + `
		    trustStorePassword: ` + securityReference.ApiSecurity.Ssl.TrustStorePassword + `
		  overrideBaseUrl: ` + securityReference.ApiSecurity.OverrideBaseUrl
	}

	if nil != securityReference.UiSecurity {
		str += `
		uiSecurity:
		  ssl:
		    enabled: ` + strconv.FormatBool(securityReference.UiSecurity.Ssl.Enabled) + `
		    keyAlias: ` + securityReference.UiSecurity.Ssl.KeyAlias + `
		    keyStore: ` + securityReference.UiSecurity.Ssl.KeyStore + `
		    keyStoreType: ` + securityReference.UiSecurity.Ssl.KeyStoreType + `
		    keyStorePassword: ` + securityReference.UiSecurity.Ssl.KeyStorePassword + `
		    trustStore: ` + securityReference.UiSecurity.Ssl.TrustStore + `
		    trustStoreType: ` + securityReference.UiSecurity.Ssl.TrustStoreType + `
		    trustStorePassword: ` + securityReference.UiSecurity.Ssl.TrustStorePassword + `
		  overrideBaseUrl: ` + securityReference.UiSecurity.OverrideBaseUrl
	}

	str = strings.Replace(str, "\t", "    ", -1)

	return str
}

func GetAuthnSecurity(securityReference *security.Security) string {
	str := ""

	if nil == securityReference.Authn {
		return str
	}

	str = `
		authn:`

	if nil != securityReference.Authn.Oauth2 {
		str += `
		  oauth2:
		    enabled: ` + strconv.FormatBool(securityReference.Authn.Oauth2.Enabled)
		if nil != securityReference.Authn.Oauth2.Client && "" != securityReference.Authn.Oauth2.Client.ClientId {
			str += `
		    client:
		      clientId: ` + securityReference.Authn.Oauth2.Client.ClientId + `
		      clientSecret: ` + securityReference.Authn.Oauth2.Client.ClientSecret + `
		      accessTokenUri: ` + securityReference.Authn.Oauth2.Client.AccessTokenUri + `
		      userAuthorizationUri: ` + securityReference.Authn.Oauth2.Client.UserAuthorizationUri + `
		      scope: ` + securityReference.Authn.Oauth2.Client.Scope
		} else {
			str += `
		    client: {}`
		}

		if nil != securityReference.Authn.Oauth2.Resource && "" != securityReference.Authn.Oauth2.Resource.UserInfoUri {
			str += `
		    resource:
		      userInfoUri: XXXX` + securityReference.Authn.Oauth2.Resource.UserInfoUri
		} else {
			str += `
			    resource: {}`
		}

		if nil != securityReference.Authn.Oauth2.UserInfoMapping && "" != securityReference.Authn.Oauth2.UserInfoMapping.Email {
			str += `
		    userInfoMapping:
		      email: ` + securityReference.Authn.Oauth2.UserInfoMapping.Email + `
		      firstName: ` + securityReference.Authn.Oauth2.UserInfoMapping.FirstName + `
		      lastName: "` + securityReference.Authn.Oauth2.UserInfoMapping.LastName + `"
		      username: ` + securityReference.Authn.Oauth2.UserInfoMapping.Username
		} else {
			str += `
		    userInfoMapping: {}`
		}

		str += `
		    provider: ` + securityReference.Authn.Oauth2.Provider + `                # One of AZURE, GITHUB, ORACLE, OTHER, GOOGLE`
	}

	if nil != securityReference.Authn.Saml {
		str += `
		  saml:
		    enabled: ` + strconv.FormatBool(securityReference.Authn.Saml.Enabled) + `
		    metadataRemote: ` + securityReference.Authn.Saml.MetadataRemote + `
		    metadataLocal: ` + /*securityReference.Authn.Saml.MetadataLocal + */ `
		    issuerId: ` + securityReference.Authn.Saml.IssuerId + `
		    keyStore: ` + securityReference.Authn.Saml.KeyStore + `
		    keyStorePassword: ` + securityReference.Authn.Saml.KeyStorePassword + `
		    keyStoreAliasName: ` + securityReference.Authn.Saml.KeyStoreAliasName + `
		    serviceAddress: ` + securityReference.Authn.Saml.ServiceAddress

		if nil != securityReference.Authn.Saml.UserInfoMapping {
			str += `
		    userAttributeMapping:
		      email: ` + securityReference.Authn.Saml.UserInfoMapping.Email + `
		      firstname: ` + securityReference.Authn.Saml.UserInfoMapping.Firstname + `
		      lastname: ` + securityReference.Authn.Saml.UserInfoMapping.Lastname + `
		      username: ` + securityReference.Authn.Saml.UserInfoMapping.Username + `
		      roles: ` + securityReference.Authn.Saml.UserInfoMapping.Roles + `
		      rolesDelimiter: ` + securityReference.Authn.Saml.UserInfoMapping.RolesDelimiter
		} else {
			str += `
		    userAttributeMapping: {}`
		}
	}

	if nil != securityReference.Authn.Ldap {
		str += `
		  ldap:
		    enabled: ` + strconv.FormatBool(securityReference.Authn.Ldap.Enabled) + `
		    url: ` + securityReference.Authn.Ldap.Url + `
		    userDnPattern: ` + securityReference.Authn.Ldap.UserDnPattern + `
		    userSearchFilter: ` + securityReference.Authn.Ldap.UserSearchFilter + `
		    managerDn: ` + securityReference.Authn.Ldap.ManagerDn + `
		    managerPassword: ` + securityReference.Authn.Ldap.ManagerPassword + `
		    groupSearchBase: ` + securityReference.Authn.Ldap.GroupSearchBase
	}

	if nil != securityReference.Authn.X509 {
		str += `
		  x509:
		    enabled: ` + strconv.FormatBool(securityReference.Authn.X509.Enabled) + `
		    subjectPrincipalRegex: ` + securityReference.Authn.X509.SubjectPrincipalRegex + `
		    roleOid: ` + securityReference.Authn.X509.RoleOid
	}

	if nil != securityReference.Authn.Iap {
		str += `
		  iap:
		    enabled: ` + strconv.FormatBool(securityReference.Authn.Iap.Enabled) + `
		    jwtHeader: ` + securityReference.Authn.Iap.JwtHeader + `
		    issuerId: ` + securityReference.Authn.Iap.IssuerId + `
		    audience: ` + securityReference.Authn.Iap.Audience + `
		    iapVerifyKeyUrl: ` + securityReference.Authn.Iap.IapVerifyKeyUrl
	}

	str = strings.Replace(str, "\t", "    ", -1)

	return str
}

func GetAuthzSecurity(securityReference *security.Security) string {
	str := ""

	if nil == securityReference.Authz.GroupMembership {
		return str
	}

	str = `
		authz:
		  groupMembership:
		    service: ` + securityReference.Authz.GroupMembership.Service + `
		    enable: ` + strconv.FormatBool(securityReference.Authz.GroupMembership.Enabled)

	if nil != securityReference.Authz.GroupMembership.Google {
		str += `
		    google:
		      roleProviderType: ` + securityReference.Authz.GroupMembership.Google.RoleProviderType + `
		      path: ` + securityReference.Authz.GroupMembership.Google.Path + `
		      credentialPath: ` + securityReference.Authz.GroupMembership.Google.CredentialPath + `
		      adminUsername: ` + securityReference.Authz.GroupMembership.Google.AdminUsername + `
		      domain: ` + securityReference.Authz.GroupMembership.Google.Domain
	}

	if nil != securityReference.Authz.GroupMembership.Github {
		str += `
		    github:
		      roleProviderType: ` + securityReference.Authz.GroupMembership.Github.RoleProviderType + `
		      path: ` + securityReference.Authz.GroupMembership.Github.Path + `
		      credentialPath: ` + securityReference.Authz.GroupMembership.Github.CredentialPath + `
		      adminUsername: ` + securityReference.Authz.GroupMembership.Github.AdminUsername + `
		      domain: ` + securityReference.Authz.GroupMembership.Github.Domain
	}
	// TODO Missing proto fields
	//     github:
	//       roleProviderType: GITHUB
	//       baseUrl: https://api.github.com
	//       accessToken: XXXX
	//       organization: XXXX

	if nil != securityReference.Authz.GroupMembership.File {
		str += `
		    file:
		      roleProviderType: ` + securityReference.Authz.GroupMembership.File.RoleProviderType + `
		      path: ` + securityReference.Authz.GroupMembership.File.Path + `
		      credentialPath: ` + securityReference.Authz.GroupMembership.File.CredentialPath + `
		      adminUsername: ` + securityReference.Authz.GroupMembership.File.AdminUsername + `
		      domain: ` + securityReference.Authz.GroupMembership.File.Domain
	}

	if nil != securityReference.Authz.GroupMembership.Ldap {
		str += `
		    ldap:
		      roleProviderType: ` + securityReference.Authz.GroupMembership.Ldap.RoleProviderType + `
		      path: ` + securityReference.Authz.GroupMembership.Ldap.Path + `
		      credentialPath: ` + securityReference.Authz.GroupMembership.Ldap.CredentialPath + `
		      adminUsername: ` + securityReference.Authz.GroupMembership.Ldap.AdminUsername + `
		      domain: ` + securityReference.Authz.GroupMembership.Ldap.Domain
	}

	str += `
	      enabled: ` + strconv.FormatBool(securityReference.Authz.Enabled)

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
