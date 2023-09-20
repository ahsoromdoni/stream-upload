// pkg/proto/upload.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: grpc/upload.proto

package uploadpb

import (
	context "context"
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

type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mime  string `protobuf:"bytes,1,opt,name=mime,proto3" json:"mime,omitempty"`
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_upload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_upload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_grpc_upload_proto_rawDescGZIP(), []int{0}
}

func (x *UploadRequest) GetMime() string {
	if x != nil {
		return x.Mime
	}
	return ""
}

func (x *UploadRequest) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type UploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_upload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_upload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_grpc_upload_proto_rawDescGZIP(), []int{1}
}

func (x *UploadResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_grpc_upload_proto protoreflect.FileDescriptor

var file_grpc_upload_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x0d, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x24, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x4a, 0x0a, 0x0d, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x06,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_upload_proto_rawDescOnce sync.Once
	file_grpc_upload_proto_rawDescData = file_grpc_upload_proto_rawDesc
)

func file_grpc_upload_proto_rawDescGZIP() []byte {
	file_grpc_upload_proto_rawDescOnce.Do(func() {
		file_grpc_upload_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_upload_proto_rawDescData)
	})
	return file_grpc_upload_proto_rawDescData
}

var file_grpc_upload_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_upload_proto_goTypes = []interface{}{
	(*UploadRequest)(nil),  // 0: proto.UploadRequest
	(*UploadResponse)(nil), // 1: proto.UploadResponse
}
var file_grpc_upload_proto_depIdxs = []int32{
	0, // 0: proto.UploadService.Upload:input_type -> proto.UploadRequest
	1, // 1: proto.UploadService.Upload:output_type -> proto.UploadResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_upload_proto_init() }
func file_grpc_upload_proto_init() {
	if File_grpc_upload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_upload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRequest); i {
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
		file_grpc_upload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadResponse); i {
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
			RawDescriptor: file_grpc_upload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_upload_proto_goTypes,
		DependencyIndexes: file_grpc_upload_proto_depIdxs,
		MessageInfos:      file_grpc_upload_proto_msgTypes,
	}.Build()
	File_grpc_upload_proto = out.File
	file_grpc_upload_proto_rawDesc = nil
	file_grpc_upload_proto_goTypes = nil
	file_grpc_upload_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UploadServiceClient is the client API for UploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UploadServiceClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (UploadService_UploadClient, error)
}

type uploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadServiceClient(cc grpc.ClientConnInterface) UploadServiceClient {
	return &uploadServiceClient{cc}
}

func (c *uploadServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (UploadService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UploadService_serviceDesc.Streams[0], "/proto.UploadService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &uploadServiceUploadClient{stream}
	return x, nil
}

type UploadService_UploadClient interface {
	Send(*UploadRequest) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type uploadServiceUploadClient struct {
	grpc.ClientStream
}

func (x *uploadServiceUploadClient) Send(m *UploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uploadServiceUploadClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UploadServiceServer is the server API for UploadService service.
type UploadServiceServer interface {
	Upload(UploadService_UploadServer) error
}

// UnimplementedUploadServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUploadServiceServer struct {
}

func (*UnimplementedUploadServiceServer) Upload(UploadService_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}

func RegisterUploadServiceServer(s *grpc.Server, srv UploadServiceServer) {
	s.RegisterService(&_UploadService_serviceDesc, srv)
}

func _UploadService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UploadServiceServer).Upload(&uploadServiceUploadServer{stream})
}

type UploadService_UploadServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*UploadRequest, error)
	grpc.ServerStream
}

type uploadServiceUploadServer struct {
	grpc.ServerStream
}

func (x *uploadServiceUploadServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uploadServiceUploadServer) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _UploadService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UploadService",
	HandlerType: (*UploadServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _UploadService_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/upload.proto",
}