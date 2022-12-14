// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.21.5
// source: proto/deploymentConfigurations/ci/Gcb.proto

package ci

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

// import "proto/deploymentConfigurations/ci/Permissions.proto";
// Configuration for the clouddriver microservice.
type Gcb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled  bool           `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Accounts []*GcbAccounts `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *Gcb) Reset() {
	*x = Gcb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_ci_Gcb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gcb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gcb) ProtoMessage() {}

func (x *Gcb) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_ci_Gcb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gcb.ProtoReflect.Descriptor instead.
func (*Gcb) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_ci_Gcb_proto_rawDescGZIP(), []int{0}
}

func (x *Gcb) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Gcb) GetAccounts() []*GcbAccounts {
	if x != nil {
		return x.Accounts
	}
	return nil
}

type GcbAccounts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Permission       string `protobuf:"bytes,2,opt,name=permission,proto3" json:"permission,omitempty"`
	Project          string `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	SubscriptionName string `protobuf:"bytes,4,opt,name=subscriptionName,proto3" json:"subscriptionName,omitempty"`
	JsonKey          string `protobuf:"bytes,5,opt,name=jsonKey,proto3" json:"jsonKey,omitempty"`
}

func (x *GcbAccounts) Reset() {
	*x = GcbAccounts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_ci_Gcb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcbAccounts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcbAccounts) ProtoMessage() {}

func (x *GcbAccounts) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_ci_Gcb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcbAccounts.ProtoReflect.Descriptor instead.
func (*GcbAccounts) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_ci_Gcb_proto_rawDescGZIP(), []int{1}
}

func (x *GcbAccounts) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GcbAccounts) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *GcbAccounts) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *GcbAccounts) GetSubscriptionName() string {
	if x != nil {
		return x.SubscriptionName
	}
	return ""
}

func (x *GcbAccounts) GetJsonKey() string {
	if x != nil {
		return x.JsonKey
	}
	return ""
}

var File_proto_deploymentConfigurations_ci_Gcb_proto protoreflect.FileDescriptor

var file_proto_deploymentConfigurations_ci_Gcb_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x63, 0x69, 0x2f, 0x47, 0x63, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x69, 0x22, 0x52, 0x0a, 0x03, 0x47, 0x63, 0x62, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x31, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x63, 0x69, 0x2e, 0x47, 0x63, 0x62, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0xa1, 0x01, 0x0a, 0x0b,
	0x47, 0x63, 0x62, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x73, 0x6f, 0x6e, 0x4b, 0x65, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x73, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x42,
	0x24, 0x5a, 0x22, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x63, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_deploymentConfigurations_ci_Gcb_proto_rawDescOnce sync.Once
	file_proto_deploymentConfigurations_ci_Gcb_proto_rawDescData = file_proto_deploymentConfigurations_ci_Gcb_proto_rawDesc
)

func file_proto_deploymentConfigurations_ci_Gcb_proto_rawDescGZIP() []byte {
	file_proto_deploymentConfigurations_ci_Gcb_proto_rawDescOnce.Do(func() {
		file_proto_deploymentConfigurations_ci_Gcb_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_deploymentConfigurations_ci_Gcb_proto_rawDescData)
	})
	return file_proto_deploymentConfigurations_ci_Gcb_proto_rawDescData
}

var file_proto_deploymentConfigurations_ci_Gcb_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_deploymentConfigurations_ci_Gcb_proto_goTypes = []interface{}{
	(*Gcb)(nil),         // 0: proto.ci.Gcb
	(*GcbAccounts)(nil), // 1: proto.ci.GcbAccounts
}
var file_proto_deploymentConfigurations_ci_Gcb_proto_depIdxs = []int32{
	1, // 0: proto.ci.Gcb.accounts:type_name -> proto.ci.GcbAccounts
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_deploymentConfigurations_ci_Gcb_proto_init() }
func file_proto_deploymentConfigurations_ci_Gcb_proto_init() {
	if File_proto_deploymentConfigurations_ci_Gcb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_deploymentConfigurations_ci_Gcb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gcb); i {
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
		file_proto_deploymentConfigurations_ci_Gcb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcbAccounts); i {
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
			RawDescriptor: file_proto_deploymentConfigurations_ci_Gcb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_deploymentConfigurations_ci_Gcb_proto_goTypes,
		DependencyIndexes: file_proto_deploymentConfigurations_ci_Gcb_proto_depIdxs,
		MessageInfos:      file_proto_deploymentConfigurations_ci_Gcb_proto_msgTypes,
	}.Build()
	File_proto_deploymentConfigurations_ci_Gcb_proto = out.File
	file_proto_deploymentConfigurations_ci_Gcb_proto_rawDesc = nil
	file_proto_deploymentConfigurations_ci_Gcb_proto_goTypes = nil
	file_proto_deploymentConfigurations_ci_Gcb_proto_depIdxs = nil
}
