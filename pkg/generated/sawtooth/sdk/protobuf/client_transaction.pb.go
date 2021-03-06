// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client_transaction.proto

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

type ClientTransactionListResponse_Status int32

const (
	ClientTransactionListResponse_STATUS_UNSET   ClientTransactionListResponse_Status = 0
	ClientTransactionListResponse_OK             ClientTransactionListResponse_Status = 1
	ClientTransactionListResponse_INTERNAL_ERROR ClientTransactionListResponse_Status = 2
	ClientTransactionListResponse_NOT_READY      ClientTransactionListResponse_Status = 3
	ClientTransactionListResponse_NO_ROOT        ClientTransactionListResponse_Status = 4
	ClientTransactionListResponse_NO_RESOURCE    ClientTransactionListResponse_Status = 5
	ClientTransactionListResponse_INVALID_PAGING ClientTransactionListResponse_Status = 6
	ClientTransactionListResponse_INVALID_SORT   ClientTransactionListResponse_Status = 7
	ClientTransactionListResponse_INVALID_ID     ClientTransactionListResponse_Status = 8
)

var ClientTransactionListResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INTERNAL_ERROR",
	3: "NOT_READY",
	4: "NO_ROOT",
	5: "NO_RESOURCE",
	6: "INVALID_PAGING",
	7: "INVALID_SORT",
	8: "INVALID_ID",
}

var ClientTransactionListResponse_Status_value = map[string]int32{
	"STATUS_UNSET":   0,
	"OK":             1,
	"INTERNAL_ERROR": 2,
	"NOT_READY":      3,
	"NO_ROOT":        4,
	"NO_RESOURCE":    5,
	"INVALID_PAGING": 6,
	"INVALID_SORT":   7,
	"INVALID_ID":     8,
}

func (x ClientTransactionListResponse_Status) String() string {
	return proto.EnumName(ClientTransactionListResponse_Status_name, int32(x))
}

func (ClientTransactionListResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b8feb52ffa146393, []int{1, 0}
}

type ClientTransactionGetResponse_Status int32

const (
	ClientTransactionGetResponse_STATUS_UNSET   ClientTransactionGetResponse_Status = 0
	ClientTransactionGetResponse_OK             ClientTransactionGetResponse_Status = 1
	ClientTransactionGetResponse_INTERNAL_ERROR ClientTransactionGetResponse_Status = 2
	ClientTransactionGetResponse_NO_RESOURCE    ClientTransactionGetResponse_Status = 5
	ClientTransactionGetResponse_INVALID_ID     ClientTransactionGetResponse_Status = 8
)

var ClientTransactionGetResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INTERNAL_ERROR",
	5: "NO_RESOURCE",
	8: "INVALID_ID",
}

var ClientTransactionGetResponse_Status_value = map[string]int32{
	"STATUS_UNSET":   0,
	"OK":             1,
	"INTERNAL_ERROR": 2,
	"NO_RESOURCE":    5,
	"INVALID_ID":     8,
}

func (x ClientTransactionGetResponse_Status) String() string {
	return proto.EnumName(ClientTransactionGetResponse_Status_name, int32(x))
}

func (ClientTransactionGetResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b8feb52ffa146393, []int{3, 0}
}

// A request to return a list of txns from the validator. May include the id
// of a particular block to be the `head` of the chain being requested. In that
// case the list will include the txns from that block, and all txns
// previous to that block on the chain. Filter with specific `transaction_ids`.
type ClientTransactionListRequest struct {
	HeadId               string                `protobuf:"bytes,1,opt,name=head_id,json=headId,proto3" json:"head_id,omitempty"`
	TransactionIds       []string              `protobuf:"bytes,2,rep,name=transaction_ids,json=transactionIds,proto3" json:"transaction_ids,omitempty"`
	Paging               *ClientPagingControls `protobuf:"bytes,3,opt,name=paging,proto3" json:"paging,omitempty"`
	Sorting              []*ClientSortControls `protobuf:"bytes,4,rep,name=sorting,proto3" json:"sorting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ClientTransactionListRequest) Reset()         { *m = ClientTransactionListRequest{} }
func (m *ClientTransactionListRequest) String() string { return proto.CompactTextString(m) }
func (*ClientTransactionListRequest) ProtoMessage()    {}
func (*ClientTransactionListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8feb52ffa146393, []int{0}
}

func (m *ClientTransactionListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientTransactionListRequest.Unmarshal(m, b)
}
func (m *ClientTransactionListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientTransactionListRequest.Marshal(b, m, deterministic)
}
func (m *ClientTransactionListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientTransactionListRequest.Merge(m, src)
}
func (m *ClientTransactionListRequest) XXX_Size() int {
	return xxx_messageInfo_ClientTransactionListRequest.Size(m)
}
func (m *ClientTransactionListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientTransactionListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientTransactionListRequest proto.InternalMessageInfo

func (m *ClientTransactionListRequest) GetHeadId() string {
	if m != nil {
		return m.HeadId
	}
	return ""
}

func (m *ClientTransactionListRequest) GetTransactionIds() []string {
	if m != nil {
		return m.TransactionIds
	}
	return nil
}

func (m *ClientTransactionListRequest) GetPaging() *ClientPagingControls {
	if m != nil {
		return m.Paging
	}
	return nil
}

func (m *ClientTransactionListRequest) GetSorting() []*ClientSortControls {
	if m != nil {
		return m.Sorting
	}
	return nil
}

// A response that lists transactions from newest to oldest.
//
// Statuses:
//   * OK - everything worked as expected
//   * INTERNAL_ERROR - general error, such as protobuf failing to deserialize
//   * NOT_READY - the validator does not yet have a genesis block
//   * NO_ROOT - the head block specified was not found
//   * NO_RESOURCE - no txns were found with the parameters specified
//   * INVALID_PAGING - the paging controls were malformed or out of range
//   * INVALID_SORT - the sorting controls were malformed or invalid
type ClientTransactionListResponse struct {
	Status               ClientTransactionListResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=sawtooth.sdk.protobuf.ClientTransactionListResponse_Status" json:"status,omitempty"`
	Transactions         []*Transaction                       `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
	HeadId               string                               `protobuf:"bytes,3,opt,name=head_id,json=headId,proto3" json:"head_id,omitempty"`
	Paging               *ClientPagingResponse                `protobuf:"bytes,4,opt,name=paging,proto3" json:"paging,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *ClientTransactionListResponse) Reset()         { *m = ClientTransactionListResponse{} }
func (m *ClientTransactionListResponse) String() string { return proto.CompactTextString(m) }
func (*ClientTransactionListResponse) ProtoMessage()    {}
func (*ClientTransactionListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8feb52ffa146393, []int{1}
}

func (m *ClientTransactionListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientTransactionListResponse.Unmarshal(m, b)
}
func (m *ClientTransactionListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientTransactionListResponse.Marshal(b, m, deterministic)
}
func (m *ClientTransactionListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientTransactionListResponse.Merge(m, src)
}
func (m *ClientTransactionListResponse) XXX_Size() int {
	return xxx_messageInfo_ClientTransactionListResponse.Size(m)
}
func (m *ClientTransactionListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientTransactionListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClientTransactionListResponse proto.InternalMessageInfo

func (m *ClientTransactionListResponse) GetStatus() ClientTransactionListResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientTransactionListResponse_STATUS_UNSET
}

func (m *ClientTransactionListResponse) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *ClientTransactionListResponse) GetHeadId() string {
	if m != nil {
		return m.HeadId
	}
	return ""
}

func (m *ClientTransactionListResponse) GetPaging() *ClientPagingResponse {
	if m != nil {
		return m.Paging
	}
	return nil
}

// Fetches a specific txn by its id (header_signature) from the blockchain.
type ClientTransactionGetRequest struct {
	TransactionId        string   `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientTransactionGetRequest) Reset()         { *m = ClientTransactionGetRequest{} }
func (m *ClientTransactionGetRequest) String() string { return proto.CompactTextString(m) }
func (*ClientTransactionGetRequest) ProtoMessage()    {}
func (*ClientTransactionGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8feb52ffa146393, []int{2}
}

func (m *ClientTransactionGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientTransactionGetRequest.Unmarshal(m, b)
}
func (m *ClientTransactionGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientTransactionGetRequest.Marshal(b, m, deterministic)
}
func (m *ClientTransactionGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientTransactionGetRequest.Merge(m, src)
}
func (m *ClientTransactionGetRequest) XXX_Size() int {
	return xxx_messageInfo_ClientTransactionGetRequest.Size(m)
}
func (m *ClientTransactionGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientTransactionGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientTransactionGetRequest proto.InternalMessageInfo

func (m *ClientTransactionGetRequest) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

// A response that returns the txn specified by a ClientTransactionGetRequest.
//
// Statuses:
//   * OK - everything worked as expected, txn has been fetched
//   * INTERNAL_ERROR - general error, such as protobuf failing to deserialize
//   * NO_RESOURCE - no txn with the specified id exists
type ClientTransactionGetResponse struct {
	Status               ClientTransactionGetResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=sawtooth.sdk.protobuf.ClientTransactionGetResponse_Status" json:"status,omitempty"`
	Transaction          *Transaction                        `protobuf:"bytes,2,opt,name=transaction,proto3" json:"transaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *ClientTransactionGetResponse) Reset()         { *m = ClientTransactionGetResponse{} }
func (m *ClientTransactionGetResponse) String() string { return proto.CompactTextString(m) }
func (*ClientTransactionGetResponse) ProtoMessage()    {}
func (*ClientTransactionGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8feb52ffa146393, []int{3}
}

func (m *ClientTransactionGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientTransactionGetResponse.Unmarshal(m, b)
}
func (m *ClientTransactionGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientTransactionGetResponse.Marshal(b, m, deterministic)
}
func (m *ClientTransactionGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientTransactionGetResponse.Merge(m, src)
}
func (m *ClientTransactionGetResponse) XXX_Size() int {
	return xxx_messageInfo_ClientTransactionGetResponse.Size(m)
}
func (m *ClientTransactionGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientTransactionGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClientTransactionGetResponse proto.InternalMessageInfo

func (m *ClientTransactionGetResponse) GetStatus() ClientTransactionGetResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientTransactionGetResponse_STATUS_UNSET
}

func (m *ClientTransactionGetResponse) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func init() {
	proto.RegisterEnum("sawtooth.sdk.protobuf.ClientTransactionListResponse_Status", ClientTransactionListResponse_Status_name, ClientTransactionListResponse_Status_value)
	proto.RegisterEnum("sawtooth.sdk.protobuf.ClientTransactionGetResponse_Status", ClientTransactionGetResponse_Status_name, ClientTransactionGetResponse_Status_value)
	proto.RegisterType((*ClientTransactionListRequest)(nil), "sawtooth.sdk.protobuf.ClientTransactionListRequest")
	proto.RegisterType((*ClientTransactionListResponse)(nil), "sawtooth.sdk.protobuf.ClientTransactionListResponse")
	proto.RegisterType((*ClientTransactionGetRequest)(nil), "sawtooth.sdk.protobuf.ClientTransactionGetRequest")
	proto.RegisterType((*ClientTransactionGetResponse)(nil), "sawtooth.sdk.protobuf.ClientTransactionGetResponse")
}

func init() { proto.RegisterFile("client_transaction.proto", fileDescriptor_b8feb52ffa146393) }

var fileDescriptor_b8feb52ffa146393 = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xdf, 0x8e, 0x93, 0x40,
	0x14, 0xc6, 0x05, 0x56, 0x6a, 0x0f, 0xbb, 0xec, 0x38, 0x89, 0x91, 0x5d, 0x35, 0x69, 0x48, 0x8c,
	0x35, 0x26, 0x5c, 0xd4, 0x3b, 0xbd, 0x42, 0xc0, 0x86, 0xd8, 0x40, 0x33, 0x50, 0x8d, 0x57, 0x13,
	0xb6, 0xe0, 0x2e, 0xb1, 0x61, 0x2a, 0x33, 0x8d, 0xcf, 0xe1, 0xa5, 0x2f, 0xe2, 0x43, 0xf9, 0x14,
	0xa6, 0xfc, 0xd9, 0xa5, 0xff, 0x36, 0x8d, 0x77, 0xcc, 0x99, 0x39, 0xbf, 0xf9, 0xe6, 0xfb, 0x4e,
	0x00, 0x63, 0xbe, 0xc8, 0xb3, 0x42, 0x50, 0x51, 0x26, 0x05, 0x4f, 0xe6, 0x22, 0x67, 0x85, 0xb5,
	0x2c, 0x99, 0x60, 0xf8, 0x09, 0x4f, 0x7e, 0x0a, 0xc6, 0xc4, 0x8d, 0xc5, 0xd3, 0xef, 0x75, 0xed,
	0x6a, 0xf5, 0xed, 0xf2, 0xf1, 0xce, 0xc9, 0xcb, 0x8b, 0x86, 0xb1, 0xc8, 0xb9, 0xa0, 0x73, 0x56,
	0x88, 0x92, 0x2d, 0xea, 0x2d, 0xf3, 0xaf, 0x04, 0xcf, 0x9d, 0x6a, 0x37, 0xbe, 0x6b, 0x9b, 0xe4,
	0x5c, 0x90, 0xec, 0xc7, 0x2a, 0xe3, 0x02, 0x3f, 0x85, 0xde, 0x4d, 0x96, 0xa4, 0x34, 0x4f, 0x0d,
	0x69, 0x20, 0x0d, 0xfb, 0x44, 0x5d, 0x2f, 0xfd, 0x14, 0xbf, 0x82, 0xf3, 0xce, 0x4d, 0x34, 0x4f,
	0xb9, 0x21, 0x0f, 0x94, 0x61, 0x9f, 0xe8, 0x9d, 0xb2, 0x9f, 0x72, 0xec, 0x80, 0xba, 0x4c, 0xae,
	0xf3, 0xe2, 0xda, 0x50, 0x06, 0xd2, 0x50, 0x1b, 0xbd, 0xb1, 0xf6, 0x0a, 0xb7, 0x6a, 0x19, 0xd3,
	0xea, 0xa8, 0x53, 0x8b, 0xe4, 0xa4, 0x69, 0xc5, 0x0e, 0xf4, 0x38, 0x2b, 0xc5, 0x9a, 0x72, 0x32,
	0x50, 0x86, 0xda, 0xe8, 0xf5, 0xbd, 0x94, 0x88, 0x95, 0xe2, 0x96, 0xd1, 0x76, 0x9a, 0x7f, 0x14,
	0x78, 0x71, 0xe0, 0xb1, 0x7c, 0xc9, 0x0a, 0x9e, 0xe1, 0x08, 0x54, 0x2e, 0x12, 0xb1, 0xe2, 0xd5,
	0x63, 0xf5, 0xd1, 0xfb, 0x7b, 0x6f, 0x39, 0x40, 0xb1, 0xa2, 0x0a, 0x41, 0x1a, 0x14, 0xfe, 0x08,
	0xa7, 0x1d, 0x4b, 0x6a, 0x9b, 0xb4, 0x91, 0x79, 0x00, 0xdd, 0x81, 0x92, 0x8d, 0xbe, 0x6e, 0x14,
	0xca, 0x46, 0x14, 0x77, 0x0e, 0x9f, 0x1c, 0xed, 0x70, 0x2b, 0xb6, 0x75, 0xd8, 0xfc, 0x2d, 0x81,
	0x5a, 0x0b, 0xc7, 0x08, 0x4e, 0xa3, 0xd8, 0x8e, 0x67, 0x11, 0x9d, 0x05, 0x91, 0x17, 0xa3, 0x07,
	0x58, 0x05, 0x39, 0xfc, 0x84, 0x24, 0x8c, 0x41, 0xf7, 0x83, 0xd8, 0x23, 0x81, 0x3d, 0xa1, 0x1e,
	0x21, 0x21, 0x41, 0x32, 0x3e, 0x83, 0x7e, 0x10, 0xc6, 0x94, 0x78, 0xb6, 0xfb, 0x15, 0x29, 0x58,
	0x83, 0x5e, 0x10, 0x52, 0x12, 0x86, 0x31, 0x3a, 0xc1, 0xe7, 0xa0, 0xad, 0x17, 0x5e, 0x14, 0xce,
	0x88, 0xe3, 0xa1, 0x87, 0x35, 0xe0, 0xb3, 0x3d, 0xf1, 0x5d, 0x3a, 0xb5, 0xc7, 0x7e, 0x30, 0x46,
	0xea, 0xfa, 0xba, 0xb6, 0x16, 0x85, 0x24, 0x46, 0x3d, 0xac, 0x03, 0xb4, 0x15, 0xdf, 0x45, 0x8f,
	0x4c, 0x17, 0x9e, 0xed, 0x38, 0x3e, 0xce, 0x6e, 0x67, 0xf4, 0x25, 0xe8, 0x9b, 0xa3, 0xd8, 0x8c,
	0xea, 0xd9, 0xc6, 0x24, 0x9a, 0xbf, 0xe4, 0x3d, 0xb3, 0x5e, 0x61, 0x9a, 0xf4, 0xc9, 0x56, 0xfa,
	0xef, 0x8e, 0x4d, 0xbf, 0x03, 0xd9, 0x0e, 0xdf, 0x05, 0xad, 0xa3, 0xc2, 0x90, 0xab, 0x80, 0x8e,
	0xc9, 0xbe, 0xdb, 0x66, 0x7e, 0xf9, 0xcf, 0x6c, 0x76, 0xfc, 0xdf, 0x72, 0xf6, 0xc3, 0x05, 0xec,
	0xff, 0x8d, 0x4c, 0xa5, 0x2b, 0xb5, 0xfa, 0x7e, 0xfb, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x74, 0xc9,
	0x65, 0xe3, 0x82, 0x04, 0x00, 0x00,
}
