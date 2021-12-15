// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: cosmos/mint/v1beta1/query.proto

package mintv1beta1

import (
	_ "github.com/gogo/protobuf/gogoproto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueryParamsRequest) Reset() {
	*x = QueryParamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_mint_v1beta1_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryParamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryParamsRequest) ProtoMessage() {}

func (x *QueryParamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_mint_v1beta1_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryParamsRequest.ProtoReflect.Descriptor instead.
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_mint_v1beta1_query_proto_rawDescGZIP(), []int{0}
}

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// params defines the parameters of the module.
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *QueryParamsResponse) Reset() {
	*x = QueryParamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_mint_v1beta1_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryParamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryParamsResponse) ProtoMessage() {}

func (x *QueryParamsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_mint_v1beta1_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryParamsResponse.ProtoReflect.Descriptor instead.
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_mint_v1beta1_query_proto_rawDescGZIP(), []int{1}
}

func (x *QueryParamsResponse) GetParams() *Params {
	if x != nil {
		return x.Params
	}
	return nil
}

// QueryInflationRequest is the request type for the Query/Inflation RPC method.
type QueryInflationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueryInflationRequest) Reset() {
	*x = QueryInflationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_mint_v1beta1_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryInflationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryInflationRequest) ProtoMessage() {}

func (x *QueryInflationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_mint_v1beta1_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryInflationRequest.ProtoReflect.Descriptor instead.
func (*QueryInflationRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_mint_v1beta1_query_proto_rawDescGZIP(), []int{2}
}

// QueryInflationResponse is the response type for the Query/Inflation RPC
// method.
type QueryInflationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// inflation is the current minting inflation value.
	Inflation []byte `protobuf:"bytes,1,opt,name=inflation,proto3" json:"inflation,omitempty"`
}

func (x *QueryInflationResponse) Reset() {
	*x = QueryInflationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_mint_v1beta1_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryInflationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryInflationResponse) ProtoMessage() {}

func (x *QueryInflationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_mint_v1beta1_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryInflationResponse.ProtoReflect.Descriptor instead.
func (*QueryInflationResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_mint_v1beta1_query_proto_rawDescGZIP(), []int{3}
}

func (x *QueryInflationResponse) GetInflation() []byte {
	if x != nil {
		return x.Inflation
	}
	return nil
}

// QueryAnnualProvisionsRequest is the request type for the
// Query/AnnualProvisions RPC method.
type QueryAnnualProvisionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueryAnnualProvisionsRequest) Reset() {
	*x = QueryAnnualProvisionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_mint_v1beta1_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAnnualProvisionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAnnualProvisionsRequest) ProtoMessage() {}

func (x *QueryAnnualProvisionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_mint_v1beta1_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAnnualProvisionsRequest.ProtoReflect.Descriptor instead.
func (*QueryAnnualProvisionsRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_mint_v1beta1_query_proto_rawDescGZIP(), []int{4}
}

// QueryAnnualProvisionsResponse is the response type for the
// Query/AnnualProvisions RPC method.
type QueryAnnualProvisionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// annual_provisions is the current minting annual provisions value.
	AnnualProvisions []byte `protobuf:"bytes,1,opt,name=annual_provisions,json=annualProvisions,proto3" json:"annual_provisions,omitempty"`
}

func (x *QueryAnnualProvisionsResponse) Reset() {
	*x = QueryAnnualProvisionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_mint_v1beta1_query_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAnnualProvisionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAnnualProvisionsResponse) ProtoMessage() {}

func (x *QueryAnnualProvisionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_mint_v1beta1_query_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAnnualProvisionsResponse.ProtoReflect.Descriptor instead.
func (*QueryAnnualProvisionsResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_mint_v1beta1_query_proto_rawDescGZIP(), []int{5}
}

func (x *QueryAnnualProvisionsResponse) GetAnnualProvisions() []byte {
	if x != nil {
		return x.AnnualProvisions
	}
	return nil
}

var File_cosmos_mint_v1beta1_query_proto protoreflect.FileDescriptor

var file_cosmos_mint_v1beta1_query_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2f, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x50, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x66, 0x0a, 0x16, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x09, 0x69, 0x6e, 0x66, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x2e, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde,
	0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x52, 0x09, 0x69, 0x6e, 0x66, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x1e, 0x0a, 0x1c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6e, 0x6e, 0x75,
	0x61, 0x6c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x7c, 0x0a, 0x1d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6e, 0x6e, 0x75,
	0x61, 0x6c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x11, 0x61, 0x6e, 0x6e, 0x75, 0x61, 0x6c, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42,
	0x2e, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x52,
	0x10, 0x61, 0x6e, 0x6e, 0x75, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x32, 0xc5, 0x03, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x80, 0x01, 0x0a, 0x06,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x27, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x28, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1d, 0x12, 0x1b, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6d, 0x69, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x8c,
	0x01, 0x0a, 0x09, 0x49, 0x6e, 0x66, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x69, 0x6e, 0x66, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xa9, 0x01,
	0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x75, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x31, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x6d, 0x69, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6e,
	0x6e, 0x75, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x6d,
	0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x41, 0x6e, 0x6e, 0x75, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x28, 0x12, 0x26, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6d, 0x69, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x75, 0x61, 0x6c, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0xd5, 0x01, 0x0a, 0x17, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6d, 0x69, 0x6e,
	0x74, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x6d, 0x69, 0x6e, 0x74, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x4d, 0x58, 0xaa, 0x02, 0x13, 0x43, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x4d, 0x69, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0xca, 0x02, 0x13, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x4d, 0x69, 0x6e, 0x74, 0x5c,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x1f, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x5c, 0x4d, 0x69, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x43, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x3a, 0x3a, 0x4d, 0x69, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_mint_v1beta1_query_proto_rawDescOnce sync.Once
	file_cosmos_mint_v1beta1_query_proto_rawDescData = file_cosmos_mint_v1beta1_query_proto_rawDesc
)

func file_cosmos_mint_v1beta1_query_proto_rawDescGZIP() []byte {
	file_cosmos_mint_v1beta1_query_proto_rawDescOnce.Do(func() {
		file_cosmos_mint_v1beta1_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_mint_v1beta1_query_proto_rawDescData)
	})
	return file_cosmos_mint_v1beta1_query_proto_rawDescData
}

var file_cosmos_mint_v1beta1_query_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cosmos_mint_v1beta1_query_proto_goTypes = []interface{}{
	(*QueryParamsRequest)(nil),            // 0: cosmos.mint.v1beta1.QueryParamsRequest
	(*QueryParamsResponse)(nil),           // 1: cosmos.mint.v1beta1.QueryParamsResponse
	(*QueryInflationRequest)(nil),         // 2: cosmos.mint.v1beta1.QueryInflationRequest
	(*QueryInflationResponse)(nil),        // 3: cosmos.mint.v1beta1.QueryInflationResponse
	(*QueryAnnualProvisionsRequest)(nil),  // 4: cosmos.mint.v1beta1.QueryAnnualProvisionsRequest
	(*QueryAnnualProvisionsResponse)(nil), // 5: cosmos.mint.v1beta1.QueryAnnualProvisionsResponse
	(*Params)(nil),                        // 6: cosmos.mint.v1beta1.Params
}
var file_cosmos_mint_v1beta1_query_proto_depIdxs = []int32{
	6, // 0: cosmos.mint.v1beta1.QueryParamsResponse.params:type_name -> cosmos.mint.v1beta1.Params
	0, // 1: cosmos.mint.v1beta1.Query.Params:input_type -> cosmos.mint.v1beta1.QueryParamsRequest
	2, // 2: cosmos.mint.v1beta1.Query.Inflation:input_type -> cosmos.mint.v1beta1.QueryInflationRequest
	4, // 3: cosmos.mint.v1beta1.Query.AnnualProvisions:input_type -> cosmos.mint.v1beta1.QueryAnnualProvisionsRequest
	1, // 4: cosmos.mint.v1beta1.Query.Params:output_type -> cosmos.mint.v1beta1.QueryParamsResponse
	3, // 5: cosmos.mint.v1beta1.Query.Inflation:output_type -> cosmos.mint.v1beta1.QueryInflationResponse
	5, // 6: cosmos.mint.v1beta1.Query.AnnualProvisions:output_type -> cosmos.mint.v1beta1.QueryAnnualProvisionsResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cosmos_mint_v1beta1_query_proto_init() }
func file_cosmos_mint_v1beta1_query_proto_init() {
	if File_cosmos_mint_v1beta1_query_proto != nil {
		return
	}
	file_cosmos_mint_v1beta1_mint_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cosmos_mint_v1beta1_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryParamsRequest); i {
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
		file_cosmos_mint_v1beta1_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryParamsResponse); i {
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
		file_cosmos_mint_v1beta1_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryInflationRequest); i {
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
		file_cosmos_mint_v1beta1_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryInflationResponse); i {
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
		file_cosmos_mint_v1beta1_query_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAnnualProvisionsRequest); i {
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
		file_cosmos_mint_v1beta1_query_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAnnualProvisionsResponse); i {
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
			RawDescriptor: file_cosmos_mint_v1beta1_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cosmos_mint_v1beta1_query_proto_goTypes,
		DependencyIndexes: file_cosmos_mint_v1beta1_query_proto_depIdxs,
		MessageInfos:      file_cosmos_mint_v1beta1_query_proto_msgTypes,
	}.Build()
	File_cosmos_mint_v1beta1_query_proto = out.File
	file_cosmos_mint_v1beta1_query_proto_rawDesc = nil
	file_cosmos_mint_v1beta1_query_proto_goTypes = nil
	file_cosmos_mint_v1beta1_query_proto_depIdxs = nil
}