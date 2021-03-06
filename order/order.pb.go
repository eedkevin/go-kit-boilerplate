// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: order.proto

package order

import (
	context "context"
	encoding_binary "encoding/binary"
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
	return fileDescriptor_cd01338c35d87077, []int{0}
}

type GeoCoordinate struct {
	Lat float32 `protobuf:"fixed32,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng float32 `protobuf:"fixed32,2,opt,name=lng,proto3" json:"lng,omitempty"`
}

func (m *GeoCoordinate) Reset()         { *m = GeoCoordinate{} }
func (m *GeoCoordinate) String() string { return proto.CompactTextString(m) }
func (*GeoCoordinate) ProtoMessage()    {}
func (*GeoCoordinate) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}
func (m *GeoCoordinate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GeoCoordinate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GeoCoordinate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GeoCoordinate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeoCoordinate.Merge(m, src)
}
func (m *GeoCoordinate) XXX_Size() int {
	return m.Size()
}
func (m *GeoCoordinate) XXX_DiscardUnknown() {
	xxx_messageInfo_GeoCoordinate.DiscardUnknown(m)
}

var xxx_messageInfo_GeoCoordinate proto.InternalMessageInfo

func (m *GeoCoordinate) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *GeoCoordinate) GetLng() float32 {
	if m != nil {
		return m.Lng
	}
	return 0
}

type StatusRequest struct {
	Full bool `protobuf:"varint,1,opt,name=full,proto3" json:"full,omitempty"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1}
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
	Status ServiceStatus `protobuf:"varint,1,opt,name=status,proto3,enum=order.ServiceStatus" json:"status,omitempty"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{2}
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

type PlaceOrderRequest struct {
	UserId      string         `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Origin      *GeoCoordinate `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	Destination *GeoCoordinate `protobuf:"bytes,3,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (m *PlaceOrderRequest) Reset()         { *m = PlaceOrderRequest{} }
func (m *PlaceOrderRequest) String() string { return proto.CompactTextString(m) }
func (*PlaceOrderRequest) ProtoMessage()    {}
func (*PlaceOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{3}
}
func (m *PlaceOrderRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PlaceOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PlaceOrderRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PlaceOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlaceOrderRequest.Merge(m, src)
}
func (m *PlaceOrderRequest) XXX_Size() int {
	return m.Size()
}
func (m *PlaceOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PlaceOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PlaceOrderRequest proto.InternalMessageInfo

func (m *PlaceOrderRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *PlaceOrderRequest) GetOrigin() *GeoCoordinate {
	if m != nil {
		return m.Origin
	}
	return nil
}

func (m *PlaceOrderRequest) GetDestination() *GeoCoordinate {
	if m != nil {
		return m.Destination
	}
	return nil
}

type PlaceOrderResponse struct {
	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (m *PlaceOrderResponse) Reset()         { *m = PlaceOrderResponse{} }
func (m *PlaceOrderResponse) String() string { return proto.CompactTextString(m) }
func (*PlaceOrderResponse) ProtoMessage()    {}
func (*PlaceOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{4}
}
func (m *PlaceOrderResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PlaceOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PlaceOrderResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PlaceOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlaceOrderResponse.Merge(m, src)
}
func (m *PlaceOrderResponse) XXX_Size() int {
	return m.Size()
}
func (m *PlaceOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PlaceOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PlaceOrderResponse proto.InternalMessageInfo

func (m *PlaceOrderResponse) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func init() {
	proto.RegisterEnum("order.ServiceStatus", ServiceStatus_name, ServiceStatus_value)
	proto.RegisterType((*GeoCoordinate)(nil), "order.GeoCoordinate")
	proto.RegisterType((*StatusRequest)(nil), "order.StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "order.StatusResponse")
	proto.RegisterType((*PlaceOrderRequest)(nil), "order.PlaceOrderRequest")
	proto.RegisterType((*PlaceOrderResponse)(nil), "order.PlaceOrderResponse")
}

func init() { proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077) }

var fileDescriptor_cd01338c35d87077 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x6b, 0x14, 0x31,
	0x14, 0xc7, 0x27, 0xd3, 0x76, 0x76, 0xfb, 0xd6, 0xad, 0xdb, 0xa0, 0x74, 0xba, 0x87, 0x41, 0xc7,
	0x8b, 0x14, 0xd9, 0x40, 0x0b, 0x1e, 0x3c, 0x08, 0x2a, 0x56, 0x16, 0x85, 0xca, 0xf4, 0xe6, 0xa5,
	0xa4, 0x9b, 0xd7, 0x69, 0x60, 0x9a, 0x8c, 0x49, 0xa6, 0xe0, 0xd5, 0x4f, 0x20, 0xfa, 0x31, 0xfc,
	0x22, 0x1e, 0x0b, 0x5e, 0x3c, 0xca, 0xae, 0x1f, 0x44, 0x26, 0x93, 0xb6, 0x5b, 0x15, 0x6f, 0x79,
	0xff, 0xfc, 0xde, 0x3f, 0xef, 0xfd, 0x09, 0x0c, 0xb4, 0x11, 0x68, 0x26, 0xb5, 0xd1, 0x4e, 0xd3,
	0x35, 0x5f, 0x8c, 0x5f, 0x96, 0xd2, 0x9d, 0x36, 0xc7, 0x93, 0x99, 0x3e, 0x63, 0x67, 0xe8, 0xf8,
	0x39, 0x1a, 0x8b, 0xcc, 0x99, 0xc6, 0x5a, 0x26, 0xf0, 0xc4, 0x19, 0x44, 0x56, 0x6a, 0x5d, 0x56,
	0xe8, 0x4e, 0xa5, 0x11, 0x35, 0x37, 0xee, 0x03, 0xe3, 0x4a, 0x69, 0xc7, 0x9d, 0xd4, 0xca, 0x76,
	0x6e, 0xf9, 0x1e, 0x0c, 0x5f, 0xa1, 0x7e, 0xa1, 0xb5, 0x11, 0x52, 0x71, 0x87, 0x74, 0x04, 0x2b,
	0x15, 0x77, 0x29, 0xb9, 0x47, 0x1e, 0xc6, 0x45, 0x7b, 0xf4, 0x8a, 0x2a, 0xd3, 0x38, 0x28, 0xaa,
	0xcc, 0x1f, 0xc0, 0xf0, 0xd0, 0x71, 0xd7, 0xd8, 0x02, 0xdf, 0x37, 0x68, 0x1d, 0xa5, 0xb0, 0x7a,
	0xd2, 0x54, 0x95, 0xef, 0xea, 0x17, 0xfe, 0x9c, 0x3f, 0x85, 0x8d, 0x4b, 0xc8, 0xd6, 0x5a, 0x59,
	0xa4, 0x8f, 0x20, 0xb1, 0x5e, 0xf1, 0xdc, 0xc6, 0xee, 0x9d, 0x49, 0xb7, 0xd7, 0x21, 0x9a, 0x73,
	0x39, 0xc3, 0x40, 0x07, 0x26, 0xff, 0x4c, 0x60, 0xf3, 0x6d, 0xc5, 0x67, 0x78, 0xd0, 0x42, 0x97,
	0x2f, 0x6d, 0x41, 0xaf, 0xb1, 0x68, 0x8e, 0xa4, 0xf0, 0x26, 0xeb, 0x45, 0xd2, 0x96, 0x53, 0xd1,
	0x9a, 0x6b, 0x23, 0x4b, 0xa9, 0xfc, 0xa0, 0x83, 0x2b, 0xf3, 0x1b, 0xdb, 0x15, 0x81, 0xa1, 0x8f,
	0x61, 0x20, 0xd0, 0xba, 0x56, 0x94, 0x5a, 0xa5, 0x2b, 0xff, 0x69, 0x59, 0x06, 0x73, 0x06, 0x74,
	0x79, 0xa6, 0xb0, 0xd8, 0x36, 0xf4, 0x7d, 0xe7, 0xf5, 0x54, 0x3d, 0x5f, 0x4f, 0xc5, 0xce, 0x7d,
	0x18, 0xde, 0x58, 0x8f, 0xf6, 0x61, 0x75, 0xff, 0xd9, 0xf4, 0xcd, 0x28, 0xa2, 0x09, 0xc4, 0x07,
	0xaf, 0x47, 0x64, 0xf7, 0x2b, 0x81, 0x35, 0xef, 0x47, 0xf7, 0x21, 0x09, 0xd4, 0x55, 0x34, 0xcb,
	0x31, 0x8f, 0xef, 0xfe, 0xa1, 0x76, 0xcf, 0xe7, 0xb7, 0x3f, 0x7e, 0xff, 0xf5, 0x25, 0x5e, 0xa7,
	0x3d, 0xd6, 0x45, 0x47, 0xdf, 0x01, 0x5c, 0x4f, 0x49, 0xd3, 0xd0, 0xf5, 0x57, 0x98, 0xe3, 0xed,
	0x7f, 0xdc, 0x04, 0xcf, 0x2d, 0xef, 0xb9, 0x99, 0xdf, 0x62, 0x75, 0x7b, 0x79, 0xe4, 0xc1, 0x27,
	0x64, 0xe7, 0x79, 0xfa, 0x6d, 0x9e, 0x91, 0x8b, 0x79, 0x46, 0x7e, 0xce, 0x33, 0xf2, 0x69, 0x91,
	0x45, 0x17, 0x8b, 0x2c, 0xfa, 0xb1, 0xc8, 0xa2, 0xe3, 0xc4, 0xff, 0xa8, 0xbd, 0xdf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x37, 0xf5, 0xf4, 0xb7, 0xae, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderClient interface {
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	PlaceOrder(ctx context.Context, in *PlaceOrderRequest, opts ...grpc.CallOption) (*PlaceOrderResponse, error)
}

type orderClient struct {
	cc *grpc.ClientConn
}

func NewOrderClient(cc *grpc.ClientConn) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/order.Order/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) PlaceOrder(ctx context.Context, in *PlaceOrderRequest, opts ...grpc.CallOption) (*PlaceOrderResponse, error) {
	out := new(PlaceOrderResponse)
	err := c.cc.Invoke(ctx, "/order.Order/PlaceOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
type OrderServer interface {
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	PlaceOrder(context.Context, *PlaceOrderRequest) (*PlaceOrderResponse, error)
}

// UnimplementedOrderServer can be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (*UnimplementedOrderServer) Status(ctx context.Context, req *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (*UnimplementedOrderServer) PlaceOrder(ctx context.Context, req *PlaceOrderRequest) (*PlaceOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")
}

func RegisterOrderServer(s *grpc.Server, srv OrderServer) {
	s.RegisterService(&_Order_serviceDesc, srv)
}

func _Order_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_PlaceOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).PlaceOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/PlaceOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).PlaceOrder(ctx, req.(*PlaceOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Order_serviceDesc = grpc.ServiceDesc{
	ServiceName: "order.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Order_Status_Handler,
		},
		{
			MethodName: "PlaceOrder",
			Handler:    _Order_PlaceOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}

func (m *GeoCoordinate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GeoCoordinate) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Lat != 0 {
		dAtA[i] = 0xd
		i++
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Lat))))
		i += 4
	}
	if m.Lng != 0 {
		dAtA[i] = 0x15
		i++
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Lng))))
		i += 4
	}
	return i, nil
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
		i = encodeVarintOrder(dAtA, i, uint64(m.Status))
	}
	return i, nil
}

func (m *PlaceOrderRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlaceOrderRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UserId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintOrder(dAtA, i, uint64(len(m.UserId)))
		i += copy(dAtA[i:], m.UserId)
	}
	if m.Origin != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintOrder(dAtA, i, uint64(m.Origin.Size()))
		n1, err1 := m.Origin.MarshalTo(dAtA[i:])
		if err1 != nil {
			return 0, err1
		}
		i += n1
	}
	if m.Destination != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintOrder(dAtA, i, uint64(m.Destination.Size()))
		n2, err2 := m.Destination.MarshalTo(dAtA[i:])
		if err2 != nil {
			return 0, err2
		}
		i += n2
	}
	return i, nil
}

func (m *PlaceOrderResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlaceOrderResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.OrderId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintOrder(dAtA, i, uint64(len(m.OrderId)))
		i += copy(dAtA[i:], m.OrderId)
	}
	return i, nil
}

func encodeVarintOrder(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GeoCoordinate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Lat != 0 {
		n += 5
	}
	if m.Lng != 0 {
		n += 5
	}
	return n
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
		n += 1 + sovOrder(uint64(m.Status))
	}
	return n
}

func (m *PlaceOrderRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UserId)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	if m.Origin != nil {
		l = m.Origin.Size()
		n += 1 + l + sovOrder(uint64(l))
	}
	if m.Destination != nil {
		l = m.Destination.Size()
		n += 1 + l + sovOrder(uint64(l))
	}
	return n
}

func (m *PlaceOrderResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OrderId)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	return n
}

func sovOrder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrder(x uint64) (n int) {
	return sovOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GeoCoordinate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
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
			return fmt.Errorf("proto: GeoCoordinate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GeoCoordinate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lat", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Lat = float32(math.Float32frombits(v))
		case 2:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lng", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Lng = float32(math.Float32frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrder
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrder
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
func (m *StatusRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
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
					return ErrIntOverflowOrder
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
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrder
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrder
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
				return ErrIntOverflowOrder
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
					return ErrIntOverflowOrder
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
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrder
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrder
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
func (m *PlaceOrderRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
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
			return fmt.Errorf("proto: PlaceOrderRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlaceOrderRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Origin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Origin == nil {
				m.Origin = &GeoCoordinate{}
			}
			if err := m.Origin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Destination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Destination == nil {
				m.Destination = &GeoCoordinate{}
			}
			if err := m.Destination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrder
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrder
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
func (m *PlaceOrderResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
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
			return fmt.Errorf("proto: PlaceOrderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlaceOrderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrder
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrder
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
func skipOrder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrder
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
					return 0, ErrIntOverflowOrder
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
					return 0, ErrIntOverflowOrder
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
				return 0, ErrInvalidLengthOrder
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthOrder
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowOrder
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
				next, err := skipOrder(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthOrder
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
	ErrInvalidLengthOrder = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrder   = fmt.Errorf("proto: integer overflow")
)
