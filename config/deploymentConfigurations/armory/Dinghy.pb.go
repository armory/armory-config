// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.14.0
// source: proto/deploymentConfigurations/armory/Dinghy.proto

package armory

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

type Dinghy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled                     bool       `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	TemplateOrg                 string     `protobuf:"bytes,2,opt,name=templateOrg,proto3" json:"templateOrg,omitempty"`
	TemplateRepo                string     `protobuf:"bytes,3,opt,name=templateRepo,proto3" json:"templateRepo,omitempty"`
	GithubToken                 string     `protobuf:"bytes,4,opt,name=githubToken,proto3" json:"githubToken,omitempty"`
	GithubEndpoint              string     `protobuf:"bytes,5,opt,name=githubEndpoint,proto3" json:"githubEndpoint,omitempty"`
	RepositoryRawdataProcessing bool       `protobuf:"varint,6,opt,name=repositoryRawdataProcessing,proto3" json:"repositoryRawdataProcessing,omitempty"`
	DinghyFilename              string     `protobuf:"bytes,7,opt,name=dinghyFilename,proto3" json:"dinghyFilename,omitempty"`
	AutoLockPipelines           bool       `protobuf:"varint,8,opt,name=autoLockPipelines,proto3" json:"autoLockPipelines,omitempty"`
	FiatUser                    string     `protobuf:"bytes,9,opt,name=fiatUser,proto3" json:"fiatUser,omitempty"`
	Notifiers                   *Notifiers `protobuf:"bytes,10,opt,name=notifiers,proto3" json:"notifiers,omitempty"`
}

func (x *Dinghy) Reset() {
	*x = Dinghy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_armory_Dinghy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dinghy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dinghy) ProtoMessage() {}

func (x *Dinghy) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_armory_Dinghy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dinghy.ProtoReflect.Descriptor instead.
func (*Dinghy) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_armory_Dinghy_proto_rawDescGZIP(), []int{0}
}

func (x *Dinghy) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Dinghy) GetTemplateOrg() string {
	if x != nil {
		return x.TemplateOrg
	}
	return ""
}

func (x *Dinghy) GetTemplateRepo() string {
	if x != nil {
		return x.TemplateRepo
	}
	return ""
}

func (x *Dinghy) GetGithubToken() string {
	if x != nil {
		return x.GithubToken
	}
	return ""
}

func (x *Dinghy) GetGithubEndpoint() string {
	if x != nil {
		return x.GithubEndpoint
	}
	return ""
}

func (x *Dinghy) GetRepositoryRawdataProcessing() bool {
	if x != nil {
		return x.RepositoryRawdataProcessing
	}
	return false
}

func (x *Dinghy) GetDinghyFilename() string {
	if x != nil {
		return x.DinghyFilename
	}
	return ""
}

func (x *Dinghy) GetAutoLockPipelines() bool {
	if x != nil {
		return x.AutoLockPipelines
	}
	return false
}

func (x *Dinghy) GetFiatUser() string {
	if x != nil {
		return x.FiatUser
	}
	return ""
}

func (x *Dinghy) GetNotifiers() *Notifiers {
	if x != nil {
		return x.Notifiers
	}
	return nil
}

type Notifiers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slack *Slack `protobuf:"bytes,1,opt,name=slack,proto3" json:"slack,omitempty"`
}

func (x *Notifiers) Reset() {
	*x = Notifiers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_armory_Dinghy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notifiers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notifiers) ProtoMessage() {}

func (x *Notifiers) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_armory_Dinghy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notifiers.ProtoReflect.Descriptor instead.
func (*Notifiers) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_armory_Dinghy_proto_rawDescGZIP(), []int{1}
}

func (x *Notifiers) GetSlack() *Slack {
	if x != nil {
		return x.Slack
	}
	return nil
}

type Slack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *Slack) Reset() {
	*x = Slack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_armory_Dinghy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Slack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Slack) ProtoMessage() {}

func (x *Slack) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_armory_Dinghy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Slack.ProtoReflect.Descriptor instead.
func (*Slack) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_armory_Dinghy_proto_rawDescGZIP(), []int{2}
}

func (x *Slack) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

var File_proto_deploymentConfigurations_armory_Dinghy_proto protoreflect.FileDescriptor

var file_proto_deploymentConfigurations_armory_Dinghy_proto_rawDesc = []byte{
	0x0a, 0x32, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x79, 0x2f, 0x44, 0x69, 0x6e, 0x67, 0x68, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x72, 0x6d, 0x6f,
	0x72, 0x79, 0x22, 0x9d, 0x03, 0x0a, 0x06, 0x44, 0x69, 0x6e, 0x67, 0x68, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x20, 0x0a,
	0x0b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x26, 0x0a, 0x0e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x1b, 0x72, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1b, 0x72, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x69, 0x6e,
	0x67, 0x68, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x64, 0x69, 0x6e, 0x67, 0x68, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x6f, 0x4c, 0x6f, 0x63, 0x6b, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x61, 0x75,
	0x74, 0x6f, 0x4c, 0x6f, 0x63, 0x6b, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x61, 0x74, 0x55, 0x73, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x61, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x09, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x52, 0x09, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x73, 0x22, 0x36, 0x0a, 0x09, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x12,
	0x29, 0x0a, 0x05, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x53, 0x6c,
	0x61, 0x63, 0x6b, 0x52, 0x05, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x22, 0x21, 0x0a, 0x05, 0x53, 0x6c,
	0x61, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42, 0x28, 0x5a,
	0x26, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_deploymentConfigurations_armory_Dinghy_proto_rawDescOnce sync.Once
	file_proto_deploymentConfigurations_armory_Dinghy_proto_rawDescData = file_proto_deploymentConfigurations_armory_Dinghy_proto_rawDesc
)

func file_proto_deploymentConfigurations_armory_Dinghy_proto_rawDescGZIP() []byte {
	file_proto_deploymentConfigurations_armory_Dinghy_proto_rawDescOnce.Do(func() {
		file_proto_deploymentConfigurations_armory_Dinghy_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_deploymentConfigurations_armory_Dinghy_proto_rawDescData)
	})
	return file_proto_deploymentConfigurations_armory_Dinghy_proto_rawDescData
}

var file_proto_deploymentConfigurations_armory_Dinghy_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_deploymentConfigurations_armory_Dinghy_proto_goTypes = []interface{}{
	(*Dinghy)(nil),    // 0: proto.armory.Dinghy
	(*Notifiers)(nil), // 1: proto.armory.Notifiers
	(*Slack)(nil),     // 2: proto.armory.Slack
}
var file_proto_deploymentConfigurations_armory_Dinghy_proto_depIdxs = []int32{
	1, // 0: proto.armory.Dinghy.notifiers:type_name -> proto.armory.Notifiers
	2, // 1: proto.armory.Notifiers.slack:type_name -> proto.armory.Slack
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_deploymentConfigurations_armory_Dinghy_proto_init() }
func file_proto_deploymentConfigurations_armory_Dinghy_proto_init() {
	if File_proto_deploymentConfigurations_armory_Dinghy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_deploymentConfigurations_armory_Dinghy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dinghy); i {
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
		file_proto_deploymentConfigurations_armory_Dinghy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notifiers); i {
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
		file_proto_deploymentConfigurations_armory_Dinghy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Slack); i {
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
			RawDescriptor: file_proto_deploymentConfigurations_armory_Dinghy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_deploymentConfigurations_armory_Dinghy_proto_goTypes,
		DependencyIndexes: file_proto_deploymentConfigurations_armory_Dinghy_proto_depIdxs,
		MessageInfos:      file_proto_deploymentConfigurations_armory_Dinghy_proto_msgTypes,
	}.Build()
	File_proto_deploymentConfigurations_armory_Dinghy_proto = out.File
	file_proto_deploymentConfigurations_armory_Dinghy_proto_rawDesc = nil
	file_proto_deploymentConfigurations_armory_Dinghy_proto_goTypes = nil
	file_proto_deploymentConfigurations_armory_Dinghy_proto_depIdxs = nil
}
