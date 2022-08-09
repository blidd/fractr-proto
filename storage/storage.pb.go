// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: storage/storage.proto

package storage

import (
	marketplace_common "github.com/blidd/fractr-proto/marketplace_common"
	service "github.com/blidd/fractr-proto/service"
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

type Type int32

const (
	Type_ARTWORK  Type = 0
	Type_BID      Type = 1
	Type_ASK      Type = 2
	Type_USER     Type = 3
	Type_USERBIDS Type = 4
	Type_USERASKS Type = 5
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "ARTWORK",
		1: "BID",
		2: "ASK",
		3: "USER",
		4: "USERBIDS",
		5: "USERASKS",
	}
	Type_value = map[string]int32{
		"ARTWORK":  0,
		"BID":      1,
		"ASK":      2,
		"USER":     3,
		"USERBIDS": 4,
		"USERASKS": 5,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_storage_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_storage_storage_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_storage_storage_proto_rawDescGZIP(), []int{0}
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type          Type     `protobuf:"varint,1,opt,name=type,proto3,enum=storage.Type" json:"type,omitempty"`
	Id            *uint32  `protobuf:"varint,2,opt,name=id,proto3,oneof" json:"id,omitempty"`                                            // optional field
	Ids           []uint32 `protobuf:"varint,3,rep,packed,name=ids,proto3" json:"ids,omitempty"`                                         // optional field
	PrimaryMarket *bool    `protobuf:"varint,4,opt,name=primary_market,json=primaryMarket,proto3,oneof" json:"primary_market,omitempty"` // optional field
	ArtistName    *string  `protobuf:"bytes,5,opt,name=artist_name,json=artistName,proto3,oneof" json:"artist_name,omitempty"`           // optional field
	ArtTitle      *string  `protobuf:"bytes,6,opt,name=art_title,json=artTitle,proto3,oneof" json:"art_title,omitempty"`                 // optional field
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_storage_storage_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_ARTWORK
}

func (x *GetRequest) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *GetRequest) GetIds() []uint32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *GetRequest) GetPrimaryMarket() bool {
	if x != nil && x.PrimaryMarket != nil {
		return *x.PrimaryMarket
	}
	return false
}

func (x *GetRequest) GetArtistName() string {
	if x != nil && x.ArtistName != nil {
		return *x.ArtistName
	}
	return ""
}

func (x *GetRequest) GetArtTitle() string {
	if x != nil && x.ArtTitle != nil {
		return *x.ArtTitle
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        Type                            `protobuf:"varint,1,opt,name=type,proto3,enum=storage.Type" json:"type,omitempty"`
	Artworks    []*service.Artwork              `protobuf:"bytes,2,rep,name=artworks,proto3" json:"artworks,omitempty"`
	Profile     *service.UserProfile            `protobuf:"bytes,3,opt,name=profile,proto3,oneof" json:"profile,omitempty"`
	BidStatuses []*marketplace_common.BidStatus `protobuf:"bytes,4,rep,name=bid_statuses,json=bidStatuses,proto3" json:"bid_statuses,omitempty"`
	AskStatuses []*marketplace_common.AskStatus `protobuf:"bytes,5,rep,name=ask_statuses,json=askStatuses,proto3" json:"ask_statuses,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_storage_storage_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_ARTWORK
}

func (x *GetResponse) GetArtworks() []*service.Artwork {
	if x != nil {
		return x.Artworks
	}
	return nil
}

func (x *GetResponse) GetProfile() *service.UserProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *GetResponse) GetBidStatuses() []*marketplace_common.BidStatus {
	if x != nil {
		return x.BidStatuses
	}
	return nil
}

func (x *GetResponse) GetAskStatuses() []*marketplace_common.AskStatus {
	if x != nil {
		return x.AskStatuses
	}
	return nil
}

type PutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type          Type                          `protobuf:"varint,1,opt,name=type,proto3,enum=storage.Type" json:"type,omitempty"`
	Id            *uint32                       `protobuf:"varint,2,opt,name=id,proto3,oneof" json:"id,omitempty"`                                            // optional field
	Ids           []uint32                      `protobuf:"varint,3,rep,packed,name=ids,proto3" json:"ids,omitempty"`                                         // optional field
	PrimaryMarket *bool                         `protobuf:"varint,4,opt,name=primary_market,json=primaryMarket,proto3,oneof" json:"primary_market,omitempty"` // optional field
	ArtistName    *string                       `protobuf:"bytes,5,opt,name=artist_name,json=artistName,proto3,oneof" json:"artist_name,omitempty"`           // optional field
	ArtTitle      *string                       `protobuf:"bytes,6,opt,name=art_title,json=artTitle,proto3,oneof" json:"art_title,omitempty"`                 // optional field
	BidStatus     *marketplace_common.BidStatus `protobuf:"bytes,7,opt,name=bid_status,json=bidStatus,proto3,oneof" json:"bid_status,omitempty"`
	AskStatus     *marketplace_common.AskStatus `protobuf:"bytes,8,opt,name=ask_status,json=askStatus,proto3,oneof" json:"ask_status,omitempty"`
}

func (x *PutRequest) Reset() {
	*x = PutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutRequest) ProtoMessage() {}

func (x *PutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutRequest.ProtoReflect.Descriptor instead.
func (*PutRequest) Descriptor() ([]byte, []int) {
	return file_storage_storage_proto_rawDescGZIP(), []int{2}
}

func (x *PutRequest) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_ARTWORK
}

func (x *PutRequest) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *PutRequest) GetIds() []uint32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *PutRequest) GetPrimaryMarket() bool {
	if x != nil && x.PrimaryMarket != nil {
		return *x.PrimaryMarket
	}
	return false
}

func (x *PutRequest) GetArtistName() string {
	if x != nil && x.ArtistName != nil {
		return *x.ArtistName
	}
	return ""
}

func (x *PutRequest) GetArtTitle() string {
	if x != nil && x.ArtTitle != nil {
		return *x.ArtTitle
	}
	return ""
}

func (x *PutRequest) GetBidStatus() *marketplace_common.BidStatus {
	if x != nil {
		return x.BidStatus
	}
	return nil
}

func (x *PutRequest) GetAskStatus() *marketplace_common.AskStatus {
	if x != nil {
		return x.AskStatus
	}
	return nil
}

type PutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *PutResponse) Reset() {
	*x = PutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_storage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutResponse) ProtoMessage() {}

func (x *PutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_storage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutResponse.ProtoReflect.Descriptor instead.
func (*PutResponse) Descriptor() ([]byte, []int) {
	return file_storage_storage_proto_rawDescGZIP(), []int{3}
}

func (x *PutResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type          Type     `protobuf:"varint,1,opt,name=type,proto3,enum=storage.Type" json:"type,omitempty"`
	Id            *uint32  `protobuf:"varint,2,opt,name=id,proto3,oneof" json:"id,omitempty"`                                            // optional field
	Ids           []uint32 `protobuf:"varint,3,rep,packed,name=ids,proto3" json:"ids,omitempty"`                                         // optional field
	PrimaryMarket *bool    `protobuf:"varint,4,opt,name=primary_market,json=primaryMarket,proto3,oneof" json:"primary_market,omitempty"` // optional field
	ArtistName    *string  `protobuf:"bytes,5,opt,name=artist_name,json=artistName,proto3,oneof" json:"artist_name,omitempty"`           // optional field
	ArtTitle      *string  `protobuf:"bytes,6,opt,name=art_title,json=artTitle,proto3,oneof" json:"art_title,omitempty"`                 // optional field
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_storage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_storage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_storage_storage_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRequest) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_ARTWORK
}

func (x *DeleteRequest) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *DeleteRequest) GetIds() []uint32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *DeleteRequest) GetPrimaryMarket() bool {
	if x != nil && x.PrimaryMarket != nil {
		return *x.PrimaryMarket
	}
	return false
}

func (x *DeleteRequest) GetArtistName() string {
	if x != nil && x.ArtistName != nil {
		return *x.ArtistName
	}
	return ""
}

func (x *DeleteRequest) GetArtTitle() string {
	if x != nil && x.ArtTitle != nil {
		return *x.ArtTitle
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_storage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_storage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_storage_storage_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_storage_storage_proto protoreflect.FileDescriptor

var file_storage_storage_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x1a, 0x2b, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x02, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x2a, 0x0a,
	0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x61, 0x72, 0x74,
	0x69, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02,
	0x52, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x20, 0x0a, 0x09, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x03, 0x52, 0x08, 0x61, 0x72, 0x74, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0xa3, 0x02, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x08,
	0x61, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x52, 0x08, 0x61, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x48, 0x00, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x40, 0x0a, 0x0c, 0x62, 0x69, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x69, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x0b, 0x62, 0x69, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65,
	0x73, 0x12, 0x40, 0x0a, 0x0c, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x73, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0b, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x65, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22,
	0xa6, 0x03, 0x0a, 0x0a, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52,
	0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x2a, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x01, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0a, 0x61, 0x72, 0x74,
	0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52,
	0x08, 0x61, 0x72, 0x74, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x41, 0x0a, 0x0a,
	0x62, 0x69, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x69, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48,
	0x04, 0x52, 0x09, 0x62, 0x69, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12,
	0x41, 0x0a, 0x0a, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x48, 0x05, 0x52, 0x09, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88,
	0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x62,
	0x69, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x61, 0x73,
	0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x27, 0x0a, 0x0b, 0x50, 0x75, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x85, 0x02, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x2a, 0x0a,
	0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x61, 0x72, 0x74,
	0x69, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02,
	0x52, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x20, 0x0a, 0x09, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x03, 0x52, 0x08, 0x61, 0x72, 0x74, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x2a, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x2a, 0x4b, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x41, 0x52, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x49,
	0x44, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x4b, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04,
	0x55, 0x53, 0x45, 0x52, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x53, 0x45, 0x52, 0x42, 0x49,
	0x44, 0x53, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x53, 0x45, 0x52, 0x41, 0x53, 0x4b, 0x53,
	0x10, 0x05, 0x32, 0xa8, 0x01, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x30,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x30, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x13, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x27, 0x5a,
	0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6c, 0x69, 0x64,
	0x64, 0x2f, 0x66, 0x72, 0x61, 0x63, 0x74, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_storage_proto_rawDescOnce sync.Once
	file_storage_storage_proto_rawDescData = file_storage_storage_proto_rawDesc
)

func file_storage_storage_proto_rawDescGZIP() []byte {
	file_storage_storage_proto_rawDescOnce.Do(func() {
		file_storage_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_storage_proto_rawDescData)
	})
	return file_storage_storage_proto_rawDescData
}

var file_storage_storage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_storage_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_storage_storage_proto_goTypes = []interface{}{
	(Type)(0),                            // 0: storage.Type
	(*GetRequest)(nil),                   // 1: storage.GetRequest
	(*GetResponse)(nil),                  // 2: storage.GetResponse
	(*PutRequest)(nil),                   // 3: storage.PutRequest
	(*PutResponse)(nil),                  // 4: storage.PutResponse
	(*DeleteRequest)(nil),                // 5: storage.DeleteRequest
	(*DeleteResponse)(nil),               // 6: storage.DeleteResponse
	(*service.Artwork)(nil),              // 7: service.Artwork
	(*service.UserProfile)(nil),          // 8: service.UserProfile
	(*marketplace_common.BidStatus)(nil), // 9: marketplace_common.BidStatus
	(*marketplace_common.AskStatus)(nil), // 10: marketplace_common.AskStatus
}
var file_storage_storage_proto_depIdxs = []int32{
	0,  // 0: storage.GetRequest.type:type_name -> storage.Type
	0,  // 1: storage.GetResponse.type:type_name -> storage.Type
	7,  // 2: storage.GetResponse.artworks:type_name -> service.Artwork
	8,  // 3: storage.GetResponse.profile:type_name -> service.UserProfile
	9,  // 4: storage.GetResponse.bid_statuses:type_name -> marketplace_common.BidStatus
	10, // 5: storage.GetResponse.ask_statuses:type_name -> marketplace_common.AskStatus
	0,  // 6: storage.PutRequest.type:type_name -> storage.Type
	9,  // 7: storage.PutRequest.bid_status:type_name -> marketplace_common.BidStatus
	10, // 8: storage.PutRequest.ask_status:type_name -> marketplace_common.AskStatus
	0,  // 9: storage.DeleteRequest.type:type_name -> storage.Type
	1,  // 10: storage.Storage.Get:input_type -> storage.GetRequest
	3,  // 11: storage.Storage.Put:input_type -> storage.PutRequest
	5,  // 12: storage.Storage.Delete:input_type -> storage.DeleteRequest
	2,  // 13: storage.Storage.Get:output_type -> storage.GetResponse
	4,  // 14: storage.Storage.Put:output_type -> storage.PutResponse
	6,  // 15: storage.Storage.Delete:output_type -> storage.DeleteResponse
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_storage_storage_proto_init() }
func file_storage_storage_proto_init() {
	if File_storage_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storage_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_storage_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_storage_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutRequest); i {
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
		file_storage_storage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutResponse); i {
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
		file_storage_storage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_storage_storage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
	file_storage_storage_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_storage_storage_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_storage_storage_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_storage_storage_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_storage_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_storage_storage_proto_goTypes,
		DependencyIndexes: file_storage_storage_proto_depIdxs,
		EnumInfos:         file_storage_storage_proto_enumTypes,
		MessageInfos:      file_storage_storage_proto_msgTypes,
	}.Build()
	File_storage_storage_proto = out.File
	file_storage_storage_proto_rawDesc = nil
	file_storage_storage_proto_goTypes = nil
	file_storage_storage_proto_depIdxs = nil
}
