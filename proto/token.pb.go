// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.14.0
// source: proto/token.proto

package steamtoken

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Memo  string `protobuf:"bytes,1,opt,name=memo,proto3" json:"memo,omitempty"`
	AppId int32  `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *CreateTokenRequest) Reset() {
	*x = CreateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokenRequest) ProtoMessage() {}

func (x *CreateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokenRequest.ProtoReflect.Descriptor instead.
func (*CreateTokenRequest) Descriptor() ([]byte, []int) {
	return file_proto_token_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTokenRequest) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *CreateTokenRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

type CreateTokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerToken string `protobuf:"bytes,1,opt,name=server_token,json=serverToken,proto3" json:"server_token,omitempty"`
}

func (x *CreateTokenReply) Reset() {
	*x = CreateTokenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokenReply) ProtoMessage() {}

func (x *CreateTokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokenReply.ProtoReflect.Descriptor instead.
func (*CreateTokenReply) Descriptor() ([]byte, []int) {
	return file_proto_token_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTokenReply) GetServerToken() string {
	if x != nil {
		return x.ServerToken
	}
	return ""
}

type DeleteTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerToken string `protobuf:"bytes,1,opt,name=server_token,json=serverToken,proto3" json:"server_token,omitempty"`
}

func (x *DeleteTokenRequest) Reset() {
	*x = DeleteTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTokenRequest) ProtoMessage() {}

func (x *DeleteTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTokenRequest.ProtoReflect.Descriptor instead.
func (*DeleteTokenRequest) Descriptor() ([]byte, []int) {
	return file_proto_token_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteTokenRequest) GetServerToken() string {
	if x != nil {
		return x.ServerToken
	}
	return ""
}

type DeleteTokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTokenReply) Reset() {
	*x = DeleteTokenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTokenReply) ProtoMessage() {}

func (x *DeleteTokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTokenReply.ProtoReflect.Descriptor instead.
func (*DeleteTokenReply) Descriptor() ([]byte, []int) {
	return file_proto_token_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteTokenReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_token_proto protoreflect.FileDescriptor

var file_proto_token_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x3f, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64,
	0x22, 0x35, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x37, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x22, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x32, 0xac, 0x01, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6e, 0x70, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x65, 0x61,
	0x6d, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3b, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_token_proto_rawDescOnce sync.Once
	file_proto_token_proto_rawDescData = file_proto_token_proto_rawDesc
)

func file_proto_token_proto_rawDescGZIP() []byte {
	file_proto_token_proto_rawDescOnce.Do(func() {
		file_proto_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_token_proto_rawDescData)
	})
	return file_proto_token_proto_rawDescData
}

var file_proto_token_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_token_proto_goTypes = []interface{}{
	(*CreateTokenRequest)(nil), // 0: steamtoken.CreateTokenRequest
	(*CreateTokenReply)(nil),   // 1: steamtoken.CreateTokenReply
	(*DeleteTokenRequest)(nil), // 2: steamtoken.DeleteTokenRequest
	(*DeleteTokenReply)(nil),   // 3: steamtoken.DeleteTokenReply
}
var file_proto_token_proto_depIdxs = []int32{
	0, // 0: steamtoken.Tokenservice.CreateToken:input_type -> steamtoken.CreateTokenRequest
	2, // 1: steamtoken.Tokenservice.DeleteToken:input_type -> steamtoken.DeleteTokenRequest
	1, // 2: steamtoken.Tokenservice.CreateToken:output_type -> steamtoken.CreateTokenReply
	3, // 3: steamtoken.Tokenservice.DeleteToken:output_type -> steamtoken.DeleteTokenReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_token_proto_init() }
func file_proto_token_proto_init() {
	if File_proto_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTokenRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTokenReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTokenRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTokenReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_token_proto_goTypes,
		DependencyIndexes: file_proto_token_proto_depIdxs,
		MessageInfos:      file_proto_token_proto_msgTypes,
	}.Build()
	File_proto_token_proto = out.File
	file_proto_token_proto_rawDesc = nil
	file_proto_token_proto_goTypes = nil
	file_proto_token_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TokenserviceClient is the client API for Tokenservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TokenserviceClient interface {
	CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenReply, error)
	DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*DeleteTokenReply, error)
}

type tokenserviceClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenserviceClient(cc grpc.ClientConnInterface) TokenserviceClient {
	return &tokenserviceClient{cc}
}

func (c *tokenserviceClient) CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenReply, error) {
	out := new(CreateTokenReply)
	err := c.cc.Invoke(ctx, "/steamtoken.Tokenservice/CreateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenserviceClient) DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*DeleteTokenReply, error) {
	out := new(DeleteTokenReply)
	err := c.cc.Invoke(ctx, "/steamtoken.Tokenservice/DeleteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenserviceServer is the server API for Tokenservice service.
type TokenserviceServer interface {
	CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenReply, error)
	DeleteToken(context.Context, *DeleteTokenRequest) (*DeleteTokenReply, error)
}

// UnimplementedTokenserviceServer can be embedded to have forward compatible implementations.
type UnimplementedTokenserviceServer struct {
}

func (*UnimplementedTokenserviceServer) CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (*UnimplementedTokenserviceServer) DeleteToken(context.Context, *DeleteTokenRequest) (*DeleteTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToken not implemented")
}

func RegisterTokenserviceServer(s *grpc.Server, srv TokenserviceServer) {
	s.RegisterService(&_Tokenservice_serviceDesc, srv)
}

func _Tokenservice_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenserviceServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/steamtoken.Tokenservice/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenserviceServer).CreateToken(ctx, req.(*CreateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tokenservice_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenserviceServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/steamtoken.Tokenservice/DeleteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenserviceServer).DeleteToken(ctx, req.(*DeleteTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tokenservice_serviceDesc = grpc.ServiceDesc{
	ServiceName: "steamtoken.Tokenservice",
	HandlerType: (*TokenserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToken",
			Handler:    _Tokenservice_CreateToken_Handler,
		},
		{
			MethodName: "DeleteToken",
			Handler:    _Tokenservice_DeleteToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/token.proto",
}
