// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: proto/reserve/reserve.proto

package yuemnoi_reserve

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

type RequestStatus int32

const (
	RequestStatus_REJECT  RequestStatus = 0
	RequestStatus_PENDING RequestStatus = 1
	RequestStatus_ACCEPT  RequestStatus = 2
)

// Enum value maps for RequestStatus.
var (
	RequestStatus_name = map[int32]string{
		0: "REJECT",
		1: "PENDING",
		2: "ACCEPT",
	}
	RequestStatus_value = map[string]int32{
		"REJECT":  0,
		"PENDING": 1,
		"ACCEPT":  2,
	}
)

func (x RequestStatus) Enum() *RequestStatus {
	p := new(RequestStatus)
	*p = x
	return p
}

func (x RequestStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_reserve_reserve_proto_enumTypes[0].Descriptor()
}

func (RequestStatus) Type() protoreflect.EnumType {
	return &file_proto_reserve_reserve_proto_enumTypes[0]
}

func (x RequestStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestStatus.Descriptor instead.
func (RequestStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_reserve_reserve_proto_rawDescGZIP(), []int{0}
}

type BorrowingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	LendingUserId   string        `protobuf:"bytes,2,opt,name=lending_user_id,json=lendingUserId,proto3" json:"lending_user_id,omitempty"`
	BorrowingUserId string        `protobuf:"bytes,3,opt,name=borrowing_user_id,json=borrowingUserId,proto3" json:"borrowing_user_id,omitempty"`
	LendingPostId   string        `protobuf:"bytes,4,opt,name=lending_post_id,json=lendingPostId,proto3" json:"lending_post_id,omitempty"`
	BorrowingPostId string        `protobuf:"bytes,5,opt,name=borrowing_post_id,json=borrowingPostId,proto3" json:"borrowing_post_id,omitempty"`
	Status          RequestStatus `protobuf:"varint,6,opt,name=status,proto3,enum=RequestStatus" json:"status,omitempty"`
	ActiveStatus    bool          `protobuf:"varint,7,opt,name=active_status,json=activeStatus,proto3" json:"active_status,omitempty"`
}

func (x *BorrowingRequest) Reset() {
	*x = BorrowingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reserve_reserve_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BorrowingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowingRequest) ProtoMessage() {}

func (x *BorrowingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reserve_reserve_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowingRequest.ProtoReflect.Descriptor instead.
func (*BorrowingRequest) Descriptor() ([]byte, []int) {
	return file_proto_reserve_reserve_proto_rawDescGZIP(), []int{0}
}

func (x *BorrowingRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BorrowingRequest) GetLendingUserId() string {
	if x != nil {
		return x.LendingUserId
	}
	return ""
}

func (x *BorrowingRequest) GetBorrowingUserId() string {
	if x != nil {
		return x.BorrowingUserId
	}
	return ""
}

func (x *BorrowingRequest) GetLendingPostId() string {
	if x != nil {
		return x.LendingPostId
	}
	return ""
}

func (x *BorrowingRequest) GetBorrowingPostId() string {
	if x != nil {
		return x.BorrowingPostId
	}
	return ""
}

func (x *BorrowingRequest) GetStatus() RequestStatus {
	if x != nil {
		return x.Status
	}
	return RequestStatus_REJECT
}

func (x *BorrowingRequest) GetActiveStatus() bool {
	if x != nil {
		return x.ActiveStatus
	}
	return false
}

type CreateRequestFromBorrowingPostInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LendingUserId   string `protobuf:"bytes,1,opt,name=lending_user_id,json=lendingUserId,proto3" json:"lending_user_id,omitempty"`
	BorrowingUserId string `protobuf:"bytes,2,opt,name=borrowing_user_id,json=borrowingUserId,proto3" json:"borrowing_user_id,omitempty"`
	LendingPostId   string `protobuf:"bytes,3,opt,name=lending_post_id,json=lendingPostId,proto3" json:"lending_post_id,omitempty"`
	BorrowingPostId string `protobuf:"bytes,4,opt,name=borrowing_post_id,json=borrowingPostId,proto3" json:"borrowing_post_id,omitempty"`
}

func (x *CreateRequestFromBorrowingPostInput) Reset() {
	*x = CreateRequestFromBorrowingPostInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reserve_reserve_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequestFromBorrowingPostInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequestFromBorrowingPostInput) ProtoMessage() {}

func (x *CreateRequestFromBorrowingPostInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reserve_reserve_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequestFromBorrowingPostInput.ProtoReflect.Descriptor instead.
func (*CreateRequestFromBorrowingPostInput) Descriptor() ([]byte, []int) {
	return file_proto_reserve_reserve_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRequestFromBorrowingPostInput) GetLendingUserId() string {
	if x != nil {
		return x.LendingUserId
	}
	return ""
}

func (x *CreateRequestFromBorrowingPostInput) GetBorrowingUserId() string {
	if x != nil {
		return x.BorrowingUserId
	}
	return ""
}

func (x *CreateRequestFromBorrowingPostInput) GetLendingPostId() string {
	if x != nil {
		return x.LendingPostId
	}
	return ""
}

func (x *CreateRequestFromBorrowingPostInput) GetBorrowingPostId() string {
	if x != nil {
		return x.BorrowingPostId
	}
	return ""
}

type CreateRequestFromBorrowingPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateRequestFromBorrowingPostResponse) Reset() {
	*x = CreateRequestFromBorrowingPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reserve_reserve_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequestFromBorrowingPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequestFromBorrowingPostResponse) ProtoMessage() {}

func (x *CreateRequestFromBorrowingPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reserve_reserve_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequestFromBorrowingPostResponse.ProtoReflect.Descriptor instead.
func (*CreateRequestFromBorrowingPostResponse) Descriptor() ([]byte, []int) {
	return file_proto_reserve_reserve_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRequestFromBorrowingPostResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetBorrowingRequestInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBorrowingRequestInput) Reset() {
	*x = GetBorrowingRequestInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reserve_reserve_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBorrowingRequestInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBorrowingRequestInput) ProtoMessage() {}

func (x *GetBorrowingRequestInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reserve_reserve_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBorrowingRequestInput.ProtoReflect.Descriptor instead.
func (*GetBorrowingRequestInput) Descriptor() ([]byte, []int) {
	return file_proto_reserve_reserve_proto_rawDescGZIP(), []int{3}
}

func (x *GetBorrowingRequestInput) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ConfirmBorrowingRequestInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ConfirmBorrowingRequestInput) Reset() {
	*x = ConfirmBorrowingRequestInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reserve_reserve_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmBorrowingRequestInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmBorrowingRequestInput) ProtoMessage() {}

func (x *ConfirmBorrowingRequestInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reserve_reserve_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmBorrowingRequestInput.ProtoReflect.Descriptor instead.
func (*ConfirmBorrowingRequestInput) Descriptor() ([]byte, []int) {
	return file_proto_reserve_reserve_proto_rawDescGZIP(), []int{4}
}

func (x *ConfirmBorrowingRequestInput) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RejectBorrowingRequestInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RejectBorrowingRequestInput) Reset() {
	*x = RejectBorrowingRequestInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reserve_reserve_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RejectBorrowingRequestInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejectBorrowingRequestInput) ProtoMessage() {}

func (x *RejectBorrowingRequestInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reserve_reserve_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejectBorrowingRequestInput.ProtoReflect.Descriptor instead.
func (*RejectBorrowingRequestInput) Descriptor() ([]byte, []int) {
	return file_proto_reserve_reserve_proto_rawDescGZIP(), []int{5}
}

func (x *RejectBorrowingRequestInput) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ReturnItemRequestInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReturnItemRequestInput) Reset() {
	*x = ReturnItemRequestInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reserve_reserve_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReturnItemRequestInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnItemRequestInput) ProtoMessage() {}

func (x *ReturnItemRequestInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reserve_reserve_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnItemRequestInput.ProtoReflect.Descriptor instead.
func (*ReturnItemRequestInput) Descriptor() ([]byte, []int) {
	return file_proto_reserve_reserve_proto_rawDescGZIP(), []int{6}
}

func (x *ReturnItemRequestInput) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_proto_reserve_reserve_proto protoreflect.FileDescriptor

var file_proto_reserve_reserve_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2f,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x02,
	0x0a, 0x10, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x62, 0x6f,
	0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6c, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x2a,
	0x0a, 0x11, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x6f, 0x72, 0x72, 0x6f,
	0x77, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xcd, 0x01, 0x0a, 0x23, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x6f, 0x72,
	0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x26, 0x0a, 0x0f, 0x6c, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x62, 0x6f, 0x72, 0x72, 0x6f,
	0x77, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x62,
	0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e,
	0x67, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x26, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x6f, 0x72, 0x72,
	0x6f, 0x77, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x1c, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x1b, 0x52, 0x65, 0x6a, 0x65, 0x63,
	0x74, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x16, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x2a, 0x34, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43,
	0x43, 0x45, 0x50, 0x54, 0x10, 0x02, 0x32, 0xa5, 0x03, 0x0a, 0x10, 0x42, 0x6f, 0x72, 0x72, 0x6f,
	0x77, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6f, 0x0a, 0x1e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x72, 0x6f, 0x6d,
	0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x24, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x72, 0x6f,
	0x6d, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x27, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x19, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x72,
	0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x11, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x11, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x49, 0x0a, 0x16, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x6f, 0x72, 0x72,
	0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x52,
	0x65, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x11, 0x2e, 0x42, 0x6f, 0x72,
	0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a,
	0x11, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x11, 0x2e, 0x42, 0x6f,
	0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x26,
	0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x4b, 0x68,
	0x69, 0x6d, 0x6d, 0x6f, 0x6f, 0x6e, 0x2f, 0x79, 0x75, 0x65, 0x6d, 0x6e, 0x6f, 0x69, 0x2d, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_reserve_reserve_proto_rawDescOnce sync.Once
	file_proto_reserve_reserve_proto_rawDescData = file_proto_reserve_reserve_proto_rawDesc
)

func file_proto_reserve_reserve_proto_rawDescGZIP() []byte {
	file_proto_reserve_reserve_proto_rawDescOnce.Do(func() {
		file_proto_reserve_reserve_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_reserve_reserve_proto_rawDescData)
	})
	return file_proto_reserve_reserve_proto_rawDescData
}

var file_proto_reserve_reserve_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_reserve_reserve_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_reserve_reserve_proto_goTypes = []any{
	(RequestStatus)(0),                             // 0: RequestStatus
	(*BorrowingRequest)(nil),                       // 1: BorrowingRequest
	(*CreateRequestFromBorrowingPostInput)(nil),    // 2: CreateRequestFromBorrowingPostInput
	(*CreateRequestFromBorrowingPostResponse)(nil), // 3: CreateRequestFromBorrowingPostResponse
	(*GetBorrowingRequestInput)(nil),               // 4: GetBorrowingRequestInput
	(*ConfirmBorrowingRequestInput)(nil),           // 5: ConfirmBorrowingRequestInput
	(*RejectBorrowingRequestInput)(nil),            // 6: RejectBorrowingRequestInput
	(*ReturnItemRequestInput)(nil),                 // 7: ReturnItemRequestInput
}
var file_proto_reserve_reserve_proto_depIdxs = []int32{
	0, // 0: BorrowingRequest.status:type_name -> RequestStatus
	2, // 1: BorrowingService.CreateRequestFromBorrowingPost:input_type -> CreateRequestFromBorrowingPostInput
	4, // 2: BorrowingService.GetBorrowingRequestById:input_type -> GetBorrowingRequestInput
	5, // 3: BorrowingService.ConfirmBorrowingRequest:input_type -> ConfirmBorrowingRequestInput
	6, // 4: BorrowingService.RejectBorrowingRequest:input_type -> RejectBorrowingRequestInput
	7, // 5: BorrowingService.ReturnItemRequest:input_type -> ReturnItemRequestInput
	3, // 6: BorrowingService.CreateRequestFromBorrowingPost:output_type -> CreateRequestFromBorrowingPostResponse
	1, // 7: BorrowingService.GetBorrowingRequestById:output_type -> BorrowingRequest
	1, // 8: BorrowingService.ConfirmBorrowingRequest:output_type -> BorrowingRequest
	1, // 9: BorrowingService.RejectBorrowingRequest:output_type -> BorrowingRequest
	1, // 10: BorrowingService.ReturnItemRequest:output_type -> BorrowingRequest
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_reserve_reserve_proto_init() }
func file_proto_reserve_reserve_proto_init() {
	if File_proto_reserve_reserve_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_reserve_reserve_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BorrowingRequest); i {
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
		file_proto_reserve_reserve_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRequestFromBorrowingPostInput); i {
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
		file_proto_reserve_reserve_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRequestFromBorrowingPostResponse); i {
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
		file_proto_reserve_reserve_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetBorrowingRequestInput); i {
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
		file_proto_reserve_reserve_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ConfirmBorrowingRequestInput); i {
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
		file_proto_reserve_reserve_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*RejectBorrowingRequestInput); i {
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
		file_proto_reserve_reserve_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ReturnItemRequestInput); i {
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
			RawDescriptor: file_proto_reserve_reserve_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_reserve_reserve_proto_goTypes,
		DependencyIndexes: file_proto_reserve_reserve_proto_depIdxs,
		EnumInfos:         file_proto_reserve_reserve_proto_enumTypes,
		MessageInfos:      file_proto_reserve_reserve_proto_msgTypes,
	}.Build()
	File_proto_reserve_reserve_proto = out.File
	file_proto_reserve_reserve_proto_rawDesc = nil
	file_proto_reserve_reserve_proto_goTypes = nil
	file_proto_reserve_reserve_proto_depIdxs = nil
}
