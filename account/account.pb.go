// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: account.proto

package account

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/metaverse/truss/deftree/googlethirdparty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ServiceStatus int32

const (
	ServiceStatus_FAIL ServiceStatus = 0
	ServiceStatus_OK   ServiceStatus = 1
)

var ServiceStatus_name = map[int32]string{
	0: "FAIL",
	1: "OK",
}

var ServiceStatus_value = map[string]int32{
	"FAIL": 0,
	"OK":   1,
}

func (x ServiceStatus) String() string {
	return proto.EnumName(ServiceStatus_name, int32(x))
}

func (ServiceStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}

type StatusRequest struct {
	Full bool `protobuf:"varint,1,opt,name=full,proto3" json:"full,omitempty"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}
func (m *StatusRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StatusRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequest.Merge(m, src)
}
func (m *StatusRequest) XXX_Size() int {
	return m.Size()
}
func (m *StatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequest proto.InternalMessageInfo

func (m *StatusRequest) GetFull() bool {
	if m != nil {
		return m.Full
	}
	return false
}

type StatusResponse struct {
	Status ServiceStatus `protobuf:"varint,1,opt,name=status,proto3,enum=account.ServiceStatus" json:"status,omitempty"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{1}
}
func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(m, src)
}
func (m *StatusResponse) XXX_Size() int {
	return m.Size()
}
func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

func (m *StatusResponse) GetStatus() ServiceStatus {
	if m != nil {
		return m.Status
	}
	return ServiceStatus_FAIL
}

type AuthTokenRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (m *AuthTokenRequest) Reset()         { *m = AuthTokenRequest{} }
func (m *AuthTokenRequest) String() string { return proto.CompactTextString(m) }
func (*AuthTokenRequest) ProtoMessage()    {}
func (*AuthTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{2}
}
func (m *AuthTokenRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthTokenRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthTokenRequest.Merge(m, src)
}
func (m *AuthTokenRequest) XXX_Size() int {
	return m.Size()
}
func (m *AuthTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthTokenRequest proto.InternalMessageInfo

func (m *AuthTokenRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthTokenRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthTokenResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *AuthTokenResponse) Reset()         { *m = AuthTokenResponse{} }
func (m *AuthTokenResponse) String() string { return proto.CompactTextString(m) }
func (*AuthTokenResponse) ProtoMessage()    {}
func (*AuthTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{3}
}
func (m *AuthTokenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthTokenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthTokenResponse.Merge(m, src)
}
func (m *AuthTokenResponse) XXX_Size() int {
	return m.Size()
}
func (m *AuthTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthTokenResponse proto.InternalMessageInfo

func (m *AuthTokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthTokenValidateRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *AuthTokenValidateRequest) Reset()         { *m = AuthTokenValidateRequest{} }
func (m *AuthTokenValidateRequest) String() string { return proto.CompactTextString(m) }
func (*AuthTokenValidateRequest) ProtoMessage()    {}
func (*AuthTokenValidateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{4}
}
func (m *AuthTokenValidateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthTokenValidateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthTokenValidateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthTokenValidateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthTokenValidateRequest.Merge(m, src)
}
func (m *AuthTokenValidateRequest) XXX_Size() int {
	return m.Size()
}
func (m *AuthTokenValidateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthTokenValidateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthTokenValidateRequest proto.InternalMessageInfo

func (m *AuthTokenValidateRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthTokenValidateResponse struct {
	IsValid bool `protobuf:"varint,2,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
}

func (m *AuthTokenValidateResponse) Reset()         { *m = AuthTokenValidateResponse{} }
func (m *AuthTokenValidateResponse) String() string { return proto.CompactTextString(m) }
func (*AuthTokenValidateResponse) ProtoMessage()    {}
func (*AuthTokenValidateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{5}
}
func (m *AuthTokenValidateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthTokenValidateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthTokenValidateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthTokenValidateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthTokenValidateResponse.Merge(m, src)
}
func (m *AuthTokenValidateResponse) XXX_Size() int {
	return m.Size()
}
func (m *AuthTokenValidateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthTokenValidateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthTokenValidateResponse proto.InternalMessageInfo

func (m *AuthTokenValidateResponse) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

func init() {
	proto.RegisterEnum("account.ServiceStatus", ServiceStatus_name, ServiceStatus_value)
	proto.RegisterType((*StatusRequest)(nil), "account.StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "account.StatusResponse")
	proto.RegisterType((*AuthTokenRequest)(nil), "account.AuthTokenRequest")
	proto.RegisterType((*AuthTokenResponse)(nil), "account.AuthTokenResponse")
	proto.RegisterType((*AuthTokenValidateRequest)(nil), "account.AuthTokenValidateRequest")
	proto.RegisterType((*AuthTokenValidateResponse)(nil), "account.AuthTokenValidateResponse")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x8a, 0x13, 0x41,
	0x10, 0x86, 0xd3, 0x61, 0x4d, 0x26, 0x25, 0x59, 0x63, 0xb3, 0xac, 0xc9, 0x1c, 0x46, 0xb7, 0xbd,
	0x68, 0x0e, 0x69, 0x59, 0xc1, 0xc3, 0x9e, 0x8c, 0xa0, 0xe0, 0x2a, 0x08, 0x59, 0xf1, 0xe0, 0x45,
	0x7a, 0x93, 0xda, 0x64, 0x30, 0x99, 0x8e, 0xdd, 0xd5, 0x11, 0xf1, 0xe6, 0x13, 0x08, 0xbe, 0x87,
	0xcf, 0xe1, 0x71, 0xc1, 0x8b, 0x47, 0x49, 0x7c, 0x10, 0x49, 0x4f, 0x4f, 0x18, 0x37, 0xc1, 0xdb,
	0x54, 0xd5, 0x57, 0xff, 0x5f, 0xf3, 0xcf, 0x40, 0x53, 0x0d, 0x87, 0xda, 0x65, 0xd4, 0x9b, 0x1b,
	0x4d, 0x9a, 0xd7, 0x43, 0x19, 0x3f, 0x1d, 0xa7, 0x34, 0x71, 0xe7, 0xbd, 0xa1, 0x9e, 0xc9, 0x19,
	0x92, 0x5a, 0xa0, 0xb1, 0x28, 0xc9, 0x38, 0x6b, 0xe5, 0x08, 0x2f, 0xc8, 0x20, 0xca, 0xb1, 0xd6,
	0xe3, 0x29, 0xd2, 0x24, 0x35, 0xa3, 0xb9, 0x32, 0xf4, 0x49, 0xaa, 0x2c, 0xd3, 0xa4, 0x28, 0xd5,
	0x99, 0xcd, 0xf5, 0xc4, 0x5d, 0x68, 0x9e, 0x91, 0x22, 0x67, 0x07, 0xf8, 0xc1, 0xa1, 0x25, 0xce,
	0x61, 0xef, 0xc2, 0x4d, 0xa7, 0x6d, 0x76, 0x87, 0xdd, 0x8b, 0x06, 0xfe, 0x59, 0x3c, 0x86, 0xfd,
	0x02, 0xb2, 0x73, 0x9d, 0x59, 0xe4, 0x3d, 0xa8, 0x59, 0xdf, 0xf1, 0xdc, 0xfe, 0xf1, 0x61, 0xaf,
	0x38, 0xf3, 0x0c, 0xcd, 0x22, 0x1d, 0x62, 0xe0, 0x03, 0x25, 0x4e, 0xa1, 0xd5, 0x77, 0x34, 0x79,
	0xad, 0xdf, 0x63, 0x56, 0x38, 0xc5, 0x10, 0x39, 0x8b, 0x26, 0x53, 0x33, 0xf4, 0x2a, 0x8d, 0xc1,
	0xa6, 0x5e, 0xcf, 0xe6, 0xca, 0xda, 0x8f, 0xda, 0x8c, 0xda, 0xd5, 0x7c, 0x56, 0xd4, 0xe2, 0x3e,
	0xdc, 0x2c, 0x69, 0x85, 0x83, 0x0e, 0xe0, 0x1a, 0xad, 0x1b, 0x41, 0x29, 0x2f, 0xc4, 0x03, 0x68,
	0x6f, 0xd0, 0x37, 0x6a, 0x9a, 0x8e, 0x14, 0x61, 0x61, 0xbf, 0x7b, 0xe3, 0x11, 0x74, 0x76, 0x6c,
	0x04, 0x93, 0x0e, 0x44, 0xa9, 0x7d, 0xb7, 0x58, 0xb7, 0xfd, 0x55, 0xd1, 0xa0, 0x9e, 0x5a, 0x4f,
	0x75, 0x8f, 0xa0, 0xf9, 0xcf, 0x9b, 0xf3, 0x08, 0xf6, 0x9e, 0xf5, 0x9f, 0xbf, 0x6c, 0x55, 0x78,
	0x0d, 0xaa, 0xaf, 0x5e, 0xb4, 0xd8, 0xf1, 0xf7, 0x2a, 0xd4, 0xfb, 0x79, 0x4a, 0xfc, 0x14, 0x6a,
	0x81, 0x2b, 0x25, 0x57, 0xfe, 0x0e, 0xf1, 0xad, 0xad, 0x7e, 0x7e, 0x84, 0xb8, 0xf1, 0xe5, 0xe7,
	0x9f, 0x6f, 0xd5, 0x06, 0xaf, 0xcb, 0x3c, 0x5b, 0xfe, 0x16, 0x1a, 0x9b, 0x93, 0x79, 0x67, 0xb3,
	0x76, 0x35, 0xef, 0x38, 0xde, 0x35, 0x0a, 0xa2, 0x87, 0x5e, 0xb4, 0x25, 0xae, 0x4b, 0xe5, 0x68,
	0x22, 0x7d, 0x16, 0x27, 0xac, 0xcb, 0x3f, 0x97, 0xb2, 0x2e, 0xe2, 0xe0, 0x47, 0xdb, 0x42, 0x57,
	0xc2, 0x8d, 0xc5, 0xff, 0x90, 0xe0, 0x79, 0xdb, 0x7b, 0x76, 0xc4, 0x41, 0xc9, 0x53, 0x2e, 0x02,
	0x75, 0xc2, 0xba, 0x4f, 0xda, 0x3f, 0x96, 0x09, 0xbb, 0x5c, 0x26, 0xec, 0xf7, 0x32, 0x61, 0x5f,
	0x57, 0x49, 0xe5, 0x72, 0x95, 0x54, 0x7e, 0xad, 0x92, 0xca, 0x79, 0xcd, 0xff, 0xbc, 0x0f, 0xff,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x67, 0x19, 0x31, 0x1d, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountClient interface {
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	AuthToken(ctx context.Context, in *AuthTokenRequest, opts ...grpc.CallOption) (*AuthTokenResponse, error)
	AuthTokenValidate(ctx context.Context, in *AuthTokenValidateRequest, opts ...grpc.CallOption) (*AuthTokenValidateResponse, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/account.Account/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) AuthToken(ctx context.Context, in *AuthTokenRequest, opts ...grpc.CallOption) (*AuthTokenResponse, error) {
	out := new(AuthTokenResponse)
	err := c.cc.Invoke(ctx, "/account.Account/AuthToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) AuthTokenValidate(ctx context.Context, in *AuthTokenValidateRequest, opts ...grpc.CallOption) (*AuthTokenValidateResponse, error) {
	out := new(AuthTokenValidateResponse)
	err := c.cc.Invoke(ctx, "/account.Account/AuthTokenValidate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
type AccountServer interface {
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	AuthToken(context.Context, *AuthTokenRequest) (*AuthTokenResponse, error)
	AuthTokenValidate(context.Context, *AuthTokenValidateRequest) (*AuthTokenValidateResponse, error)
}

// UnimplementedAccountServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (*UnimplementedAccountServer) Status(ctx context.Context, req *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (*UnimplementedAccountServer) AuthToken(ctx context.Context, req *AuthTokenRequest) (*AuthTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthToken not implemented")
}
func (*UnimplementedAccountServer) AuthTokenValidate(ctx context.Context, req *AuthTokenValidateRequest) (*AuthTokenValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthTokenValidate not implemented")
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_AuthToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).AuthToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/AuthToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).AuthToken(ctx, req.(*AuthTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_AuthTokenValidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthTokenValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).AuthTokenValidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/AuthTokenValidate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).AuthTokenValidate(ctx, req.(*AuthTokenValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Account_Status_Handler,
		},
		{
			MethodName: "AuthToken",
			Handler:    _Account_AuthToken_Handler,
		},
		{
			MethodName: "AuthTokenValidate",
			Handler:    _Account_AuthTokenValidate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}

func (m *StatusRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatusRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Full {
		dAtA[i] = 0x8
		i++
		if m.Full {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *StatusResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatusResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintAccount(dAtA, i, uint64(m.Status))
	}
	return i, nil
}

func (m *AuthTokenRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthTokenRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Username) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Username)))
		i += copy(dAtA[i:], m.Username)
	}
	if len(m.Password) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Password)))
		i += copy(dAtA[i:], m.Password)
	}
	return i, nil
}

func (m *AuthTokenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthTokenResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Token) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	return i, nil
}

func (m *AuthTokenValidateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthTokenValidateRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Token) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	return i, nil
}

func (m *AuthTokenValidateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthTokenValidateResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.IsValid {
		dAtA[i] = 0x10
		i++
		if m.IsValid {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintAccount(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *StatusRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Full {
		n += 2
	}
	return n
}

func (m *StatusResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovAccount(uint64(m.Status))
	}
	return n
}

func (m *AuthTokenRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	return n
}

func (m *AuthTokenResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	return n
}

func (m *AuthTokenValidateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	return n
}

func (m *AuthTokenValidateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsValid {
		n += 2
	}
	return n
}

func sovAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccount(x uint64) (n int) {
	return sovAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StatusRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatusRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Full", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Full = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StatusResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatusResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= ServiceStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthTokenRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AuthTokenRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthTokenRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthTokenResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AuthTokenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthTokenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthTokenValidateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AuthTokenValidateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthTokenValidateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthTokenValidateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AuthTokenValidateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthTokenValidateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsValid", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsValid = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccount
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAccount
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthAccount
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAccount
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAccount(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthAccount
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAccount = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccount   = fmt.Errorf("proto: integer overflow")
)
