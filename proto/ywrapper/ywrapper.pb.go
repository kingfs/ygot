// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/openconfig/ygot/proto/ywrapper/ywrapper.proto

/*
Package ywrapper is a generated protocol buffer package.

It is generated from these files:
	github.com/openconfig/ygot/proto/ywrapper/ywrapper.proto

It has these top-level messages:
	BytesValue
	BoolValue
	Decimal64Value
	IntValue
	StringValue
	UintValue
*/
package ywrapper

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

// BytesValue is used to store a value which is a byte array, particularly
// the YANG binary type.
type BytesValue struct {
	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *BytesValue) Reset()                    { *m = BytesValue{} }
func (m *BytesValue) String() string            { return proto.CompactTextString(m) }
func (*BytesValue) ProtoMessage()               {}
func (*BytesValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BytesValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// BoolValue is used to store a value which is a boolean, particularly the
// YANG bool and empty types.
type BoolValue struct {
	Value bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *BoolValue) Reset()                    { *m = BoolValue{} }
func (m *BoolValue) String() string            { return proto.CompactTextString(m) }
func (*BoolValue) ProtoMessage()               {}
func (*BoolValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BoolValue) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

// Decimal64Value is used to store a value which is a decimal64, split into
// a digits field, and a precision field. The precision indicates the number
// of digits that occur after the decimal point in the digits field.
type Decimal64Value struct {
	Digits    int64  `protobuf:"varint,1,opt,name=digits" json:"digits,omitempty"`
	Precision uint32 `protobuf:"varint,2,opt,name=precision" json:"precision,omitempty"`
}

func (m *Decimal64Value) Reset()                    { *m = Decimal64Value{} }
func (m *Decimal64Value) String() string            { return proto.CompactTextString(m) }
func (*Decimal64Value) ProtoMessage()               {}
func (*Decimal64Value) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Decimal64Value) GetDigits() int64 {
	if m != nil {
		return m.Digits
	}
	return 0
}

func (m *Decimal64Value) GetPrecision() uint32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

// IntValue stores a value which a signed integer, particularly the YANG
// int8, int16, int32, and int64 types.
type IntValue struct {
	Value int64 `protobuf:"zigzag64,1,opt,name=value" json:"value,omitempty"`
}

func (m *IntValue) Reset()                    { *m = IntValue{} }
func (m *IntValue) String() string            { return proto.CompactTextString(m) }
func (*IntValue) ProtoMessage()               {}
func (*IntValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IntValue) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// StringValue stores a value which is a string, particularly the YANG
// string type.
type StringValue struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *StringValue) Reset()                    { *m = StringValue{} }
func (m *StringValue) String() string            { return proto.CompactTextString(m) }
func (*StringValue) ProtoMessage()               {}
func (*StringValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StringValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// UintVal is used to store a value which an unsigned integer, particularly
// the YANG uint8, uint16, uint32 and uint64 types.
type UintValue struct {
	Value uint64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *UintValue) Reset()                    { *m = UintValue{} }
func (m *UintValue) String() string            { return proto.CompactTextString(m) }
func (*UintValue) ProtoMessage()               {}
func (*UintValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UintValue) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*BytesValue)(nil), "ywrapper.BytesValue")
	proto.RegisterType((*BoolValue)(nil), "ywrapper.BoolValue")
	proto.RegisterType((*Decimal64Value)(nil), "ywrapper.Decimal64Value")
	proto.RegisterType((*IntValue)(nil), "ywrapper.IntValue")
	proto.RegisterType((*StringValue)(nil), "ywrapper.StringValue")
	proto.RegisterType((*UintValue)(nil), "ywrapper.UintValue")
}

func init() {
	proto.RegisterFile("github.com/openconfig/ygot/proto/ywrapper/ywrapper.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x2f, 0x48, 0xcd, 0x4b, 0xce, 0xcf, 0x4b, 0xcb,
	0x4c, 0xd7, 0xaf, 0x4c, 0xcf, 0x2f, 0xd1, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xaf, 0x2c, 0x2f,
	0x4a, 0x2c, 0x28, 0x48, 0x2d, 0x82, 0x33, 0xf4, 0xc0, 0xe2, 0x42, 0x1c, 0x30, 0xbe, 0x92, 0x12,
	0x17, 0x97, 0x53, 0x65, 0x49, 0x6a, 0x71, 0x58, 0x62, 0x4e, 0x69, 0xaa, 0x90, 0x08, 0x17, 0x6b,
	0x19, 0x88, 0x21, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x13, 0x04, 0xe1, 0x28, 0x29, 0x72, 0x71, 0x3a,
	0xe5, 0xe7, 0xe7, 0x60, 0x51, 0xc2, 0x01, 0x53, 0xe2, 0xc6, 0xc5, 0xe7, 0x92, 0x9a, 0x9c, 0x99,
	0x9b, 0x98, 0x63, 0x66, 0x02, 0x51, 0x27, 0xc6, 0xc5, 0x96, 0x92, 0x99, 0x9e, 0x59, 0x52, 0x0c,
	0x56, 0xc8, 0x1c, 0x04, 0xe5, 0x09, 0xc9, 0x70, 0x71, 0x16, 0x14, 0xa5, 0x26, 0x67, 0x16, 0x67,
	0xe6, 0xe7, 0x49, 0x30, 0x29, 0x30, 0x6a, 0xf0, 0x06, 0x21, 0x04, 0x94, 0x14, 0xb8, 0x38, 0x3c,
	0xf3, 0x4a, 0xb0, 0xd8, 0x24, 0x04, 0xb3, 0x49, 0x99, 0x8b, 0x3b, 0xb8, 0xa4, 0x28, 0x33, 0x2f,
	0x1d, 0x8b, 0x22, 0x4e, 0x24, 0x17, 0x87, 0x66, 0x62, 0x35, 0x87, 0x05, 0xaa, 0x24, 0x89, 0x0d,
	0x1c, 0x12, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x71, 0x13, 0x9f, 0xa4, 0x45, 0x01, 0x00,
	0x00,
}
