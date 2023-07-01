// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nemo/earn/v1beta1/strategy.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// StrategyType is the type of strategy that a vault uses to optimize yields.
type StrategyType int32

const (
	// STRATEGY_TYPE_UNSPECIFIED represents an unspecified or invalid strategy type.
	STRATEGY_TYPE_UNSPECIFIED StrategyType = 0
	// STRATEGY_TYPE_HARD represents the strategy that deposits assets in the Hard
	// module.
	STRATEGY_TYPE_HARD StrategyType = 1
	// STRATEGY_TYPE_SAVINGS represents the strategy that deposits assets in the
	// Savings module.
	STRATEGY_TYPE_SAVINGS StrategyType = 2
)

var StrategyType_name = map[int32]string{
	0: "STRATEGY_TYPE_UNSPECIFIED",
	1: "STRATEGY_TYPE_HARD",
	2: "STRATEGY_TYPE_SAVINGS",
}

var StrategyType_value = map[string]int32{
	"STRATEGY_TYPE_UNSPECIFIED": 0,
	"STRATEGY_TYPE_HARD":        1,
	"STRATEGY_TYPE_SAVINGS":     2,
}

func (x StrategyType) String() string {
	return proto.EnumName(StrategyType_name, int32(x))
}

func (StrategyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_257c4968dd48fa09, []int{0}
}

func init() {
	proto.RegisterEnum("nemo.earn.v1beta1.StrategyType", StrategyType_name, StrategyType_value)
}

func init() { proto.RegisterFile("nemo/earn/v1beta1/strategy.proto", fileDescriptor_257c4968dd48fa09) }

var fileDescriptor_257c4968dd48fa09 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc8, 0x4e, 0x2c, 0x4b,
	0xd4, 0x4f, 0x4d, 0x2c, 0xca, 0xd3, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f, 0x2e,
	0x29, 0x4a, 0x2c, 0x49, 0x4d, 0xaf, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x04, 0xa9,
	0xd0, 0x03, 0xa9, 0xd0, 0x83, 0xaa, 0x90, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xcb, 0xea, 0x83,
	0x58, 0x10, 0x85, 0x5a, 0x69, 0x5c, 0x3c, 0xc1, 0x50, 0xad, 0x21, 0x95, 0x05, 0xa9, 0x42, 0xb2,
	0x5c, 0x92, 0xc1, 0x21, 0x41, 0x8e, 0x21, 0xae, 0xee, 0x91, 0xf1, 0x21, 0x91, 0x01, 0xae, 0xf1,
	0xa1, 0x7e, 0xc1, 0x01, 0xae, 0xce, 0x9e, 0x6e, 0x9e, 0xae, 0x2e, 0x02, 0x0c, 0x42, 0x62, 0x5c,
	0x42, 0xa8, 0xd2, 0x1e, 0x8e, 0x41, 0x2e, 0x02, 0x8c, 0x42, 0x92, 0x5c, 0xa2, 0xa8, 0xe2, 0xc1,
	0x8e, 0x61, 0x9e, 0x7e, 0xee, 0xc1, 0x02, 0x4c, 0x52, 0x2c, 0x1d, 0x8b, 0xe5, 0x18, 0x9c, 0x1c,
	0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5,
	0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x2d, 0x3d, 0xb3, 0x24, 0xa3,
	0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0xe4, 0x6a, 0xdd, 0x9c, 0xc4, 0xa4, 0x62, 0x30, 0x4b,
	0xbf, 0x02, 0xe2, 0xc7, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0x83, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x81, 0x9e, 0x23, 0x1c, 0xfd, 0x00, 0x00, 0x00,
}
