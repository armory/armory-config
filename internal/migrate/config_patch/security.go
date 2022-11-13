package config_patch

import (
	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/security"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
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
		  ssl:` +
			helpers.PrintFmtBool(`enabled: `, securityReference.ApiSecurity.Ssl.Enabled, 6, true) +
			helpers.PrintFmtStr(`keyAlias: `, securityReference.ApiSecurity.Ssl.KeyAlias, 6, true) +
			helpers.PrintFmtStr(`keyStore: `, securityReference.ApiSecurity.Ssl.KeyStore, 6, true) +
			helpers.PrintFmtStr(`keyStoreType: `, securityReference.ApiSecurity.Ssl.KeyStoreType, 6, true) +
			helpers.PrintFmtStr(`keyStorePassword: `, securityReference.ApiSecurity.Ssl.KeyStorePassword, 6, true) +
			helpers.PrintFmtStr(`trustStore: `, securityReference.ApiSecurity.Ssl.TrustStore, 6, true) +
			helpers.PrintFmtStr(`trustStoreType: `, securityReference.ApiSecurity.Ssl.TrustStoreType, 6, true) +
			helpers.PrintFmtStr(`trustStorePassword: `, securityReference.ApiSecurity.Ssl.TrustStorePassword, 6, true) +
			helpers.PrintFmtStr(`overrideBaseUrl: `, securityReference.ApiSecurity.OverrideBaseUrl, 5, true)
	}

	if nil != securityReference.UiSecurity {
		str += `
		uiSecurity:
		  ssl:` +
			helpers.PrintFmtBool(`enabled: `, securityReference.UiSecurity.Ssl.Enabled, 6, true) +
			helpers.PrintFmtStr(`keyAlias: `, securityReference.UiSecurity.Ssl.KeyAlias, 6, true) +
			helpers.PrintFmtStr(`keyStore: `, securityReference.UiSecurity.Ssl.KeyStore, 6, true) +
			helpers.PrintFmtStr(`keyStoreType: `, securityReference.UiSecurity.Ssl.KeyStoreType, 6, true) +
			helpers.PrintFmtStr(`keyStorePassword: `, securityReference.UiSecurity.Ssl.KeyStorePassword, 6, true) +
			helpers.PrintFmtStr(`trustStore: `, securityReference.UiSecurity.Ssl.TrustStore, 6, true) +
			helpers.PrintFmtStr(`trustStoreType: `, securityReference.UiSecurity.Ssl.TrustStoreType, 6, true) +
			helpers.PrintFmtStr(`trustStorePassword: `, securityReference.UiSecurity.Ssl.TrustStorePassword, 6, true) +
			helpers.PrintFmtStr(`overrideBaseUrl: `, securityReference.UiSecurity.OverrideBaseUrl, 5, true)
	}

	return str
}

func GetAuthnSecurity(securityReference *security.Security) string {
	str := ""

	if nil == securityReference.Authn {
		return str
	}

	str = `
		authn:` +
		helpers.PrintFmtBool(`enabled: `, securityReference.Authn.Enabled, 5, true)

	if nil != securityReference.Authn.Oauth2 {
		str += `
		  oauth2:` +
			helpers.PrintFmtBool(`enabled: `, securityReference.Authn.Oauth2.Enabled, 6, true)
		if nil != securityReference.Authn.Oauth2.Client && "" != securityReference.Authn.Oauth2.Client.ClientId {
			str += `
		    client:` +
				helpers.PrintFmtStr(`clientId: `, securityReference.Authn.Oauth2.Client.ClientId, 7, true) +
				helpers.PrintFmtStr(`clientSecret: `, securityReference.Authn.Oauth2.Client.ClientSecret, 7, true) +
				helpers.PrintFmtStr(`accessTokenUri: `, securityReference.Authn.Oauth2.Client.AccessTokenUri, 7, true) +
				helpers.PrintFmtStr(`userAuthorizationUri: `, securityReference.Authn.Oauth2.Client.UserAuthorizationUri, 7, true) +
				helpers.PrintFmtStr(`scope: `, securityReference.Authn.Oauth2.Client.Scope, 7, true)
		} else {
			str += `
		    client: {}`
		}

		if nil != securityReference.Authn.Oauth2.Resource && "" != securityReference.Authn.Oauth2.Resource.UserInfoUri {
			str += `
			resource:` +
				helpers.PrintFmtStr(`userInfoUri: `, securityReference.Authn.Oauth2.Resource.UserInfoUri, 7, true)
		} else {
			str += `
			resource: {}`
		}

		if nil != securityReference.Authn.Oauth2.UserInfoMapping && "" != securityReference.Authn.Oauth2.UserInfoMapping.Email {
			str += `
		    userInfoMapping:` +
				helpers.PrintFmtStr(`email: `, securityReference.Authn.Oauth2.UserInfoMapping.Email, 7, true) +
				helpers.PrintFmtStr(`firstName: `, securityReference.Authn.Oauth2.UserInfoMapping.FirstName, 7, true) +
				helpers.PrintFmtStrDoubleQuote(`lastName: `, securityReference.Authn.Oauth2.UserInfoMapping.LastName, 7, true) +
				helpers.PrintFmtStr(`username: `, securityReference.Authn.Oauth2.UserInfoMapping.Username, 7, true)
		} else {
			str += `
		    userInfoMapping: {}`
		}

		str += helpers.PrintFmtStr(`provider: `, securityReference.Authn.Oauth2.Provider, 6, true)
	}

	if nil != securityReference.Authn.Saml {
		str += `
		  saml:` +
			helpers.PrintFmtBool(`enabled: `, securityReference.Authn.Saml.Enabled, 6, true) +
			helpers.PrintFmtStr(`metadataRemote: `, securityReference.Authn.Saml.MetadataRemote, 6, true) +
			helpers.PrintFmtStr(`metadataLocal: `, securityReference.Authn.Saml.MetadataLocal, 6, true) +
			helpers.PrintFmtStr(`issuerId: `, securityReference.Authn.Saml.IssuerId, 6, true) +
			helpers.PrintFmtStr(`keyStore: `, securityReference.Authn.Saml.KeyStore, 6, true) +
			helpers.PrintFmtStr(`keyStorePassword: `, securityReference.Authn.Saml.KeyStorePassword, 6, true) +
			helpers.PrintFmtStr(`keyStoreAliasName: `, securityReference.Authn.Saml.KeyStoreAliasName, 6, true) +
			helpers.PrintFmtStr(`serviceAddress: `, securityReference.Authn.Saml.ServiceAddress, 6, true)

		if nil != securityReference.Authn.Saml.UserInfoMapping {
			str += `
		    userAttributeMapping:` +
				helpers.PrintFmtStr(`email: `, securityReference.Authn.Saml.UserInfoMapping.Email, 7, true) +
				helpers.PrintFmtStr(`firstname: `, securityReference.Authn.Saml.UserInfoMapping.Firstname, 7, true) +
				helpers.PrintFmtStr(`lastname: `, securityReference.Authn.Saml.UserInfoMapping.Lastname, 7, true) +
				helpers.PrintFmtStr(`username: `, securityReference.Authn.Saml.UserInfoMapping.Username, 7, true) +
				helpers.PrintFmtStr(`roles: `, securityReference.Authn.Saml.UserInfoMapping.Roles, 7, true) +
				helpers.PrintFmtStr(`rolesDelimiter: `, securityReference.Authn.Saml.UserInfoMapping.RolesDelimiter, 7, true)
		} else {
			str += `
		    userAttributeMapping: {}`
		}
	}

	if nil != securityReference.Authn.Ldap {
		str += `
		  ldap:` +
			helpers.PrintFmtBool(`enabled: `, securityReference.Authn.Ldap.Enabled, 6, true) +
			helpers.PrintFmtStr(`url: `, securityReference.Authn.Ldap.Url, 6, true) +
			helpers.PrintFmtStr(`userDnPattern: `, securityReference.Authn.Ldap.UserDnPattern, 6, true) +
			helpers.PrintFmtStr(`userSearchFilter: `, securityReference.Authn.Ldap.UserSearchFilter, 6, true) +
			helpers.PrintFmtStr(`managerDn: `, securityReference.Authn.Ldap.ManagerDn, 6, true) +
			helpers.PrintFmtStr(`managerPassword: `, securityReference.Authn.Ldap.ManagerPassword, 6, true) +
			helpers.PrintFmtStr(`groupSearchBase: `, securityReference.Authn.Ldap.GroupSearchBase, 6, true)
	}

	if nil != securityReference.Authn.X509 {
		str += `
		  x509:` +
			helpers.PrintFmtBool(`enabled: `, securityReference.Authn.X509.Enabled, 6, true) +
			helpers.PrintFmtStr(`subjectPrincipalRegex: `, securityReference.Authn.X509.SubjectPrincipalRegex, 6, true) +
			helpers.PrintFmtStr(`roleOid: `, securityReference.Authn.X509.RoleOid, 6, true)
	}

	if nil != securityReference.Authn.Iap {
		str += `
		  iap:` +
			helpers.PrintFmtBool(`enabled: `, securityReference.Authn.Iap.Enabled, 6, true) +
			helpers.PrintFmtStr(`jwtHeader: `, securityReference.Authn.Iap.JwtHeader, 6, true) +
			helpers.PrintFmtStr(`issuerId: `, securityReference.Authn.Iap.IssuerId, 6, true) +
			helpers.PrintFmtStr(`audience: `, securityReference.Authn.Iap.Audience, 6, true) +
			helpers.PrintFmtStr(`iapVerifyKeyUrl: `, securityReference.Authn.Iap.IapVerifyKeyUrl, 6, true)
	}

	return str
}

func GetAuthzSecurity(securityReference *security.Security) string {
	str := ""

	if nil == securityReference.Authz.GroupMembership {
		return str
	}

	str = `
		authz:
		  groupMembership:` +
		helpers.PrintFmtStr(`service: `, securityReference.Authz.GroupMembership.Service, 6, true)

	if nil != securityReference.Authz.GroupMembership.Google {
		str += `
		    google:` +
			helpers.PrintFmtStr(`roleProviderType: `, securityReference.Authz.GroupMembership.Google.RoleProviderType, 7, true) +
			helpers.PrintFmtStr(`path: `, securityReference.Authz.GroupMembership.Google.Path, 7, true) +
			helpers.PrintFmtStr(`credentialPath: `, securityReference.Authz.GroupMembership.Google.CredentialPath, 7, true) +
			helpers.PrintFmtStr(`adminUsername: `, securityReference.Authz.GroupMembership.Google.AdminUsername, 7, true) +
			helpers.PrintFmtStr(`domain: `, securityReference.Authz.GroupMembership.Google.Domain, 7, true)
	}

	if nil != securityReference.Authz.GroupMembership.Github {
		str += `
		    github:` +
			helpers.PrintFmtStr(`roleProviderType: `, securityReference.Authz.GroupMembership.Github.RoleProviderType, 7, true) +
			helpers.PrintFmtStr(`path: `, securityReference.Authz.GroupMembership.Github.Path, 7, true) +
			helpers.PrintFmtStr(`credentialPath: `, securityReference.Authz.GroupMembership.Github.CredentialPath, 7, true) +
			helpers.PrintFmtStr(`adminUsername: `, securityReference.Authz.GroupMembership.Github.AdminUsername, 7, true) +
			helpers.PrintFmtStr(`domain: `, securityReference.Authz.GroupMembership.Github.Domain, 7, true)
	}
	// TODO Missing proto fields
	//     github:
	//       roleProviderType: GITHUB
	//       baseUrl: https://api.github.com
	//       accessToken: XXXX
	//       organization: XXXX

	if nil != securityReference.Authz.GroupMembership.File {
		str += `
		    file:` +
			helpers.PrintFmtStr(`roleProviderType: `, securityReference.Authz.GroupMembership.File.RoleProviderType, 7, true) +
			helpers.PrintFmtStr(`path: `, securityReference.Authz.GroupMembership.File.Path, 7, true) +
			helpers.PrintFmtStr(`credentialPath: `, securityReference.Authz.GroupMembership.File.CredentialPath, 7, true) +
			helpers.PrintFmtStr(`adminUsername: `, securityReference.Authz.GroupMembership.File.AdminUsername, 7, true) +
			helpers.PrintFmtStr(`domain: `, securityReference.Authz.GroupMembership.File.Domain, 7, true)
	}

	if nil != securityReference.Authz.GroupMembership.Ldap {
		str += `
		    ldap:` +
			helpers.PrintFmtStr(`roleProviderType: `, securityReference.Authz.GroupMembership.Ldap.RoleProviderType, 7, true) +
			helpers.PrintFmtStr(`path: `, securityReference.Authz.GroupMembership.Ldap.Path, 7, true) +
			helpers.PrintFmtStr(`credentialPath: `, securityReference.Authz.GroupMembership.Ldap.CredentialPath, 7, true) +
			helpers.PrintFmtStr(`adminUsername: `, securityReference.Authz.GroupMembership.Ldap.AdminUsername, 7, true) +
			helpers.PrintFmtStr(`domain: `, securityReference.Authz.GroupMembership.Ldap.Domain, 7, true)
	}

	str += helpers.PrintFmtBool(`enabled: `, securityReference.Authz.Enabled, 5, true)

	return str
}
