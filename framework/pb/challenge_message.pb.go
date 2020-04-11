// Code generated by protoc-gen-go. DO NOT EDIT.
// source: challenge_message.proto

package pb

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

type Challenge struct {
	Title                string           `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Slug                 string           `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	Description          string           `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Tags                 string           `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
	Requirements         string           `protobuf:"bytes,5,opt,name=requirements,proto3" json:"requirements,omitempty"`
	File                 []*ChallengeFile `protobuf:"bytes,6,rep,name=file,proto3" json:"file,omitempty"`
	Level                string           `protobuf:"bytes,7,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Challenge) Reset()         { *m = Challenge{} }
func (m *Challenge) String() string { return proto.CompactTextString(m) }
func (*Challenge) ProtoMessage()    {}
func (*Challenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e18aabac6b3a996, []int{0}
}

func (m *Challenge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Challenge.Unmarshal(m, b)
}
func (m *Challenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Challenge.Marshal(b, m, deterministic)
}
func (m *Challenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Challenge.Merge(m, src)
}
func (m *Challenge) XXX_Size() int {
	return xxx_messageInfo_Challenge.Size(m)
}
func (m *Challenge) XXX_DiscardUnknown() {
	xxx_messageInfo_Challenge.DiscardUnknown(m)
}

var xxx_messageInfo_Challenge proto.InternalMessageInfo

func (m *Challenge) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Challenge) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Challenge) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Challenge) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

func (m *Challenge) GetRequirements() string {
	if m != nil {
		return m.Requirements
	}
	return ""
}

func (m *Challenge) GetFile() []*ChallengeFile {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *Challenge) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

type ChallengeFile struct {
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChallengeFile) Reset()         { *m = ChallengeFile{} }
func (m *ChallengeFile) String() string { return proto.CompactTextString(m) }
func (*ChallengeFile) ProtoMessage()    {}
func (*ChallengeFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e18aabac6b3a996, []int{1}
}

func (m *ChallengeFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChallengeFile.Unmarshal(m, b)
}
func (m *ChallengeFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChallengeFile.Marshal(b, m, deterministic)
}
func (m *ChallengeFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChallengeFile.Merge(m, src)
}
func (m *ChallengeFile) XXX_Size() int {
	return xxx_messageInfo_ChallengeFile.Size(m)
}
func (m *ChallengeFile) XXX_DiscardUnknown() {
	xxx_messageInfo_ChallengeFile.DiscardUnknown(m)
}

var xxx_messageInfo_ChallengeFile proto.InternalMessageInfo

func (m *ChallengeFile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChallengeFile) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type NewChallengeRequest struct {
	Challenge            *Challenge `protobuf:"bytes,1,opt,name=challenge,proto3" json:"challenge,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *NewChallengeRequest) Reset()         { *m = NewChallengeRequest{} }
func (m *NewChallengeRequest) String() string { return proto.CompactTextString(m) }
func (*NewChallengeRequest) ProtoMessage()    {}
func (*NewChallengeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e18aabac6b3a996, []int{2}
}

func (m *NewChallengeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewChallengeRequest.Unmarshal(m, b)
}
func (m *NewChallengeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewChallengeRequest.Marshal(b, m, deterministic)
}
func (m *NewChallengeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewChallengeRequest.Merge(m, src)
}
func (m *NewChallengeRequest) XXX_Size() int {
	return xxx_messageInfo_NewChallengeRequest.Size(m)
}
func (m *NewChallengeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewChallengeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewChallengeRequest proto.InternalMessageInfo

func (m *NewChallengeRequest) GetChallenge() *Challenge {
	if m != nil {
		return m.Challenge
	}
	return nil
}

type NewChallengeResponse struct {
	ChallengeId          string   `protobuf:"bytes,1,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewChallengeResponse) Reset()         { *m = NewChallengeResponse{} }
func (m *NewChallengeResponse) String() string { return proto.CompactTextString(m) }
func (*NewChallengeResponse) ProtoMessage()    {}
func (*NewChallengeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e18aabac6b3a996, []int{3}
}

func (m *NewChallengeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewChallengeResponse.Unmarshal(m, b)
}
func (m *NewChallengeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewChallengeResponse.Marshal(b, m, deterministic)
}
func (m *NewChallengeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewChallengeResponse.Merge(m, src)
}
func (m *NewChallengeResponse) XXX_Size() int {
	return xxx_messageInfo_NewChallengeResponse.Size(m)
}
func (m *NewChallengeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewChallengeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewChallengeResponse proto.InternalMessageInfo

func (m *NewChallengeResponse) GetChallengeId() string {
	if m != nil {
		return m.ChallengeId
	}
	return ""
}

func init() {
	proto.RegisterType((*Challenge)(nil), "education.code.codeedu.Challenge")
	proto.RegisterType((*ChallengeFile)(nil), "education.code.codeedu.ChallengeFile")
	proto.RegisterType((*NewChallengeRequest)(nil), "education.code.codeedu.NewChallengeRequest")
	proto.RegisterType((*NewChallengeResponse)(nil), "education.code.codeedu.NewChallengeResponse")
}

func init() {
	proto.RegisterFile("challenge_message.proto", fileDescriptor_9e18aabac6b3a996)
}

var fileDescriptor_9e18aabac6b3a996 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x4d, 0x9b, 0x56, 0x3a, 0xa9, 0x58, 0xd6, 0xa2, 0x8b, 0xa7, 0x36, 0x20, 0x14, 0x94,
	0x1c, 0x2a, 0x1e, 0x7a, 0x12, 0x2c, 0x08, 0x5e, 0x3c, 0x44, 0xf0, 0xe0, 0x45, 0xd2, 0x64, 0x8c,
	0x0b, 0xdb, 0x4d, 0xba, 0x7f, 0xea, 0xdd, 0x4f, 0xea, 0x47, 0x91, 0xdd, 0xb4, 0xa9, 0x81, 0x8a,
	0x5e, 0xc2, 0xdb, 0xb7, 0xfb, 0x86, 0xdf, 0xcc, 0x04, 0xce, 0xd2, 0xf7, 0x84, 0x73, 0x14, 0x39,
	0xbe, 0x2e, 0x51, 0xa9, 0x24, 0xc7, 0xa8, 0x94, 0x85, 0x2e, 0xc8, 0x29, 0x66, 0x26, 0x4d, 0x34,
	0x2b, 0x44, 0x94, 0x16, 0x19, 0xba, 0x0f, 0x66, 0x26, 0xfc, 0xf2, 0xa0, 0x37, 0xdf, 0x66, 0xc8,
	0x10, 0x3a, 0x9a, 0x69, 0x8e, 0xd4, 0x1b, 0x79, 0x93, 0x5e, 0x5c, 0x1d, 0x08, 0x01, 0x5f, 0x71,
	0x93, 0xd3, 0x96, 0x33, 0x9d, 0x26, 0x23, 0x08, 0x32, 0x54, 0xa9, 0x64, 0xa5, 0xad, 0x49, 0xdb,
	0xee, 0xea, 0xa7, 0x65, 0x53, 0x3a, 0xc9, 0x15, 0xf5, 0xab, 0x94, 0xd5, 0x24, 0x84, 0xbe, 0xc4,
	0x95, 0x61, 0x12, 0x97, 0x28, 0xb4, 0xa2, 0x1d, 0x77, 0xd7, 0xf0, 0xc8, 0x0c, 0xfc, 0x37, 0xc6,
	0x91, 0x76, 0x47, 0xed, 0x49, 0x30, 0xbd, 0x88, 0xf6, 0x83, 0x47, 0x35, 0xf4, 0x3d, 0xe3, 0x18,
	0xbb, 0x88, 0xc5, 0xe7, 0xb8, 0x46, 0x4e, 0x0f, 0x2b, 0x7c, 0x77, 0x08, 0x6f, 0xe0, 0xa8, 0xf1,
	0xd8, 0x92, 0x89, 0x64, 0x89, 0xdb, 0x7e, 0xac, 0x26, 0x03, 0x68, 0x1b, 0xc9, 0x37, 0x7d, 0x58,
	0x19, 0x3e, 0xc3, 0xc9, 0x23, 0x7e, 0xd4, 0xc9, 0x18, 0x57, 0x06, 0x95, 0x26, 0xb7, 0xd0, 0xab,
	0x67, 0xec, 0xc6, 0x14, 0x4c, 0xc7, 0x7f, 0x32, 0xc6, 0xbb, 0x4c, 0x38, 0x83, 0x61, 0xb3, 0xae,
	0x2a, 0x0b, 0xa1, 0x90, 0x8c, 0xa1, 0xbf, 0x5b, 0x1e, 0xcb, 0x36, 0x2b, 0x08, 0x6a, 0xef, 0x21,
	0x9b, 0x7e, 0x7a, 0x30, 0xa8, 0x83, 0x4f, 0x28, 0xd7, 0x2c, 0x45, 0x22, 0xe0, 0x78, 0x2e, 0x31,
	0xd1, 0xb8, 0x5b, 0xe3, 0xe5, 0x6f, 0x40, 0x7b, 0x1a, 0x3a, 0xbf, 0xfa, 0xdf, 0xe3, 0x8a, 0x32,
	0x3c, 0xb8, 0xf3, 0x5f, 0x5a, 0xe5, 0x62, 0xd1, 0x75, 0xbf, 0xd5, 0xf5, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x4c, 0x9a, 0x04, 0xac, 0x71, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChallengeServiceClient is the client API for ChallengeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChallengeServiceClient interface {
	CreateChallenge(ctx context.Context, in *NewChallengeRequest, opts ...grpc.CallOption) (*NewChallengeResponse, error)
}

type challengeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChallengeServiceClient(cc grpc.ClientConnInterface) ChallengeServiceClient {
	return &challengeServiceClient{cc}
}

func (c *challengeServiceClient) CreateChallenge(ctx context.Context, in *NewChallengeRequest, opts ...grpc.CallOption) (*NewChallengeResponse, error) {
	out := new(NewChallengeResponse)
	err := c.cc.Invoke(ctx, "/education.code.codeedu.ChallengeService/CreateChallenge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChallengeServiceServer is the server API for ChallengeService service.
type ChallengeServiceServer interface {
	CreateChallenge(context.Context, *NewChallengeRequest) (*NewChallengeResponse, error)
}

// UnimplementedChallengeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChallengeServiceServer struct {
}

func (*UnimplementedChallengeServiceServer) CreateChallenge(ctx context.Context, req *NewChallengeRequest) (*NewChallengeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChallenge not implemented")
}

func RegisterChallengeServiceServer(s *grpc.Server, srv ChallengeServiceServer) {
	s.RegisterService(&_ChallengeService_serviceDesc, srv)
}

func _ChallengeService_CreateChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewChallengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChallengeServiceServer).CreateChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/education.code.codeedu.ChallengeService/CreateChallenge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChallengeServiceServer).CreateChallenge(ctx, req.(*NewChallengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChallengeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "education.code.codeedu.ChallengeService",
	HandlerType: (*ChallengeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChallenge",
			Handler:    _ChallengeService_CreateChallenge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "challenge_message.proto",
}