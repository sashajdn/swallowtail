// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: s.ftx/proto/ftx.proto

package ftxproto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ListAccountDepositsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorId string `protobuf:"bytes,1,opt,name=actor_id,json=actorId,proto3" json:"actor_id,omitempty"`
	// Seconds in unix time
	Start int64 `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	// Seconds in unix time
	End int64 `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *ListAccountDepositsRequest) Reset() {
	*x = ListAccountDepositsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_ftx_proto_ftx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAccountDepositsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountDepositsRequest) ProtoMessage() {}

func (x *ListAccountDepositsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_s_ftx_proto_ftx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountDepositsRequest.ProtoReflect.Descriptor instead.
func (*ListAccountDepositsRequest) Descriptor() ([]byte, []int) {
	return file_s_ftx_proto_ftx_proto_rawDescGZIP(), []int{0}
}

func (x *ListAccountDepositsRequest) GetActorId() string {
	if x != nil {
		return x.ActorId
	}
	return ""
}

func (x *ListAccountDepositsRequest) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *ListAccountDepositsRequest) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

type FTXCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKey     string `protobuf:"bytes,1,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	SecretKey  string `protobuf:"bytes,2,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	Subaccount string `protobuf:"bytes,3,opt,name=subaccount,proto3" json:"subaccount,omitempty"`
}

func (x *FTXCredentials) Reset() {
	*x = FTXCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_ftx_proto_ftx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FTXCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FTXCredentials) ProtoMessage() {}

func (x *FTXCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_s_ftx_proto_ftx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FTXCredentials.ProtoReflect.Descriptor instead.
func (*FTXCredentials) Descriptor() ([]byte, []int) {
	return file_s_ftx_proto_ftx_proto_rawDescGZIP(), []int{1}
}

func (x *FTXCredentials) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *FTXCredentials) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *FTXCredentials) GetSubaccount() string {
	if x != nil {
		return x.Subaccount
	}
	return ""
}

type ListAccountDepositsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deposits []*DepositRecord `protobuf:"bytes,1,rep,name=deposits,proto3" json:"deposits,omitempty"`
}

func (x *ListAccountDepositsResponse) Reset() {
	*x = ListAccountDepositsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_ftx_proto_ftx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAccountDepositsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountDepositsResponse) ProtoMessage() {}

func (x *ListAccountDepositsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_s_ftx_proto_ftx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountDepositsResponse.ProtoReflect.Descriptor instead.
func (*ListAccountDepositsResponse) Descriptor() ([]byte, []int) {
	return file_s_ftx_proto_ftx_proto_rawDescGZIP(), []int{2}
}

func (x *ListAccountDepositsResponse) GetDeposits() []*DepositRecord {
	if x != nil {
		return x.Deposits
	}
	return nil
}

type DepositRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coin          string                 `protobuf:"bytes,1,opt,name=coin,proto3" json:"coin,omitempty"`
	Confirmations int64                  `protobuf:"varint,2,opt,name=confirmations,proto3" json:"confirmations,omitempty"`
	ConfirmedTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=confirmed_time,json=confirmedTime,proto3" json:"confirmed_time,omitempty"`
	Fee           float32                `protobuf:"fixed32,4,opt,name=fee,proto3" json:"fee,omitempty"`
	Id            int64                  `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	SentTime      *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=sent_time,json=sentTime,proto3" json:"sent_time,omitempty"`
	Size          float32                `protobuf:"fixed32,7,opt,name=size,proto3" json:"size,omitempty"`
	Status        string                 `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	Time          *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=time,proto3" json:"time,omitempty"`
	TransactionId string                 `protobuf:"bytes,10,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
}

func (x *DepositRecord) Reset() {
	*x = DepositRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_ftx_proto_ftx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositRecord) ProtoMessage() {}

func (x *DepositRecord) ProtoReflect() protoreflect.Message {
	mi := &file_s_ftx_proto_ftx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositRecord.ProtoReflect.Descriptor instead.
func (*DepositRecord) Descriptor() ([]byte, []int) {
	return file_s_ftx_proto_ftx_proto_rawDescGZIP(), []int{3}
}

func (x *DepositRecord) GetCoin() string {
	if x != nil {
		return x.Coin
	}
	return ""
}

func (x *DepositRecord) GetConfirmations() int64 {
	if x != nil {
		return x.Confirmations
	}
	return 0
}

func (x *DepositRecord) GetConfirmedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ConfirmedTime
	}
	return nil
}

func (x *DepositRecord) GetFee() float32 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *DepositRecord) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DepositRecord) GetSentTime() *timestamppb.Timestamp {
	if x != nil {
		return x.SentTime
	}
	return nil
}

func (x *DepositRecord) GetSize() float32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *DepositRecord) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DepositRecord) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *DepositRecord) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

type GetFTXStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFTXStatusRequest) Reset() {
	*x = GetFTXStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_ftx_proto_ftx_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFTXStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFTXStatusRequest) ProtoMessage() {}

func (x *GetFTXStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_s_ftx_proto_ftx_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFTXStatusRequest.ProtoReflect.Descriptor instead.
func (*GetFTXStatusRequest) Descriptor() ([]byte, []int) {
	return file_s_ftx_proto_ftx_proto_rawDescGZIP(), []int{4}
}

type GetFTXStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerTime     int64 `protobuf:"varint,1,opt,name=server_time,json=serverTime,proto3" json:"server_time,omitempty"`
	RequestLatency int64 `protobuf:"varint,2,opt,name=request_latency,json=requestLatency,proto3" json:"request_latency,omitempty"`
}

func (x *GetFTXStatusResponse) Reset() {
	*x = GetFTXStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_ftx_proto_ftx_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFTXStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFTXStatusResponse) ProtoMessage() {}

func (x *GetFTXStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_s_ftx_proto_ftx_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFTXStatusResponse.ProtoReflect.Descriptor instead.
func (*GetFTXStatusResponse) Descriptor() ([]byte, []int) {
	return file_s_ftx_proto_ftx_proto_rawDescGZIP(), []int{5}
}

func (x *GetFTXStatusResponse) GetServerTime() int64 {
	if x != nil {
		return x.ServerTime
	}
	return 0
}

func (x *GetFTXStatusResponse) GetRequestLatency() int64 {
	if x != nil {
		return x.RequestLatency
	}
	return 0
}

var File_s_ftx_proto_ftx_proto protoreflect.FileDescriptor

var file_s_ftx_proto_ftx_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x2e, 0x66, 0x74, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x74,
	0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x68, 0x0a, 0x0e, 0x46, 0x54, 0x58,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x61,
	0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70,
	0x69, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x49, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x08, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x22, 0xea,
	0x02, 0x0a, 0x0d, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x69, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x41, 0x0a, 0x0e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x66, 0x65, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x37, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08,
	0x73, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x46, 0x54, 0x58, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x60, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x54, 0x58, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x61, 0x74,
	0x65, 0x6e, 0x63, 0x79, 0x32, 0x98, 0x01, 0x0a, 0x03, 0x66, 0x74, 0x78, 0x12, 0x52, 0x0a, 0x13,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x73, 0x12, 0x1b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x54, 0x58, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x54, 0x58, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x54, 0x58, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x66, 0x74, 0x78, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_s_ftx_proto_ftx_proto_rawDescOnce sync.Once
	file_s_ftx_proto_ftx_proto_rawDescData = file_s_ftx_proto_ftx_proto_rawDesc
)

func file_s_ftx_proto_ftx_proto_rawDescGZIP() []byte {
	file_s_ftx_proto_ftx_proto_rawDescOnce.Do(func() {
		file_s_ftx_proto_ftx_proto_rawDescData = protoimpl.X.CompressGZIP(file_s_ftx_proto_ftx_proto_rawDescData)
	})
	return file_s_ftx_proto_ftx_proto_rawDescData
}

var file_s_ftx_proto_ftx_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_s_ftx_proto_ftx_proto_goTypes = []interface{}{
	(*ListAccountDepositsRequest)(nil),  // 0: ListAccountDepositsRequest
	(*FTXCredentials)(nil),              // 1: FTXCredentials
	(*ListAccountDepositsResponse)(nil), // 2: ListAccountDepositsResponse
	(*DepositRecord)(nil),               // 3: DepositRecord
	(*GetFTXStatusRequest)(nil),         // 4: GetFTXStatusRequest
	(*GetFTXStatusResponse)(nil),        // 5: GetFTXStatusResponse
	(*timestamppb.Timestamp)(nil),       // 6: google.protobuf.Timestamp
}
var file_s_ftx_proto_ftx_proto_depIdxs = []int32{
	3, // 0: ListAccountDepositsResponse.deposits:type_name -> DepositRecord
	6, // 1: DepositRecord.confirmed_time:type_name -> google.protobuf.Timestamp
	6, // 2: DepositRecord.sent_time:type_name -> google.protobuf.Timestamp
	6, // 3: DepositRecord.time:type_name -> google.protobuf.Timestamp
	0, // 4: ftx.ListAccountDeposits:input_type -> ListAccountDepositsRequest
	4, // 5: ftx.GetFTXStatus:input_type -> GetFTXStatusRequest
	2, // 6: ftx.ListAccountDeposits:output_type -> ListAccountDepositsResponse
	5, // 7: ftx.GetFTXStatus:output_type -> GetFTXStatusResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_s_ftx_proto_ftx_proto_init() }
func file_s_ftx_proto_ftx_proto_init() {
	if File_s_ftx_proto_ftx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_s_ftx_proto_ftx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAccountDepositsRequest); i {
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
		file_s_ftx_proto_ftx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FTXCredentials); i {
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
		file_s_ftx_proto_ftx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAccountDepositsResponse); i {
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
		file_s_ftx_proto_ftx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositRecord); i {
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
		file_s_ftx_proto_ftx_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFTXStatusRequest); i {
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
		file_s_ftx_proto_ftx_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFTXStatusResponse); i {
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
			RawDescriptor: file_s_ftx_proto_ftx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_s_ftx_proto_ftx_proto_goTypes,
		DependencyIndexes: file_s_ftx_proto_ftx_proto_depIdxs,
		MessageInfos:      file_s_ftx_proto_ftx_proto_msgTypes,
	}.Build()
	File_s_ftx_proto_ftx_proto = out.File
	file_s_ftx_proto_ftx_proto_rawDesc = nil
	file_s_ftx_proto_ftx_proto_goTypes = nil
	file_s_ftx_proto_ftx_proto_depIdxs = nil
}
