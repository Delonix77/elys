// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/masterchef/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params defines the parameters for the module.
type Params struct {
	LpIncentives *IncentiveInfo `protobuf:"bytes,1,opt,name=lp_incentives,json=lpIncentives,proto3" json:"lp_incentives,omitempty"`
	// gas fees and swap fees portion for lps, `100 - reward_portion_for_lps - reward_portion_for_stakers = revenue percent for protocol`.
	RewardPortionForLps github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=reward_portion_for_lps,json=rewardPortionForLps,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reward_portion_for_lps"`
	// gas fees and swap fees portion for stakers, `100 - reward_portion_for_lps - reward_portion_for_stakers = revenue percent for protocol`.
	RewardPortionForStakers github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=reward_portion_for_stakers,json=rewardPortionForStakers,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reward_portion_for_stakers"`
	// Tracking dex rewards given to LPs
	DexRewardsLps DexRewardsTracker `protobuf:"bytes,4,opt,name=dex_rewards_lps,json=dexRewardsLps,proto3" json:"dex_rewards_lps"`
	// Maximum eden reward apr for lps - [0 - 0.3]
	MaxEdenRewardAprLps    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=max_eden_reward_apr_lps,json=maxEdenRewardAprLps,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"max_eden_reward_apr_lps"`
	SupportedRewardDenoms  []*SupportedRewardDenom                `protobuf:"bytes,6,rep,name=supported_reward_denoms,json=supportedRewardDenoms,proto3" json:"supported_reward_denoms,omitempty"`
	ProtocolRevenueAddress string                                 `protobuf:"bytes,7,opt,name=protocol_revenue_address,json=protocolRevenueAddress,proto3" json:"protocol_revenue_address,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc83f5a7f5a55e20, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetLpIncentives() *IncentiveInfo {
	if m != nil {
		return m.LpIncentives
	}
	return nil
}

func (m *Params) GetDexRewardsLps() DexRewardsTracker {
	if m != nil {
		return m.DexRewardsLps
	}
	return DexRewardsTracker{}
}

func (m *Params) GetSupportedRewardDenoms() []*SupportedRewardDenom {
	if m != nil {
		return m.SupportedRewardDenoms
	}
	return nil
}

func (m *Params) GetProtocolRevenueAddress() string {
	if m != nil {
		return m.ProtocolRevenueAddress
	}
	return ""
}

type SupportedRewardDenom struct {
	Denom     string                                 `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	MinAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=min_amount,json=minAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_amount"`
}

func (m *SupportedRewardDenom) Reset()         { *m = SupportedRewardDenom{} }
func (m *SupportedRewardDenom) String() string { return proto.CompactTextString(m) }
func (*SupportedRewardDenom) ProtoMessage()    {}
func (*SupportedRewardDenom) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc83f5a7f5a55e20, []int{1}
}
func (m *SupportedRewardDenom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SupportedRewardDenom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SupportedRewardDenom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SupportedRewardDenom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupportedRewardDenom.Merge(m, src)
}
func (m *SupportedRewardDenom) XXX_Size() int {
	return m.Size()
}
func (m *SupportedRewardDenom) XXX_DiscardUnknown() {
	xxx_messageInfo_SupportedRewardDenom.DiscardUnknown(m)
}

var xxx_messageInfo_SupportedRewardDenom proto.InternalMessageInfo

func (m *SupportedRewardDenom) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "elys.masterchef.Params")
	proto.RegisterType((*SupportedRewardDenom)(nil), "elys.masterchef.SupportedRewardDenom")
}

func init() { proto.RegisterFile("elys/masterchef/params.proto", fileDescriptor_bc83f5a7f5a55e20) }

var fileDescriptor_bc83f5a7f5a55e20 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x41, 0x6b, 0x1a, 0x41,
	0x18, 0x75, 0x1b, 0x63, 0x71, 0xd2, 0x10, 0xd8, 0xda, 0xb8, 0x48, 0x59, 0x45, 0x68, 0xf1, 0x92,
	0x5d, 0x48, 0x2f, 0xa5, 0x37, 0xad, 0x2d, 0x08, 0x2d, 0xc8, 0xa6, 0xa7, 0x42, 0x19, 0x26, 0x3b,
	0x9f, 0x66, 0x71, 0x77, 0x66, 0x98, 0x19, 0x93, 0x0d, 0xfd, 0x13, 0x3d, 0xf6, 0xd8, 0x9f, 0x93,
	0x63, 0x8e, 0xa5, 0x87, 0x50, 0xf4, 0x3f, 0xf4, 0x5c, 0x76, 0x46, 0x8d, 0xa8, 0x27, 0x4f, 0xbb,
	0x33, 0xef, 0xf1, 0xde, 0xfb, 0x1e, 0xdf, 0xa0, 0x97, 0x90, 0xde, 0xaa, 0x30, 0x23, 0x4a, 0x83,
	0x8c, 0xaf, 0x60, 0x14, 0x0a, 0x22, 0x49, 0xa6, 0x02, 0x21, 0xb9, 0xe6, 0xee, 0x49, 0x81, 0x06,
	0x8f, 0x68, 0xa3, 0x36, 0xe6, 0x63, 0x6e, 0xb0, 0xb0, 0xf8, 0xb3, 0xb4, 0x46, 0x73, 0x53, 0x24,
	0x61, 0x31, 0x30, 0x9d, 0x5c, 0xc3, 0x82, 0xd0, 0xd8, 0x72, 0xe1, 0x3c, 0x5d, 0x60, 0x9d, 0x4d,
	0x8c, 0x42, 0x8e, 0x25, 0xdc, 0x10, 0x49, 0x15, 0xd6, 0x92, 0x4c, 0x40, 0x5a, 0x66, 0xfb, 0x5f,
	0x19, 0x55, 0x86, 0x26, 0x9e, 0xfb, 0x1e, 0x1d, 0xa7, 0x02, 0xaf, 0x6c, 0x94, 0xe7, 0xb4, 0x9c,
	0xce, 0xd1, 0xb9, 0x1f, 0x6c, 0x04, 0x0e, 0x06, 0x4b, 0xca, 0x80, 0x8d, 0x78, 0xf4, 0x2c, 0x15,
	0xab, 0x0b, 0xe5, 0xc6, 0xe8, 0xd4, 0xfa, 0x60, 0xc1, 0xa5, 0x4e, 0x38, 0xc3, 0x23, 0x2e, 0x71,
	0x2a, 0x94, 0xf7, 0xa4, 0xe5, 0x74, 0xaa, 0xbd, 0xe0, 0xee, 0xa1, 0x59, 0xfa, 0xf3, 0xd0, 0x7c,
	0x3d, 0x4e, 0xf4, 0xd5, 0xf4, 0x32, 0x88, 0x79, 0x16, 0xc6, 0x5c, 0x65, 0x5c, 0x2d, 0x3e, 0x67,
	0x8a, 0x4e, 0x42, 0x7d, 0x2b, 0x40, 0x05, 0x7d, 0x88, 0xa3, 0xe7, 0x56, 0x6d, 0x68, 0xc5, 0x3e,
	0x72, 0xf9, 0x49, 0x28, 0x77, 0x82, 0x1a, 0x3b, 0x4c, 0x94, 0x2e, 0xe6, 0x52, 0xde, 0xc1, 0x5e,
	0x46, 0xf5, 0x4d, 0xa3, 0x0b, 0x2b, 0xe7, 0x0e, 0xd1, 0xc9, 0x7a, 0x7b, 0xc5, 0x28, 0x65, 0x53,
	0x4c, 0x7b, 0xab, 0x98, 0x3e, 0xe4, 0x91, 0xa5, 0x7d, 0x91, 0x24, 0x9e, 0x80, 0xec, 0x95, 0x8b,
	0x14, 0xd1, 0x31, 0x5d, 0x01, 0x45, 0x7c, 0x8a, 0xea, 0x19, 0xc9, 0x31, 0x50, 0x60, 0x0b, 0x59,
	0x4c, 0x84, 0x2d, 0xe9, 0x70, 0xbf, 0x92, 0x32, 0x92, 0x7f, 0xa0, 0xc0, 0xac, 0x47, 0x57, 0x98,
	0x92, 0xbe, 0xa1, 0xba, 0x9a, 0x8a, 0xa2, 0x20, 0xa0, 0x4b, 0x1b, 0x0a, 0x8c, 0x67, 0xca, 0xab,
	0xb4, 0x0e, 0x3a, 0x47, 0xe7, 0xaf, 0xb6, 0xf2, 0x5f, 0x2c, 0xf9, 0x56, 0xa8, 0x5f, 0xb0, 0xa3,
	0x17, 0x6a, 0xc7, 0xad, 0x72, 0xdf, 0x22, 0xcf, 0x6c, 0x50, 0xcc, 0x53, 0x2c, 0xe1, 0x1a, 0xd8,
	0x14, 0x30, 0xa1, 0x54, 0x82, 0x52, 0xde, 0xd3, 0x62, 0x8a, 0xe8, 0x74, 0x89, 0x47, 0x16, 0xee,
	0x5a, 0xf4, 0x5d, 0xf9, 0xe7, 0xaf, 0x66, 0xa9, 0xfd, 0x1d, 0xd5, 0x76, 0xd9, 0xb9, 0x35, 0x74,
	0x68, 0x52, 0x9a, 0xed, 0xab, 0x46, 0xf6, 0xe0, 0x7e, 0x46, 0x28, 0x4b, 0x18, 0x26, 0x19, 0x9f,
	0x32, 0xbd, 0xc7, 0x2a, 0x0d, 0x98, 0x8e, 0xaa, 0x59, 0xc2, 0xba, 0x46, 0xa0, 0x37, 0xb8, 0x9b,
	0xf9, 0xce, 0xfd, 0xcc, 0x77, 0xfe, 0xce, 0x7c, 0xe7, 0xc7, 0xdc, 0x2f, 0xdd, 0xcf, 0xfd, 0xd2,
	0xef, 0xb9, 0x5f, 0xfa, 0x1a, 0xae, 0x89, 0x15, 0xf5, 0x9c, 0x31, 0xd0, 0x37, 0x5c, 0x4e, 0xcc,
	0x21, 0xcc, 0xd7, 0xdf, 0x94, 0x51, 0xbe, 0xac, 0x98, 0x29, 0xdf, 0xfc, 0x0f, 0x00, 0x00, 0xff,
	0xff, 0xcf, 0x0d, 0x11, 0x52, 0xf5, 0x03, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ProtocolRevenueAddress) > 0 {
		i -= len(m.ProtocolRevenueAddress)
		copy(dAtA[i:], m.ProtocolRevenueAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ProtocolRevenueAddress)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.SupportedRewardDenoms) > 0 {
		for iNdEx := len(m.SupportedRewardDenoms) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SupportedRewardDenoms[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	{
		size := m.MaxEdenRewardAprLps.Size()
		i -= size
		if _, err := m.MaxEdenRewardAprLps.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.DexRewardsLps.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.RewardPortionForStakers.Size()
		i -= size
		if _, err := m.RewardPortionForStakers.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.RewardPortionForLps.Size()
		i -= size
		if _, err := m.RewardPortionForLps.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.LpIncentives != nil {
		{
			size, err := m.LpIncentives.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SupportedRewardDenom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SupportedRewardDenom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SupportedRewardDenom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MinAmount.Size()
		i -= size
		if _, err := m.MinAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LpIncentives != nil {
		l = m.LpIncentives.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.RewardPortionForLps.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.RewardPortionForStakers.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.DexRewardsLps.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MaxEdenRewardAprLps.Size()
	n += 1 + l + sovParams(uint64(l))
	if len(m.SupportedRewardDenoms) > 0 {
		for _, e := range m.SupportedRewardDenoms {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = len(m.ProtocolRevenueAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func (m *SupportedRewardDenom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.MinAmount.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LpIncentives", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LpIncentives == nil {
				m.LpIncentives = &IncentiveInfo{}
			}
			if err := m.LpIncentives.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardPortionForLps", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardPortionForLps.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardPortionForStakers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardPortionForStakers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DexRewardsLps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DexRewardsLps.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxEdenRewardAprLps", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxEdenRewardAprLps.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SupportedRewardDenoms", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SupportedRewardDenoms = append(m.SupportedRewardDenoms, &SupportedRewardDenom{})
			if err := m.SupportedRewardDenoms[len(m.SupportedRewardDenoms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtocolRevenueAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProtocolRevenueAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *SupportedRewardDenom) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: SupportedRewardDenom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SupportedRewardDenom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)