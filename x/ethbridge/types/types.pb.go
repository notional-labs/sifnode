// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sifnode/ethbridge/v1/types.proto

package types

import (
	fmt "fmt"
	types "github.com/Sifchain/sifnode/x/oracle/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Claim type enum
type ClaimType int32

const (
	// Unspecified claim type
	ClaimType_CLAIM_TYPE_UNSPECIFIED ClaimType = 0
	// Burn claim type
	ClaimType_CLAIM_TYPE_BURN ClaimType = 1
	// Lock claim type
	ClaimType_CLAIM_TYPE_LOCK ClaimType = 2
)

var ClaimType_name = map[int32]string{
	0: "CLAIM_TYPE_UNSPECIFIED",
	1: "CLAIM_TYPE_BURN",
	2: "CLAIM_TYPE_LOCK",
}

var ClaimType_value = map[string]int32{
	"CLAIM_TYPE_UNSPECIFIED": 0,
	"CLAIM_TYPE_BURN":        1,
	"CLAIM_TYPE_LOCK":        2,
}

func (x ClaimType) String() string {
	return proto.EnumName(ClaimType_name, int32(x))
}

func (ClaimType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4cb34f678c9ed59f, []int{0}
}

// EthBridgeClaim is a structure that contains all the data for a particular
// bridge claim
type EthBridgeClaim struct {
	NetworkDescriptor types.NetworkDescriptor `protobuf:"varint,1,opt,name=network_descriptor,json=networkDescriptor,proto3,enum=sifnode.oracle.v1.NetworkDescriptor" json:"network_descriptor,omitempty" yaml:"network_descriptor"`
	// bridge_contract_address is an EthereumAddress
	BridgeContractAddress string `protobuf:"bytes,2,opt,name=bridge_contract_address,json=bridgeContractAddress,proto3" json:"bridge_contract_address,omitempty" yaml:"bridge_contract_address"`
	Nonce                 int64  `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty" yaml:"nonce"`
	Symbol                string `protobuf:"bytes,4,opt,name=symbol,proto3" json:"symbol,omitempty" yaml:"symbol"`
	// token_contract_address is an EthereumAddress
	TokenContractAddress string `protobuf:"bytes,5,opt,name=token_contract_address,json=tokenContractAddress,proto3" json:"token_contract_address,omitempty" yaml:"token_contract_address"`
	// ethereum_sender is an EthereumAddress
	EthereumSender string `protobuf:"bytes,6,opt,name=ethereum_sender,json=ethereumSender,proto3" json:"ethereum_sender,omitempty" yaml:"ethereum_sender"`
	// cosmos_receiver is an sdk.AccAddress
	CosmosReceiver string `protobuf:"bytes,7,opt,name=cosmos_receiver,json=cosmosReceiver,proto3" json:"cosmos_receiver,omitempty" yaml:"cosmos_receiver"`
	// validator_address is an sdk.ValAddress
	ValidatorAddress string                                 `protobuf:"bytes,8,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty" yaml:"validator_address"`
	Amount           github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,9,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount" yaml:"amount"`
	ClaimType        ClaimType                              `protobuf:"varint,10,opt,name=claim_type,json=claimType,proto3,enum=sifnode.ethbridge.v1.ClaimType" json:"claim_type,omitempty"`
	// Name of coin different then symbol
	TokenName string `protobuf:"bytes,11,opt,name=token_name,json=tokenName,proto3" json:"token_name,omitempty" yaml:"token_name"`
	// Number of decimals in ERC20 token
	Decimals int32 `protobuf:"varint,12,opt,name=decimals,proto3" json:"decimals,omitempty" yaml:"token_decimals"`
	// The hashed concatination of ERC20 address, network descriptor
	// Token decimals, name and symbol
	// Marked optional because messages sent by the relayer may not be hashed
	DenomHash string `protobuf:"bytes,13,opt,name=denom_hash,json=denomHash,proto3" json:"denom_hash,omitempty" yaml:"denom_hash"`
}

func (m *EthBridgeClaim) Reset()         { *m = EthBridgeClaim{} }
func (m *EthBridgeClaim) String() string { return proto.CompactTextString(m) }
func (*EthBridgeClaim) ProtoMessage()    {}
func (*EthBridgeClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cb34f678c9ed59f, []int{0}
}
func (m *EthBridgeClaim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EthBridgeClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EthBridgeClaim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EthBridgeClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthBridgeClaim.Merge(m, src)
}
func (m *EthBridgeClaim) XXX_Size() int {
	return m.Size()
}
func (m *EthBridgeClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_EthBridgeClaim.DiscardUnknown(m)
}

var xxx_messageInfo_EthBridgeClaim proto.InternalMessageInfo

func (m *EthBridgeClaim) GetNetworkDescriptor() types.NetworkDescriptor {
	if m != nil {
		return m.NetworkDescriptor
	}
	return types.NetworkDescriptor_NETWORK_DESCRIPTOR_UNSPECIFIED
}

func (m *EthBridgeClaim) GetBridgeContractAddress() string {
	if m != nil {
		return m.BridgeContractAddress
	}
	return ""
}

func (m *EthBridgeClaim) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *EthBridgeClaim) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *EthBridgeClaim) GetTokenContractAddress() string {
	if m != nil {
		return m.TokenContractAddress
	}
	return ""
}

func (m *EthBridgeClaim) GetEthereumSender() string {
	if m != nil {
		return m.EthereumSender
	}
	return ""
}

func (m *EthBridgeClaim) GetCosmosReceiver() string {
	if m != nil {
		return m.CosmosReceiver
	}
	return ""
}

func (m *EthBridgeClaim) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *EthBridgeClaim) GetClaimType() ClaimType {
	if m != nil {
		return m.ClaimType
	}
	return ClaimType_CLAIM_TYPE_UNSPECIFIED
}

func (m *EthBridgeClaim) GetTokenName() string {
	if m != nil {
		return m.TokenName
	}
	return ""
}

func (m *EthBridgeClaim) GetDecimals() int32 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *EthBridgeClaim) GetDenomHash() string {
	if m != nil {
		return m.DenomHash
	}
	return ""
}

type PeggyTokens struct {
	Tokens []string `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
}

func (m *PeggyTokens) Reset()         { *m = PeggyTokens{} }
func (m *PeggyTokens) String() string { return proto.CompactTextString(m) }
func (*PeggyTokens) ProtoMessage()    {}
func (*PeggyTokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cb34f678c9ed59f, []int{1}
}
func (m *PeggyTokens) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PeggyTokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PeggyTokens.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PeggyTokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeggyTokens.Merge(m, src)
}
func (m *PeggyTokens) XXX_Size() int {
	return m.Size()
}
func (m *PeggyTokens) XXX_DiscardUnknown() {
	xxx_messageInfo_PeggyTokens.DiscardUnknown(m)
}

var xxx_messageInfo_PeggyTokens proto.InternalMessageInfo

func (m *PeggyTokens) GetTokens() []string {
	if m != nil {
		return m.Tokens
	}
	return nil
}

// GenesisState for ethbridge
type GenesisState struct {
	CrosschainFeeReceiveAccount string   `protobuf:"bytes,1,opt,name=crosschain_fee_receive_account,json=crosschainFeeReceiveAccount,proto3" json:"crosschain_fee_receive_account,omitempty"`
	PeggyTokens                 []string `protobuf:"bytes,2,rep,name=peggy_tokens,json=peggyTokens,proto3" json:"peggy_tokens,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cb34f678c9ed59f, []int{2}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetCrosschainFeeReceiveAccount() string {
	if m != nil {
		return m.CrosschainFeeReceiveAccount
	}
	return ""
}

func (m *GenesisState) GetPeggyTokens() []string {
	if m != nil {
		return m.PeggyTokens
	}
	return nil
}

func init() {
	proto.RegisterEnum("sifnode.ethbridge.v1.ClaimType", ClaimType_name, ClaimType_value)
	proto.RegisterType((*EthBridgeClaim)(nil), "sifnode.ethbridge.v1.EthBridgeClaim")
	proto.RegisterType((*PeggyTokens)(nil), "sifnode.ethbridge.v1.PeggyTokens")
	proto.RegisterType((*GenesisState)(nil), "sifnode.ethbridge.v1.GenesisState")
}

func init() { proto.RegisterFile("sifnode/ethbridge/v1/types.proto", fileDescriptor_4cb34f678c9ed59f) }

var fileDescriptor_4cb34f678c9ed59f = []byte{
	// 746 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xdf, 0x4e, 0xfb, 0x36,
	0x14, 0xc7, 0x1b, 0x18, 0xfd, 0x11, 0x53, 0x4a, 0xeb, 0x41, 0x17, 0xba, 0x91, 0x14, 0x6b, 0x43,
	0x1d, 0x12, 0xa9, 0xd8, 0x9f, 0x9b, 0x5d, 0x6c, 0xa2, 0xa5, 0xb0, 0x6a, 0xac, 0x63, 0x2e, 0x08,
	0x8d, 0x9b, 0x28, 0x4d, 0x4c, 0x1b, 0xd1, 0xc4, 0x55, 0x6c, 0xba, 0xf5, 0x2d, 0xf6, 0x02, 0x7b,
	0x1f, 0x2e, 0xb9, 0x9c, 0x76, 0x11, 0x4d, 0xf0, 0x06, 0x79, 0x82, 0x29, 0x76, 0xd2, 0x56, 0x2d,
	0xbb, 0xaa, 0xfd, 0x3d, 0x1f, 0x7f, 0xcf, 0x71, 0xcf, 0x89, 0x41, 0x8d, 0x79, 0x0f, 0x01, 0x75,
	0x49, 0x83, 0xf0, 0x61, 0x3f, 0xf4, 0xdc, 0x01, 0x69, 0x4c, 0x4e, 0x1b, 0x7c, 0x3a, 0x26, 0xcc,
	0x1c, 0x87, 0x94, 0x53, 0xb8, 0x9b, 0x12, 0xe6, 0x8c, 0x30, 0x27, 0xa7, 0xd5, 0xdd, 0x01, 0x1d,
	0x50, 0x01, 0x34, 0x92, 0x95, 0x64, 0xab, 0xc7, 0x99, 0x1b, 0x0d, 0x6d, 0x67, 0x24, 0xac, 0x02,
	0xc2, 0x7f, 0xa7, 0xe1, 0xa3, 0xe5, 0x12, 0xe6, 0x84, 0xde, 0x98, 0xd3, 0x50, 0xb2, 0xe8, 0xaf,
	0x0f, 0xa0, 0xd8, 0xe6, 0xc3, 0xa6, 0xb0, 0x6c, 0x8d, 0x6c, 0xcf, 0x87, 0x21, 0x80, 0xab, 0xb8,
	0xa6, 0xd4, 0x94, 0x7a, 0xf1, 0xab, 0xcf, 0xcd, 0xac, 0x0e, 0xe9, 0x6d, 0x4e, 0x4e, 0xcd, 0xae,
	0x84, 0xcf, 0x67, 0x6c, 0xf3, 0x20, 0x8e, 0x8c, 0xfd, 0xa9, 0xed, 0x8f, 0xbe, 0x43, 0xab, 0x4e,
	0x08, 0x97, 0x83, 0xe5, 0x13, 0xf0, 0x1e, 0x7c, 0x22, 0x6f, 0x65, 0x39, 0x34, 0xe0, 0xa1, 0xed,
	0x70, 0xcb, 0x76, 0xdd, 0x90, 0x30, 0xa6, 0xad, 0xd5, 0x94, 0xba, 0xda, 0x44, 0x71, 0x64, 0xe8,
	0xd2, 0xf2, 0x7f, 0x40, 0x84, 0xf7, 0x64, 0xa4, 0x95, 0x06, 0xce, 0xa4, 0x0e, 0x8f, 0xc0, 0x46,
	0x40, 0x03, 0x87, 0x68, 0xeb, 0x35, 0xa5, 0xbe, 0xde, 0x2c, 0xc5, 0x91, 0x51, 0x48, 0x8b, 0x4b,
	0x64, 0x84, 0x65, 0x18, 0x7e, 0x09, 0xf2, 0x6c, 0xea, 0xf7, 0xe9, 0x48, 0xfb, 0x48, 0xa4, 0x2c,
	0xc7, 0x91, 0xb1, 0x2d, 0x41, 0xa9, 0x23, 0x9c, 0x02, 0xf0, 0x0e, 0x54, 0x38, 0x7d, 0x24, 0xc1,
	0x6a, 0xb5, 0x1b, 0xe2, 0xe8, 0x61, 0x1c, 0x19, 0x07, 0xf2, 0xe8, 0xfb, 0x1c, 0xc2, 0xbb, 0x22,
	0xb0, 0x5c, 0x6b, 0x0b, 0xec, 0x10, 0x3e, 0x24, 0x21, 0x79, 0xf2, 0x2d, 0x46, 0x02, 0x97, 0x84,
	0x5a, 0x5e, 0x38, 0x56, 0xe3, 0xc8, 0xa8, 0x48, 0xc7, 0x25, 0x00, 0xe1, 0x62, 0xa6, 0xf4, 0x84,
	0x90, 0x98, 0x38, 0x94, 0xf9, 0x94, 0x59, 0x21, 0x71, 0x88, 0x37, 0x21, 0xa1, 0xf6, 0x61, 0xd9,
	0x64, 0x09, 0x40, 0xb8, 0x28, 0x15, 0x9c, 0x0a, 0xb0, 0x03, 0xca, 0x13, 0x7b, 0xe4, 0xb9, 0x36,
	0xa7, 0xe1, 0xec, 0x76, 0x9b, 0xc2, 0xe6, 0xb3, 0x38, 0x32, 0x34, 0x69, 0xb3, 0x82, 0x20, 0x5c,
	0x9a, 0x69, 0xd9, 0xa5, 0xee, 0x40, 0xde, 0xf6, 0xe9, 0x53, 0xc0, 0x35, 0x55, 0x9c, 0xff, 0xe1,
	0x39, 0x32, 0x72, 0xff, 0x44, 0xc6, 0xd1, 0xc0, 0xe3, 0xc3, 0xa7, 0xbe, 0xe9, 0x50, 0xbf, 0x21,
	0xb3, 0xa7, 0x3f, 0x27, 0xcc, 0x7d, 0x4c, 0xa7, 0xbf, 0x13, 0xf0, 0x79, 0x1b, 0xa4, 0x0b, 0xc2,
	0xa9, 0x1d, 0xfc, 0x1e, 0x00, 0x27, 0x19, 0x59, 0x2b, 0x61, 0x35, 0x20, 0x26, 0xd4, 0x30, 0xdf,
	0xfb, 0x52, 0x4c, 0x31, 0xda, 0x37, 0xd3, 0x31, 0xc1, 0xaa, 0x93, 0x2d, 0xe1, 0x37, 0x00, 0xc8,
	0xf6, 0x04, 0xb6, 0x4f, 0xb4, 0x2d, 0x51, 0xdc, 0x5e, 0x1c, 0x19, 0xe5, 0xc5, 0xd6, 0x25, 0x31,
	0x84, 0x55, 0xb1, 0xe9, 0xda, 0x3e, 0x81, 0xdf, 0x82, 0x4d, 0x97, 0x38, 0x9e, 0x6f, 0x8f, 0x98,
	0x56, 0xa8, 0x29, 0xf5, 0x8d, 0xe6, 0x7e, 0x1c, 0x19, 0x7b, 0x8b, 0x67, 0xb2, 0x38, 0xc2, 0x33,
	0x34, 0x49, 0xe6, 0x92, 0x80, 0xfa, 0xd6, 0xd0, 0x66, 0x43, 0x6d, 0x7b, 0x39, 0xd9, 0x3c, 0x86,
	0xb0, 0x2a, 0x36, 0x3f, 0x26, 0xeb, 0x2f, 0xc0, 0xd6, 0x35, 0x19, 0x0c, 0xa6, 0x37, 0x89, 0x2f,
	0x83, 0x15, 0x90, 0x17, 0x19, 0x98, 0xa6, 0xd4, 0xd6, 0xeb, 0x2a, 0x4e, 0x77, 0x68, 0x02, 0x0a,
	0x97, 0x24, 0x20, 0xcc, 0x63, 0x3d, 0x6e, 0x73, 0x02, 0x5b, 0x40, 0x77, 0x42, 0xca, 0x98, 0x33,
	0xb4, 0xbd, 0xc0, 0x7a, 0x20, 0x24, 0xeb, 0xb4, 0x65, 0x3b, 0x8e, 0x68, 0x45, 0xf2, 0x3d, 0xab,
	0xf8, 0xd3, 0x39, 0x75, 0x41, 0x48, 0xda, 0xfc, 0x33, 0x89, 0xc0, 0x43, 0x50, 0x18, 0x27, 0xb9,
	0xad, 0x34, 0xe5, 0x9a, 0x48, 0xb9, 0x35, 0x9e, 0xd7, 0x73, 0xfc, 0x2b, 0x50, 0x67, 0xff, 0x2c,
	0xac, 0x82, 0x4a, 0xeb, 0xea, 0xac, 0xf3, 0xb3, 0x75, 0xf3, 0xdb, 0x75, 0xdb, 0xba, 0xed, 0xf6,
	0xae, 0xdb, 0xad, 0xce, 0x45, 0xa7, 0x7d, 0x5e, 0xca, 0xc1, 0x8f, 0xc1, 0xce, 0x42, 0xac, 0x79,
	0x8b, 0xbb, 0x25, 0x65, 0x49, 0xbc, 0xfa, 0xa5, 0xf5, 0x53, 0x69, 0xad, 0x79, 0xf9, 0xfc, 0xaa,
	0x2b, 0x2f, 0xaf, 0xba, 0xf2, 0xef, 0xab, 0xae, 0xfc, 0xf9, 0xa6, 0xe7, 0x5e, 0xde, 0xf4, 0xdc,
	0xdf, 0x6f, 0x7a, 0xee, 0xfe, 0x64, 0x61, 0x5e, 0x7a, 0xde, 0x83, 0xa8, 0xba, 0x91, 0xbd, 0x75,
	0x7f, 0x2c, 0xbc, 0x9d, 0x62, 0x74, 0xfa, 0x79, 0xf1, 0xc2, 0x7d, 0xfd, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xba, 0xef, 0x78, 0x63, 0x5d, 0x05, 0x00, 0x00,
}

func (m *EthBridgeClaim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EthBridgeClaim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EthBridgeClaim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DenomHash) > 0 {
		i -= len(m.DenomHash)
		copy(dAtA[i:], m.DenomHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.DenomHash)))
		i--
		dAtA[i] = 0x6a
	}
	if m.Decimals != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x60
	}
	if len(m.TokenName) > 0 {
		i -= len(m.TokenName)
		copy(dAtA[i:], m.TokenName)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.TokenName)))
		i--
		dAtA[i] = 0x5a
	}
	if m.ClaimType != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ClaimType))
		i--
		dAtA[i] = 0x50
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.CosmosReceiver) > 0 {
		i -= len(m.CosmosReceiver)
		copy(dAtA[i:], m.CosmosReceiver)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.CosmosReceiver)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.EthereumSender) > 0 {
		i -= len(m.EthereumSender)
		copy(dAtA[i:], m.EthereumSender)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.EthereumSender)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.TokenContractAddress) > 0 {
		i -= len(m.TokenContractAddress)
		copy(dAtA[i:], m.TokenContractAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.TokenContractAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x22
	}
	if m.Nonce != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x18
	}
	if len(m.BridgeContractAddress) > 0 {
		i -= len(m.BridgeContractAddress)
		copy(dAtA[i:], m.BridgeContractAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.BridgeContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.NetworkDescriptor != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.NetworkDescriptor))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PeggyTokens) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PeggyTokens) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PeggyTokens) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tokens) > 0 {
		for iNdEx := len(m.Tokens) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Tokens[iNdEx])
			copy(dAtA[i:], m.Tokens[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Tokens[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PeggyTokens) > 0 {
		for iNdEx := len(m.PeggyTokens) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PeggyTokens[iNdEx])
			copy(dAtA[i:], m.PeggyTokens[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.PeggyTokens[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.CrosschainFeeReceiveAccount) > 0 {
		i -= len(m.CrosschainFeeReceiveAccount)
		copy(dAtA[i:], m.CrosschainFeeReceiveAccount)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.CrosschainFeeReceiveAccount)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EthBridgeClaim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NetworkDescriptor != 0 {
		n += 1 + sovTypes(uint64(m.NetworkDescriptor))
	}
	l = len(m.BridgeContractAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Nonce != 0 {
		n += 1 + sovTypes(uint64(m.Nonce))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.TokenContractAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.EthereumSender)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.CosmosReceiver)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.ClaimType != 0 {
		n += 1 + sovTypes(uint64(m.ClaimType))
	}
	l = len(m.TokenName)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Decimals != 0 {
		n += 1 + sovTypes(uint64(m.Decimals))
	}
	l = len(m.DenomHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *PeggyTokens) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Tokens) > 0 {
		for _, s := range m.Tokens {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CrosschainFeeReceiveAccount)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.PeggyTokens) > 0 {
		for _, s := range m.PeggyTokens {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EthBridgeClaim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EthBridgeClaim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EthBridgeClaim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkDescriptor", wireType)
			}
			m.NetworkDescriptor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NetworkDescriptor |= types.NetworkDescriptor(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BridgeContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BridgeContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthereumSender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthereumSender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CosmosReceiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CosmosReceiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimType", wireType)
			}
			m.ClaimType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClaimType |= ClaimType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PeggyTokens) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PeggyTokens: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PeggyTokens: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tokens = append(m.Tokens, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CrosschainFeeReceiveAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CrosschainFeeReceiveAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeggyTokens", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeggyTokens = append(m.PeggyTokens, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
