// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HealthResponse struct {
	Healthy bool `protobuf:"varint,1,opt,name=healthy" json:"healthy,omitempty"`
}

func (m *HealthResponse) Reset()                    { *m = HealthResponse{} }
func (m *HealthResponse) String() string            { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()               {}
func (*HealthResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *HealthResponse) GetHealthy() bool {
	if m != nil {
		return m.Healthy
	}
	return false
}

type HealthRequest struct {
}

func (m *HealthRequest) Reset()                    { *m = HealthRequest{} }
func (m *HealthRequest) String() string            { return proto.CompactTextString(m) }
func (*HealthRequest) ProtoMessage()               {}
func (*HealthRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func init() {
	proto.RegisterType((*HealthResponse)(nil), "api.HealthResponse")
	proto.RegisterType((*HealthRequest)(nil), "api.HealthRequest")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 94 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0xd2, 0xe2,
	0xe2, 0xf3, 0x48, 0x4d, 0xcc, 0x29, 0xc9, 0x08, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15,
	0x92, 0xe0, 0x62, 0xcf, 0x00, 0x8b, 0x54, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x04, 0xc1, 0xb8,
	0x4a, 0xfc, 0x5c, 0xbc, 0x30, 0xb5, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x49, 0x6c, 0x60, 0x83, 0x8c,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x55, 0x23, 0x59, 0x44, 0x58, 0x00, 0x00, 0x00,
}