// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greetpb/greet.proto

package greetpb

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

// define your messages
// define a Greeting
type Greeting struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd67c47c0cf51822, []int{0}
}

func (m *Greeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greeting.Unmarshal(m, b)
}
func (m *Greeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greeting.Marshal(b, m, deterministic)
}
func (m *Greeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting.Merge(m, src)
}
func (m *Greeting) XXX_Size() int {
	return xxx_messageInfo_Greeting.Size(m)
}
func (m *Greeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting proto.InternalMessageInfo

func (m *Greeting) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Greeting) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type GreetRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetRequest) Reset()         { *m = GreetRequest{} }
func (m *GreetRequest) String() string { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()    {}
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd67c47c0cf51822, []int{1}
}

func (m *GreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetRequest.Unmarshal(m, b)
}
func (m *GreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetRequest.Marshal(b, m, deterministic)
}
func (m *GreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRequest.Merge(m, src)
}
func (m *GreetRequest) XXX_Size() int {
	return xxx_messageInfo_GreetRequest.Size(m)
}
func (m *GreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRequest proto.InternalMessageInfo

func (m *GreetRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetResponse) Reset()         { *m = GreetResponse{} }
func (m *GreetResponse) String() string { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()    {}
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd67c47c0cf51822, []int{2}
}

func (m *GreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetResponse.Unmarshal(m, b)
}
func (m *GreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetResponse.Marshal(b, m, deterministic)
}
func (m *GreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetResponse.Merge(m, src)
}
func (m *GreetResponse) XXX_Size() int {
	return xxx_messageInfo_GreetResponse.Size(m)
}
func (m *GreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetResponse proto.InternalMessageInfo

func (m *GreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type Sum struct {
	Number1              int32    `protobuf:"varint,1,opt,name=number1,proto3" json:"number1,omitempty"`
	Number2              int32    `protobuf:"varint,2,opt,name=number2,proto3" json:"number2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sum) Reset()         { *m = Sum{} }
func (m *Sum) String() string { return proto.CompactTextString(m) }
func (*Sum) ProtoMessage()    {}
func (*Sum) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd67c47c0cf51822, []int{3}
}

func (m *Sum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sum.Unmarshal(m, b)
}
func (m *Sum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sum.Marshal(b, m, deterministic)
}
func (m *Sum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sum.Merge(m, src)
}
func (m *Sum) XXX_Size() int {
	return xxx_messageInfo_Sum.Size(m)
}
func (m *Sum) XXX_DiscardUnknown() {
	xxx_messageInfo_Sum.DiscardUnknown(m)
}

var xxx_messageInfo_Sum proto.InternalMessageInfo

func (m *Sum) GetNumber1() int32 {
	if m != nil {
		return m.Number1
	}
	return 0
}

func (m *Sum) GetNumber2() int32 {
	if m != nil {
		return m.Number2
	}
	return 0
}

type GreetManyTimesRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetManyTimesRequest) Reset()         { *m = GreetManyTimesRequest{} }
func (m *GreetManyTimesRequest) String() string { return proto.CompactTextString(m) }
func (*GreetManyTimesRequest) ProtoMessage()    {}
func (*GreetManyTimesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd67c47c0cf51822, []int{4}
}

func (m *GreetManyTimesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetManyTimesRequest.Unmarshal(m, b)
}
func (m *GreetManyTimesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetManyTimesRequest.Marshal(b, m, deterministic)
}
func (m *GreetManyTimesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetManyTimesRequest.Merge(m, src)
}
func (m *GreetManyTimesRequest) XXX_Size() int {
	return xxx_messageInfo_GreetManyTimesRequest.Size(m)
}
func (m *GreetManyTimesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetManyTimesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetManyTimesRequest proto.InternalMessageInfo

func (m *GreetManyTimesRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetManyTimesResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetManyTimesResponse) Reset()         { *m = GreetManyTimesResponse{} }
func (m *GreetManyTimesResponse) String() string { return proto.CompactTextString(m) }
func (*GreetManyTimesResponse) ProtoMessage()    {}
func (*GreetManyTimesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd67c47c0cf51822, []int{5}
}

func (m *GreetManyTimesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetManyTimesResponse.Unmarshal(m, b)
}
func (m *GreetManyTimesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetManyTimesResponse.Marshal(b, m, deterministic)
}
func (m *GreetManyTimesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetManyTimesResponse.Merge(m, src)
}
func (m *GreetManyTimesResponse) XXX_Size() int {
	return xxx_messageInfo_GreetManyTimesResponse.Size(m)
}
func (m *GreetManyTimesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetManyTimesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetManyTimesResponse proto.InternalMessageInfo

func (m *GreetManyTimesResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeting)(nil), "greet.Greeting")
	proto.RegisterType((*GreetRequest)(nil), "greet.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "greet.GreetResponse")
	proto.RegisterType((*Sum)(nil), "greet.Sum")
	proto.RegisterType((*GreetManyTimesRequest)(nil), "greet.GreetManyTimesRequest")
	proto.RegisterType((*GreetManyTimesResponse)(nil), "greet.GreetManyTimesResponse")
}

func init() {
	proto.RegisterFile("greetpb/greet.proto", fileDescriptor_cd67c47c0cf51822)
}

var fileDescriptor_cd67c47c0cf51822 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x2f, 0x4a, 0x4d,
	0x2d, 0x29, 0x48, 0xd2, 0x07, 0xd3, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x8e,
	0x92, 0x1b, 0x17, 0x87, 0x3b, 0x88, 0x91, 0x99, 0x97, 0x2e, 0x24, 0xcb, 0xc5, 0x95, 0x96, 0x59,
	0x54, 0x5c, 0x12, 0x9f, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0xc4, 0x09,
	0x16, 0xf1, 0x4b, 0xcc, 0x4d, 0x15, 0x92, 0xe6, 0xe2, 0xcc, 0x49, 0x84, 0xc9, 0x32, 0x81, 0x65,
	0x39, 0x40, 0x02, 0x20, 0x49, 0x25, 0x6b, 0x2e, 0x1e, 0xb0, 0x39, 0x41, 0xa9, 0x85, 0xa5, 0xa9,
	0xc5, 0x25, 0x42, 0xda, 0x5c, 0x1c, 0xe9, 0x50, 0x73, 0xc1, 0x26, 0x71, 0x1b, 0xf1, 0xeb, 0x41,
	0xac, 0x87, 0x59, 0x17, 0x04, 0x57, 0xa0, 0xa4, 0xce, 0xc5, 0x0b, 0xd5, 0x5c, 0x5c, 0x90, 0x9f,
	0x57, 0x9c, 0x2a, 0x24, 0xc6, 0xc5, 0x56, 0x94, 0x5a, 0x5c, 0x9a, 0x53, 0x02, 0x75, 0x05, 0x94,
	0xa7, 0x64, 0xc9, 0xc5, 0x1c, 0x5c, 0x9a, 0x2b, 0x24, 0xc1, 0xc5, 0x9e, 0x57, 0x9a, 0x9b, 0x94,
	0x5a, 0x64, 0x08, 0x96, 0x67, 0x0d, 0x82, 0x71, 0x11, 0x32, 0x46, 0x60, 0x17, 0xc2, 0x65, 0x8c,
	0x94, 0x5c, 0xb8, 0x44, 0xc1, 0x76, 0xf8, 0x26, 0xe6, 0x55, 0x86, 0x64, 0xe6, 0xa6, 0x16, 0x93,
	0xe5, 0x52, 0x03, 0x2e, 0x31, 0x74, 0x53, 0xf0, 0x3b, 0xd9, 0x68, 0x3a, 0x23, 0x34, 0x64, 0x82,
	0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0x4c, 0xb8, 0x58, 0xc1, 0x7c, 0x21, 0x61, 0x64, 0x6b,
	0xa0, 0xae, 0x91, 0x12, 0x41, 0x15, 0x84, 0x18, 0xae, 0xc4, 0x20, 0x14, 0xc8, 0xc5, 0x87, 0x6a,
	0xb1, 0x90, 0x0c, 0xb2, 0x4a, 0x74, 0x5f, 0x49, 0xc9, 0xe2, 0x90, 0x85, 0x19, 0x68, 0xc0, 0xe8,
	0xc4, 0x19, 0xc5, 0x0e, 0x4d, 0x18, 0x49, 0x6c, 0xe0, 0x34, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0xf9, 0xa6, 0x6a, 0xaa, 0x2a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	// Unary
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	// GreetManyTimes server stream
	GreetManyTimes(ctx context.Context, in *GreetManyTimesRequest, opts ...grpc.CallOption) (GreetService_GreetManyTimesClient, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetService/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) GreetManyTimes(ctx context.Context, in *GreetManyTimesRequest, opts ...grpc.CallOption) (GreetService_GreetManyTimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[0], "/greet.GreetService/GreetManyTimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetManyTimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_GreetManyTimesClient interface {
	Recv() (*GreetManyTimesResponse, error)
	grpc.ClientStream
}

type greetServiceGreetManyTimesClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetManyTimesClient) Recv() (*GreetManyTimesResponse, error) {
	m := new(GreetManyTimesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	// Unary
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
	// GreetManyTimes server stream
	GreetManyTimes(*GreetManyTimesRequest, GreetService_GreetManyTimesServer) error
}

// UnimplementedGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (*UnimplementedGreetServiceServer) Greet(ctx context.Context, req *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (*UnimplementedGreetServiceServer) GreetManyTimes(req *GreetManyTimesRequest, srv GreetService_GreetManyTimesServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetManyTimes not implemented")
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_GreetManyTimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetManyTimesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).GreetManyTimes(m, &greetServiceGreetManyTimesServer{stream})
}

type GreetService_GreetManyTimesServer interface {
	Send(*GreetManyTimesResponse) error
	grpc.ServerStream
}

type greetServiceGreetManyTimesServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetManyTimesServer) Send(m *GreetManyTimesResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetService_Greet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetManyTimes",
			Handler:       _GreetService_GreetManyTimes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "greetpb/greet.proto",
}
