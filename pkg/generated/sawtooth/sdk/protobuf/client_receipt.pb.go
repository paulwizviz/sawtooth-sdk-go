// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client_receipt.proto

package batch_pb2

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ClientReceiptGetResponse_Status int32

const (
	ClientReceiptGetResponse_STATUS_UNSET   ClientReceiptGetResponse_Status = 0
	ClientReceiptGetResponse_OK             ClientReceiptGetResponse_Status = 1
	ClientReceiptGetResponse_INTERNAL_ERROR ClientReceiptGetResponse_Status = 2
	ClientReceiptGetResponse_NO_RESOURCE    ClientReceiptGetResponse_Status = 5
	ClientReceiptGetResponse_INVALID_ID     ClientReceiptGetResponse_Status = 8
)

var ClientReceiptGetResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INTERNAL_ERROR",
	5: "NO_RESOURCE",
	8: "INVALID_ID",
}

var ClientReceiptGetResponse_Status_value = map[string]int32{
	"STATUS_UNSET":   0,
	"OK":             1,
	"INTERNAL_ERROR": 2,
	"NO_RESOURCE":    5,
	"INVALID_ID":     8,
}

func (x ClientReceiptGetResponse_Status) String() string {
	return proto.EnumName(ClientReceiptGetResponse_Status_name, int32(x))
}

func (ClientReceiptGetResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3b95700b113cc8dd, []int{1, 0}
}

// Fetches a specific txn by its id (header_signature) from the blockchain.
type ClientReceiptGetRequest struct {
	TransactionIds       []string `protobuf:"bytes,1,rep,name=transaction_ids,json=transactionIds,proto3" json:"transaction_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientReceiptGetRequest) Reset()         { *m = ClientReceiptGetRequest{} }
func (m *ClientReceiptGetRequest) String() string { return proto.CompactTextString(m) }
func (*ClientReceiptGetRequest) ProtoMessage()    {}
func (*ClientReceiptGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b95700b113cc8dd, []int{0}
}

func (m *ClientReceiptGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientReceiptGetRequest.Unmarshal(m, b)
}
func (m *ClientReceiptGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientReceiptGetRequest.Marshal(b, m, deterministic)
}
func (m *ClientReceiptGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientReceiptGetRequest.Merge(m, src)
}
func (m *ClientReceiptGetRequest) XXX_Size() int {
	return xxx_messageInfo_ClientReceiptGetRequest.Size(m)
}
func (m *ClientReceiptGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientReceiptGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientReceiptGetRequest proto.InternalMessageInfo

func (m *ClientReceiptGetRequest) GetTransactionIds() []string {
	if m != nil {
		return m.TransactionIds
	}
	return nil
}

// A response that returns the txn receipt specified by a
// ClientTransactionReceiptGetRequest.
//
// Statuses:
//   * OK - everything worked as expected, txn receipt has been fetched
//   * INTERNAL_ERROR - general error, such as protobuf failing to deserialize
//   * NO_RESOURCE - no receipt exists for the transaction id specified
type ClientReceiptGetResponse struct {
	Status               ClientReceiptGetResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=sawtooth.sdk.protobuf.ClientReceiptGetResponse_Status" json:"status,omitempty"`
	Receipts             []*TransactionReceipt           `protobuf:"bytes,2,rep,name=receipts,proto3" json:"receipts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ClientReceiptGetResponse) Reset()         { *m = ClientReceiptGetResponse{} }
func (m *ClientReceiptGetResponse) String() string { return proto.CompactTextString(m) }
func (*ClientReceiptGetResponse) ProtoMessage()    {}
func (*ClientReceiptGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b95700b113cc8dd, []int{1}
}

func (m *ClientReceiptGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientReceiptGetResponse.Unmarshal(m, b)
}
func (m *ClientReceiptGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientReceiptGetResponse.Marshal(b, m, deterministic)
}
func (m *ClientReceiptGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientReceiptGetResponse.Merge(m, src)
}
func (m *ClientReceiptGetResponse) XXX_Size() int {
	return xxx_messageInfo_ClientReceiptGetResponse.Size(m)
}
func (m *ClientReceiptGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientReceiptGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClientReceiptGetResponse proto.InternalMessageInfo

func (m *ClientReceiptGetResponse) GetStatus() ClientReceiptGetResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientReceiptGetResponse_STATUS_UNSET
}

func (m *ClientReceiptGetResponse) GetReceipts() []*TransactionReceipt {
	if m != nil {
		return m.Receipts
	}
	return nil
}

func init() {
	proto.RegisterEnum("sawtooth.sdk.protobuf.ClientReceiptGetResponse_Status", ClientReceiptGetResponse_Status_name, ClientReceiptGetResponse_Status_value)
	proto.RegisterType((*ClientReceiptGetRequest)(nil), "sawtooth.sdk.protobuf.ClientReceiptGetRequest")
	proto.RegisterType((*ClientReceiptGetResponse)(nil), "sawtooth.sdk.protobuf.ClientReceiptGetResponse")
}

func init() { proto.RegisterFile("client_receipt.proto", fileDescriptor_3b95700b113cc8dd) }

var fileDescriptor_3b95700b113cc8dd = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0x87, 0x6d, 0x87, 0x65, 0xbe, 0x49, 0x17, 0x82, 0xe2, 0xe6, 0x69, 0xec, 0x62, 0xbd, 0xf4,
	0x30, 0xc1, 0x7b, 0xb7, 0x05, 0x29, 0x8e, 0x54, 0xd2, 0x56, 0x8f, 0xa5, 0x6b, 0x23, 0x16, 0xa5,
	0xa9, 0xcd, 0x2b, 0xfe, 0x0f, 0xfe, 0xd5, 0x62, 0x5b, 0x74, 0xe2, 0x76, 0x0b, 0x1f, 0x7c, 0x1f,
	0xbf, 0xf0, 0xe0, 0x2c, 0x7b, 0x2b, 0x64, 0x89, 0x49, 0x2d, 0x33, 0x59, 0x54, 0xe8, 0x56, 0xb5,
	0x42, 0x45, 0xcf, 0x75, 0xfa, 0x81, 0x4a, 0xe1, 0x8b, 0xab, 0xf3, 0xd7, 0x8e, 0x6d, 0x9b, 0xe7,
	0xcb, 0x29, 0xd6, 0x69, 0xa9, 0xd3, 0x0c, 0x0b, 0x55, 0xfe, 0x35, 0xe6, 0x4b, 0xb8, 0x58, 0xb5,
	0x25, 0xd1, 0xe1, 0x3b, 0x89, 0x42, 0xbe, 0x37, 0x52, 0x23, 0xbd, 0x82, 0xf1, 0xae, 0x57, 0xe4,
	0x7a, 0x62, 0xcc, 0x06, 0xce, 0x89, 0xb0, 0x77, 0xb0, 0x9f, 0xeb, 0xf9, 0xa7, 0x09, 0x93, 0xff,
	0x11, 0x5d, 0xa9, 0x52, 0x4b, 0xca, 0xc1, 0xd2, 0x98, 0x62, 0xf3, 0x2d, 0x1b, 0x8e, 0xbd, 0xb8,
	0x75, 0xf7, 0x6e, 0x74, 0x0f, 0x05, 0xdc, 0xb0, 0xb5, 0x45, 0x5f, 0xa1, 0x0c, 0x86, 0xfd, 0x0f,
	0xf4, 0xc4, 0x9c, 0x0d, 0x9c, 0xd1, 0xe2, 0xfa, 0x40, 0x31, 0xfa, 0x5d, 0xd9, 0x67, 0xc5, 0x8f,
	0x3a, 0x7f, 0x02, 0xab, 0x0b, 0x53, 0x02, 0xa7, 0x61, 0xe4, 0x45, 0x71, 0x98, 0xc4, 0x3c, 0x64,
	0x11, 0x39, 0xa2, 0x16, 0x98, 0xc1, 0x3d, 0x31, 0x28, 0x05, 0xdb, 0xe7, 0x11, 0x13, 0xdc, 0xdb,
	0x24, 0x4c, 0x88, 0x40, 0x10, 0x93, 0x8e, 0x61, 0xc4, 0x83, 0x44, 0xb0, 0x30, 0x88, 0xc5, 0x8a,
	0x91, 0x63, 0x6a, 0x03, 0xf8, 0xfc, 0xd1, 0xdb, 0xf8, 0xeb, 0xc4, 0x5f, 0x93, 0xe1, 0x72, 0x0a,
	0xfb, 0x8f, 0xf0, 0x60, 0x6c, 0xad, 0xf6, 0x7d, 0xf3, 0x15, 0x00, 0x00, 0xff, 0xff, 0x87, 0xfa,
	0xc2, 0xd0, 0xbc, 0x01, 0x00, 0x00,
}