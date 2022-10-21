// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.21.5
// source: proto/deploymentConfigurations/security/Oauth2.proto

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
type Oauth2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled  bool          `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Client   *Oauth2Client `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	Provider string        `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *Oauth2) Reset() {
	*x = Oauth2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_security_Oauth2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Oauth2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Oauth2) ProtoMessage() {}

func (x *Oauth2) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_security_Oauth2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Oauth2.ProtoReflect.Descriptor instead.
func (*Oauth2) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_security_Oauth2_proto_rawDescGZIP(), []int{0}
}

func (x *Oauth2) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Oauth2) GetClient() *Oauth2Client {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *Oauth2) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

type Oauth2Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId                   string `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	ClientSecret               string `protobuf:"bytes,2,opt,name=clientSecret,proto3" json:"clientSecret,omitempty"`
	AccessTokenUri             string `protobuf:"bytes,3,opt,name=accessTokenUri,proto3" json:"accessTokenUri,omitempty"`
	UserAuthorizationUri       string `protobuf:"bytes,4,opt,name=userAuthorizationUri,proto3" json:"userAuthorizationUri,omitempty"`
	ClientAuthenticationScheme string `protobuf:"bytes,5,opt,name=clientAuthenticationScheme,proto3" json:"clientAuthenticationScheme,omitempty"`
	Scope                      string `protobuf:"bytes,6,opt,name=scope,proto3" json:"scope,omitempty"`
	PreEstablishedRedirectUri  string `protobuf:"bytes,7,opt,name=preEstablishedRedirectUri,proto3" json:"preEstablishedRedirectUri,omitempty"`
	UseCurrentUri              bool   `protobuf:"varint,8,opt,name=useCurrentUri,proto3" json:"useCurrentUri,omitempty"`
}

func (x *Oauth2Client) Reset() {
	*x = Oauth2Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_security_Oauth2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Oauth2Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Oauth2Client) ProtoMessage() {}

func (x *Oauth2Client) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_security_Oauth2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Oauth2Client.ProtoReflect.Descriptor instead.
func (*Oauth2Client) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_security_Oauth2_proto_rawDescGZIP(), []int{1}
}

func (x *Oauth2Client) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Oauth2Client) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *Oauth2Client) GetAccessTokenUri() string {
	if x != nil {
		return x.AccessTokenUri
	}
	return ""
}

func (x *Oauth2Client) GetUserAuthorizationUri() string {
	if x != nil {
		return x.UserAuthorizationUri
	}
	return ""
}

func (x *Oauth2Client) GetClientAuthenticationScheme() string {
	if x != nil {
		return x.ClientAuthenticationScheme
	}
	return ""
}

func (x *Oauth2Client) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *Oauth2Client) GetPreEstablishedRedirectUri() string {
	if x != nil {
		return x.PreEstablishedRedirectUri
	}
	return ""
}

func (x *Oauth2Client) GetUseCurrentUri() bool {
	if x != nil {
		return x.UseCurrentUri
	}
	return false
}

type Oauth2UserInfoRequirements struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Need to revisit this struct
	Kvpair string `protobuf:"bytes,1,opt,name=kvpair,proto3" json:"kvpair,omitempty"`
}

func (x *Oauth2UserInfoRequirements) Reset() {
	*x = Oauth2UserInfoRequirements{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_security_Oauth2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Oauth2UserInfoRequirements) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Oauth2UserInfoRequirements) ProtoMessage() {}

func (x *Oauth2UserInfoRequirements) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_security_Oauth2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Oauth2UserInfoRequirements.ProtoReflect.Descriptor instead.
func (*Oauth2UserInfoRequirements) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_security_Oauth2_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Oauth2UserInfoRequirements) GetKvpair() string {
	if x != nil {
		return x.Kvpair
	}
	return ""
}

type Oauth2Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfoUri string `protobuf:"bytes,1,opt,name=userInfoUri,proto3" json:"userInfoUri,omitempty"`
}

func (x *Oauth2Resource) Reset() {
	*x = Oauth2Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_security_Oauth2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Oauth2Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Oauth2Resource) ProtoMessage() {}

func (x *Oauth2Resource) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_security_Oauth2_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Oauth2Resource.ProtoReflect.Descriptor instead.
func (*Oauth2Resource) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_security_Oauth2_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Oauth2Resource) GetUserInfoUri() string {
	if x != nil {
		return x.UserInfoUri
	}
	return ""
}

type Oauth2UserInfoMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email     string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Firstname string `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Username  string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *Oauth2UserInfoMapping) Reset() {
	*x = Oauth2UserInfoMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_security_Oauth2_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Oauth2UserInfoMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Oauth2UserInfoMapping) ProtoMessage() {}

func (x *Oauth2UserInfoMapping) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_security_Oauth2_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Oauth2UserInfoMapping.ProtoReflect.Descriptor instead.
func (*Oauth2UserInfoMapping) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_security_Oauth2_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Oauth2UserInfoMapping) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Oauth2UserInfoMapping) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *Oauth2UserInfoMapping) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *Oauth2UserInfoMapping) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_proto_deploymentConfigurations_security_Oauth2_proto protoreflect.FileDescriptor

var file_proto_deploymentConfigurations_security_Oauth2_proto_rawDesc = []byte{
	0x0a, 0x34, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x32,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x22, 0xd1, 0x02, 0x0a, 0x06, 0x4f, 0x61, 0x75, 0x74, 0x68,
	0x32, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x06, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x4f, 0x61, 0x75,
	0x74, 0x68, 0x32, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x1a, 0x2e, 0x0a,
	0x14, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6b, 0x76, 0x70, 0x61, 0x69, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6b, 0x76, 0x70, 0x61, 0x69, 0x72, 0x1a, 0x2c, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x55, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x55, 0x72, 0x69, 0x1a, 0x7d, 0x0a, 0x0f, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xe4, 0x02, 0x0a, 0x0c, 0x4f,
	0x61, 0x75, 0x74, 0x68, 0x32, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x55, 0x72, 0x69, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x55, 0x72, 0x69, 0x12, 0x32, 0x0a, 0x14, 0x75, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x14, 0x75, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x69, 0x12, 0x3e, 0x0a, 0x1a, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1a, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x3c, 0x0a,
	0x19, 0x70, 0x72, 0x65, 0x45, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x52,
	0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x69, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x19, 0x70, 0x72, 0x65, 0x45, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x69, 0x12, 0x24, 0x0a, 0x0d, 0x75,
	0x73, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x72, 0x69, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x72,
	0x69, 0x42, 0x2a, 0x5a, 0x28, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_deploymentConfigurations_security_Oauth2_proto_rawDescOnce sync.Once
	file_proto_deploymentConfigurations_security_Oauth2_proto_rawDescData = file_proto_deploymentConfigurations_security_Oauth2_proto_rawDesc
)

func file_proto_deploymentConfigurations_security_Oauth2_proto_rawDescGZIP() []byte {
	file_proto_deploymentConfigurations_security_Oauth2_proto_rawDescOnce.Do(func() {
		file_proto_deploymentConfigurations_security_Oauth2_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_deploymentConfigurations_security_Oauth2_proto_rawDescData)
	})
	return file_proto_deploymentConfigurations_security_Oauth2_proto_rawDescData
}

var file_proto_deploymentConfigurations_security_Oauth2_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_deploymentConfigurations_security_Oauth2_proto_goTypes = []interface{}{
	(*Oauth2)(nil),                     // 0: proto.security.Oauth2
	(*Oauth2Client)(nil),               // 1: proto.security.Oauth2Client
	(*Oauth2UserInfoRequirements)(nil), // 2: proto.security.Oauth2.userInfoRequirements
	(*Oauth2Resource)(nil),             // 3: proto.security.Oauth2.resource
	(*Oauth2UserInfoMapping)(nil),      // 4: proto.security.Oauth2.userInfoMapping
}
var file_proto_deploymentConfigurations_security_Oauth2_proto_depIdxs = []int32{
	1, // 0: proto.security.Oauth2.client:type_name -> proto.security.Oauth2Client
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_deploymentConfigurations_security_Oauth2_proto_init() }
func file_proto_deploymentConfigurations_security_Oauth2_proto_init() {
	if File_proto_deploymentConfigurations_security_Oauth2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_deploymentConfigurations_security_Oauth2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Oauth2); i {
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
		file_proto_deploymentConfigurations_security_Oauth2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Oauth2Client); i {
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
		file_proto_deploymentConfigurations_security_Oauth2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Oauth2UserInfoRequirements); i {
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
		file_proto_deploymentConfigurations_security_Oauth2_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Oauth2Resource); i {
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
		file_proto_deploymentConfigurations_security_Oauth2_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Oauth2UserInfoMapping); i {
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
			RawDescriptor: file_proto_deploymentConfigurations_security_Oauth2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_deploymentConfigurations_security_Oauth2_proto_goTypes,
		DependencyIndexes: file_proto_deploymentConfigurations_security_Oauth2_proto_depIdxs,
		MessageInfos:      file_proto_deploymentConfigurations_security_Oauth2_proto_msgTypes,
	}.Build()
	File_proto_deploymentConfigurations_security_Oauth2_proto = out.File
	file_proto_deploymentConfigurations_security_Oauth2_proto_rawDesc = nil
	file_proto_deploymentConfigurations_security_Oauth2_proto_goTypes = nil
	file_proto_deploymentConfigurations_security_Oauth2_proto_depIdxs = nil
}