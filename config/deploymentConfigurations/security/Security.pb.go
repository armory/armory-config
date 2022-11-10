// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.14.0
// source: proto/deploymentConfigurations/security/Security.proto

package security

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Configuration for the clouddriver microservice.
type Security struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiSecurity *ApiSecurity `protobuf:"bytes,1,opt,name=apiSecurity,proto3" json:"apiSecurity,omitempty"`
	UiSecurity  *UiSecurity  `protobuf:"bytes,2,opt,name=uiSecurity,proto3" json:"uiSecurity,omitempty"`
	Authn       *Authn       `protobuf:"bytes,3,opt,name=authn,proto3" json:"authn,omitempty"`
	Authz       *Authz       `protobuf:"bytes,4,opt,name=authz,proto3" json:"authz,omitempty"`
}

func (x *Security) Reset() {
	*x = Security{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_security_Security_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Security) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Security) ProtoMessage() {}

func (x *Security) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_security_Security_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Security.ProtoReflect.Descriptor instead.
func (*Security) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_security_Security_proto_rawDescGZIP(), []int{0}
}

func (x *Security) GetApiSecurity() *ApiSecurity {
	if x != nil {
		return x.ApiSecurity
	}
	return nil
}

func (x *Security) GetUiSecurity() *UiSecurity {
	if x != nil {
		return x.UiSecurity
	}
	return nil
}

func (x *Security) GetAuthn() *Authn {
	if x != nil {
		return x.Authn
	}
	return nil
}

func (x *Security) GetAuthz() *Authz {
	if x != nil {
		return x.Authz
	}
	return nil
}

type ApiSecurity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ssl             *Ssl   `protobuf:"bytes,1,opt,name=ssl,proto3" json:"ssl,omitempty"`
	OverrideBaseUrl string `protobuf:"bytes,2,opt,name=overrideBaseUrl,proto3" json:"overrideBaseUrl,omitempty"`
}

func (x *ApiSecurity) Reset() {
	*x = ApiSecurity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_security_Security_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiSecurity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiSecurity) ProtoMessage() {}

func (x *ApiSecurity) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_security_Security_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiSecurity.ProtoReflect.Descriptor instead.
func (*ApiSecurity) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_security_Security_proto_rawDescGZIP(), []int{1}
}

func (x *ApiSecurity) GetSsl() *Ssl {
	if x != nil {
		return x.Ssl
	}
	return nil
}

func (x *ApiSecurity) GetOverrideBaseUrl() string {
	if x != nil {
		return x.OverrideBaseUrl
	}
	return ""
}

type UiSecurity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ssl             *Ssl   `protobuf:"bytes,1,opt,name=ssl,proto3" json:"ssl,omitempty"`
	OverrideBaseUrl string `protobuf:"bytes,2,opt,name=overrideBaseUrl,proto3" json:"overrideBaseUrl,omitempty"`
}

func (x *UiSecurity) Reset() {
	*x = UiSecurity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_security_Security_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UiSecurity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UiSecurity) ProtoMessage() {}

func (x *UiSecurity) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_security_Security_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UiSecurity.ProtoReflect.Descriptor instead.
func (*UiSecurity) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_security_Security_proto_rawDescGZIP(), []int{2}
}

func (x *UiSecurity) GetSsl() *Ssl {
	if x != nil {
		return x.Ssl
	}
	return nil
}

func (x *UiSecurity) GetOverrideBaseUrl() string {
	if x != nil {
		return x.OverrideBaseUrl
	}
	return ""
}

type Ssl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled            bool   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	KeyAlias           string `protobuf:"bytes,2,opt,name=keyAlias,proto3" json:"keyAlias,omitempty"`
	KeyStore           string `protobuf:"bytes,3,opt,name=keyStore,proto3" json:"keyStore,omitempty"`
	KeyStoreType       string `protobuf:"bytes,4,opt,name=keyStoreType,proto3" json:"keyStoreType,omitempty"`
	KeyStorePassword   string `protobuf:"bytes,5,opt,name=keyStorePassword,proto3" json:"keyStorePassword,omitempty"`
	TrustStore         string `protobuf:"bytes,6,opt,name=trustStore,proto3" json:"trustStore,omitempty"`
	TrustStoreType     string `protobuf:"bytes,7,opt,name=trustStoreType,proto3" json:"trustStoreType,omitempty"`
	TrustStorePassword string `protobuf:"bytes,8,opt,name=trustStorePassword,proto3" json:"trustStorePassword,omitempty"`
}

func (x *Ssl) Reset() {
	*x = Ssl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_security_Security_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ssl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ssl) ProtoMessage() {}

func (x *Ssl) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_security_Security_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ssl.ProtoReflect.Descriptor instead.
func (*Ssl) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_security_Security_proto_rawDescGZIP(), []int{3}
}

func (x *Ssl) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Ssl) GetKeyAlias() string {
	if x != nil {
		return x.KeyAlias
	}
	return ""
}

func (x *Ssl) GetKeyStore() string {
	if x != nil {
		return x.KeyStore
	}
	return ""
}

func (x *Ssl) GetKeyStoreType() string {
	if x != nil {
		return x.KeyStoreType
	}
	return ""
}

func (x *Ssl) GetKeyStorePassword() string {
	if x != nil {
		return x.KeyStorePassword
	}
	return ""
}

func (x *Ssl) GetTrustStore() string {
	if x != nil {
		return x.TrustStore
	}
	return ""
}

func (x *Ssl) GetTrustStoreType() string {
	if x != nil {
		return x.TrustStoreType
	}
	return ""
}

func (x *Ssl) GetTrustStorePassword() string {
	if x != nil {
		return x.TrustStorePassword
	}
	return ""
}

type Authn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Oauth2  *Oauth2 `protobuf:"bytes,1,opt,name=oauth2,proto3" json:"oauth2,omitempty"`
	Saml    *Saml   `protobuf:"bytes,2,opt,name=saml,proto3" json:"saml,omitempty"`
	Ldap    *Ldap   `protobuf:"bytes,3,opt,name=ldap,proto3" json:"ldap,omitempty"`
	X509    *X509   `protobuf:"bytes,4,opt,name=x509,proto3" json:"x509,omitempty"`
	Iap     *Iap    `protobuf:"bytes,5,opt,name=iap,proto3" json:"iap,omitempty"`
	Enabled bool    `protobuf:"varint,6,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *Authn) Reset() {
	*x = Authn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_security_Security_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authn) ProtoMessage() {}

func (x *Authn) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_security_Security_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authn.ProtoReflect.Descriptor instead.
func (*Authn) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_security_Security_proto_rawDescGZIP(), []int{4}
}

func (x *Authn) GetOauth2() *Oauth2 {
	if x != nil {
		return x.Oauth2
	}
	return nil
}

func (x *Authn) GetSaml() *Saml {
	if x != nil {
		return x.Saml
	}
	return nil
}

func (x *Authn) GetLdap() *Ldap {
	if x != nil {
		return x.Ldap
	}
	return nil
}

func (x *Authn) GetX509() *X509 {
	if x != nil {
		return x.X509
	}
	return nil
}

func (x *Authn) GetIap() *Iap {
	if x != nil {
		return x.Iap
	}
	return nil
}

func (x *Authn) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type Authz struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupMembership *GroupMembership `protobuf:"bytes,1,opt,name=groupMembership,proto3" json:"groupMembership,omitempty"`
	Enabled         bool             `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *Authz) Reset() {
	*x = Authz{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_security_Security_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authz) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authz) ProtoMessage() {}

func (x *Authz) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_security_Security_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authz.ProtoReflect.Descriptor instead.
func (*Authz) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_security_Security_proto_rawDescGZIP(), []int{5}
}

func (x *Authz) GetGroupMembership() *GroupMembership {
	if x != nil {
		return x.GroupMembership
	}
	return nil
}

func (x *Authz) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type GroupMembership struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string    `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Google  *RoleType `protobuf:"bytes,2,opt,name=google,proto3" json:"google,omitempty"`
	Github  *RoleType `protobuf:"bytes,3,opt,name=github,proto3" json:"github,omitempty"`
	File    *RoleType `protobuf:"bytes,4,opt,name=file,proto3" json:"file,omitempty"`
	Ldap    *RoleType `protobuf:"bytes,5,opt,name=ldap,proto3" json:"ldap,omitempty"`
	Enabled bool      `protobuf:"varint,6,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *GroupMembership) Reset() {
	*x = GroupMembership{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_security_Security_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupMembership) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupMembership) ProtoMessage() {}

func (x *GroupMembership) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_security_Security_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupMembership.ProtoReflect.Descriptor instead.
func (*GroupMembership) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_security_Security_proto_rawDescGZIP(), []int{6}
}

func (x *GroupMembership) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *GroupMembership) GetGoogle() *RoleType {
	if x != nil {
		return x.Google
	}
	return nil
}

func (x *GroupMembership) GetGithub() *RoleType {
	if x != nil {
		return x.Github
	}
	return nil
}

func (x *GroupMembership) GetFile() *RoleType {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *GroupMembership) GetLdap() *RoleType {
	if x != nil {
		return x.Ldap
	}
	return nil
}

func (x *GroupMembership) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type RoleType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleProviderType string `protobuf:"bytes,1,opt,name=roleProviderType,proto3" json:"roleProviderType,omitempty"`
	Path             string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	CredentialPath   string `protobuf:"bytes,3,opt,name=credentialPath,proto3" json:"credentialPath,omitempty"`
	AdminUsername    string `protobuf:"bytes,4,opt,name=adminUsername,proto3" json:"adminUsername,omitempty"`
	Domain           string `protobuf:"bytes,5,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *RoleType) Reset() {
	*x = RoleType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_security_Security_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleType) ProtoMessage() {}

func (x *RoleType) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_security_Security_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleType.ProtoReflect.Descriptor instead.
func (*RoleType) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_security_Security_proto_rawDescGZIP(), []int{7}
}

func (x *RoleType) GetRoleProviderType() string {
	if x != nil {
		return x.RoleProviderType
	}
	return ""
}

func (x *RoleType) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *RoleType) GetCredentialPath() string {
	if x != nil {
		return x.CredentialPath
	}
	return ""
}

func (x *RoleType) GetAdminUsername() string {
	if x != nil {
		return x.AdminUsername
	}
	return ""
}

func (x *RoleType) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

var File_proto_deploymentConfigurations_security_Security_proto protoreflect.FileDescriptor

var file_proto_deploymentConfigurations_security_Security_proto_rawDesc = []byte{
	0x0a, 0x36, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x1a, 0x34, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x2f, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x53, 0x61, 0x6d, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x32, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x4c, 0x64, 0x61, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f,
	0x58, 0x35, 0x30, 0x39, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x2f, 0x49, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x01,
	0x0a, 0x08, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x3d, 0x0a, 0x0b, 0x61, 0x70,
	0x69, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x2e, 0x41, 0x70, 0x69, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x0b, 0x61, 0x70,
	0x69, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x3a, 0x0a, 0x0a, 0x75, 0x69, 0x53,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x55,
	0x69, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x75, 0x69, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x52, 0x05, 0x61, 0x75, 0x74,
	0x68, 0x6e, 0x12, 0x2b, 0x0a, 0x05, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x52, 0x05, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x22,
	0x5e, 0x0a, 0x0b, 0x41, 0x70, 0x69, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x25,
	0x0a, 0x03, 0x73, 0x73, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x73, 0x6c,
	0x52, 0x03, 0x73, 0x73, 0x6c, 0x12, 0x28, 0x0a, 0x0f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64,
	0x65, 0x42, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x42, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x22,
	0x5d, 0x0a, 0x0a, 0x55, 0x69, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x25, 0x0a,
	0x03, 0x73, 0x73, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x73, 0x6c, 0x52,
	0x03, 0x73, 0x73, 0x6c, 0x12, 0x28, 0x0a, 0x0f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65,
	0x42, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f,
	0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x42, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x9f,
	0x02, 0x0a, 0x03, 0x53, 0x73, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x6b, 0x65, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6b, 0x65, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6b, 0x65, 0x79, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6b, 0x65, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x10,
	0x6b, 0x65, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6b, 0x65, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x72, 0x75, 0x73,
	0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72,
	0x75, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x72, 0x75, 0x73,
	0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x2e, 0x0a, 0x12, 0x74, 0x72, 0x75, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x74, 0x72,
	0x75, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0xf6, 0x01, 0x0a, 0x05, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x12, 0x2e, 0x0a, 0x06, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x4f, 0x61, 0x75, 0x74,
	0x68, 0x32, 0x52, 0x06, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x12, 0x28, 0x0a, 0x04, 0x73, 0x61,
	0x6d, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x61, 0x6d, 0x6c, 0x52, 0x04,
	0x73, 0x61, 0x6d, 0x6c, 0x12, 0x28, 0x0a, 0x04, 0x6c, 0x64, 0x61, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x2e, 0x4c, 0x64, 0x61, 0x70, 0x52, 0x04, 0x6c, 0x64, 0x61, 0x70, 0x12, 0x28,
	0x0a, 0x04, 0x78, 0x35, 0x30, 0x39, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x58, 0x35,
	0x30, 0x39, 0x52, 0x04, 0x78, 0x35, 0x30, 0x39, 0x12, 0x25, 0x0a, 0x03, 0x69, 0x61, 0x70, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x49, 0x61, 0x70, 0x52, 0x03, 0x69, 0x61, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x6c, 0x0a, 0x05, 0x41, 0x75, 0x74,
	0x68, 0x7a, 0x12, 0x49, 0x0a, 0x0f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x0f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x85, 0x02, 0x0a, 0x0f, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x06, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x12, 0x2c, 0x0a, 0x04, 0x66, 0x69, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x6c, 0x64, 0x61, 0x70, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x6c, 0x64, 0x61, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22,
	0xb0, 0x01, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x10,
	0x72, 0x6f, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x6f, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x26, 0x0a, 0x0e,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x42, 0x2a, 0x5a, 0x28, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_deploymentConfigurations_security_Security_proto_rawDescOnce sync.Once
	file_proto_deploymentConfigurations_security_Security_proto_rawDescData = file_proto_deploymentConfigurations_security_Security_proto_rawDesc
)

func file_proto_deploymentConfigurations_security_Security_proto_rawDescGZIP() []byte {
	file_proto_deploymentConfigurations_security_Security_proto_rawDescOnce.Do(func() {
		file_proto_deploymentConfigurations_security_Security_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_deploymentConfigurations_security_Security_proto_rawDescData)
	})
	return file_proto_deploymentConfigurations_security_Security_proto_rawDescData
}

var file_proto_deploymentConfigurations_security_Security_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_deploymentConfigurations_security_Security_proto_goTypes = []interface{}{
	(*Security)(nil),        // 0: proto.security.Security
	(*ApiSecurity)(nil),     // 1: proto.security.ApiSecurity
	(*UiSecurity)(nil),      // 2: proto.security.UiSecurity
	(*Ssl)(nil),             // 3: proto.security.Ssl
	(*Authn)(nil),           // 4: proto.security.Authn
	(*Authz)(nil),           // 5: proto.security.Authz
	(*GroupMembership)(nil), // 6: proto.security.GroupMembership
	(*RoleType)(nil),        // 7: proto.security.RoleType
	(*Oauth2)(nil),          // 8: proto.security.Oauth2
	(*Saml)(nil),            // 9: proto.security.Saml
	(*Ldap)(nil),            // 10: proto.security.Ldap
	(*X509)(nil),            // 11: proto.security.X509
	(*Iap)(nil),             // 12: proto.security.Iap
}
var file_proto_deploymentConfigurations_security_Security_proto_depIdxs = []int32{
	1,  // 0: proto.security.Security.apiSecurity:type_name -> proto.security.ApiSecurity
	2,  // 1: proto.security.Security.uiSecurity:type_name -> proto.security.UiSecurity
	4,  // 2: proto.security.Security.authn:type_name -> proto.security.Authn
	5,  // 3: proto.security.Security.authz:type_name -> proto.security.Authz
	3,  // 4: proto.security.ApiSecurity.ssl:type_name -> proto.security.Ssl
	3,  // 5: proto.security.UiSecurity.ssl:type_name -> proto.security.Ssl
	8,  // 6: proto.security.Authn.oauth2:type_name -> proto.security.Oauth2
	9,  // 7: proto.security.Authn.saml:type_name -> proto.security.Saml
	10, // 8: proto.security.Authn.ldap:type_name -> proto.security.Ldap
	11, // 9: proto.security.Authn.x509:type_name -> proto.security.X509
	12, // 10: proto.security.Authn.iap:type_name -> proto.security.Iap
	6,  // 11: proto.security.Authz.groupMembership:type_name -> proto.security.GroupMembership
	7,  // 12: proto.security.GroupMembership.google:type_name -> proto.security.RoleType
	7,  // 13: proto.security.GroupMembership.github:type_name -> proto.security.RoleType
	7,  // 14: proto.security.GroupMembership.file:type_name -> proto.security.RoleType
	7,  // 15: proto.security.GroupMembership.ldap:type_name -> proto.security.RoleType
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_proto_deploymentConfigurations_security_Security_proto_init() }
func file_proto_deploymentConfigurations_security_Security_proto_init() {
	if File_proto_deploymentConfigurations_security_Security_proto != nil {
		return
	}
	file_proto_deploymentConfigurations_security_Oauth2_proto_init()
	file_proto_deploymentConfigurations_security_Saml_proto_init()
	file_proto_deploymentConfigurations_security_Ldap_proto_init()
	file_proto_deploymentConfigurations_security_X509_proto_init()
	file_proto_deploymentConfigurations_security_Iap_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_deploymentConfigurations_security_Security_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Security); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_deploymentConfigurations_security_Security_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiSecurity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_deploymentConfigurations_security_Security_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UiSecurity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_deploymentConfigurations_security_Security_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ssl); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_deploymentConfigurations_security_Security_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_deploymentConfigurations_security_Security_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authz); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_deploymentConfigurations_security_Security_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupMembership); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_deploymentConfigurations_security_Security_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_deploymentConfigurations_security_Security_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_deploymentConfigurations_security_Security_proto_goTypes,
		DependencyIndexes: file_proto_deploymentConfigurations_security_Security_proto_depIdxs,
		MessageInfos:      file_proto_deploymentConfigurations_security_Security_proto_msgTypes,
	}.Build()
	File_proto_deploymentConfigurations_security_Security_proto = out.File
	file_proto_deploymentConfigurations_security_Security_proto_rawDesc = nil
	file_proto_deploymentConfigurations_security_Security_proto_goTypes = nil
	file_proto_deploymentConfigurations_security_Security_proto_depIdxs = nil
}
