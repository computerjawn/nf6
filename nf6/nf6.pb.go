// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: nf6.proto

package nf6

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

type CreateRepoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateRepoRequest) Reset() {
	*x = CreateRepoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nf6_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRepoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRepoRequest) ProtoMessage() {}

func (x *CreateRepoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nf6_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRepoRequest.ProtoReflect.Descriptor instead.
func (*CreateRepoRequest) Descriptor() ([]byte, []int) {
	return file_nf6_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRepoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateRepoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CreateRepoReply) Reset() {
	*x = CreateRepoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nf6_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRepoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRepoReply) ProtoMessage() {}

func (x *CreateRepoReply) ProtoReflect() protoreflect.Message {
	mi := &file_nf6_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRepoReply.ProtoReflect.Descriptor instead.
func (*CreateRepoReply) Descriptor() ([]byte, []int) {
	return file_nf6_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRepoReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetCaCertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCaCertRequest) Reset() {
	*x = GetCaCertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nf6_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCaCertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCaCertRequest) ProtoMessage() {}

func (x *GetCaCertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nf6_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCaCertRequest.ProtoReflect.Descriptor instead.
func (*GetCaCertRequest) Descriptor() ([]byte, []int) {
	return file_nf6_proto_rawDescGZIP(), []int{2}
}

type GetCaCertReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cert []byte `protobuf:"bytes,1,opt,name=cert,proto3" json:"cert,omitempty"`
}

func (x *GetCaCertReply) Reset() {
	*x = GetCaCertReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nf6_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCaCertReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCaCertReply) ProtoMessage() {}

func (x *GetCaCertReply) ProtoReflect() protoreflect.Message {
	mi := &file_nf6_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCaCertReply.ProtoReflect.Descriptor instead.
func (*GetCaCertReply) Descriptor() ([]byte, []int) {
	return file_nf6_proto_rawDescGZIP(), []int{3}
}

func (x *GetCaCertReply) GetCert() []byte {
	if x != nil {
		return x.Cert
	}
	return nil
}

type ListReposRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListReposRequest) Reset() {
	*x = ListReposRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nf6_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReposRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReposRequest) ProtoMessage() {}

func (x *ListReposRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nf6_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReposRequest.ProtoReflect.Descriptor instead.
func (*ListReposRequest) Descriptor() ([]byte, []int) {
	return file_nf6_proto_rawDescGZIP(), []int{4}
}

type ListReposReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *ListReposReply) Reset() {
	*x = ListReposReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nf6_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReposReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReposReply) ProtoMessage() {}

func (x *ListReposReply) ProtoReflect() protoreflect.Message {
	mi := &file_nf6_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReposReply.ProtoReflect.Descriptor instead.
func (*ListReposReply) Descriptor() ([]byte, []int) {
	return file_nf6_proto_rawDescGZIP(), []int{5}
}

func (x *ListReposReply) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping bool `protobuf:"varint,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nf6_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nf6_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_nf6_proto_rawDescGZIP(), []int{6}
}

func (x *PingRequest) GetPing() bool {
	if x != nil {
		return x.Ping
	}
	return false
}

type PingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong bool `protobuf:"varint,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *PingReply) Reset() {
	*x = PingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nf6_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReply) ProtoMessage() {}

func (x *PingReply) ProtoReflect() protoreflect.Message {
	mi := &file_nf6_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReply.ProtoReflect.Descriptor instead.
func (*PingReply) Descriptor() ([]byte, []int) {
	return file_nf6_proto_rawDescGZIP(), []int{7}
}

func (x *PingReply) GetPong() bool {
	if x != nil {
		return x.Pong
	}
	return false
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email        string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	SslPublicKey []byte `protobuf:"bytes,2,opt,name=sslPublicKey,proto3" json:"sslPublicKey,omitempty"`
	SshPublicKey []byte `protobuf:"bytes,3,opt,name=sshPublicKey,proto3" json:"sshPublicKey,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nf6_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nf6_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_nf6_proto_rawDescGZIP(), []int{8}
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetSslPublicKey() []byte {
	if x != nil {
		return x.SslPublicKey
	}
	return nil
}

func (x *RegisterRequest) GetSshPublicKey() []byte {
	if x != nil {
		return x.SshPublicKey
	}
	return nil
}

type RegisterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SslCert []byte `protobuf:"bytes,1,opt,name=sslCert,proto3" json:"sslCert,omitempty"`
}

func (x *RegisterReply) Reset() {
	*x = RegisterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nf6_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReply) ProtoMessage() {}

func (x *RegisterReply) ProtoReflect() protoreflect.Message {
	mi := &file_nf6_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReply.ProtoReflect.Descriptor instead.
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return file_nf6_proto_rawDescGZIP(), []int{9}
}

func (x *RegisterReply) GetSslCert() []byte {
	if x != nil {
		return x.SslCert
	}
	return nil
}

type RenameRepoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldName string `protobuf:"bytes,1,opt,name=oldName,proto3" json:"oldName,omitempty"`
	NewName string `protobuf:"bytes,2,opt,name=newName,proto3" json:"newName,omitempty"`
}

func (x *RenameRepoRequest) Reset() {
	*x = RenameRepoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nf6_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenameRepoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameRepoRequest) ProtoMessage() {}

func (x *RenameRepoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nf6_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameRepoRequest.ProtoReflect.Descriptor instead.
func (*RenameRepoRequest) Descriptor() ([]byte, []int) {
	return file_nf6_proto_rawDescGZIP(), []int{10}
}

func (x *RenameRepoRequest) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

func (x *RenameRepoRequest) GetNewName() string {
	if x != nil {
		return x.NewName
	}
	return ""
}

type RenameRepoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RenameRepoReply) Reset() {
	*x = RenameRepoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nf6_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenameRepoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameRepoReply) ProtoMessage() {}

func (x *RenameRepoReply) ProtoReflect() protoreflect.Message {
	mi := &file_nf6_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameRepoReply.ProtoReflect.Descriptor instead.
func (*RenameRepoReply) Descriptor() ([]byte, []int) {
	return file_nf6_proto_rawDescGZIP(), []int{11}
}

func (x *RenameRepoReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type UpdateSshPublicKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SshPublicKey []byte `protobuf:"bytes,1,opt,name=sshPublicKey,proto3" json:"sshPublicKey,omitempty"`
}

func (x *UpdateSshPublicKeyRequest) Reset() {
	*x = UpdateSshPublicKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nf6_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSshPublicKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSshPublicKeyRequest) ProtoMessage() {}

func (x *UpdateSshPublicKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nf6_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSshPublicKeyRequest.ProtoReflect.Descriptor instead.
func (*UpdateSshPublicKeyRequest) Descriptor() ([]byte, []int) {
	return file_nf6_proto_rawDescGZIP(), []int{12}
}

func (x *UpdateSshPublicKeyRequest) GetSshPublicKey() []byte {
	if x != nil {
		return x.SshPublicKey
	}
	return nil
}

type UpdateSshPublicKeyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UpdateSshPublicKeyReply) Reset() {
	*x = UpdateSshPublicKeyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nf6_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSshPublicKeyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSshPublicKeyReply) ProtoMessage() {}

func (x *UpdateSshPublicKeyReply) ProtoReflect() protoreflect.Message {
	mi := &file_nf6_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSshPublicKeyReply.ProtoReflect.Descriptor instead.
func (*UpdateSshPublicKeyReply) Descriptor() ([]byte, []int) {
	return file_nf6_proto_rawDescGZIP(), []int{13}
}

func (x *UpdateSshPublicKeyReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type WhoAmIRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WhoAmIRequest) Reset() {
	*x = WhoAmIRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nf6_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhoAmIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhoAmIRequest) ProtoMessage() {}

func (x *WhoAmIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nf6_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhoAmIRequest.ProtoReflect.Descriptor instead.
func (*WhoAmIRequest) Descriptor() ([]byte, []int) {
	return file_nf6_proto_rawDescGZIP(), []int{14}
}

type WhoAmIReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email        string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	SslPublicKey string `protobuf:"bytes,2,opt,name=sslPublicKey,proto3" json:"sslPublicKey,omitempty"`
	SshPublicKey string `protobuf:"bytes,3,opt,name=sshPublicKey,proto3" json:"sshPublicKey,omitempty"`
}

func (x *WhoAmIReply) Reset() {
	*x = WhoAmIReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nf6_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhoAmIReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhoAmIReply) ProtoMessage() {}

func (x *WhoAmIReply) ProtoReflect() protoreflect.Message {
	mi := &file_nf6_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhoAmIReply.ProtoReflect.Descriptor instead.
func (*WhoAmIReply) Descriptor() ([]byte, []int) {
	return file_nf6_proto_rawDescGZIP(), []int{15}
}

func (x *WhoAmIReply) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *WhoAmIReply) GetSslPublicKey() string {
	if x != nil {
		return x.SslPublicKey
	}
	return ""
}

func (x *WhoAmIReply) GetSshPublicKey() string {
	if x != nil {
		return x.SshPublicKey
	}
	return ""
}

var File_nf6_proto protoreflect.FileDescriptor

var file_nf6_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6e, 0x66, 0x36, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e, 0x66, 0x36,
	0x22, 0x27, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x0f, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x61, 0x43,
	0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x24, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x63, 0x65, 0x72, 0x74,
	0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x26, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x21, 0x0a, 0x0b,
	0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x22,
	0x1f, 0x0a, 0x09, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67,
	0x22, 0x6f, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x73, 0x6c,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0c, 0x73, 0x73, 0x6c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x22, 0x0a,
	0x0c, 0x73, 0x73, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0c, 0x73, 0x73, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x22, 0x29, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x73, 0x6c, 0x43, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x73, 0x73, 0x6c, 0x43, 0x65, 0x72, 0x74, 0x22, 0x47, 0x0a, 0x11,
	0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6e,
	0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65,
	0x77, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x0f, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x70, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x3f, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x73, 0x68, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x73, 0x73, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x73, 0x73, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x22, 0x33, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x73, 0x68,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x0f, 0x0a, 0x0d, 0x57, 0x68, 0x6f, 0x41,
	0x6d, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6b, 0x0a, 0x0b, 0x57, 0x68, 0x6f,
	0x41, 0x6d, 0x49, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x22,
	0x0a, 0x0c, 0x73, 0x73, 0x6c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x73, 0x6c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x73, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x73, 0x68, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x32, 0xac, 0x01, 0x0a, 0x0b, 0x4e, 0x66, 0x36, 0x49, 0x6e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x61, 0x43,
	0x65, 0x72, 0x74, 0x12, 0x15, 0x2e, 0x6e, 0x66, 0x36, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x43,
	0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6e, 0x66, 0x36,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x2a, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x2e, 0x6e, 0x66, 0x36, 0x2e,
	0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6e, 0x66,
	0x36, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a,
	0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x6e, 0x66, 0x36, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x6e, 0x66, 0x36, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x32, 0xca, 0x02, 0x0a, 0x09, 0x4e, 0x66, 0x36, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70,
	0x6f, 0x12, 0x16, 0x2e, 0x6e, 0x66, 0x36, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x70, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6e, 0x66, 0x36, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x39, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x12, 0x15,
	0x2e, 0x6e, 0x66, 0x36, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6e, 0x66, 0x36, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x6f, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0a,
	0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x16, 0x2e, 0x6e, 0x66, 0x36,
	0x2e, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6e, 0x66, 0x36, 0x2e, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x70, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x73, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x12, 0x1e, 0x2e, 0x6e, 0x66, 0x36, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x73, 0x68,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x6e, 0x66, 0x36, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x73, 0x68,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x30, 0x0a, 0x06, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x12, 0x12, 0x2e, 0x6e, 0x66, 0x36,
	0x2e, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x6e, 0x66, 0x36, 0x2e, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x64, 0x61, 0x6e, 0x65, 0x2f, 0x6e, 0x66,
	0x36, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nf6_proto_rawDescOnce sync.Once
	file_nf6_proto_rawDescData = file_nf6_proto_rawDesc
)

func file_nf6_proto_rawDescGZIP() []byte {
	file_nf6_proto_rawDescOnce.Do(func() {
		file_nf6_proto_rawDescData = protoimpl.X.CompressGZIP(file_nf6_proto_rawDescData)
	})
	return file_nf6_proto_rawDescData
}

var file_nf6_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_nf6_proto_goTypes = []interface{}{
	(*CreateRepoRequest)(nil),         // 0: nf6.CreateRepoRequest
	(*CreateRepoReply)(nil),           // 1: nf6.CreateRepoReply
	(*GetCaCertRequest)(nil),          // 2: nf6.GetCaCertRequest
	(*GetCaCertReply)(nil),            // 3: nf6.GetCaCertReply
	(*ListReposRequest)(nil),          // 4: nf6.ListReposRequest
	(*ListReposReply)(nil),            // 5: nf6.ListReposReply
	(*PingRequest)(nil),               // 6: nf6.PingRequest
	(*PingReply)(nil),                 // 7: nf6.PingReply
	(*RegisterRequest)(nil),           // 8: nf6.RegisterRequest
	(*RegisterReply)(nil),             // 9: nf6.RegisterReply
	(*RenameRepoRequest)(nil),         // 10: nf6.RenameRepoRequest
	(*RenameRepoReply)(nil),           // 11: nf6.RenameRepoReply
	(*UpdateSshPublicKeyRequest)(nil), // 12: nf6.UpdateSshPublicKeyRequest
	(*UpdateSshPublicKeyReply)(nil),   // 13: nf6.UpdateSshPublicKeyReply
	(*WhoAmIRequest)(nil),             // 14: nf6.WhoAmIRequest
	(*WhoAmIReply)(nil),               // 15: nf6.WhoAmIReply
}
var file_nf6_proto_depIdxs = []int32{
	2,  // 0: nf6.Nf6Insecure.GetCaCert:input_type -> nf6.GetCaCertRequest
	6,  // 1: nf6.Nf6Insecure.Ping:input_type -> nf6.PingRequest
	8,  // 2: nf6.Nf6Insecure.Register:input_type -> nf6.RegisterRequest
	0,  // 3: nf6.Nf6Secure.CreateRepo:input_type -> nf6.CreateRepoRequest
	4,  // 4: nf6.Nf6Secure.ListRepos:input_type -> nf6.ListReposRequest
	10, // 5: nf6.Nf6Secure.RenameRepo:input_type -> nf6.RenameRepoRequest
	12, // 6: nf6.Nf6Secure.UpdateSshPublicKey:input_type -> nf6.UpdateSshPublicKeyRequest
	14, // 7: nf6.Nf6Secure.WhoAmI:input_type -> nf6.WhoAmIRequest
	3,  // 8: nf6.Nf6Insecure.GetCaCert:output_type -> nf6.GetCaCertReply
	7,  // 9: nf6.Nf6Insecure.Ping:output_type -> nf6.PingReply
	9,  // 10: nf6.Nf6Insecure.Register:output_type -> nf6.RegisterReply
	1,  // 11: nf6.Nf6Secure.CreateRepo:output_type -> nf6.CreateRepoReply
	5,  // 12: nf6.Nf6Secure.ListRepos:output_type -> nf6.ListReposReply
	11, // 13: nf6.Nf6Secure.RenameRepo:output_type -> nf6.RenameRepoReply
	13, // 14: nf6.Nf6Secure.UpdateSshPublicKey:output_type -> nf6.UpdateSshPublicKeyReply
	15, // 15: nf6.Nf6Secure.WhoAmI:output_type -> nf6.WhoAmIReply
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_nf6_proto_init() }
func file_nf6_proto_init() {
	if File_nf6_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nf6_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRepoRequest); i {
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
		file_nf6_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRepoReply); i {
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
		file_nf6_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCaCertRequest); i {
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
		file_nf6_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCaCertReply); i {
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
		file_nf6_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReposRequest); i {
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
		file_nf6_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReposReply); i {
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
		file_nf6_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_nf6_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingReply); i {
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
		file_nf6_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_nf6_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReply); i {
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
		file_nf6_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenameRepoRequest); i {
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
		file_nf6_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenameRepoReply); i {
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
		file_nf6_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSshPublicKeyRequest); i {
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
		file_nf6_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSshPublicKeyReply); i {
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
		file_nf6_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhoAmIRequest); i {
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
		file_nf6_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhoAmIReply); i {
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
			RawDescriptor: file_nf6_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_nf6_proto_goTypes,
		DependencyIndexes: file_nf6_proto_depIdxs,
		MessageInfos:      file_nf6_proto_msgTypes,
	}.Build()
	File_nf6_proto = out.File
	file_nf6_proto_rawDesc = nil
	file_nf6_proto_goTypes = nil
	file_nf6_proto_depIdxs = nil
}
