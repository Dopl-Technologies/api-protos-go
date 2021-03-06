// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: catheter_path_service.proto

package dtprotos

import (
	context "context"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type CreateCatheterPathRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path *CatheterPath `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *CreateCatheterPathRequest) Reset() {
	*x = CreateCatheterPathRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catheter_path_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCatheterPathRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCatheterPathRequest) ProtoMessage() {}

func (x *CreateCatheterPathRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catheter_path_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCatheterPathRequest.ProtoReflect.Descriptor instead.
func (*CreateCatheterPathRequest) Descriptor() ([]byte, []int) {
	return file_catheter_path_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCatheterPathRequest) GetPath() *CatheterPath {
	if x != nil {
		return x.Path
	}
	return nil
}

type CreateCatheterPathResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateCatheterPathResponse) Reset() {
	*x = CreateCatheterPathResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catheter_path_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCatheterPathResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCatheterPathResponse) ProtoMessage() {}

func (x *CreateCatheterPathResponse) ProtoReflect() protoreflect.Message {
	mi := &file_catheter_path_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCatheterPathResponse.ProtoReflect.Descriptor instead.
func (*CreateCatheterPathResponse) Descriptor() ([]byte, []int) {
	return file_catheter_path_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCatheterPathResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCatheterPathRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCatheterPathRequest) Reset() {
	*x = GetCatheterPathRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catheter_path_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCatheterPathRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCatheterPathRequest) ProtoMessage() {}

func (x *GetCatheterPathRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catheter_path_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCatheterPathRequest.ProtoReflect.Descriptor instead.
func (*GetCatheterPathRequest) Descriptor() ([]byte, []int) {
	return file_catheter_path_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetCatheterPathRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCatheterPathResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path *CatheterPath `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *GetCatheterPathResponse) Reset() {
	*x = GetCatheterPathResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catheter_path_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCatheterPathResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCatheterPathResponse) ProtoMessage() {}

func (x *GetCatheterPathResponse) ProtoReflect() protoreflect.Message {
	mi := &file_catheter_path_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCatheterPathResponse.ProtoReflect.Descriptor instead.
func (*GetCatheterPathResponse) Descriptor() ([]byte, []int) {
	return file_catheter_path_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetCatheterPathResponse) GetPath() *CatheterPath {
	if x != nil {
		return x.Path
	}
	return nil
}

type GetCoordinatesCatheterPathRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCoordinatesCatheterPathRequest) Reset() {
	*x = GetCoordinatesCatheterPathRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catheter_path_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCoordinatesCatheterPathRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoordinatesCatheterPathRequest) ProtoMessage() {}

func (x *GetCoordinatesCatheterPathRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catheter_path_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCoordinatesCatheterPathRequest.ProtoReflect.Descriptor instead.
func (*GetCoordinatesCatheterPathRequest) Descriptor() ([]byte, []int) {
	return file_catheter_path_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetCoordinatesCatheterPathRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCoordinatesCatheterPathResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coordinate *CatheterCoordinates `protobuf:"bytes,1,opt,name=coordinate,proto3" json:"coordinate,omitempty"`
}

func (x *GetCoordinatesCatheterPathResponse) Reset() {
	*x = GetCoordinatesCatheterPathResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catheter_path_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCoordinatesCatheterPathResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoordinatesCatheterPathResponse) ProtoMessage() {}

func (x *GetCoordinatesCatheterPathResponse) ProtoReflect() protoreflect.Message {
	mi := &file_catheter_path_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCoordinatesCatheterPathResponse.ProtoReflect.Descriptor instead.
func (*GetCoordinatesCatheterPathResponse) Descriptor() ([]byte, []int) {
	return file_catheter_path_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetCoordinatesCatheterPathResponse) GetCoordinate() *CatheterCoordinates {
	if x != nil {
		return x.Coordinate
	}
	return nil
}

type ListCatheterPathsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCatheterPathsRequest) Reset() {
	*x = ListCatheterPathsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catheter_path_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCatheterPathsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCatheterPathsRequest) ProtoMessage() {}

func (x *ListCatheterPathsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catheter_path_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCatheterPathsRequest.ProtoReflect.Descriptor instead.
func (*ListCatheterPathsRequest) Descriptor() ([]byte, []int) {
	return file_catheter_path_service_proto_rawDescGZIP(), []int{6}
}

type ListCatheterPathsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path *CatheterPath `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *ListCatheterPathsResponse) Reset() {
	*x = ListCatheterPathsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catheter_path_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCatheterPathsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCatheterPathsResponse) ProtoMessage() {}

func (x *ListCatheterPathsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_catheter_path_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCatheterPathsResponse.ProtoReflect.Descriptor instead.
func (*ListCatheterPathsResponse) Descriptor() ([]byte, []int) {
	return file_catheter_path_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListCatheterPathsResponse) GetPath() *CatheterPath {
	if x != nil {
		return x.Path
	}
	return nil
}

type CatheterPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StartCatheterDataID uint64               `protobuf:"varint,2,opt,name=startCatheterDataID,proto3" json:"startCatheterDataID,omitempty"`
	EndCatheterDataID   uint64               `protobuf:"varint,3,opt,name=endCatheterDataID,proto3" json:"endCatheterDataID,omitempty"`
	Name                string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Created             *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *CatheterPath) Reset() {
	*x = CatheterPath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catheter_path_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CatheterPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CatheterPath) ProtoMessage() {}

func (x *CatheterPath) ProtoReflect() protoreflect.Message {
	mi := &file_catheter_path_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CatheterPath.ProtoReflect.Descriptor instead.
func (*CatheterPath) Descriptor() ([]byte, []int) {
	return file_catheter_path_service_proto_rawDescGZIP(), []int{8}
}

func (x *CatheterPath) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CatheterPath) GetStartCatheterDataID() uint64 {
	if x != nil {
		return x.StartCatheterDataID
	}
	return 0
}

func (x *CatheterPath) GetEndCatheterDataID() uint64 {
	if x != nil {
		return x.EndCatheterDataID
	}
	return 0
}

func (x *CatheterPath) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CatheterPath) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

var File_catheter_path_service_proto protoreflect.FileDescriptor

var file_catheter_path_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x61, 0x74, 0x68, 0x65, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x64,
	0x6f, 0x70, 0x6c, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x74, 0x68, 0x65, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x39, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x64, 0x6f, 0x70, 0x6c, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67,
	0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x68, 0x65,
	0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x2c, 0x0a,
	0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x68, 0x65, 0x74, 0x65, 0x72, 0x50,
	0x61, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x74, 0x68, 0x65, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x54, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x68,
	0x65, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x39, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x64, 0x6f, 0x70, 0x6c, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x68, 0x65, 0x74, 0x65,
	0x72, 0x50, 0x61, 0x74, 0x68, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x33, 0x0a, 0x21, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x43, 0x61, 0x74,
	0x68, 0x65, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x72, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x73, 0x43, 0x61, 0x74, 0x68, 0x65, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0a, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x64, 0x6f, 0x70,
	0x6c, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x68, 0x65, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x52, 0x0a, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x22, 0x1a, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x68,
	0x65, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x56, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x68, 0x65, 0x74, 0x65, 0x72,
	0x50, 0x61, 0x74, 0x68, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x64, 0x6f,
	0x70, 0x6c, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x68, 0x65, 0x74, 0x65, 0x72, 0x50, 0x61,
	0x74, 0x68, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0xc8, 0x01, 0x0a, 0x0c, 0x43, 0x61, 0x74,
	0x68, 0x65, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x43, 0x61, 0x74, 0x68, 0x65, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x61, 0x74,
	0x68, 0x65, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x11, 0x65,
	0x6e, 0x64, 0x43, 0x61, 0x74, 0x68, 0x65, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x49, 0x44,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x65, 0x6e, 0x64, 0x43, 0x61, 0x74, 0x68, 0x65,
	0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x32, 0xf1, 0x03, 0x0a, 0x13, 0x43, 0x61, 0x74, 0x68, 0x65, 0x74, 0x65, 0x72,
	0x50, 0x61, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x71, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x32, 0x2e, 0x64, 0x6f, 0x70, 0x6c, 0x74, 0x65, 0x63, 0x68,
	0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x68, 0x65, 0x74, 0x65, 0x72, 0x50, 0x61,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x64, 0x6f, 0x70, 0x6c,
	0x74, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x68, 0x65, 0x74,
	0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x68,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x2f, 0x2e, 0x64, 0x6f, 0x70, 0x6c, 0x74, 0x65, 0x63, 0x68,
	0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x68, 0x65, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x64, 0x6f, 0x70, 0x6c, 0x74, 0x65, 0x63,
	0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x68, 0x65, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x8b, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x3a, 0x2e, 0x64, 0x6f,
	0x70, 0x6c, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x73, 0x43, 0x61, 0x74, 0x68, 0x65, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x64, 0x6f, 0x70, 0x6c, 0x74, 0x65,
	0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73,
	0x43, 0x61, 0x74, 0x68, 0x65, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x6f, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31,
	0x2e, 0x64, 0x6f, 0x70, 0x6c, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74,
	0x68, 0x65, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x32, 0x2e, 0x64, 0x6f, 0x70, 0x6c, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f,
	0x67, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x61, 0x74, 0x68, 0x65, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x4f, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x70, 0x6c, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x6e,
	0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2d, 0x67, 0x6f, 0x3b, 0x64, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0xaa, 0x02,
	0x17, 0x44, 0x6f, 0x70, 0x6c, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x65,
	0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_catheter_path_service_proto_rawDescOnce sync.Once
	file_catheter_path_service_proto_rawDescData = file_catheter_path_service_proto_rawDesc
)

func file_catheter_path_service_proto_rawDescGZIP() []byte {
	file_catheter_path_service_proto_rawDescOnce.Do(func() {
		file_catheter_path_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_catheter_path_service_proto_rawDescData)
	})
	return file_catheter_path_service_proto_rawDescData
}

var file_catheter_path_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_catheter_path_service_proto_goTypes = []interface{}{
	(*CreateCatheterPathRequest)(nil),          // 0: dopltechnologies.protos.CreateCatheterPathRequest
	(*CreateCatheterPathResponse)(nil),         // 1: dopltechnologies.protos.CreateCatheterPathResponse
	(*GetCatheterPathRequest)(nil),             // 2: dopltechnologies.protos.GetCatheterPathRequest
	(*GetCatheterPathResponse)(nil),            // 3: dopltechnologies.protos.GetCatheterPathResponse
	(*GetCoordinatesCatheterPathRequest)(nil),  // 4: dopltechnologies.protos.GetCoordinatesCatheterPathRequest
	(*GetCoordinatesCatheterPathResponse)(nil), // 5: dopltechnologies.protos.GetCoordinatesCatheterPathResponse
	(*ListCatheterPathsRequest)(nil),           // 6: dopltechnologies.protos.ListCatheterPathsRequest
	(*ListCatheterPathsResponse)(nil),          // 7: dopltechnologies.protos.ListCatheterPathsResponse
	(*CatheterPath)(nil),                       // 8: dopltechnologies.protos.CatheterPath
	(*CatheterCoordinates)(nil),                // 9: dopltechnologies.protos.CatheterCoordinates
	(*timestamp.Timestamp)(nil),                // 10: google.protobuf.Timestamp
}
var file_catheter_path_service_proto_depIdxs = []int32{
	8,  // 0: dopltechnologies.protos.CreateCatheterPathRequest.path:type_name -> dopltechnologies.protos.CatheterPath
	8,  // 1: dopltechnologies.protos.GetCatheterPathResponse.path:type_name -> dopltechnologies.protos.CatheterPath
	9,  // 2: dopltechnologies.protos.GetCoordinatesCatheterPathResponse.coordinate:type_name -> dopltechnologies.protos.CatheterCoordinates
	8,  // 3: dopltechnologies.protos.ListCatheterPathsResponse.path:type_name -> dopltechnologies.protos.CatheterPath
	10, // 4: dopltechnologies.protos.CatheterPath.created:type_name -> google.protobuf.Timestamp
	0,  // 5: dopltechnologies.protos.CatheterPathService.Create:input_type -> dopltechnologies.protos.CreateCatheterPathRequest
	2,  // 6: dopltechnologies.protos.CatheterPathService.Get:input_type -> dopltechnologies.protos.GetCatheterPathRequest
	4,  // 7: dopltechnologies.protos.CatheterPathService.GetCoordinates:input_type -> dopltechnologies.protos.GetCoordinatesCatheterPathRequest
	6,  // 8: dopltechnologies.protos.CatheterPathService.List:input_type -> dopltechnologies.protos.ListCatheterPathsRequest
	1,  // 9: dopltechnologies.protos.CatheterPathService.Create:output_type -> dopltechnologies.protos.CreateCatheterPathResponse
	3,  // 10: dopltechnologies.protos.CatheterPathService.Get:output_type -> dopltechnologies.protos.GetCatheterPathResponse
	5,  // 11: dopltechnologies.protos.CatheterPathService.GetCoordinates:output_type -> dopltechnologies.protos.GetCoordinatesCatheterPathResponse
	7,  // 12: dopltechnologies.protos.CatheterPathService.List:output_type -> dopltechnologies.protos.ListCatheterPathsResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_catheter_path_service_proto_init() }
func file_catheter_path_service_proto_init() {
	if File_catheter_path_service_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_catheter_path_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCatheterPathRequest); i {
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
		file_catheter_path_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCatheterPathResponse); i {
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
		file_catheter_path_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCatheterPathRequest); i {
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
		file_catheter_path_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCatheterPathResponse); i {
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
		file_catheter_path_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCoordinatesCatheterPathRequest); i {
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
		file_catheter_path_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCoordinatesCatheterPathResponse); i {
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
		file_catheter_path_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCatheterPathsRequest); i {
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
		file_catheter_path_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCatheterPathsResponse); i {
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
		file_catheter_path_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CatheterPath); i {
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
			RawDescriptor: file_catheter_path_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_catheter_path_service_proto_goTypes,
		DependencyIndexes: file_catheter_path_service_proto_depIdxs,
		MessageInfos:      file_catheter_path_service_proto_msgTypes,
	}.Build()
	File_catheter_path_service_proto = out.File
	file_catheter_path_service_proto_rawDesc = nil
	file_catheter_path_service_proto_goTypes = nil
	file_catheter_path_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CatheterPathServiceClient is the client API for CatheterPathService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CatheterPathServiceClient interface {
	// Create a catheter path
	Create(ctx context.Context, in *CreateCatheterPathRequest, opts ...grpc.CallOption) (*CreateCatheterPathResponse, error)
	// Gets a catheter path
	Get(ctx context.Context, in *GetCatheterPathRequest, opts ...grpc.CallOption) (*GetCatheterPathResponse, error)
	// Gets a catheter path coordinates
	GetCoordinates(ctx context.Context, in *GetCoordinatesCatheterPathRequest, opts ...grpc.CallOption) (CatheterPathService_GetCoordinatesClient, error)
	// Lists catheter paths
	List(ctx context.Context, in *ListCatheterPathsRequest, opts ...grpc.CallOption) (CatheterPathService_ListClient, error)
}

type catheterPathServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCatheterPathServiceClient(cc grpc.ClientConnInterface) CatheterPathServiceClient {
	return &catheterPathServiceClient{cc}
}

func (c *catheterPathServiceClient) Create(ctx context.Context, in *CreateCatheterPathRequest, opts ...grpc.CallOption) (*CreateCatheterPathResponse, error) {
	out := new(CreateCatheterPathResponse)
	err := c.cc.Invoke(ctx, "/dopltechnologies.protos.CatheterPathService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catheterPathServiceClient) Get(ctx context.Context, in *GetCatheterPathRequest, opts ...grpc.CallOption) (*GetCatheterPathResponse, error) {
	out := new(GetCatheterPathResponse)
	err := c.cc.Invoke(ctx, "/dopltechnologies.protos.CatheterPathService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catheterPathServiceClient) GetCoordinates(ctx context.Context, in *GetCoordinatesCatheterPathRequest, opts ...grpc.CallOption) (CatheterPathService_GetCoordinatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CatheterPathService_serviceDesc.Streams[0], "/dopltechnologies.protos.CatheterPathService/GetCoordinates", opts...)
	if err != nil {
		return nil, err
	}
	x := &catheterPathServiceGetCoordinatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CatheterPathService_GetCoordinatesClient interface {
	Recv() (*GetCoordinatesCatheterPathResponse, error)
	grpc.ClientStream
}

type catheterPathServiceGetCoordinatesClient struct {
	grpc.ClientStream
}

func (x *catheterPathServiceGetCoordinatesClient) Recv() (*GetCoordinatesCatheterPathResponse, error) {
	m := new(GetCoordinatesCatheterPathResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *catheterPathServiceClient) List(ctx context.Context, in *ListCatheterPathsRequest, opts ...grpc.CallOption) (CatheterPathService_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CatheterPathService_serviceDesc.Streams[1], "/dopltechnologies.protos.CatheterPathService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &catheterPathServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CatheterPathService_ListClient interface {
	Recv() (*ListCatheterPathsResponse, error)
	grpc.ClientStream
}

type catheterPathServiceListClient struct {
	grpc.ClientStream
}

func (x *catheterPathServiceListClient) Recv() (*ListCatheterPathsResponse, error) {
	m := new(ListCatheterPathsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CatheterPathServiceServer is the server API for CatheterPathService service.
type CatheterPathServiceServer interface {
	// Create a catheter path
	Create(context.Context, *CreateCatheterPathRequest) (*CreateCatheterPathResponse, error)
	// Gets a catheter path
	Get(context.Context, *GetCatheterPathRequest) (*GetCatheterPathResponse, error)
	// Gets a catheter path coordinates
	GetCoordinates(*GetCoordinatesCatheterPathRequest, CatheterPathService_GetCoordinatesServer) error
	// Lists catheter paths
	List(*ListCatheterPathsRequest, CatheterPathService_ListServer) error
}

// UnimplementedCatheterPathServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCatheterPathServiceServer struct {
}

func (*UnimplementedCatheterPathServiceServer) Create(context.Context, *CreateCatheterPathRequest) (*CreateCatheterPathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedCatheterPathServiceServer) Get(context.Context, *GetCatheterPathRequest) (*GetCatheterPathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCatheterPathServiceServer) GetCoordinates(*GetCoordinatesCatheterPathRequest, CatheterPathService_GetCoordinatesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCoordinates not implemented")
}
func (*UnimplementedCatheterPathServiceServer) List(*ListCatheterPathsRequest, CatheterPathService_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterCatheterPathServiceServer(s *grpc.Server, srv CatheterPathServiceServer) {
	s.RegisterService(&_CatheterPathService_serviceDesc, srv)
}

func _CatheterPathService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCatheterPathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatheterPathServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dopltechnologies.protos.CatheterPathService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatheterPathServiceServer).Create(ctx, req.(*CreateCatheterPathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatheterPathService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCatheterPathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatheterPathServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dopltechnologies.protos.CatheterPathService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatheterPathServiceServer).Get(ctx, req.(*GetCatheterPathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatheterPathService_GetCoordinates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetCoordinatesCatheterPathRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CatheterPathServiceServer).GetCoordinates(m, &catheterPathServiceGetCoordinatesServer{stream})
}

type CatheterPathService_GetCoordinatesServer interface {
	Send(*GetCoordinatesCatheterPathResponse) error
	grpc.ServerStream
}

type catheterPathServiceGetCoordinatesServer struct {
	grpc.ServerStream
}

func (x *catheterPathServiceGetCoordinatesServer) Send(m *GetCoordinatesCatheterPathResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CatheterPathService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListCatheterPathsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CatheterPathServiceServer).List(m, &catheterPathServiceListServer{stream})
}

type CatheterPathService_ListServer interface {
	Send(*ListCatheterPathsResponse) error
	grpc.ServerStream
}

type catheterPathServiceListServer struct {
	grpc.ServerStream
}

func (x *catheterPathServiceListServer) Send(m *ListCatheterPathsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _CatheterPathService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dopltechnologies.protos.CatheterPathService",
	HandlerType: (*CatheterPathServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CatheterPathService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CatheterPathService_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCoordinates",
			Handler:       _CatheterPathService_GetCoordinates_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "List",
			Handler:       _CatheterPathService_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "catheter_path_service.proto",
}
