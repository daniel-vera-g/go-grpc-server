// Code generated by protoc-gen-go. DO NOT EDIT.
// source: birthday.proto

package birthday

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

type Date struct {
	Day                  int64    `protobuf:"varint,1,opt,name=day,proto3" json:"day,omitempty"`
	Month                int64    `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	Year                 int64    `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Date) Reset()         { *m = Date{} }
func (m *Date) String() string { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()    {}
func (*Date) Descriptor() ([]byte, []int) {
	return fileDescriptor_b57480d75fb8881b, []int{0}
}

func (m *Date) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Date.Unmarshal(m, b)
}
func (m *Date) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Date.Marshal(b, m, deterministic)
}
func (m *Date) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Date.Merge(m, src)
}
func (m *Date) XXX_Size() int {
	return xxx_messageInfo_Date.Size(m)
}
func (m *Date) XXX_DiscardUnknown() {
	xxx_messageInfo_Date.DiscardUnknown(m)
}

var xxx_messageInfo_Date proto.InternalMessageInfo

func (m *Date) GetDay() int64 {
	if m != nil {
		return m.Day
	}
	return 0
}

func (m *Date) GetMonth() int64 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *Date) GetYear() int64 {
	if m != nil {
		return m.Year
	}
	return 0
}

type BirthdayStatus struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Age                  int64    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BirthdayStatus) Reset()         { *m = BirthdayStatus{} }
func (m *BirthdayStatus) String() string { return proto.CompactTextString(m) }
func (*BirthdayStatus) ProtoMessage()    {}
func (*BirthdayStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_b57480d75fb8881b, []int{1}
}

func (m *BirthdayStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BirthdayStatus.Unmarshal(m, b)
}
func (m *BirthdayStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BirthdayStatus.Marshal(b, m, deterministic)
}
func (m *BirthdayStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BirthdayStatus.Merge(m, src)
}
func (m *BirthdayStatus) XXX_Size() int {
	return xxx_messageInfo_BirthdayStatus.Size(m)
}
func (m *BirthdayStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_BirthdayStatus.DiscardUnknown(m)
}

var xxx_messageInfo_BirthdayStatus proto.InternalMessageInfo

func (m *BirthdayStatus) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *BirthdayStatus) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func init() {
	proto.RegisterType((*Date)(nil), "birthday.Date")
	proto.RegisterType((*BirthdayStatus)(nil), "birthday.BirthdayStatus")
}

func init() { proto.RegisterFile("birthday.proto", fileDescriptor_b57480d75fb8881b) }

var fileDescriptor_b57480d75fb8881b = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xca, 0x2c, 0x2a,
	0xc9, 0x48, 0x49, 0xac, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x9c,
	0xb8, 0x58, 0x5c, 0x12, 0x4b, 0x52, 0x85, 0x04, 0xb8, 0x98, 0x53, 0x12, 0x2b, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x98, 0x83, 0x40, 0x4c, 0x21, 0x11, 0x2e, 0xd6, 0xdc, 0xfc, 0xbc, 0x92, 0x0c, 0x09,
	0x26, 0xb0, 0x18, 0x84, 0x23, 0x24, 0xc4, 0xc5, 0x52, 0x99, 0x9a, 0x58, 0x24, 0xc1, 0x0c, 0x16,
	0x04, 0xb3, 0x95, 0xac, 0xb8, 0xf8, 0x9c, 0xa0, 0xe6, 0x05, 0x97, 0x24, 0x96, 0x94, 0x16, 0x0b,
	0x89, 0x71, 0xb1, 0x15, 0x83, 0x59, 0x60, 0x03, 0x39, 0x82, 0xa0, 0x3c, 0x90, 0x2d, 0x89, 0xe9,
	0xa9, 0x50, 0x13, 0x41, 0x4c, 0x23, 0x1f, 0x2e, 0x7e, 0x98, 0x5e, 0xe7, 0x8c, 0xd4, 0xe4, 0xec,
	0xd4, 0x22, 0x21, 0x4b, 0x2e, 0xde, 0x64, 0x10, 0x13, 0x26, 0x2e, 0xc4, 0xa7, 0x07, 0x77, 0x3e,
	0xc8, 0xad, 0x52, 0x12, 0x08, 0x3e, 0xaa, 0xbd, 0x49, 0x6c, 0x60, 0xef, 0x19, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x5c, 0x98, 0x26, 0x9c, 0xf0, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BirthdayCheckerClient is the client API for BirthdayChecker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BirthdayCheckerClient interface {
	CheckBirthday(ctx context.Context, in *Date, opts ...grpc.CallOption) (*BirthdayStatus, error)
}

type birthdayCheckerClient struct {
	cc *grpc.ClientConn
}

func NewBirthdayCheckerClient(cc *grpc.ClientConn) BirthdayCheckerClient {
	return &birthdayCheckerClient{cc}
}

func (c *birthdayCheckerClient) CheckBirthday(ctx context.Context, in *Date, opts ...grpc.CallOption) (*BirthdayStatus, error) {
	out := new(BirthdayStatus)
	err := c.cc.Invoke(ctx, "/birthday.BirthdayChecker/checkBirthday", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BirthdayCheckerServer is the server API for BirthdayChecker service.
type BirthdayCheckerServer interface {
	CheckBirthday(context.Context, *Date) (*BirthdayStatus, error)
}

// UnimplementedBirthdayCheckerServer can be embedded to have forward compatible implementations.
type UnimplementedBirthdayCheckerServer struct {
}

func (*UnimplementedBirthdayCheckerServer) CheckBirthday(ctx context.Context, req *Date) (*BirthdayStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckBirthday not implemented")
}

func RegisterBirthdayCheckerServer(s *grpc.Server, srv BirthdayCheckerServer) {
	s.RegisterService(&_BirthdayChecker_serviceDesc, srv)
}

func _BirthdayChecker_CheckBirthday_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Date)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirthdayCheckerServer).CheckBirthday(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/birthday.BirthdayChecker/CheckBirthday",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirthdayCheckerServer).CheckBirthday(ctx, req.(*Date))
	}
	return interceptor(ctx, in, info, handler)
}

var _BirthdayChecker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "birthday.BirthdayChecker",
	HandlerType: (*BirthdayCheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "checkBirthday",
			Handler:    _BirthdayChecker_CheckBirthday_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "birthday.proto",
}
