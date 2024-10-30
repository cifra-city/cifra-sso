// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: resources/grpc/proto/sso.proto

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

type Events int32

const (
	Events_LOGIN           Events = 0
	Events_CHANGE_PASSWORD Events = 1
	Events_CHANGE_USERNAME Events = 2
	Events_CHANGE_EMAIL    Events = 3
	Events_CONFIRM_EMAIL   Events = 4
)

// Enum value maps for Events.
var (
	Events_name = map[int32]string{
		0: "LOGIN",
		1: "CHANGE_PASSWORD",
		2: "CHANGE_USERNAME",
		3: "CHANGE_EMAIL",
		4: "CONFIRM_EMAIL",
	}
	Events_value = map[string]int32{
		"LOGIN":           0,
		"CHANGE_PASSWORD": 1,
		"CHANGE_USERNAME": 2,
		"CHANGE_EMAIL":    3,
		"CONFIRM_EMAIL":   4,
	}
)

func (x Events) Enum() *Events {
	p := new(Events)
	*p = x
	return p
}

func (x Events) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Events) Descriptor() protoreflect.EnumDescriptor {
	return file_resources_grpc_proto_sso_proto_enumTypes[0].Descriptor()
}

func (Events) Type() protoreflect.EnumType {
	return &file_resources_grpc_proto_sso_proto_enumTypes[0]
}

func (x Events) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Events.Descriptor instead.
func (Events) EnumDescriptor() ([]byte, []int) {
	return file_resources_grpc_proto_sso_proto_rawDescGZIP(), []int{0}
}

type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_grpc_proto_sso_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_resources_grpc_proto_sso_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_resources_grpc_proto_sso_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    *string `protobuf:"bytes,1,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Username *string `protobuf:"bytes,2,opt,name=username,proto3,oneof" json:"username,omitempty"`
	Password string  `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_grpc_proto_sso_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_resources_grpc_proto_sso_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_resources_grpc_proto_sso_proto_rawDescGZIP(), []int{1}
}

func (x *LoginReq) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *LoginReq) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ChangePassReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewPassword string `protobuf:"bytes,1,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
}

func (x *ChangePassReq) Reset() {
	*x = ChangePassReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_grpc_proto_sso_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePassReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePassReq) ProtoMessage() {}

func (x *ChangePassReq) ProtoReflect() protoreflect.Message {
	mi := &file_resources_grpc_proto_sso_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePassReq.ProtoReflect.Descriptor instead.
func (*ChangePassReq) Descriptor() ([]byte, []int) {
	return file_resources_grpc_proto_sso_proto_rawDescGZIP(), []int{2}
}

func (x *ChangePassReq) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type ChangeEmailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	NewEmail string `protobuf:"bytes,3,opt,name=new_email,json=newEmail,proto3" json:"new_email,omitempty"`
}

func (x *ChangeEmailReq) Reset() {
	*x = ChangeEmailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_grpc_proto_sso_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeEmailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeEmailReq) ProtoMessage() {}

func (x *ChangeEmailReq) ProtoReflect() protoreflect.Message {
	mi := &file_resources_grpc_proto_sso_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeEmailReq.ProtoReflect.Descriptor instead.
func (*ChangeEmailReq) Descriptor() ([]byte, []int) {
	return file_resources_grpc_proto_sso_proto_rawDescGZIP(), []int{3}
}

func (x *ChangeEmailReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ChangeEmailReq) GetNewEmail() string {
	if x != nil {
		return x.NewEmail
	}
	return ""
}

type ChangeUsernameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password    string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	OldUsername string `protobuf:"bytes,2,opt,name=old_username,json=oldUsername,proto3" json:"old_username,omitempty"`
	NewUsername string `protobuf:"bytes,3,opt,name=new_username,json=newUsername,proto3" json:"new_username,omitempty"`
}

func (x *ChangeUsernameReq) Reset() {
	*x = ChangeUsernameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_grpc_proto_sso_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeUsernameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeUsernameReq) ProtoMessage() {}

func (x *ChangeUsernameReq) ProtoReflect() protoreflect.Message {
	mi := &file_resources_grpc_proto_sso_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeUsernameReq.ProtoReflect.Descriptor instead.
func (*ChangeUsernameReq) Descriptor() ([]byte, []int) {
	return file_resources_grpc_proto_sso_proto_rawDescGZIP(), []int{4}
}

func (x *ChangeUsernameReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ChangeUsernameReq) GetOldUsername() string {
	if x != nil {
		return x.OldUsername
	}
	return ""
}

func (x *ChangeUsernameReq) GetNewUsername() string {
	if x != nil {
		return x.NewUsername
	}
	return ""
}

type AccessReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eve  Events `protobuf:"varint,1,opt,name=eve,proto3,enum=auth.Events" json:"eve,omitempty"`
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *AccessReq) Reset() {
	*x = AccessReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_grpc_proto_sso_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessReq) ProtoMessage() {}

func (x *AccessReq) ProtoReflect() protoreflect.Message {
	mi := &file_resources_grpc_proto_sso_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessReq.ProtoReflect.Descriptor instead.
func (*AccessReq) Descriptor() ([]byte, []int) {
	return file_resources_grpc_proto_sso_proto_rawDescGZIP(), []int{5}
}

func (x *AccessReq) GetEve() Events {
	if x != nil {
		return x.Eve
	}
	return Events_LOGIN
}

func (x *AccessReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type InquiryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eve Events `protobuf:"varint,1,opt,name=eve,proto3,enum=auth.Events" json:"eve,omitempty"`
}

func (x *InquiryReq) Reset() {
	*x = InquiryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_grpc_proto_sso_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InquiryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InquiryReq) ProtoMessage() {}

func (x *InquiryReq) ProtoReflect() protoreflect.Message {
	mi := &file_resources_grpc_proto_sso_proto_msgTypes[6]
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
	return file_resources_grpc_proto_sso_proto_rawDescGZIP(), []int{6}
}

func (x *InquiryReq) GetEve() Events {
	if x != nil {
		return x.Eve
	}
	return Events_LOGIN
}

type LoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_grpc_proto_sso_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_resources_grpc_proto_sso_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_resources_grpc_proto_sso_proto_rawDescGZIP(), []int{7}
}

func (x *LoginResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AccessResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eve Events `protobuf:"varint,1,opt,name=eve,proto3,enum=auth.Events" json:"eve,omitempty"`
}

func (x *AccessResp) Reset() {
	*x = AccessResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_grpc_proto_sso_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessResp) ProtoMessage() {}

func (x *AccessResp) ProtoReflect() protoreflect.Message {
	mi := &file_resources_grpc_proto_sso_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessResp.ProtoReflect.Descriptor instead.
func (*AccessResp) Descriptor() ([]byte, []int) {
	return file_resources_grpc_proto_sso_proto_rawDescGZIP(), []int{8}
}

func (x *AccessResp) GetEve() Events {
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
		mi := &file_resources_grpc_proto_sso_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InquiryResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InquiryResp) ProtoMessage() {}

func (x *InquiryResp) ProtoReflect() protoreflect.Message {
	mi := &file_resources_grpc_proto_sso_proto_msgTypes[9]
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
	return file_resources_grpc_proto_sso_proto_rawDescGZIP(), []int{9}
}

func (x *InquiryResp) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_grpc_proto_sso_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_resources_grpc_proto_sso_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_resources_grpc_proto_sso_proto_rawDescGZIP(), []int{10}
}

var File_resources_grpc_proto_sso_proto protoreflect.FileDescriptor

var file_resources_grpc_proto_sso_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x73, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x5b, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x79, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12,
	0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x32,
	0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x12,
	0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x49, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x75, 0x0a,
	0x11, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x6f, 0x6c, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x6c, 0x64, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3f, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x12, 0x1e, 0x0a, 0x03, 0x65, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x03, 0x65, 0x76,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x2c, 0x0a, 0x0a, 0x49, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x03, 0x65, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x03,
	0x65, 0x76, 0x65, 0x22, 0x21, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2c, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x03, 0x65, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x03, 0x65, 0x76, 0x65, 0x22, 0x21, 0x0a, 0x0b, 0x49, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x2a, 0x62, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x4f,
	0x47, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f,
	0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x48,
	0x41, 0x4e, 0x47, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x02, 0x12,
	0x10, 0x0a, 0x0c, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10,
	0x03, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x5f, 0x45, 0x4d, 0x41,
	0x49, 0x4c, 0x10, 0x04, 0x32, 0xe5, 0x02, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x22, 0x0a,
	0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x3e, 0x0a, 0x17, 0x49, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x46, 0x6f, 0x72, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x49, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x11,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x49, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x35, 0x0a, 0x10, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x46, 0x6f, 0x72, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x0f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x1a, 0x0f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x2e, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73,
	0x12, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x36, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x30, 0x0a, 0x0b, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a,
	0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x31, 0x0a, 0x03,
	0x52, 0x65, 0x67, 0x12, 0x2a, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32,
	0x31, 0x0a, 0x06, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x27, 0x0a, 0x0b, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x69, 0x66, 0x72, 0x61, 0x2d, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x69, 0x66, 0x72,
	0x61, 0x2d, 0x73, 0x73, 0x6f, 0x3b, 0x73, 0x73, 0x6f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_resources_grpc_proto_sso_proto_rawDescOnce sync.Once
	file_resources_grpc_proto_sso_proto_rawDescData = file_resources_grpc_proto_sso_proto_rawDesc
)

func file_resources_grpc_proto_sso_proto_rawDescGZIP() []byte {
	file_resources_grpc_proto_sso_proto_rawDescOnce.Do(func() {
		file_resources_grpc_proto_sso_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_grpc_proto_sso_proto_rawDescData)
	})
	return file_resources_grpc_proto_sso_proto_rawDescData
}

var file_resources_grpc_proto_sso_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_resources_grpc_proto_sso_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_resources_grpc_proto_sso_proto_goTypes = []any{
	(Events)(0),               // 0: auth.events
	(*RegisterReq)(nil),       // 1: auth.RegisterReq
	(*LoginReq)(nil),          // 2: auth.LoginReq
	(*ChangePassReq)(nil),     // 3: auth.ChangePassReq
	(*ChangeEmailReq)(nil),    // 4: auth.ChangeEmailReq
	(*ChangeUsernameReq)(nil), // 5: auth.ChangeUsernameReq
	(*AccessReq)(nil),         // 6: auth.AccessReq
	(*InquiryReq)(nil),        // 7: auth.InquiryReq
	(*LoginResp)(nil),         // 8: auth.LoginResp
	(*AccessResp)(nil),        // 9: auth.AccessResp
	(*InquiryResp)(nil),       // 10: auth.InquiryResp
	(*Empty)(nil),             // 11: auth.Empty
}
var file_resources_grpc_proto_sso_proto_depIdxs = []int32{
	0,  // 0: auth.AccessReq.eve:type_name -> auth.events
	0,  // 1: auth.InquiryReq.eve:type_name -> auth.events
	0,  // 2: auth.AccessResp.eve:type_name -> auth.events
	11, // 3: auth.Auth.Logout:input_type -> auth.Empty
	7,  // 4: auth.Auth.InquiryForChangeRequest:input_type -> auth.InquiryReq
	6,  // 5: auth.Auth.AccessForChanges:input_type -> auth.AccessReq
	2,  // 6: auth.Auth.Login:input_type -> auth.LoginReq
	3,  // 7: auth.Auth.ChangePass:input_type -> auth.ChangePassReq
	5,  // 8: auth.Auth.ChangeUsername:input_type -> auth.ChangeUsernameReq
	4,  // 9: auth.Auth.ChangeEmail:input_type -> auth.ChangeEmailReq
	1,  // 10: auth.Reg.Register:input_type -> auth.RegisterReq
	11, // 11: auth.Verify.VerifyEmail:input_type -> auth.Empty
	11, // 12: auth.Auth.Logout:output_type -> auth.Empty
	10, // 13: auth.Auth.InquiryForChangeRequest:output_type -> auth.InquiryResp
	9,  // 14: auth.Auth.AccessForChanges:output_type -> auth.AccessResp
	8,  // 15: auth.Auth.Login:output_type -> auth.LoginResp
	11, // 16: auth.Auth.ChangePass:output_type -> auth.Empty
	11, // 17: auth.Auth.ChangeUsername:output_type -> auth.Empty
	11, // 18: auth.Auth.ChangeEmail:output_type -> auth.Empty
	11, // 19: auth.Reg.Register:output_type -> auth.Empty
	11, // 20: auth.Verify.VerifyEmail:output_type -> auth.Empty
	12, // [12:21] is the sub-list for method output_type
	3,  // [3:12] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_resources_grpc_proto_sso_proto_init() }
func file_resources_grpc_proto_sso_proto_init() {
	if File_resources_grpc_proto_sso_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_grpc_proto_sso_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterReq); i {
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
		file_resources_grpc_proto_sso_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LoginReq); i {
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
		file_resources_grpc_proto_sso_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ChangePassReq); i {
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
		file_resources_grpc_proto_sso_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ChangeEmailReq); i {
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
		file_resources_grpc_proto_sso_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ChangeUsernameReq); i {
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
		file_resources_grpc_proto_sso_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*AccessReq); i {
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
		file_resources_grpc_proto_sso_proto_msgTypes[6].Exporter = func(v any, i int) any {
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
		file_resources_grpc_proto_sso_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*LoginResp); i {
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
		file_resources_grpc_proto_sso_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*AccessResp); i {
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
		file_resources_grpc_proto_sso_proto_msgTypes[9].Exporter = func(v any, i int) any {
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
		file_resources_grpc_proto_sso_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*Empty); i {
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
	file_resources_grpc_proto_sso_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_grpc_proto_sso_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_resources_grpc_proto_sso_proto_goTypes,
		DependencyIndexes: file_resources_grpc_proto_sso_proto_depIdxs,
		EnumInfos:         file_resources_grpc_proto_sso_proto_enumTypes,
		MessageInfos:      file_resources_grpc_proto_sso_proto_msgTypes,
	}.Build()
	File_resources_grpc_proto_sso_proto = out.File
	file_resources_grpc_proto_sso_proto_rawDesc = nil
	file_resources_grpc_proto_sso_proto_goTypes = nil
	file_resources_grpc_proto_sso_proto_depIdxs = nil
}
