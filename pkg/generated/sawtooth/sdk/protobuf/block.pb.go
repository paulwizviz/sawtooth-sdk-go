// Code generated by protoc-gen-go. DO NOT EDIT.
// source: block.proto

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

type BlockHeader struct {
	// Block number in the chain
	BlockNum uint64 `protobuf:"varint,1,opt,name=block_num,json=blockNum,proto3" json:"block_num,omitempty"`
	// The header_signature of the previous block that was added to the chain.
	PreviousBlockId string `protobuf:"bytes,2,opt,name=previous_block_id,json=previousBlockId,proto3" json:"previous_block_id,omitempty"`
	// Public key for the component internal to the validator that
	// signed the BlockHeader
	SignerPublicKey string `protobuf:"bytes,3,opt,name=signer_public_key,json=signerPublicKey,proto3" json:"signer_public_key,omitempty"`
	// List of batch.header_signatures that match the order of batches
	// required for the block
	BatchIds []string `protobuf:"bytes,4,rep,name=batch_ids,json=batchIds,proto3" json:"batch_ids,omitempty"`
	// Bytes that are set and verified by the consensus algorithm used to
	// create and validate the block
	Consensus []byte `protobuf:"bytes,5,opt,name=consensus,proto3" json:"consensus,omitempty"`
	// The state_root_hash should match the final state_root after all
	// transactions in the batches have been applied, otherwise the block
	// is not valid
	StateRootHash        string   `protobuf:"bytes,6,opt,name=state_root_hash,json=stateRootHash,proto3" json:"state_root_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e550b1f5926e92d, []int{0}
}

func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (m *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(m, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetBlockNum() uint64 {
	if m != nil {
		return m.BlockNum
	}
	return 0
}

func (m *BlockHeader) GetPreviousBlockId() string {
	if m != nil {
		return m.PreviousBlockId
	}
	return ""
}

func (m *BlockHeader) GetSignerPublicKey() string {
	if m != nil {
		return m.SignerPublicKey
	}
	return ""
}

func (m *BlockHeader) GetBatchIds() []string {
	if m != nil {
		return m.BatchIds
	}
	return nil
}

func (m *BlockHeader) GetConsensus() []byte {
	if m != nil {
		return m.Consensus
	}
	return nil
}

func (m *BlockHeader) GetStateRootHash() string {
	if m != nil {
		return m.StateRootHash
	}
	return ""
}

type Block struct {
	// The serialized version of the BlockHeader
	Header []byte `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The signature derived from signing the header
	HeaderSignature string `protobuf:"bytes,2,opt,name=header_signature,json=headerSignature,proto3" json:"header_signature,omitempty"`
	// A list of batches. The batches may contain new batches that other
	// validators may not have received yet, or they will be all batches needed
	// for block validation when passed to the journal
	Batches              []*Batch `protobuf:"bytes,3,rep,name=batches,proto3" json:"batches,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e550b1f5926e92d, []int{1}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHeader() []byte {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetHeaderSignature() string {
	if m != nil {
		return m.HeaderSignature
	}
	return ""
}

func (m *Block) GetBatches() []*Batch {
	if m != nil {
		return m.Batches
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockHeader)(nil), "sawtooth.sdk.protobuf.BlockHeader")
	proto.RegisterType((*Block)(nil), "sawtooth.sdk.protobuf.Block")
}

func init() { proto.RegisterFile("block.proto", fileDescriptor_8e550b1f5926e92d) }

var fileDescriptor_8e550b1f5926e92d = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4e, 0xf3, 0x30,
	0x10, 0x85, 0xe5, 0x3f, 0x6d, 0xff, 0xc6, 0x29, 0x2a, 0x58, 0x02, 0x19, 0xe8, 0x22, 0xea, 0x02,
	0x05, 0x16, 0x59, 0x80, 0xc4, 0x01, 0xba, 0x6a, 0x85, 0x84, 0x2a, 0x73, 0x00, 0xcb, 0x49, 0x0c,
	0x89, 0xda, 0xc6, 0x95, 0xc7, 0x06, 0x75, 0xcb, 0x5d, 0xb9, 0x07, 0xca, 0xb8, 0x11, 0x9b, 0xee,
	0x66, 0xbe, 0x79, 0x99, 0x79, 0x2f, 0xa6, 0x49, 0xb1, 0x35, 0xe5, 0x26, 0xdf, 0x5b, 0xe3, 0x0c,
	0xbb, 0x04, 0xf5, 0xe5, 0x8c, 0x71, 0x75, 0x0e, 0xd5, 0x91, 0x15, 0xfe, 0xfd, 0x26, 0x29, 0x94,
	0x2b, 0xeb, 0xd0, 0xcf, 0x7f, 0x08, 0x4d, 0x16, 0xdd, 0x37, 0x4b, 0xad, 0x2a, 0x6d, 0xd9, 0x2d,
	0x8d, 0x71, 0x85, 0x6c, 0xfd, 0x8e, 0x93, 0x94, 0x64, 0x03, 0x31, 0x46, 0xf0, 0xea, 0x77, 0xec,
	0x81, 0x5e, 0xec, 0xad, 0xfe, 0x6c, 0x8c, 0x07, 0x19, 0x54, 0x4d, 0xc5, 0xff, 0xa5, 0x24, 0x8b,
	0xc5, 0xb4, 0x1f, 0xe0, 0xb2, 0x55, 0xd5, 0x69, 0xa1, 0xf9, 0x68, 0xb5, 0x95, 0x7b, 0x5f, 0x6c,
	0x9b, 0x52, 0x6e, 0xf4, 0x81, 0x47, 0x41, 0x1b, 0x06, 0x6b, 0xe4, 0x2f, 0xfa, 0x80, 0x47, 0x3b,
	0x4f, 0xb2, 0xa9, 0x80, 0x0f, 0xd2, 0x28, 0x8b, 0xc5, 0x18, 0xc1, 0xaa, 0x02, 0x36, 0xa3, 0x71,
	0x69, 0x5a, 0xd0, 0x2d, 0x78, 0xe0, 0xc3, 0x94, 0x64, 0x13, 0xf1, 0x07, 0xd8, 0x1d, 0x9d, 0x82,
	0x53, 0x4e, 0x4b, 0x6b, 0x8c, 0x93, 0xb5, 0x82, 0x9a, 0x8f, 0xf0, 0xc8, 0x19, 0x62, 0x61, 0x8c,
	0x5b, 0x2a, 0xa8, 0xe7, 0xdf, 0x84, 0x0e, 0xd1, 0x1a, 0xbb, 0xa2, 0xa3, 0x1a, 0xb3, 0x62, 0xbc,
	0x89, 0x38, 0x76, 0xec, 0x9e, 0x9e, 0x87, 0x4a, 0x76, 0xf6, 0x94, 0xf3, 0x56, 0xf7, 0xd9, 0x02,
	0x7f, 0xeb, 0x31, 0x7b, 0xa6, 0xff, 0xd1, 0x9e, 0x06, 0x1e, 0xa5, 0x51, 0x96, 0x3c, 0xce, 0xf2,
	0x93, 0xbf, 0x3a, 0x5f, 0x74, 0x2a, 0xd1, 0x8b, 0x17, 0xd7, 0xf4, 0xf4, 0x93, 0xac, 0x49, 0x31,
	0xc2, 0xfa, 0xe9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x9a, 0xf7, 0xa9, 0xc1, 0x01, 0x00, 0x00,
}