// Code generated by protoc-gen-go.
// source: Oauth.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UserLogin struct {
	UserId       string `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
	Accesstocken string `protobuf:"bytes,2,opt,name=accesstocken" json:"accesstocken,omitempty"`
}

func (m *UserLogin) Reset()                    { *m = UserLogin{} }
func (m *UserLogin) String() string            { return proto.CompactTextString(m) }
func (*UserLogin) ProtoMessage()               {}
func (*UserLogin) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func init() {
	proto.RegisterType((*UserLogin)(nil), "pb.UserLogin")
}

func init() { proto.RegisterFile("Oauth.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xf6, 0x4f, 0x2c, 0x2d,
	0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x72, 0xe7, 0xe2, 0x0c,
	0x2d, 0x4e, 0x2d, 0xf2, 0xc9, 0x4f, 0xcf, 0xcc, 0x13, 0x12, 0xe3, 0x62, 0x2b, 0x05, 0x72, 0x3c,
	0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xa0, 0x3c, 0x21, 0x25, 0x2e, 0x9e, 0xc4, 0xe4,
	0xe4, 0xd4, 0xe2, 0xe2, 0x92, 0xfc, 0xe4, 0xec, 0xd4, 0x3c, 0x09, 0x26, 0xb0, 0x2c, 0x8a, 0x98,
	0x13, 0xc7, 0x2a, 0x26, 0xd6, 0x00, 0x90, 0xb1, 0x49, 0x6c, 0x60, 0xd3, 0x8d, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x37, 0x65, 0xb3, 0x3f, 0x6c, 0x00, 0x00, 0x00,
}