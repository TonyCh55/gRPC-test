// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: countries.proto

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_countries_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_countries_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_countries_proto_rawDescGZIP(), []int{0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_countries_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_countries_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_countries_proto_rawDescGZIP(), []int{1}
}

type Response_Country struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Response_Country) Reset() {
	*x = Response_Country{}
	if protoimpl.UnsafeEnabled {
		mi := &file_countries_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_Country) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_Country) ProtoMessage() {}

func (x *Response_Country) ProtoReflect() protoreflect.Message {
	mi := &file_countries_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_Country.ProtoReflect.Descriptor instead.
func (*Response_Country) Descriptor() ([]byte, []int) {
	return file_countries_proto_rawDescGZIP(), []int{0, 0}
}

type Response_Country_Single struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Likes int32  `protobuf:"varint,3,opt,name=likes,proto3" json:"likes,omitempty"`
}

func (x *Response_Country_Single) Reset() {
	*x = Response_Country_Single{}
	if protoimpl.UnsafeEnabled {
		mi := &file_countries_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_Country_Single) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_Country_Single) ProtoMessage() {}

func (x *Response_Country_Single) ProtoReflect() protoreflect.Message {
	mi := &file_countries_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_Country_Single.ProtoReflect.Descriptor instead.
func (*Response_Country_Single) Descriptor() ([]byte, []int) {
	return file_countries_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Response_Country_Single) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Response_Country_Single) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Response_Country_Single) GetLikes() int32 {
	if x != nil {
		return x.Likes
	}
	return 0
}

type Response_Country_List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Countries []*Response_Country_Single `protobuf:"bytes,1,rep,name=countries,proto3" json:"countries,omitempty"`
}

func (x *Response_Country_List) Reset() {
	*x = Response_Country_List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_countries_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_Country_List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_Country_List) ProtoMessage() {}

func (x *Response_Country_List) ProtoReflect() protoreflect.Message {
	mi := &file_countries_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_Country_List.ProtoReflect.Descriptor instead.
func (*Response_Country_List) Descriptor() ([]byte, []int) {
	return file_countries_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *Response_Country_List) GetCountries() []*Response_Country_Single {
	if x != nil {
		return x.Countries
	}
	return nil
}

type Request_Country struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Request_Country) Reset() {
	*x = Request_Country{}
	if protoimpl.UnsafeEnabled {
		mi := &file_countries_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Country) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Country) ProtoMessage() {}

func (x *Request_Country) ProtoReflect() protoreflect.Message {
	mi := &file_countries_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_Country.ProtoReflect.Descriptor instead.
func (*Request_Country) Descriptor() ([]byte, []int) {
	return file_countries_proto_rawDescGZIP(), []int{1, 0}
}

type Request_Country_Single struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Request_Country_Single) Reset() {
	*x = Request_Country_Single{}
	if protoimpl.UnsafeEnabled {
		mi := &file_countries_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Country_Single) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Country_Single) ProtoMessage() {}

func (x *Request_Country_Single) ProtoReflect() protoreflect.Message {
	mi := &file_countries_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_Country_Single.ProtoReflect.Descriptor instead.
func (*Request_Country_Single) Descriptor() ([]byte, []int) {
	return file_countries_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *Request_Country_Single) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Request_Country_List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset uint32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *Request_Country_List) Reset() {
	*x = Request_Country_List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_countries_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Country_List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Country_List) ProtoMessage() {}

func (x *Request_Country_List) ProtoReflect() protoreflect.Message {
	mi := &file_countries_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_Country_List.ProtoReflect.Descriptor instead.
func (*Request_Country_List) Descriptor() ([]byte, []int) {
	return file_countries_proto_rawDescGZIP(), []int{1, 0, 1}
}

func (x *Request_Country_List) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Request_Country_List) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type Request_Country_Like struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Request_Country_Like) Reset() {
	*x = Request_Country_Like{}
	if protoimpl.UnsafeEnabled {
		mi := &file_countries_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Country_Like) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Country_Like) ProtoMessage() {}

func (x *Request_Country_Like) ProtoReflect() protoreflect.Message {
	mi := &file_countries_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_Country_Like.ProtoReflect.Descriptor instead.
func (*Request_Country_Like) Descriptor() ([]byte, []int) {
	return file_countries_proto_rawDescGZIP(), []int{1, 0, 2}
}

func (x *Request_Country_Like) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Request_Country_Dislike struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Request_Country_Dislike) Reset() {
	*x = Request_Country_Dislike{}
	if protoimpl.UnsafeEnabled {
		mi := &file_countries_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Country_Dislike) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Country_Dislike) ProtoMessage() {}

func (x *Request_Country_Dislike) ProtoReflect() protoreflect.Message {
	mi := &file_countries_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_Country_Dislike.ProtoReflect.Descriptor instead.
func (*Request_Country_Dislike) Descriptor() ([]byte, []int) {
	return file_countries_proto_rawDescGZIP(), []int{1, 0, 3}
}

func (x *Request_Country_Dislike) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Request_Country_LikeDislike struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Action:
	//	*Request_Country_LikeDislike_Like
	//	*Request_Country_LikeDislike_Dislike
	Action isRequest_Country_LikeDislike_Action `protobuf_oneof:"action"`
}

func (x *Request_Country_LikeDislike) Reset() {
	*x = Request_Country_LikeDislike{}
	if protoimpl.UnsafeEnabled {
		mi := &file_countries_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Country_LikeDislike) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Country_LikeDislike) ProtoMessage() {}

func (x *Request_Country_LikeDislike) ProtoReflect() protoreflect.Message {
	mi := &file_countries_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_Country_LikeDislike.ProtoReflect.Descriptor instead.
func (*Request_Country_LikeDislike) Descriptor() ([]byte, []int) {
	return file_countries_proto_rawDescGZIP(), []int{1, 0, 4}
}

func (m *Request_Country_LikeDislike) GetAction() isRequest_Country_LikeDislike_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *Request_Country_LikeDislike) GetLike() *Request_Country_Like {
	if x, ok := x.GetAction().(*Request_Country_LikeDislike_Like); ok {
		return x.Like
	}
	return nil
}

func (x *Request_Country_LikeDislike) GetDislike() *Request_Country_Dislike {
	if x, ok := x.GetAction().(*Request_Country_LikeDislike_Dislike); ok {
		return x.Dislike
	}
	return nil
}

type isRequest_Country_LikeDislike_Action interface {
	isRequest_Country_LikeDislike_Action()
}

type Request_Country_LikeDislike_Like struct {
	Like *Request_Country_Like `protobuf:"bytes,1,opt,name=like,proto3,oneof"`
}

type Request_Country_LikeDislike_Dislike struct {
	Dislike *Request_Country_Dislike `protobuf:"bytes,2,opt,name=dislike,proto3,oneof"`
}

func (*Request_Country_LikeDislike_Like) isRequest_Country_LikeDislike_Action() {}

func (*Request_Country_LikeDislike_Dislike) isRequest_Country_LikeDislike_Action() {}

var File_countries_proto protoreflect.FileDescriptor

var file_countries_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xa2, 0x01, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x95, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x1a, 0x42, 0x0a, 0x06, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x1a, 0x46, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x3e, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22,
	0xa5, 0x02, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x99, 0x02, 0x0a, 0x07,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x1a, 0x18, 0x0a, 0x06, 0x53, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x1a, 0x34, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x1a, 0x16, 0x0a, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x1a,
	0x19, 0x0a, 0x07, 0x44, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x1a, 0x8a, 0x01, 0x0a, 0x0b, 0x4c,
	0x69, 0x6b, 0x65, 0x44, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x6c, 0x69,
	0x6b, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x48, 0x00, 0x52, 0x04, 0x6c, 0x69, 0x6b, 0x65, 0x12,
	0x3c, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x44, 0x69, 0x73, 0x6c, 0x69,
	0x6b, 0x65, 0x48, 0x00, 0x52, 0x07, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x42, 0x08, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xc1, 0x02, 0x0a, 0x09, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x4f,
	0x0a, 0x06, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x49, 0x0a, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x07, 0x44, 0x69,
	0x73, 0x6c, 0x69, 0x6b, 0x65, 0x12, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e,
	0x44, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x70, 0x61, 0x74, 0x68, 0x2f, 0x67, 0x65, 0x6e, 0x3b, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_countries_proto_rawDescOnce sync.Once
	file_countries_proto_rawDescData = file_countries_proto_rawDesc
)

func file_countries_proto_rawDescGZIP() []byte {
	file_countries_proto_rawDescOnce.Do(func() {
		file_countries_proto_rawDescData = protoimpl.X.CompressGZIP(file_countries_proto_rawDescData)
	})
	return file_countries_proto_rawDescData
}

var file_countries_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_countries_proto_goTypes = []interface{}{
	(*Response)(nil),                    // 0: service.Response
	(*Request)(nil),                     // 1: service.Request
	(*Response_Country)(nil),            // 2: service.Response.Country
	(*Response_Country_Single)(nil),     // 3: service.Response.Country.Single
	(*Response_Country_List)(nil),       // 4: service.Response.Country.List
	(*Request_Country)(nil),             // 5: service.Request.Country
	(*Request_Country_Single)(nil),      // 6: service.Request.Country.Single
	(*Request_Country_List)(nil),        // 7: service.Request.Country.List
	(*Request_Country_Like)(nil),        // 8: service.Request.Country.Like
	(*Request_Country_Dislike)(nil),     // 9: service.Request.Country.Dislike
	(*Request_Country_LikeDislike)(nil), // 10: service.Request.Country.LikeDislike
}
var file_countries_proto_depIdxs = []int32{
	3, // 0: service.Response.Country.List.countries:type_name -> service.Response.Country.Single
	8, // 1: service.Request.Country.LikeDislike.like:type_name -> service.Request.Country.Like
	9, // 2: service.Request.Country.LikeDislike.dislike:type_name -> service.Request.Country.Dislike
	7, // 3: service.Countries.List:input_type -> service.Request.Country.List
	6, // 4: service.Countries.Single:input_type -> service.Request.Country.Single
	8, // 5: service.Countries.Like:input_type -> service.Request.Country.Like
	9, // 6: service.Countries.Dislike:input_type -> service.Request.Country.Dislike
	4, // 7: service.Countries.List:output_type -> service.Response.Country.List
	3, // 8: service.Countries.Single:output_type -> service.Response.Country.Single
	3, // 9: service.Countries.Like:output_type -> service.Response.Country.Single
	3, // 10: service.Countries.Dislike:output_type -> service.Response.Country.Single
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_countries_proto_init() }
func file_countries_proto_init() {
	if File_countries_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_countries_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_countries_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_countries_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_Country); i {
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
		file_countries_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_Country_Single); i {
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
		file_countries_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_Country_List); i {
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
		file_countries_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_Country); i {
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
		file_countries_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_Country_Single); i {
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
		file_countries_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_Country_List); i {
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
		file_countries_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_Country_Like); i {
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
		file_countries_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_Country_Dislike); i {
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
		file_countries_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_Country_LikeDislike); i {
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
	file_countries_proto_msgTypes[10].OneofWrappers = []interface{}{
		(*Request_Country_LikeDislike_Like)(nil),
		(*Request_Country_LikeDislike_Dislike)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_countries_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_countries_proto_goTypes,
		DependencyIndexes: file_countries_proto_depIdxs,
		MessageInfos:      file_countries_proto_msgTypes,
	}.Build()
	File_countries_proto = out.File
	file_countries_proto_rawDesc = nil
	file_countries_proto_goTypes = nil
	file_countries_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CountriesClient is the client API for Countries service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CountriesClient interface {
	List(ctx context.Context, in *Request_Country_List, opts ...grpc.CallOption) (*Response_Country_List, error)
	Single(ctx context.Context, in *Request_Country_Single, opts ...grpc.CallOption) (Countries_SingleClient, error)
	Like(ctx context.Context, in *Request_Country_Like, opts ...grpc.CallOption) (*Response_Country_Single, error)
	Dislike(ctx context.Context, in *Request_Country_Dislike, opts ...grpc.CallOption) (*Response_Country_Single, error)
}

type countriesClient struct {
	cc grpc.ClientConnInterface
}

func NewCountriesClient(cc grpc.ClientConnInterface) CountriesClient {
	return &countriesClient{cc}
}

func (c *countriesClient) List(ctx context.Context, in *Request_Country_List, opts ...grpc.CallOption) (*Response_Country_List, error) {
	out := new(Response_Country_List)
	err := c.cc.Invoke(ctx, "/service.Countries/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countriesClient) Single(ctx context.Context, in *Request_Country_Single, opts ...grpc.CallOption) (Countries_SingleClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Countries_serviceDesc.Streams[0], "/service.Countries/Single", opts...)
	if err != nil {
		return nil, err
	}
	x := &countriesSingleClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Countries_SingleClient interface {
	Recv() (*Response_Country_Single, error)
	grpc.ClientStream
}

type countriesSingleClient struct {
	grpc.ClientStream
}

func (x *countriesSingleClient) Recv() (*Response_Country_Single, error) {
	m := new(Response_Country_Single)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *countriesClient) Like(ctx context.Context, in *Request_Country_Like, opts ...grpc.CallOption) (*Response_Country_Single, error) {
	out := new(Response_Country_Single)
	err := c.cc.Invoke(ctx, "/service.Countries/Like", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countriesClient) Dislike(ctx context.Context, in *Request_Country_Dislike, opts ...grpc.CallOption) (*Response_Country_Single, error) {
	out := new(Response_Country_Single)
	err := c.cc.Invoke(ctx, "/service.Countries/Dislike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CountriesServer is the server API for Countries service.
type CountriesServer interface {
	List(context.Context, *Request_Country_List) (*Response_Country_List, error)
	Single(*Request_Country_Single, Countries_SingleServer) error
	Like(context.Context, *Request_Country_Like) (*Response_Country_Single, error)
	Dislike(context.Context, *Request_Country_Dislike) (*Response_Country_Single, error)
}

// UnimplementedCountriesServer can be embedded to have forward compatible implementations.
type UnimplementedCountriesServer struct {
}

func (*UnimplementedCountriesServer) List(context.Context, *Request_Country_List) (*Response_Country_List, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedCountriesServer) Single(*Request_Country_Single, Countries_SingleServer) error {
	return status.Errorf(codes.Unimplemented, "method Single not implemented")
}
func (*UnimplementedCountriesServer) Like(context.Context, *Request_Country_Like) (*Response_Country_Single, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Like not implemented")
}
func (*UnimplementedCountriesServer) Dislike(context.Context, *Request_Country_Dislike) (*Response_Country_Single, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dislike not implemented")
}

func RegisterCountriesServer(s *grpc.Server, srv CountriesServer) {
	s.RegisterService(&_Countries_serviceDesc, srv)
}

func _Countries_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request_Country_List)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountriesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Countries/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountriesServer).List(ctx, req.(*Request_Country_List))
	}
	return interceptor(ctx, in, info, handler)
}

func _Countries_Single_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request_Country_Single)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CountriesServer).Single(m, &countriesSingleServer{stream})
}

type Countries_SingleServer interface {
	Send(*Response_Country_Single) error
	grpc.ServerStream
}

type countriesSingleServer struct {
	grpc.ServerStream
}

func (x *countriesSingleServer) Send(m *Response_Country_Single) error {
	return x.ServerStream.SendMsg(m)
}

func _Countries_Like_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request_Country_Like)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountriesServer).Like(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Countries/Like",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountriesServer).Like(ctx, req.(*Request_Country_Like))
	}
	return interceptor(ctx, in, info, handler)
}

func _Countries_Dislike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request_Country_Dislike)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountriesServer).Dislike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Countries/Dislike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountriesServer).Dislike(ctx, req.(*Request_Country_Dislike))
	}
	return interceptor(ctx, in, info, handler)
}

var _Countries_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Countries",
	HandlerType: (*CountriesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Countries_List_Handler,
		},
		{
			MethodName: "Like",
			Handler:    _Countries_Like_Handler,
		},
		{
			MethodName: "Dislike",
			Handler:    _Countries_Dislike_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Single",
			Handler:       _Countries_Single_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "countries.proto",
}
