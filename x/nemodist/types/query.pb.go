// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nemo/nemodist/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest defines the request type for querying x/nemodist parameters.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08142b3a0a4f2f78, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse defines the response type for querying x/nemodist parameters.
type QueryParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_08142b3a0a4f2f78, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// QueryBalanceRequest defines the request type for querying x/nemodist balance.
type QueryBalanceRequest struct {
}

func (m *QueryBalanceRequest) Reset()         { *m = QueryBalanceRequest{} }
func (m *QueryBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*QueryBalanceRequest) ProtoMessage()    {}
func (*QueryBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08142b3a0a4f2f78, []int{2}
}
func (m *QueryBalanceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBalanceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBalanceRequest.Merge(m, src)
}
func (m *QueryBalanceRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBalanceRequest proto.InternalMessageInfo

// QueryBalanceResponse defines the response type for querying x/nemodist balance.
type QueryBalanceResponse struct {
	Coins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *QueryBalanceResponse) Reset()         { *m = QueryBalanceResponse{} }
func (m *QueryBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBalanceResponse) ProtoMessage()    {}
func (*QueryBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_08142b3a0a4f2f78, []int{3}
}
func (m *QueryBalanceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBalanceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBalanceResponse.Merge(m, src)
}
func (m *QueryBalanceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBalanceResponse proto.InternalMessageInfo

func (m *QueryBalanceResponse) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "nemo.nemodist.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "nemo.nemodist.v1beta1.QueryParamsResponse")
	proto.RegisterType((*QueryBalanceRequest)(nil), "nemo.nemodist.v1beta1.QueryBalanceRequest")
	proto.RegisterType((*QueryBalanceResponse)(nil), "nemo.nemodist.v1beta1.QueryBalanceResponse")
}

func init() { proto.RegisterFile("nemo/nemodist/v1beta1/query.proto", fileDescriptor_08142b3a0a4f2f78) }

var fileDescriptor_08142b3a0a4f2f78 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcf, 0x6e, 0xda, 0x30,
	0x1c, 0x4e, 0xd8, 0x60, 0x92, 0xb9, 0x79, 0x20, 0x6d, 0xd1, 0x66, 0x20, 0x93, 0x26, 0x60, 0xc2,
	0x1e, 0xec, 0xb8, 0x5b, 0xb6, 0x07, 0xd8, 0x72, 0xec, 0xcd, 0x09, 0x56, 0x1a, 0x01, 0x71, 0x88,
	0x0d, 0x2a, 0xd7, 0x1e, 0x7b, 0xaa, 0xd4, 0x27, 0xe8, 0xb5, 0x4f, 0xc2, 0x11, 0xa9, 0x97, 0x9e,
	0xda, 0x0a, 0xfa, 0x20, 0x95, 0x63, 0x43, 0x8b, 0x0a, 0x88, 0x4b, 0x62, 0xfd, 0xfc, 0x7d, 0xbf,
	0xef, 0x4f, 0x02, 0x1a, 0x03, 0x3a, 0xa5, 0x44, 0x3d, 0xfa, 0xb1, 0x90, 0x64, 0xda, 0x0d, 0x98,
	0xa4, 0x5d, 0x32, 0x9e, 0xb0, 0x6c, 0x86, 0xd3, 0x8c, 0x4b, 0x0e, 0xab, 0xea, 0x16, 0xaf, 0x21,
	0xd8, 0x40, 0x1c, 0x14, 0x72, 0x31, 0xe2, 0x82, 0x04, 0x54, 0xb0, 0x0d, 0x2f, 0xe4, 0x71, 0xa2,
	0x69, 0x4e, 0x25, 0xe2, 0x11, 0xcf, 0x8f, 0x44, 0x9d, 0xcc, 0xf4, 0x4b, 0xc4, 0x79, 0x34, 0x64,
	0x84, 0xa6, 0x31, 0xa1, 0x49, 0xc2, 0x25, 0x95, 0x31, 0x4f, 0x84, 0xb9, 0x75, 0x77, 0xbb, 0x49,
	0x69, 0x46, 0x47, 0x06, 0xe3, 0x56, 0x00, 0xfc, 0xaf, 0xdc, 0xfd, 0xcb, 0x87, 0x3e, 0x1b, 0x4f,
	0x98, 0x90, 0xae, 0x0f, 0x3e, 0x6e, 0x4d, 0x45, 0xca, 0x13, 0xc1, 0xe0, 0x6f, 0x50, 0xd2, 0xe4,
	0x4f, 0x76, 0xdd, 0x6e, 0x96, 0x7b, 0x5f, 0xf1, 0xce, 0x30, 0x58, 0xd3, 0xbc, 0xf7, 0xf3, 0xfb,
	0x9a, 0xe5, 0x1b, 0x8a, 0x5b, 0x35, 0x3b, 0x3d, 0x3a, 0xa4, 0x49, 0xc8, 0xd6, 0x52, 0x33, 0x50,
	0xd9, 0x1e, 0x1b, 0x2d, 0x0a, 0x8a, 0x2a, 0xbe, 0x92, 0x7a, 0xd7, 0x2c, 0xf7, 0x3e, 0x63, 0x5d,
	0x10, 0x56, 0x05, 0x6d, 0x84, 0xfe, 0xf0, 0x38, 0xf1, 0x7e, 0x2a, 0x99, 0x9b, 0x87, 0x5a, 0x33,
	0x8a, 0xe5, 0xe9, 0x24, 0xc0, 0x21, 0x1f, 0x11, 0xd3, 0xa6, 0x7e, 0x75, 0x44, 0x7f, 0x40, 0xe4,
	0x2c, 0x65, 0x22, 0x27, 0x08, 0x5f, 0x6f, 0xee, 0x5d, 0x17, 0x40, 0x31, 0xd7, 0x86, 0x17, 0x36,
	0x28, 0x69, 0xd3, 0xb0, 0xb5, 0x27, 0xd3, 0xdb, 0x96, 0x9c, 0xf6, 0x31, 0x50, 0x1d, 0xc7, 0x6d,
	0x9d, 0xdf, 0x3e, 0x5d, 0x15, 0xbe, 0xc1, 0x06, 0x39, 0xf0, 0x51, 0x98, 0x64, 0x99, 0x50, 0x66,
	0x3e, 0x98, 0x36, 0xe0, 0x41, 0x89, 0xed, 0x26, 0x9d, 0x1f, 0x47, 0x61, 0x8d, 0x9f, 0xef, 0xb9,
	0x9f, 0x3a, 0x44, 0x7b, 0xfc, 0x04, 0x1a, 0xef, 0xfd, 0x9d, 0x2f, 0x91, 0xbd, 0x58, 0x22, 0xfb,
	0x71, 0x89, 0xec, 0xcb, 0x15, 0xb2, 0x16, 0x2b, 0x64, 0xdd, 0xad, 0x90, 0x75, 0xd2, 0x7e, 0x55,
	0xb7, 0xa2, 0x77, 0x86, 0x34, 0x10, 0x7a, 0xdb, 0xd9, 0xcb, 0xbe, 0xbc, 0xf6, 0xa0, 0x94, 0xff,
	0x6c, 0xbf, 0x9e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x28, 0xdc, 0x24, 0xfa, 0x20, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Params queries the parameters of x/nemodist module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Balance queries the balance of all coins of x/nemodist module.
	Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/nemo.nemodist.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error) {
	out := new(QueryBalanceResponse)
	err := c.cc.Invoke(ctx, "/nemo.nemodist.v1beta1.Query/Balance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params queries the parameters of x/nemodist module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Balance queries the balance of all coins of x/nemodist module.
	Balance(context.Context, *QueryBalanceRequest) (*QueryBalanceResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Balance(ctx context.Context, req *QueryBalanceRequest) (*QueryBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nemo.nemodist.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nemo.nemodist.v1beta1.Query/Balance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Balance(ctx, req.(*QueryBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nemo.nemodist.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Balance",
			Handler:    _Query_Balance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nemo/nemodist/v1beta1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryBalanceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBalanceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBalanceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryBalanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBalanceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBalanceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryBalanceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryBalanceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryBalanceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryBalanceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBalanceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryBalanceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryBalanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBalanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
