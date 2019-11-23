// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transaction.proto

package transaction

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

type LineItem struct {
	Accountname          string   `protobuf:"bytes,1,opt,name=accountname,proto3" json:"accountname,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Amount               int64    `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LineItem) Reset()         { *m = LineItem{} }
func (m *LineItem) String() string { return proto.CompactTextString(m) }
func (*LineItem) ProtoMessage()    {}
func (*LineItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{0}
}

func (m *LineItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LineItem.Unmarshal(m, b)
}
func (m *LineItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LineItem.Marshal(b, m, deterministic)
}
func (m *LineItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LineItem.Merge(m, src)
}
func (m *LineItem) XXX_Size() int {
	return xxx_messageInfo_LineItem.Size(m)
}
func (m *LineItem) XXX_DiscardUnknown() {
	xxx_messageInfo_LineItem.DiscardUnknown(m)
}

var xxx_messageInfo_LineItem proto.InternalMessageInfo

func (m *LineItem) GetAccountname() string {
	if m != nil {
		return m.Accountname
	}
	return ""
}

func (m *LineItem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *LineItem) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type TransactionRequest struct {
	Date                 string      `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Description          string      `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Lines                []*LineItem `protobuf:"bytes,3,rep,name=lines,proto3" json:"lines,omitempty"`
	Signature            string      `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TransactionRequest) Reset()         { *m = TransactionRequest{} }
func (m *TransactionRequest) String() string { return proto.CompactTextString(m) }
func (*TransactionRequest) ProtoMessage()    {}
func (*TransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{1}
}

func (m *TransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionRequest.Unmarshal(m, b)
}
func (m *TransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionRequest.Marshal(b, m, deterministic)
}
func (m *TransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionRequest.Merge(m, src)
}
func (m *TransactionRequest) XXX_Size() int {
	return xxx_messageInfo_TransactionRequest.Size(m)
}
func (m *TransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionRequest proto.InternalMessageInfo

func (m *TransactionRequest) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *TransactionRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TransactionRequest) GetLines() []*LineItem {
	if m != nil {
		return m.Lines
	}
	return nil
}

func (m *TransactionRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type TransactionResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{2}
}

func (m *TransactionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionResponse.Unmarshal(m, b)
}
func (m *TransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionResponse.Marshal(b, m, deterministic)
}
func (m *TransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionResponse.Merge(m, src)
}
func (m *TransactionResponse) XXX_Size() int {
	return xxx_messageInfo_TransactionResponse.Size(m)
}
func (m *TransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionResponse proto.InternalMessageInfo

func (m *TransactionResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type VersionRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{3}
}

func (m *VersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionRequest.Unmarshal(m, b)
}
func (m *VersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionRequest.Marshal(b, m, deterministic)
}
func (m *VersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRequest.Merge(m, src)
}
func (m *VersionRequest) XXX_Size() int {
	return xxx_messageInfo_VersionRequest.Size(m)
}
func (m *VersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRequest proto.InternalMessageInfo

func (m *VersionRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type VersionResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{4}
}

func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*LineItem)(nil), "transaction.LineItem")
	proto.RegisterType((*TransactionRequest)(nil), "transaction.TransactionRequest")
	proto.RegisterType((*TransactionResponse)(nil), "transaction.TransactionResponse")
	proto.RegisterType((*VersionRequest)(nil), "transaction.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "transaction.VersionResponse")
}

func init() { proto.RegisterFile("transaction.proto", fileDescriptor_2cc4e03d2c28c490) }

var fileDescriptor_2cc4e03d2c28c490 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xd1, 0x4a, 0xf3, 0x40,
	0x10, 0x85, 0xff, 0xfc, 0xa9, 0xd5, 0x4e, 0xa0, 0xe2, 0x88, 0xb2, 0xd4, 0x82, 0x21, 0x57, 0xc1,
	0x42, 0x85, 0xfa, 0x04, 0x5e, 0x2a, 0xe2, 0x45, 0x50, 0xef, 0xd7, 0x64, 0x2c, 0x01, 0xb3, 0x1b,
	0x77, 0x36, 0x8f, 0xe2, 0x73, 0xf8, 0x8a, 0x92, 0xa4, 0x6b, 0x37, 0x68, 0xd1, 0xbb, 0xcc, 0xc9,
	0x37, 0x73, 0xce, 0x21, 0x81, 0x23, 0x6b, 0xa4, 0x62, 0x99, 0xdb, 0x52, 0xab, 0x65, 0x6d, 0xb4,
	0xd5, 0x18, 0x79, 0x52, 0xf2, 0x02, 0x07, 0x77, 0xa5, 0xa2, 0x1b, 0x4b, 0x15, 0xc6, 0x10, 0xc9,
	0x3c, 0xd7, 0x8d, 0xb2, 0x4a, 0x56, 0x24, 0x82, 0x38, 0x48, 0x27, 0x99, 0x2f, 0xb5, 0x44, 0x41,
	0x9c, 0x9b, 0xb2, 0x6e, 0x97, 0xc5, 0xff, 0x9e, 0xf0, 0x24, 0x3c, 0x85, 0xb1, 0xac, 0x5a, 0x5e,
	0x84, 0x71, 0x90, 0x86, 0xd9, 0x66, 0x4a, 0xde, 0x03, 0xc0, 0x87, 0xad, 0x6f, 0x46, 0x6f, 0x0d,
	0xb1, 0x45, 0x84, 0x51, 0x21, 0xad, 0xf3, 0xea, 0x9e, 0xff, 0x60, 0xb2, 0x80, 0xbd, 0xd7, 0x52,
	0x11, 0x8b, 0x30, 0x0e, 0xd3, 0x68, 0x75, 0xb2, 0xf4, 0x4b, 0xba, 0x3a, 0x59, 0xcf, 0xe0, 0x1c,
	0x26, 0x5c, 0xae, 0x95, 0xb4, 0x8d, 0x21, 0x31, 0xea, 0x8e, 0x6d, 0x85, 0xe4, 0x12, 0x8e, 0x07,
	0xb1, 0xb8, 0xd6, 0x8a, 0x09, 0x05, 0xec, 0x57, 0xc4, 0x2c, 0xd7, 0x2e, 0x9a, 0x1b, 0x93, 0x0b,
	0x98, 0x3e, 0x91, 0x61, 0xaf, 0xc3, 0x6e, 0x76, 0x01, 0x87, 0x5f, 0xec, 0x6f, 0x87, 0x57, 0x1f,
	0x01, 0x80, 0x8b, 0xa2, 0x0d, 0xde, 0x42, 0x74, 0xaf, 0x0b, 0xda, 0xec, 0xe3, 0xd9, 0xa0, 0xe3,
	0x30, 0xc1, 0x6c, 0xfe, 0xf3, 0xcb, 0xde, 0x32, 0xf9, 0x87, 0x8f, 0x30, 0xbd, 0x2e, 0x0a, 0xaf,
	0x27, 0x9e, 0x0f, 0x36, 0xbe, 0x7f, 0x98, 0x59, 0xbc, 0x1b, 0x70, 0x67, 0x9f, 0xc7, 0xdd, 0xff,
	0x74, 0xf5, 0x19, 0x00, 0x00, 0xff, 0xff, 0x3d, 0xcf, 0xc3, 0xf5, 0x64, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransactorClient is the client API for Transactor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactorClient interface {
	NodeVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	AddTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type transactorClient struct {
	cc *grpc.ClientConn
}

func NewTransactorClient(cc *grpc.ClientConn) TransactorClient {
	return &transactorClient{cc}
}

func (c *transactorClient) NodeVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/transaction.Transactor/NodeVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactorClient) AddTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/transaction.Transactor/AddTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactorServer is the server API for Transactor service.
type TransactorServer interface {
	NodeVersion(context.Context, *VersionRequest) (*VersionResponse, error)
	AddTransaction(context.Context, *TransactionRequest) (*TransactionResponse, error)
}

// UnimplementedTransactorServer can be embedded to have forward compatible implementations.
type UnimplementedTransactorServer struct {
}

func (*UnimplementedTransactorServer) NodeVersion(ctx context.Context, req *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeVersion not implemented")
}
func (*UnimplementedTransactorServer) AddTransaction(ctx context.Context, req *TransactionRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTransaction not implemented")
}

func RegisterTransactorServer(s *grpc.Server, srv TransactorServer) {
	s.RegisterService(&_Transactor_serviceDesc, srv)
}

func _Transactor_NodeVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactorServer).NodeVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.Transactor/NodeVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactorServer).NodeVersion(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transactor_AddTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactorServer).AddTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.Transactor/AddTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactorServer).AddTransaction(ctx, req.(*TransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transactor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transaction.Transactor",
	HandlerType: (*TransactorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NodeVersion",
			Handler:    _Transactor_NodeVersion_Handler,
		},
		{
			MethodName: "AddTransaction",
			Handler:    _Transactor_AddTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transaction.proto",
}
