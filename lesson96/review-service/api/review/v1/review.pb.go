// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.1
// source: api/review/v1/review.proto

package v1

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

type CreateReviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateReviewRequest) Reset() {
	*x = CreateReviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_review_v1_review_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReviewRequest) ProtoMessage() {}

func (x *CreateReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReviewRequest.ProtoReflect.Descriptor instead.
func (*CreateReviewRequest) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{0}
}

type CreateReviewReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateReviewReply) Reset() {
	*x = CreateReviewReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_review_v1_review_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReviewReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReviewReply) ProtoMessage() {}

func (x *CreateReviewReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReviewReply.ProtoReflect.Descriptor instead.
func (*CreateReviewReply) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{1}
}

type UpdateReviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateReviewRequest) Reset() {
	*x = UpdateReviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_review_v1_review_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReviewRequest) ProtoMessage() {}

func (x *UpdateReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReviewRequest.ProtoReflect.Descriptor instead.
func (*UpdateReviewRequest) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{2}
}

type UpdateReviewReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateReviewReply) Reset() {
	*x = UpdateReviewReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_review_v1_review_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReviewReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReviewReply) ProtoMessage() {}

func (x *UpdateReviewReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReviewReply.ProtoReflect.Descriptor instead.
func (*UpdateReviewReply) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{3}
}

type DeleteReviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteReviewRequest) Reset() {
	*x = DeleteReviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_review_v1_review_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReviewRequest) ProtoMessage() {}

func (x *DeleteReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReviewRequest.ProtoReflect.Descriptor instead.
func (*DeleteReviewRequest) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{4}
}

type DeleteReviewReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteReviewReply) Reset() {
	*x = DeleteReviewReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_review_v1_review_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReviewReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReviewReply) ProtoMessage() {}

func (x *DeleteReviewReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReviewReply.ProtoReflect.Descriptor instead.
func (*DeleteReviewReply) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{5}
}

type GetReviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetReviewRequest) Reset() {
	*x = GetReviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_review_v1_review_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReviewRequest) ProtoMessage() {}

func (x *GetReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReviewRequest.ProtoReflect.Descriptor instead.
func (*GetReviewRequest) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{6}
}

type GetReviewReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetReviewReply) Reset() {
	*x = GetReviewReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_review_v1_review_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReviewReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReviewReply) ProtoMessage() {}

func (x *GetReviewReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReviewReply.ProtoReflect.Descriptor instead.
func (*GetReviewReply) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{7}
}

type ListReviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListReviewRequest) Reset() {
	*x = ListReviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_review_v1_review_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReviewRequest) ProtoMessage() {}

func (x *ListReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReviewRequest.ProtoReflect.Descriptor instead.
func (*ListReviewRequest) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{8}
}

type ListReviewReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListReviewReply) Reset() {
	*x = ListReviewReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_review_v1_review_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReviewReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReviewReply) ProtoMessage() {}

func (x *ListReviewReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReviewReply.ProtoReflect.Descriptor instead.
func (*ListReviewReply) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{9}
}

var File_api_review_v1_review_proto protoreflect.FileDescriptor

var file_api_review_v1_review_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70,
	0x69, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x76, 0x31, 0x22, 0x15, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xa7, 0x03,
	0x0a, 0x06, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x54, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x54,
	0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x22,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x54, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4b, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4e, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x32, 0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1f, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_review_v1_review_proto_rawDescOnce sync.Once
	file_api_review_v1_review_proto_rawDescData = file_api_review_v1_review_proto_rawDesc
)

func file_api_review_v1_review_proto_rawDescGZIP() []byte {
	file_api_review_v1_review_proto_rawDescOnce.Do(func() {
		file_api_review_v1_review_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_review_v1_review_proto_rawDescData)
	})
	return file_api_review_v1_review_proto_rawDescData
}

var file_api_review_v1_review_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_review_v1_review_proto_goTypes = []interface{}{
	(*CreateReviewRequest)(nil), // 0: api.review.v1.CreateReviewRequest
	(*CreateReviewReply)(nil),   // 1: api.review.v1.CreateReviewReply
	(*UpdateReviewRequest)(nil), // 2: api.review.v1.UpdateReviewRequest
	(*UpdateReviewReply)(nil),   // 3: api.review.v1.UpdateReviewReply
	(*DeleteReviewRequest)(nil), // 4: api.review.v1.DeleteReviewRequest
	(*DeleteReviewReply)(nil),   // 5: api.review.v1.DeleteReviewReply
	(*GetReviewRequest)(nil),    // 6: api.review.v1.GetReviewRequest
	(*GetReviewReply)(nil),      // 7: api.review.v1.GetReviewReply
	(*ListReviewRequest)(nil),   // 8: api.review.v1.ListReviewRequest
	(*ListReviewReply)(nil),     // 9: api.review.v1.ListReviewReply
}
var file_api_review_v1_review_proto_depIdxs = []int32{
	0, // 0: api.review.v1.Review.CreateReview:input_type -> api.review.v1.CreateReviewRequest
	2, // 1: api.review.v1.Review.UpdateReview:input_type -> api.review.v1.UpdateReviewRequest
	4, // 2: api.review.v1.Review.DeleteReview:input_type -> api.review.v1.DeleteReviewRequest
	6, // 3: api.review.v1.Review.GetReview:input_type -> api.review.v1.GetReviewRequest
	8, // 4: api.review.v1.Review.ListReview:input_type -> api.review.v1.ListReviewRequest
	1, // 5: api.review.v1.Review.CreateReview:output_type -> api.review.v1.CreateReviewReply
	3, // 6: api.review.v1.Review.UpdateReview:output_type -> api.review.v1.UpdateReviewReply
	5, // 7: api.review.v1.Review.DeleteReview:output_type -> api.review.v1.DeleteReviewReply
	7, // 8: api.review.v1.Review.GetReview:output_type -> api.review.v1.GetReviewReply
	9, // 9: api.review.v1.Review.ListReview:output_type -> api.review.v1.ListReviewReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_review_v1_review_proto_init() }
func file_api_review_v1_review_proto_init() {
	if File_api_review_v1_review_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_review_v1_review_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReviewRequest); i {
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
		file_api_review_v1_review_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReviewReply); i {
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
		file_api_review_v1_review_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReviewRequest); i {
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
		file_api_review_v1_review_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReviewReply); i {
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
		file_api_review_v1_review_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReviewRequest); i {
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
		file_api_review_v1_review_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReviewReply); i {
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
		file_api_review_v1_review_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReviewRequest); i {
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
		file_api_review_v1_review_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReviewReply); i {
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
		file_api_review_v1_review_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReviewRequest); i {
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
		file_api_review_v1_review_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReviewReply); i {
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
			RawDescriptor: file_api_review_v1_review_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_review_v1_review_proto_goTypes,
		DependencyIndexes: file_api_review_v1_review_proto_depIdxs,
		MessageInfos:      file_api_review_v1_review_proto_msgTypes,
	}.Build()
	File_api_review_v1_review_proto = out.File
	file_api_review_v1_review_proto_rawDesc = nil
	file_api_review_v1_review_proto_goTypes = nil
	file_api_review_v1_review_proto_depIdxs = nil
}