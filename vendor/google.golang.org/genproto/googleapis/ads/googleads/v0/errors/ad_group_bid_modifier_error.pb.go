// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/ad_group_bid_modifier_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible ad group bid modifier errors.
type AdGroupBidModifierErrorEnum_AdGroupBidModifierError int32

const (
	// Enum unspecified.
	AdGroupBidModifierErrorEnum_UNSPECIFIED AdGroupBidModifierErrorEnum_AdGroupBidModifierError = 0
	// The received error code is not known in this version.
	AdGroupBidModifierErrorEnum_UNKNOWN AdGroupBidModifierErrorEnum_AdGroupBidModifierError = 1
	// The criterion ID does not support bid modification.
	AdGroupBidModifierErrorEnum_CRITERION_ID_NOT_SUPPORTED AdGroupBidModifierErrorEnum_AdGroupBidModifierError = 2
	// Cannot override the bid modifier for the given criterion ID if the parent
	// campaign is opted out of the same criterion.
	AdGroupBidModifierErrorEnum_CANNOT_OVERRIDE_OPTED_OUT_CAMPAIGN_CRITERION_BID_MODIFIER AdGroupBidModifierErrorEnum_AdGroupBidModifierError = 3
)

var AdGroupBidModifierErrorEnum_AdGroupBidModifierError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CRITERION_ID_NOT_SUPPORTED",
	3: "CANNOT_OVERRIDE_OPTED_OUT_CAMPAIGN_CRITERION_BID_MODIFIER",
}
var AdGroupBidModifierErrorEnum_AdGroupBidModifierError_value = map[string]int32{
	"UNSPECIFIED":                0,
	"UNKNOWN":                    1,
	"CRITERION_ID_NOT_SUPPORTED": 2,
	"CANNOT_OVERRIDE_OPTED_OUT_CAMPAIGN_CRITERION_BID_MODIFIER": 3,
}

func (x AdGroupBidModifierErrorEnum_AdGroupBidModifierError) String() string {
	return proto.EnumName(AdGroupBidModifierErrorEnum_AdGroupBidModifierError_name, int32(x))
}
func (AdGroupBidModifierErrorEnum_AdGroupBidModifierError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_bid_modifier_error_3bb7ac3c0a520d5f, []int{0, 0}
}

// Container for enum describing possible ad group bid modifier errors.
type AdGroupBidModifierErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupBidModifierErrorEnum) Reset()         { *m = AdGroupBidModifierErrorEnum{} }
func (m *AdGroupBidModifierErrorEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupBidModifierErrorEnum) ProtoMessage()    {}
func (*AdGroupBidModifierErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_bid_modifier_error_3bb7ac3c0a520d5f, []int{0}
}
func (m *AdGroupBidModifierErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupBidModifierErrorEnum.Unmarshal(m, b)
}
func (m *AdGroupBidModifierErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupBidModifierErrorEnum.Marshal(b, m, deterministic)
}
func (dst *AdGroupBidModifierErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupBidModifierErrorEnum.Merge(dst, src)
}
func (m *AdGroupBidModifierErrorEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupBidModifierErrorEnum.Size(m)
}
func (m *AdGroupBidModifierErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupBidModifierErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupBidModifierErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdGroupBidModifierErrorEnum)(nil), "google.ads.googleads.v0.errors.AdGroupBidModifierErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.AdGroupBidModifierErrorEnum_AdGroupBidModifierError", AdGroupBidModifierErrorEnum_AdGroupBidModifierError_name, AdGroupBidModifierErrorEnum_AdGroupBidModifierError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/ad_group_bid_modifier_error.proto", fileDescriptor_ad_group_bid_modifier_error_3bb7ac3c0a520d5f)
}

var fileDescriptor_ad_group_bid_modifier_error_3bb7ac3c0a520d5f = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x6a, 0xe3, 0x30,
	0x18, 0x85, 0xc7, 0x09, 0xcc, 0x80, 0xb2, 0x18, 0xe3, 0xcd, 0xc0, 0xb4, 0x64, 0x91, 0x03, 0xc8,
	0x86, 0xae, 0x4a, 0x29, 0x54, 0xb6, 0x54, 0x23, 0x4a, 0x24, 0xe3, 0xd8, 0x29, 0x14, 0x83, 0x70,
	0x2a, 0xd7, 0x18, 0x92, 0x28, 0x48, 0x4d, 0x8e, 0xd2, 0x03, 0x74, 0xd9, 0x03, 0xf4, 0x10, 0x5d,
	0xf6, 0x44, 0x45, 0x56, 0x92, 0xae, 0xd2, 0x95, 0x1f, 0xfe, 0xdf, 0xff, 0x3d, 0xe9, 0x09, 0xdc,
	0xb4, 0x4a, 0xb5, 0xcb, 0x26, 0xac, 0xa5, 0x09, 0x9d, 0xb4, 0x6a, 0x17, 0x85, 0x8d, 0xd6, 0x4a,
	0x9b, 0xb0, 0x96, 0xa2, 0xd5, 0x6a, 0xbb, 0x11, 0x8b, 0x4e, 0x8a, 0x95, 0x92, 0xdd, 0x53, 0xd7,
	0x68, 0xd1, 0x0f, 0xe1, 0x46, 0xab, 0x67, 0x15, 0x8c, 0xdd, 0x1a, 0xac, 0xa5, 0x81, 0x47, 0x02,
	0xdc, 0x45, 0xd0, 0x11, 0x26, 0xef, 0x1e, 0x38, 0x43, 0x32, 0xb5, 0x90, 0xb8, 0x93, 0xd3, 0x3d,
	0x82, 0xd8, 0x21, 0x59, 0x6f, 0x57, 0x93, 0x17, 0x0f, 0xfc, 0x3b, 0x31, 0x0f, 0xfe, 0x82, 0x51,
	0xc9, 0x66, 0x19, 0x49, 0xe8, 0x2d, 0x25, 0xd8, 0xff, 0x15, 0x8c, 0xc0, 0x9f, 0x92, 0xdd, 0x31,
	0x7e, 0xcf, 0x7c, 0x2f, 0x18, 0x83, 0xff, 0x49, 0x4e, 0x0b, 0x92, 0x53, 0xce, 0x04, 0xc5, 0x82,
	0xf1, 0x42, 0xcc, 0xca, 0x2c, 0xe3, 0x79, 0x41, 0xb0, 0x3f, 0x08, 0xae, 0xc1, 0x65, 0x82, 0x98,
	0xfd, 0xcb, 0xe7, 0x24, 0xcf, 0x29, 0x26, 0x82, 0x67, 0x05, 0xc1, 0x82, 0x97, 0x85, 0x48, 0xd0,
	0x34, 0x43, 0x34, 0x65, 0xe2, 0x1b, 0x11, 0x53, 0x2c, 0xa6, 0x1c, 0xdb, 0xac, 0xdc, 0x1f, 0xc6,
	0x9f, 0x1e, 0x98, 0x3c, 0xaa, 0x15, 0xfc, 0xf9, 0x7e, 0xf1, 0xf9, 0x89, 0xc3, 0x67, 0xb6, 0x9d,
	0xcc, 0x7b, 0xc0, 0xfb, 0xfd, 0x56, 0x2d, 0xeb, 0x75, 0x0b, 0x95, 0x6e, 0xc3, 0xb6, 0x59, 0xf7,
	0xdd, 0x1d, 0x1a, 0xdf, 0x74, 0xe6, 0xd4, 0x03, 0x5c, 0xb9, 0xcf, 0xeb, 0x60, 0x98, 0x22, 0xf4,
	0x36, 0x18, 0xa7, 0x0e, 0x86, 0xa4, 0x81, 0x4e, 0x5a, 0x35, 0x8f, 0x60, 0x1f, 0x69, 0x3e, 0x0e,
	0x86, 0x0a, 0x49, 0x53, 0x1d, 0x0d, 0xd5, 0x3c, 0xaa, 0x9c, 0x61, 0xf1, 0xbb, 0x0f, 0xbe, 0xf8,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x0a, 0x9b, 0xd3, 0xf8, 0x01, 0x00, 0x00,
}