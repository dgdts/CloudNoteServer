// idl/share/share.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v5.29.0--rc1
// source: share.proto

package share

import (
	_ "github.com/dgdts/UniversalServer/biz/model/api"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetShareNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShareId string `protobuf:"bytes,1,opt,name=share_id,json=shareId,proto3" form:"share_id" json:"share_id,omitempty" query:"share_id"`
}

func (x *GetShareNoteRequest) Reset() {
	*x = GetShareNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShareNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShareNoteRequest) ProtoMessage() {}

func (x *GetShareNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShareNoteRequest.ProtoReflect.Descriptor instead.
func (*GetShareNoteRequest) Descriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{0}
}

func (x *GetShareNoteRequest) GetShareId() string {
	if x != nil {
		return x.ShareId
	}
	return ""
}

type GetShareNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetShareNoteResponse) Reset() {
	*x = GetShareNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShareNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShareNoteResponse) ProtoMessage() {}

func (x *GetShareNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShareNoteResponse.ProtoReflect.Descriptor instead.
func (*GetShareNoteResponse) Descriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{1}
}

type ListShareNoteCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShareId  string `protobuf:"bytes,1,opt,name=share_id,json=shareId,proto3" form:"share_id" json:"share_id,omitempty" query:"share_id"`
	Page     int64  `protobuf:"varint,2,opt,name=page,proto3" form:"page" json:"page,omitempty" query:"page"`
	PageSize int64  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" form:"page_size" json:"page_size,omitempty" query:"page_size"`
}

func (x *ListShareNoteCommentsRequest) Reset() {
	*x = ListShareNoteCommentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListShareNoteCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShareNoteCommentsRequest) ProtoMessage() {}

func (x *ListShareNoteCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShareNoteCommentsRequest.ProtoReflect.Descriptor instead.
func (*ListShareNoteCommentsRequest) Descriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{2}
}

func (x *ListShareNoteCommentsRequest) GetShareId() string {
	if x != nil {
		return x.ShareId
	}
	return ""
}

func (x *ListShareNoteCommentsRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListShareNoteCommentsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ShareNoteComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShareId   string                 `protobuf:"bytes,1,opt,name=share_id,json=shareId,proto3" form:"share_id" json:"share_id,omitempty" query:"share_id"`
	Alias     string                 `protobuf:"bytes,2,opt,name=alias,proto3" form:"alias" json:"alias,omitempty" query:"alias"`
	Content   string                 `protobuf:"bytes,3,opt,name=content,proto3" form:"content" json:"content,omitempty" query:"content"`
	Ip        string                 `protobuf:"bytes,4,opt,name=ip,proto3" form:"ip" json:"ip,omitempty" query:"ip"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" form:"created_at" json:"created_at,omitempty" query:"created_at"`
}

func (x *ShareNoteComment) Reset() {
	*x = ShareNoteComment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareNoteComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareNoteComment) ProtoMessage() {}

func (x *ShareNoteComment) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareNoteComment.ProtoReflect.Descriptor instead.
func (*ShareNoteComment) Descriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{3}
}

func (x *ShareNoteComment) GetShareId() string {
	if x != nil {
		return x.ShareId
	}
	return ""
}

func (x *ShareNoteComment) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *ShareNoteComment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ShareNoteComment) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *ShareNoteComment) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type ListShareNoteCommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*ShareNoteComment `protobuf:"bytes,1,rep,name=comments,proto3" form:"comments" json:"comments,omitempty" query:"comments"`
	Total    int64               `protobuf:"varint,2,opt,name=total,proto3" form:"total" json:"total,omitempty" query:"total"`
}

func (x *ListShareNoteCommentsResponse) Reset() {
	*x = ListShareNoteCommentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListShareNoteCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShareNoteCommentsResponse) ProtoMessage() {}

func (x *ListShareNoteCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShareNoteCommentsResponse.ProtoReflect.Descriptor instead.
func (*ListShareNoteCommentsResponse) Descriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{4}
}

func (x *ListShareNoteCommentsResponse) GetComments() []*ShareNoteComment {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *ListShareNoteCommentsResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type CreateShareNoteCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *ShareNoteComment `protobuf:"bytes,1,opt,name=comment,proto3" form:"comment" json:"comment,omitempty" query:"comment"`
}

func (x *CreateShareNoteCommentRequest) Reset() {
	*x = CreateShareNoteCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShareNoteCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShareNoteCommentRequest) ProtoMessage() {}

func (x *CreateShareNoteCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShareNoteCommentRequest.ProtoReflect.Descriptor instead.
func (*CreateShareNoteCommentRequest) Descriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{5}
}

func (x *CreateShareNoteCommentRequest) GetComment() *ShareNoteComment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type CreateShareNoteCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateShareNoteCommentResponse) Reset() {
	*x = CreateShareNoteCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShareNoteCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShareNoteCommentResponse) ProtoMessage() {}

func (x *CreateShareNoteCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShareNoteCommentResponse.ProtoReflect.Descriptor instead.
func (*CreateShareNoteCommentResponse) Descriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{6}
}

var File_share_proto protoreflect.FileDescriptor

var file_share_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x30, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4e, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4e, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6a, 0x0a, 0x1c, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xa8, 0x01, 0x0a, 0x10, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x4e, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x6a, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4e, 0x6f,
	0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x52, 0x0a,
	0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4e, 0x6f, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4e, 0x6f, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x22, 0x20, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x4e, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xfd, 0x02, 0x0a, 0x0c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x4e, 0x6f, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0xca,
	0xc1, 0x18, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x83, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x23, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x4e, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0xca, 0xc1, 0x18, 0x1b,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2f, 0x6e, 0x6f,
	0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x85, 0x01, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x4e, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1e, 0xd2, 0xc1, 0x18, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x67, 0x64, 0x74, 0x73, 0x2f, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_share_proto_rawDescOnce sync.Once
	file_share_proto_rawDescData = file_share_proto_rawDesc
)

func file_share_proto_rawDescGZIP() []byte {
	file_share_proto_rawDescOnce.Do(func() {
		file_share_proto_rawDescData = protoimpl.X.CompressGZIP(file_share_proto_rawDescData)
	})
	return file_share_proto_rawDescData
}

var file_share_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_share_proto_goTypes = []interface{}{
	(*GetShareNoteRequest)(nil),            // 0: share.GetShareNoteRequest
	(*GetShareNoteResponse)(nil),           // 1: share.GetShareNoteResponse
	(*ListShareNoteCommentsRequest)(nil),   // 2: share.ListShareNoteCommentsRequest
	(*ShareNoteComment)(nil),               // 3: share.ShareNoteComment
	(*ListShareNoteCommentsResponse)(nil),  // 4: share.ListShareNoteCommentsResponse
	(*CreateShareNoteCommentRequest)(nil),  // 5: share.CreateShareNoteCommentRequest
	(*CreateShareNoteCommentResponse)(nil), // 6: share.CreateShareNoteCommentResponse
	(*timestamppb.Timestamp)(nil),          // 7: google.protobuf.Timestamp
}
var file_share_proto_depIdxs = []int32{
	7, // 0: share.ShareNoteComment.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: share.ListShareNoteCommentsResponse.comments:type_name -> share.ShareNoteComment
	3, // 2: share.CreateShareNoteCommentRequest.comment:type_name -> share.ShareNoteComment
	0, // 3: share.ShareService.GetShareNote:input_type -> share.GetShareNoteRequest
	2, // 4: share.ShareService.ListShareNoteComments:input_type -> share.ListShareNoteCommentsRequest
	5, // 5: share.ShareService.CreateShareNoteComment:input_type -> share.CreateShareNoteCommentRequest
	1, // 6: share.ShareService.GetShareNote:output_type -> share.GetShareNoteResponse
	4, // 7: share.ShareService.ListShareNoteComments:output_type -> share.ListShareNoteCommentsResponse
	6, // 8: share.ShareService.CreateShareNoteComment:output_type -> share.CreateShareNoteCommentResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_share_proto_init() }
func file_share_proto_init() {
	if File_share_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_share_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShareNoteRequest); i {
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
		file_share_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShareNoteResponse); i {
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
		file_share_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListShareNoteCommentsRequest); i {
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
		file_share_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShareNoteComment); i {
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
		file_share_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListShareNoteCommentsResponse); i {
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
		file_share_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShareNoteCommentRequest); i {
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
		file_share_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShareNoteCommentResponse); i {
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
			RawDescriptor: file_share_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_share_proto_goTypes,
		DependencyIndexes: file_share_proto_depIdxs,
		MessageInfos:      file_share_proto_msgTypes,
	}.Build()
	File_share_proto = out.File
	file_share_proto_rawDesc = nil
	file_share_proto_goTypes = nil
	file_share_proto_depIdxs = nil
}
