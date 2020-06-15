// Code generated by protoc-gen-go. DO NOT EDIT.
// source: changer.proto

package changer

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ChangeRequest struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeRequest) Reset()         { *m = ChangeRequest{} }
func (m *ChangeRequest) String() string { return proto.CompactTextString(m) }
func (*ChangeRequest) ProtoMessage()    {}
func (*ChangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ba7d01e1b27820c, []int{0}
}

func (m *ChangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeRequest.Unmarshal(m, b)
}
func (m *ChangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeRequest.Marshal(b, m, deterministic)
}
func (m *ChangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeRequest.Merge(m, src)
}
func (m *ChangeRequest) XXX_Size() int {
	return xxx_messageInfo_ChangeRequest.Size(m)
}
func (m *ChangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeRequest proto.InternalMessageInfo

func (m *ChangeRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type ChangeReply struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeReply) Reset()         { *m = ChangeReply{} }
func (m *ChangeReply) String() string { return proto.CompactTextString(m) }
func (*ChangeReply) ProtoMessage()    {}
func (*ChangeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ba7d01e1b27820c, []int{1}
}

func (m *ChangeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeReply.Unmarshal(m, b)
}
func (m *ChangeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeReply.Marshal(b, m, deterministic)
}
func (m *ChangeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeReply.Merge(m, src)
}
func (m *ChangeReply) XXX_Size() int {
	return xxx_messageInfo_ChangeReply.Size(m)
}
func (m *ChangeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeReply.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeReply proto.InternalMessageInfo

func (m *ChangeReply) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func init() {
	proto.RegisterType((*ChangeRequest)(nil), "ChangeRequest")
	proto.RegisterType((*ChangeReply)(nil), "ChangeReply")
}

func init() {
	proto.RegisterFile("changer.proto", fileDescriptor_3ba7d01e1b27820c)
}

var fileDescriptor_3ba7d01e1b27820c = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0x48, 0xcc,
	0x4b, 0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe7, 0xe2, 0x75, 0x06, 0x0b,
	0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xf1, 0x71, 0x31, 0x65, 0x16, 0x48, 0x30, 0x2a,
	0x30, 0x6a, 0x70, 0x06, 0x01, 0x59, 0x4a, 0xaa, 0x5c, 0xdc, 0x30, 0x05, 0x05, 0x39, 0x95, 0x42,
	0x62, 0x5c, 0x6c, 0x45, 0xa9, 0xc5, 0xa5, 0x39, 0x25, 0x60, 0x25, 0x1c, 0x41, 0x50, 0x9e, 0x91,
	0x31, 0x17, 0x3b, 0x44, 0x59, 0x91, 0x90, 0x06, 0x17, 0x1b, 0x84, 0x29, 0xc4, 0xa7, 0x87, 0x62,
	0xb6, 0x14, 0x8f, 0x1e, 0x92, 0x51, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x37, 0x18, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x3d, 0x48, 0x6e, 0x95, 0x94, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChangerClient is the client API for Changer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChangerClient interface {
	// Change changes the DNS records
	Change(ctx context.Context, in *ChangeRequest, opts ...grpc.CallOption) (*ChangeReply, error)
}

type changerClient struct {
	cc grpc.ClientConnInterface
}

func NewChangerClient(cc grpc.ClientConnInterface) ChangerClient {
	return &changerClient{cc}
}

func (c *changerClient) Change(ctx context.Context, in *ChangeRequest, opts ...grpc.CallOption) (*ChangeReply, error) {
	out := new(ChangeReply)
	err := c.cc.Invoke(ctx, "/Changer/Change", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChangerServer is the server API for Changer service.
type ChangerServer interface {
	// Change changes the DNS records
	Change(context.Context, *ChangeRequest) (*ChangeReply, error)
}

// UnimplementedChangerServer can be embedded to have forward compatible implementations.
type UnimplementedChangerServer struct {
}

func (*UnimplementedChangerServer) Change(ctx context.Context, req *ChangeRequest) (*ChangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Change not implemented")
}

func RegisterChangerServer(s *grpc.Server, srv ChangerServer) {
	s.RegisterService(&_Changer_serviceDesc, srv)
}

func _Changer_Change_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChangerServer).Change(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Changer/Change",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChangerServer).Change(ctx, req.(*ChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Changer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Changer",
	HandlerType: (*ChangerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Change",
			Handler:    _Changer_Change_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "changer.proto",
}