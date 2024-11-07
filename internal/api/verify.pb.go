// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: api/proto/verify.proto

package ssov1

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

type InquiryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eve Events `protobuf:"varint,1,opt,name=eve,proto3,enum=auth.Events" json:"eve,omitempty"`
}

func (x *InquiryReq) Reset() {
	*x = InquiryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_verify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InquiryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InquiryReq) ProtoMessage() {}

func (x *InquiryReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_verify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InquiryReq.ProtoReflect.Descriptor instead.
func (*InquiryReq) Descriptor() ([]byte, []int) {
	return file_api_proto_verify_proto_rawDescGZIP(), []int{0}
}

func (x *InquiryReq) GetEve() Events {
	if x != nil {
		return x.Eve
	}
	return Events_LOGIN
}

type CheckConfirmCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eve  Events `protobuf:"varint,1,opt,name=eve,proto3,enum=auth.Events" json:"eve,omitempty"`
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *CheckConfirmCodeReq) Reset() {
	*x = CheckConfirmCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_verify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckConfirmCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckConfirmCodeReq) ProtoMessage() {}

func (x *CheckConfirmCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_verify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckConfirmCodeReq.ProtoReflect.Descriptor instead.
func (*CheckConfirmCodeReq) Descriptor() ([]byte, []int) {
	return file_api_proto_verify_proto_rawDescGZIP(), []int{1}
}

func (x *CheckConfirmCodeReq) GetEve() Events {
	if x != nil {
		return x.Eve
	}
	return Events_LOGIN
}

func (x *CheckConfirmCodeReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type CheckConfirmCodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eve Events `protobuf:"varint,1,opt,name=eve,proto3,enum=auth.Events" json:"eve,omitempty"`
}

func (x *CheckConfirmCodeResp) Reset() {
	*x = CheckConfirmCodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_verify_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckConfirmCodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckConfirmCodeResp) ProtoMessage() {}

func (x *CheckConfirmCodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_verify_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckConfirmCodeResp.ProtoReflect.Descriptor instead.
func (*CheckConfirmCodeResp) Descriptor() ([]byte, []int) {
	return file_api_proto_verify_proto_rawDescGZIP(), []int{2}
}

func (x *CheckConfirmCodeResp) GetEve() Events {
	if x != nil {
		return x.Eve
	}
	return Events_LOGIN
}

type InquiryResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link string `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *InquiryResp) Reset() {
	*x = InquiryResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_verify_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InquiryResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InquiryResp) ProtoMessage() {}

func (x *InquiryResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_verify_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InquiryResp.ProtoReflect.Descriptor instead.
func (*InquiryResp) Descriptor() ([]byte, []int) {
	return file_api_proto_verify_proto_rawDescGZIP(), []int{3}
}

func (x *InquiryResp) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type VerifyEmailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *VerifyEmailReq) Reset() {
	*x = VerifyEmailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_verify_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyEmailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyEmailReq) ProtoMessage() {}

func (x *VerifyEmailReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_verify_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyEmailReq.ProtoReflect.Descriptor instead.
func (*VerifyEmailReq) Descriptor() ([]byte, []int) {
	return file_api_proto_verify_proto_rawDescGZIP(), []int{4}
}

func (x *VerifyEmailReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_api_proto_verify_proto protoreflect.FileDescriptor

var file_api_proto_verify_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x14,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x0a, 0x49, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x12, 0x1e, 0x0a, 0x03, 0x65, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x03, 0x65,
	0x76, 0x65, 0x22, 0x49, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x03, 0x65, 0x76, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x03, 0x65, 0x76, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x36, 0x0a,
	0x14, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x03, 0x65, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x03, 0x65, 0x76, 0x65, 0x22, 0x21, 0x0a, 0x0b, 0x49, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x24, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0xaf,
	0x01, 0x0a, 0x06, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x27, 0x0a, 0x0b, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x31, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x49, 0x6e, 0x71, 0x75, 0x69, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x49, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x69, 0x66, 0x72, 0x61, 0x2d, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x69, 0x66, 0x72, 0x61, 0x2d,
	0x73, 0x73, 0x6f, 0x3b, 0x73, 0x73, 0x6f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_proto_verify_proto_rawDescOnce sync.Once
	file_api_proto_verify_proto_rawDescData = file_api_proto_verify_proto_rawDesc
)

func file_api_proto_verify_proto_rawDescGZIP() []byte {
	file_api_proto_verify_proto_rawDescOnce.Do(func() {
		file_api_proto_verify_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_verify_proto_rawDescData)
	})
	return file_api_proto_verify_proto_rawDescData
}

var file_api_proto_verify_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_proto_verify_proto_goTypes = []any{
	(*InquiryReq)(nil),           // 0: auth.InquiryReq
	(*CheckConfirmCodeReq)(nil),  // 1: auth.CheckConfirmCodeReq
	(*CheckConfirmCodeResp)(nil), // 2: auth.CheckConfirmCodeResp
	(*InquiryResp)(nil),          // 3: auth.InquiryResp
	(*VerifyEmailReq)(nil),       // 4: auth.VerifyEmailReq
	(Events)(0),                  // 5: auth.events
	(*Empty)(nil),                // 6: auth.Empty
}
var file_api_proto_verify_proto_depIdxs = []int32{
	5, // 0: auth.InquiryReq.eve:type_name -> auth.events
	5, // 1: auth.CheckConfirmCodeReq.eve:type_name -> auth.events
	5, // 2: auth.CheckConfirmCodeResp.eve:type_name -> auth.events
	6, // 3: auth.Verify.VerifyEmail:input_type -> auth.Empty
	6, // 4: auth.Verify.SendConfirmCode:input_type -> auth.Empty
	1, // 5: auth.Verify.CheckConfirmCode:input_type -> auth.CheckConfirmCodeReq
	6, // 6: auth.Verify.VerifyEmail:output_type -> auth.Empty
	3, // 7: auth.Verify.SendConfirmCode:output_type -> auth.InquiryResp
	2, // 8: auth.Verify.CheckConfirmCode:output_type -> auth.CheckConfirmCodeResp
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_proto_verify_proto_init() }
func file_api_proto_verify_proto_init() {
	if File_api_proto_verify_proto != nil {
		return
	}
	file_api_proto_auth_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_proto_verify_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*InquiryReq); i {
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
		file_api_proto_verify_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CheckConfirmCodeReq); i {
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
		file_api_proto_verify_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CheckConfirmCodeResp); i {
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
		file_api_proto_verify_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*InquiryResp); i {
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
		file_api_proto_verify_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*VerifyEmailReq); i {
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
			RawDescriptor: file_api_proto_verify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_verify_proto_goTypes,
		DependencyIndexes: file_api_proto_verify_proto_depIdxs,
		MessageInfos:      file_api_proto_verify_proto_msgTypes,
	}.Build()
	File_api_proto_verify_proto = out.File
	file_api_proto_verify_proto_rawDesc = nil
	file_api_proto_verify_proto_goTypes = nil
	file_api_proto_verify_proto_depIdxs = nil
}
