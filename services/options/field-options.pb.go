// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.1
// source: options/field-options.proto

package options

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AnnotationOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Is this annotation active
	Active bool `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
	// Custom annotation message to provide
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AnnotationOption) Reset() {
	*x = AnnotationOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_field_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnotationOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnotationOption) ProtoMessage() {}

func (x *AnnotationOption) ProtoReflect() protoreflect.Message {
	mi := &file_options_field_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnotationOption.ProtoReflect.Descriptor instead.
func (*AnnotationOption) Descriptor() ([]byte, []int) {
	return file_options_field_options_proto_rawDescGZIP(), []int{0}
}

func (x *AnnotationOption) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *AnnotationOption) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SdkTemplateOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the service endpoint allows anonymous (no auth token necessary) authentication
	// This is used by the `protoc-gen-trinsic-sdk` plugin for metadata.
	Anonymous bool `protobuf:"varint,1,opt,name=anonymous,proto3" json:"anonymous,omitempty"`
	// Whether the SDK template generator should ignore this method. This method will
	// be wrapped manually.
	Ignore bool `protobuf:"varint,2,opt,name=ignore,proto3" json:"ignore,omitempty"`
	// Whether the SDK template generator should generate this method without arguments, eg
	// ProviderService.GetEcosystemInfo() where the request object is empty
	NoArguments bool `protobuf:"varint,3,opt,name=no_arguments,json=noArguments,proto3" json:"no_arguments,omitempty"`
	// This endpoint is experimental. Consider it in beta, so documentation may be incomplete or incorrect.
	Experimental *AnnotationOption `protobuf:"bytes,4,opt,name=experimental,proto3" json:"experimental,omitempty"`
	// This endpoint is deprecated. It will be removed in the future.
	Deprecated *AnnotationOption `protobuf:"bytes,5,opt,name=deprecated,proto3" json:"deprecated,omitempty"`
}

func (x *SdkTemplateOption) Reset() {
	*x = SdkTemplateOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_field_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SdkTemplateOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SdkTemplateOption) ProtoMessage() {}

func (x *SdkTemplateOption) ProtoReflect() protoreflect.Message {
	mi := &file_options_field_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SdkTemplateOption.ProtoReflect.Descriptor instead.
func (*SdkTemplateOption) Descriptor() ([]byte, []int) {
	return file_options_field_options_proto_rawDescGZIP(), []int{1}
}

func (x *SdkTemplateOption) GetAnonymous() bool {
	if x != nil {
		return x.Anonymous
	}
	return false
}

func (x *SdkTemplateOption) GetIgnore() bool {
	if x != nil {
		return x.Ignore
	}
	return false
}

func (x *SdkTemplateOption) GetNoArguments() bool {
	if x != nil {
		return x.NoArguments
	}
	return false
}

func (x *SdkTemplateOption) GetExperimental() *AnnotationOption {
	if x != nil {
		return x.Experimental
	}
	return nil
}

func (x *SdkTemplateOption) GetDeprecated() *AnnotationOption {
	if x != nil {
		return x.Deprecated
	}
	return nil
}

var file_options_field_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         60000,
		Name:          "services.options.optional",
		Tag:           "varint,60000,opt,name=optional",
		Filename:      "options/field-options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*SdkTemplateOption)(nil),
		Field:         60001,
		Name:          "services.options.sdk_template_option",
		Tag:           "bytes,60001,opt,name=sdk_template_option",
		Filename:      "options/field-options.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// Whether field is optional in Trinsic's backend.
	// This is not the same as an `optional` protobuf label;
	// it only impacts documentation generation for the field.
	//
	// optional bool optional = 60000;
	E_Optional = &file_options_field_options_proto_extTypes[0]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional services.options.SdkTemplateOption sdk_template_option = 60001;
	E_SdkTemplateOption = &file_options_field_options_proto_extTypes[1]
)

var File_options_field_options_proto protoreflect.FileDescriptor

var file_options_field_options_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2d,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x44, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xf8, 0x01, 0x0a, 0x11, 0x53, 0x64, 0x6b, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x67, 0x6e, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x67, 0x6e,
	0x6f, 0x72, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x6f, 0x5f, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x6e, 0x6f, 0x41, 0x72, 0x67,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x46, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0c, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x12, 0x42,
	0x0a, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x3a, 0x3b, 0x0a, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe0, 0xd4,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x3a,
	0x75, 0x0a, 0x13, 0x73, 0x64, 0x6b, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe1, 0xd4, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x53, 0x64, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x11, 0x73, 0x64, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x5b, 0x0a, 0x21, 0x74, 0x72, 0x69, 0x6e, 0x73, 0x69,
	0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x01, 0x5a, 0x10, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xaa,
	0x02, 0x21, 0x54, 0x72, 0x69, 0x6e, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_options_field_options_proto_rawDescOnce sync.Once
	file_options_field_options_proto_rawDescData = file_options_field_options_proto_rawDesc
)

func file_options_field_options_proto_rawDescGZIP() []byte {
	file_options_field_options_proto_rawDescOnce.Do(func() {
		file_options_field_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_options_field_options_proto_rawDescData)
	})
	return file_options_field_options_proto_rawDescData
}

var file_options_field_options_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_options_field_options_proto_goTypes = []interface{}{
	(*AnnotationOption)(nil),           // 0: services.options.AnnotationOption
	(*SdkTemplateOption)(nil),          // 1: services.options.SdkTemplateOption
	(*descriptorpb.FieldOptions)(nil),  // 2: google.protobuf.FieldOptions
	(*descriptorpb.MethodOptions)(nil), // 3: google.protobuf.MethodOptions
}
var file_options_field_options_proto_depIdxs = []int32{
	0, // 0: services.options.SdkTemplateOption.experimental:type_name -> services.options.AnnotationOption
	0, // 1: services.options.SdkTemplateOption.deprecated:type_name -> services.options.AnnotationOption
	2, // 2: services.options.optional:extendee -> google.protobuf.FieldOptions
	3, // 3: services.options.sdk_template_option:extendee -> google.protobuf.MethodOptions
	1, // 4: services.options.sdk_template_option:type_name -> services.options.SdkTemplateOption
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	4, // [4:5] is the sub-list for extension type_name
	2, // [2:4] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_options_field_options_proto_init() }
func file_options_field_options_proto_init() {
	if File_options_field_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_options_field_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnotationOption); i {
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
		file_options_field_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SdkTemplateOption); i {
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
			RawDescriptor: file_options_field_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_options_field_options_proto_goTypes,
		DependencyIndexes: file_options_field_options_proto_depIdxs,
		MessageInfos:      file_options_field_options_proto_msgTypes,
		ExtensionInfos:    file_options_field_options_proto_extTypes,
	}.Build()
	File_options_field_options_proto = out.File
	file_options_field_options_proto_rawDesc = nil
	file_options_field_options_proto_goTypes = nil
	file_options_field_options_proto_depIdxs = nil
}
