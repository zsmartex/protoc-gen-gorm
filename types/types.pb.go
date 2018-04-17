// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types/types.proto

/*
Package gormable_types is a generated protocol buffer package.

It is generated from these files:
	types/types.proto

It has these top-level messages:
	UUIDValue
*/
package gormable_types

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

type UUIDValue struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *UUIDValue) Reset()                    { *m = UUIDValue{} }
func (m *UUIDValue) String() string            { return proto.CompactTextString(m) }
func (*UUIDValue) ProtoMessage()               {}
func (*UUIDValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UUIDValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*UUIDValue)(nil), "gormable_types.UUIDValue")
}

func init() { proto.RegisterFile("types/types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 86 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x7c, 0xe9, 0xf9, 0x45, 0xb9,
	0x89, 0x49, 0x39, 0xa9, 0xf1, 0x60, 0x51, 0x25, 0x45, 0x2e, 0xce, 0xd0, 0x50, 0x4f, 0x97, 0xb0,
	0xc4, 0x9c, 0xd2, 0x54, 0x21, 0x11, 0x2e, 0xd6, 0x32, 0x10, 0x43, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x33, 0x08, 0xc2, 0x49, 0x62, 0x03, 0xeb, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xed, 0xd1,
	0x11, 0xd8, 0x4e, 0x00, 0x00, 0x00,
}
