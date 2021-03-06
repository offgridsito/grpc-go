// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calcpb/calc.proto

package calcpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type SumRequest struct {
	FirstAddend          int32    `protobuf:"varint,1,opt,name=first_addend,json=firstAddend,proto3" json:"first_addend,omitempty"`
	SecondAddend         int32    `protobuf:"varint,2,opt,name=second_addend,json=secondAddend,proto3" json:"second_addend,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{0}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetFirstAddend() int32 {
	if m != nil {
		return m.FirstAddend
	}
	return 0
}

func (m *SumRequest) GetSecondAddend() int32 {
	if m != nil {
		return m.SecondAddend
	}
	return 0
}

type SumResponse struct {
	SumResult            int32    `protobuf:"varint,1,opt,name=sum_result,json=sumResult,proto3" json:"sum_result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{1}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetSumResult() int32 {
	if m != nil {
		return m.SumResult
	}
	return 0
}

type PrimeNumberDecompRequest struct {
	FirstNum             int64    `protobuf:"varint,1,opt,name=first_num,json=firstNum,proto3" json:"first_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberDecompRequest) Reset()         { *m = PrimeNumberDecompRequest{} }
func (m *PrimeNumberDecompRequest) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDecompRequest) ProtoMessage()    {}
func (*PrimeNumberDecompRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{2}
}

func (m *PrimeNumberDecompRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberDecompRequest.Unmarshal(m, b)
}
func (m *PrimeNumberDecompRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberDecompRequest.Marshal(b, m, deterministic)
}
func (m *PrimeNumberDecompRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberDecompRequest.Merge(m, src)
}
func (m *PrimeNumberDecompRequest) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberDecompRequest.Size(m)
}
func (m *PrimeNumberDecompRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberDecompRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberDecompRequest proto.InternalMessageInfo

func (m *PrimeNumberDecompRequest) GetFirstNum() int64 {
	if m != nil {
		return m.FirstNum
	}
	return 0
}

type PrimeNumberDecompResponse struct {
	PrimeFactor          int64    `protobuf:"varint,1,opt,name=prime_factor,json=primeFactor,proto3" json:"prime_factor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberDecompResponse) Reset()         { *m = PrimeNumberDecompResponse{} }
func (m *PrimeNumberDecompResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDecompResponse) ProtoMessage()    {}
func (*PrimeNumberDecompResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{3}
}

func (m *PrimeNumberDecompResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberDecompResponse.Unmarshal(m, b)
}
func (m *PrimeNumberDecompResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberDecompResponse.Marshal(b, m, deterministic)
}
func (m *PrimeNumberDecompResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberDecompResponse.Merge(m, src)
}
func (m *PrimeNumberDecompResponse) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberDecompResponse.Size(m)
}
func (m *PrimeNumberDecompResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberDecompResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberDecompResponse proto.InternalMessageInfo

func (m *PrimeNumberDecompResponse) GetPrimeFactor() int64 {
	if m != nil {
		return m.PrimeFactor
	}
	return 0
}

type ComputeAverageRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAverageRequest) Reset()         { *m = ComputeAverageRequest{} }
func (m *ComputeAverageRequest) String() string { return proto.CompactTextString(m) }
func (*ComputeAverageRequest) ProtoMessage()    {}
func (*ComputeAverageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{4}
}

func (m *ComputeAverageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAverageRequest.Unmarshal(m, b)
}
func (m *ComputeAverageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAverageRequest.Marshal(b, m, deterministic)
}
func (m *ComputeAverageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAverageRequest.Merge(m, src)
}
func (m *ComputeAverageRequest) XXX_Size() int {
	return xxx_messageInfo_ComputeAverageRequest.Size(m)
}
func (m *ComputeAverageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAverageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAverageRequest proto.InternalMessageInfo

func (m *ComputeAverageRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type ComputeAverageResponse struct {
	Average              float64  `protobuf:"fixed64,1,opt,name=average,proto3" json:"average,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAverageResponse) Reset()         { *m = ComputeAverageResponse{} }
func (m *ComputeAverageResponse) String() string { return proto.CompactTextString(m) }
func (*ComputeAverageResponse) ProtoMessage()    {}
func (*ComputeAverageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{5}
}

func (m *ComputeAverageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAverageResponse.Unmarshal(m, b)
}
func (m *ComputeAverageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAverageResponse.Marshal(b, m, deterministic)
}
func (m *ComputeAverageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAverageResponse.Merge(m, src)
}
func (m *ComputeAverageResponse) XXX_Size() int {
	return xxx_messageInfo_ComputeAverageResponse.Size(m)
}
func (m *ComputeAverageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAverageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAverageResponse proto.InternalMessageInfo

func (m *ComputeAverageResponse) GetAverage() float64 {
	if m != nil {
		return m.Average
	}
	return 0
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "calc.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calc.SumResponse")
	proto.RegisterType((*PrimeNumberDecompRequest)(nil), "calc.PrimeNumberDecompRequest")
	proto.RegisterType((*PrimeNumberDecompResponse)(nil), "calc.PrimeNumberDecompResponse")
	proto.RegisterType((*ComputeAverageRequest)(nil), "calc.ComputeAverageRequest")
	proto.RegisterType((*ComputeAverageResponse)(nil), "calc.ComputeAverageResponse")
}

func init() { proto.RegisterFile("calculator/calcpb/calc.proto", fileDescriptor_cfc74e58bc0fa04b) }

var fileDescriptor_cfc74e58bc0fa04b = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4b, 0x4f, 0x32, 0x41,
	0x10, 0xfc, 0xf6, 0x43, 0x11, 0x1a, 0x34, 0x32, 0x89, 0x04, 0x01, 0x1f, 0x8c, 0x17, 0x0e, 0x04,
	0x0c, 0x1e, 0xbc, 0x99, 0x20, 0xc6, 0x23, 0x9a, 0xc5, 0x83, 0xf1, 0x42, 0x96, 0xdd, 0xc6, 0x90,
	0xec, 0xec, 0xac, 0xf3, 0xe0, 0x57, 0xfb, 0x23, 0x0c, 0x3d, 0xbb, 0x3e, 0xe1, 0x34, 0x33, 0xd5,
	0x55, 0xd5, 0x3d, 0xdd, 0x0d, 0xed, 0x30, 0x88, 0x43, 0x1b, 0x07, 0x46, 0xaa, 0xc1, 0xfa, 0x9a,
	0xce, 0xe9, 0xe8, 0xa7, 0x4a, 0x1a, 0xc9, 0x76, 0xd6, 0x77, 0xfe, 0x04, 0x30, 0xb5, 0xc2, 0xc7,
	0x37, 0x8b, 0xda, 0xb0, 0x0e, 0x54, 0x17, 0x4b, 0xa5, 0xcd, 0x2c, 0x88, 0x22, 0x4c, 0xa2, 0x86,
	0x77, 0xee, 0x75, 0x77, 0xfd, 0x0a, 0x61, 0x23, 0x82, 0xd8, 0x05, 0xec, 0x6b, 0x0c, 0x65, 0x12,
	0xe5, 0x9c, 0xff, 0xc4, 0xa9, 0x3a, 0xd0, 0x91, 0x78, 0x0f, 0x2a, 0xe4, 0xaa, 0x53, 0x99, 0x68,
	0x64, 0x27, 0x00, 0xda, 0x8a, 0x99, 0x42, 0x6d, 0x63, 0x93, 0x99, 0x96, 0x35, 0x11, 0x6c, 0x6c,
	0xf8, 0x35, 0x34, 0x1e, 0xd5, 0x52, 0xe0, 0xc4, 0x8a, 0x39, 0xaa, 0x3b, 0x0c, 0xa5, 0x48, 0xf3,
	0x8a, 0x5a, 0x50, 0x76, 0x15, 0x25, 0x56, 0x90, 0xb2, 0xe0, 0x97, 0x08, 0x98, 0x58, 0xc1, 0x6f,
	0xe0, 0x78, 0x83, 0x30, 0x4b, 0xda, 0x81, 0x6a, 0xba, 0x0e, 0xce, 0x16, 0x41, 0x68, 0xa4, 0xca,
	0xc4, 0x15, 0xc2, 0xee, 0x09, 0xe2, 0x03, 0x38, 0x1a, 0x4b, 0x91, 0x5a, 0x83, 0xa3, 0x15, 0xaa,
	0xe0, 0x15, 0xf3, 0xac, 0x75, 0x28, 0x26, 0xe4, 0x99, 0x15, 0x9b, 0xbd, 0xf8, 0x10, 0xea, 0xbf,
	0x05, 0x59, 0xb6, 0x06, 0xec, 0x05, 0x0e, 0x22, 0x89, 0xe7, 0xe7, 0xcf, 0xe1, 0xbb, 0x07, 0xb5,
	0xf1, 0xe7, 0x20, 0xa6, 0xa8, 0x56, 0xcb, 0x10, 0x59, 0x0f, 0x0a, 0x53, 0x2b, 0xd8, 0x61, 0x9f,
	0x26, 0xf2, 0x35, 0x82, 0x66, 0xed, 0x1b, 0xe2, 0xbc, 0xf9, 0x3f, 0xf6, 0x0c, 0xb5, 0x3f, 0x1f,
	0x65, 0xa7, 0x8e, 0xb9, 0xad, 0x75, 0xcd, 0xb3, 0xad, 0xf1, 0xdc, 0xf7, 0xd2, 0x63, 0x0f, 0x70,
	0xf0, 0xf3, 0x47, 0xac, 0xe5, 0x64, 0x1b, 0x1b, 0xd3, 0x6c, 0x6f, 0x0e, 0xe6, 0x86, 0x5d, 0xef,
	0xb6, 0xf4, 0x52, 0x74, 0xbb, 0x36, 0x2f, 0xd2, 0x9e, 0x5d, 0x7d, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x6d, 0xf3, 0xdb, 0xce, 0x87, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	PrimeNumberDecomp(ctx context.Context, in *PrimeNumberDecompRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompClient, error)
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calc.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) PrimeNumberDecomp(ctx context.Context, in *PrimeNumberDecompRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calc.CalculatorService/PrimeNumberDecomp", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeNumberDecompClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeNumberDecompClient interface {
	Recv() (*PrimeNumberDecompResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeNumberDecompClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeNumberDecompClient) Recv() (*PrimeNumberDecompResponse, error) {
	m := new(PrimeNumberDecompResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[1], "/calc.CalculatorService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceComputeAverageClient{stream}
	return x, nil
}

type CalculatorService_ComputeAverageClient interface {
	Send(*ComputeAverageRequest) error
	CloseAndRecv() (*ComputeAverageResponse, error)
	grpc.ClientStream
}

type calculatorServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceComputeAverageClient) Send(m *ComputeAverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageClient) CloseAndRecv() (*ComputeAverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputeAverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	PrimeNumberDecomp(*PrimeNumberDecompRequest, CalculatorService_PrimeNumberDecompServer) error
	ComputeAverage(CalculatorService_ComputeAverageServer) error
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_PrimeNumberDecomp_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberDecompRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeNumberDecomp(m, &calculatorServicePrimeNumberDecompServer{stream})
}

type CalculatorService_PrimeNumberDecompServer interface {
	Send(*PrimeNumberDecompResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeNumberDecompServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeNumberDecompServer) Send(m *PrimeNumberDecompResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).ComputeAverage(&calculatorServiceComputeAverageServer{stream})
}

type CalculatorService_ComputeAverageServer interface {
	SendAndClose(*ComputeAverageResponse) error
	Recv() (*ComputeAverageRequest, error)
	grpc.ServerStream
}

type calculatorServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceComputeAverageServer) SendAndClose(m *ComputeAverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageServer) Recv() (*ComputeAverageRequest, error) {
	m := new(ComputeAverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calc.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumberDecomp",
			Handler:       _CalculatorService_PrimeNumberDecomp_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _CalculatorService_ComputeAverage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calcpb/calc.proto",
}
