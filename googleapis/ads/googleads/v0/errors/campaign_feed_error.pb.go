// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/campaign_feed_error.proto

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

// Enum describing possible campaign feed errors.
type CampaignFeedErrorEnum_CampaignFeedError int32

const (
	// Enum unspecified.
	CampaignFeedErrorEnum_UNSPECIFIED CampaignFeedErrorEnum_CampaignFeedError = 0
	// The received error code is not known in this version.
	CampaignFeedErrorEnum_UNKNOWN CampaignFeedErrorEnum_CampaignFeedError = 1
	// An active feed already exists for this campaign and placeholder type.
	CampaignFeedErrorEnum_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE CampaignFeedErrorEnum_CampaignFeedError = 2
	// The specified feed is removed.
	CampaignFeedErrorEnum_CANNOT_CREATE_FOR_REMOVED_FEED CampaignFeedErrorEnum_CampaignFeedError = 4
	// The CampaignFeed already exists. UPDATE should be used to modify the
	// existing CampaignFeed.
	CampaignFeedErrorEnum_CANNOT_CREATE_ALREADY_EXISTING_CAMPAIGN_FEED CampaignFeedErrorEnum_CampaignFeedError = 5
	// Cannot update removed campaign feed.
	CampaignFeedErrorEnum_CANNOT_MODIFY_REMOVED_CAMPAIGN_FEED CampaignFeedErrorEnum_CampaignFeedError = 6
	// Invalid placeholder type.
	CampaignFeedErrorEnum_INVALID_PLACEHOLDER_TYPE CampaignFeedErrorEnum_CampaignFeedError = 7
	// Feed mapping for this placeholder type does not exist.
	CampaignFeedErrorEnum_MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE CampaignFeedErrorEnum_CampaignFeedError = 8
)

var CampaignFeedErrorEnum_CampaignFeedError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE",
	4: "CANNOT_CREATE_FOR_REMOVED_FEED",
	5: "CANNOT_CREATE_ALREADY_EXISTING_CAMPAIGN_FEED",
	6: "CANNOT_MODIFY_REMOVED_CAMPAIGN_FEED",
	7: "INVALID_PLACEHOLDER_TYPE",
	8: "MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE",
}
var CampaignFeedErrorEnum_CampaignFeedError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE":     2,
	"CANNOT_CREATE_FOR_REMOVED_FEED":               4,
	"CANNOT_CREATE_ALREADY_EXISTING_CAMPAIGN_FEED": 5,
	"CANNOT_MODIFY_REMOVED_CAMPAIGN_FEED":          6,
	"INVALID_PLACEHOLDER_TYPE":                     7,
	"MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE":     8,
}

func (x CampaignFeedErrorEnum_CampaignFeedError) String() string {
	return proto.EnumName(CampaignFeedErrorEnum_CampaignFeedError_name, int32(x))
}
func (CampaignFeedErrorEnum_CampaignFeedError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_campaign_feed_error_8ad65cf273383026, []int{0, 0}
}

// Container for enum describing possible campaign feed errors.
type CampaignFeedErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignFeedErrorEnum) Reset()         { *m = CampaignFeedErrorEnum{} }
func (m *CampaignFeedErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignFeedErrorEnum) ProtoMessage()    {}
func (*CampaignFeedErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_campaign_feed_error_8ad65cf273383026, []int{0}
}
func (m *CampaignFeedErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignFeedErrorEnum.Unmarshal(m, b)
}
func (m *CampaignFeedErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignFeedErrorEnum.Marshal(b, m, deterministic)
}
func (dst *CampaignFeedErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignFeedErrorEnum.Merge(dst, src)
}
func (m *CampaignFeedErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignFeedErrorEnum.Size(m)
}
func (m *CampaignFeedErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignFeedErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignFeedErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CampaignFeedErrorEnum)(nil), "google.ads.googleads.v0.errors.CampaignFeedErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.CampaignFeedErrorEnum_CampaignFeedError", CampaignFeedErrorEnum_CampaignFeedError_name, CampaignFeedErrorEnum_CampaignFeedError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/campaign_feed_error.proto", fileDescriptor_campaign_feed_error_8ad65cf273383026)
}

var fileDescriptor_campaign_feed_error_8ad65cf273383026 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0xae, 0x93, 0x40,
	0x14, 0x86, 0x2d, 0xea, 0xbd, 0x66, 0xee, 0x42, 0x9c, 0x44, 0xe3, 0xc2, 0x74, 0x81, 0x0b, 0x5d,
	0x34, 0x03, 0x89, 0x1b, 0x13, 0x57, 0x53, 0x66, 0xc0, 0x89, 0x30, 0x4c, 0x80, 0xa2, 0x35, 0x24,
	0x13, 0x2c, 0x48, 0x9a, 0xb4, 0x4c, 0x03, 0xda, 0x07, 0x72, 0xd9, 0xe7, 0x70, 0x65, 0x7c, 0x28,
	0x33, 0x4c, 0xdb, 0xa4, 0xa9, 0xba, 0xe2, 0xcf, 0xe1, 0xfb, 0xff, 0xc3, 0x39, 0x07, 0xf0, 0xb6,
	0x55, 0xaa, 0xdd, 0x34, 0x6e, 0x55, 0x0f, 0xae, 0x91, 0x5a, 0xed, 0x3d, 0xb7, 0xe9, 0x7b, 0xd5,
	0x0f, 0xee, 0xaa, 0xda, 0xee, 0xaa, 0x75, 0xdb, 0xc9, 0xaf, 0x4d, 0x53, 0xcb, 0xb1, 0x88, 0x76,
	0xbd, 0xfa, 0xa6, 0xe0, 0xd4, 0xe0, 0xa8, 0xaa, 0x07, 0x74, 0x76, 0xa2, 0xbd, 0x87, 0x8c, 0xd3,
	0xf9, 0x69, 0x81, 0xa7, 0xfe, 0xd1, 0x1d, 0x34, 0x4d, 0x4d, 0x75, 0x99, 0x76, 0xdf, 0xb7, 0xce,
	0xc1, 0x02, 0x4f, 0xae, 0xde, 0xc0, 0xc7, 0xe0, 0x6e, 0xc1, 0x33, 0x41, 0x7d, 0x16, 0x30, 0x4a,
	0xec, 0x7b, 0xf0, 0x0e, 0xdc, 0x2e, 0xf8, 0x07, 0x9e, 0x7c, 0xe4, 0xf6, 0x04, 0xce, 0xc0, 0xeb,
	0x80, 0x52, 0x22, 0x71, 0x94, 0x52, 0x4c, 0x96, 0x92, 0x7e, 0x62, 0x59, 0x9e, 0xc9, 0x20, 0x49,
	0xa5, 0x88, 0xb0, 0x4f, 0xdf, 0x27, 0x11, 0xa1, 0xa9, 0xcc, 0x97, 0x82, 0xda, 0x16, 0x74, 0xc0,
	0xd4, 0xc7, 0x9c, 0x27, 0xb9, 0xf4, 0x53, 0x8a, 0x73, 0x3a, 0x72, 0x29, 0x8d, 0x93, 0x82, 0x12,
	0xa9, 0x73, 0xec, 0x07, 0xd0, 0x03, 0xb3, 0x4b, 0xe6, 0x22, 0x9a, 0xf1, 0x50, 0xfa, 0x38, 0x16,
	0x98, 0x85, 0xdc, 0x38, 0x1e, 0xc2, 0x57, 0xe0, 0xe5, 0xd1, 0x11, 0x27, 0x84, 0x05, 0xcb, 0x73,
	0xe2, 0x25, 0x78, 0x03, 0x5f, 0x80, 0xe7, 0x8c, 0x17, 0x38, 0x62, 0xe4, 0xfa, 0xe3, 0x6e, 0xf5,
	0x28, 0x31, 0xcb, 0x32, 0xdd, 0x41, 0xf3, 0x31, 0x16, 0x62, 0xd4, 0x7f, 0x1b, 0xe5, 0xd1, 0xfc,
	0xf7, 0x04, 0x38, 0x2b, 0xb5, 0x45, 0xff, 0xdf, 0xf6, 0xfc, 0xd9, 0xd5, 0x42, 0x85, 0xbe, 0x92,
	0x98, 0x7c, 0x26, 0x47, 0x67, 0xab, 0x36, 0x55, 0xd7, 0x22, 0xd5, 0xb7, 0x6e, 0xdb, 0x74, 0xe3,
	0x0d, 0x4f, 0x17, 0xdf, 0xad, 0x87, 0x7f, 0xfd, 0x00, 0xef, 0xcc, 0xe3, 0x87, 0x75, 0x3f, 0xc4,
	0xf8, 0x60, 0x4d, 0x43, 0x13, 0x86, 0xeb, 0x01, 0x19, 0xa9, 0x55, 0xe1, 0xa1, 0xb1, 0xe5, 0xf0,
	0xeb, 0x04, 0x94, 0xb8, 0x1e, 0xca, 0x33, 0x50, 0x16, 0x5e, 0x69, 0x80, 0x2f, 0x37, 0x63, 0xe3,
	0x37, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x6a, 0xa0, 0x54, 0x78, 0x02, 0x00, 0x00,
}
