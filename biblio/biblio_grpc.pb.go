// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package biblio

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// BiblioServiceClient is the client API for BiblioService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BiblioServiceClient interface {
	Chunks(ctx context.Context, in *Chunk, opts ...grpc.CallOption) (*Mensaje, error)
	UnChunks(ctx context.Context, in *Mensaje, opts ...grpc.CallOption) (*Chunk, error)
	UploadLibro(ctx context.Context, opts ...grpc.CallOption) (BiblioService_UploadLibroClient, error)
}

type biblioServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBiblioServiceClient(cc grpc.ClientConnInterface) BiblioServiceClient {
	return &biblioServiceClient{cc}
}

func (c *biblioServiceClient) Chunks(ctx context.Context, in *Chunk, opts ...grpc.CallOption) (*Mensaje, error) {
	out := new(Mensaje)
	err := c.cc.Invoke(ctx, "/BiblioService/Chunks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *biblioServiceClient) UnChunks(ctx context.Context, in *Mensaje, opts ...grpc.CallOption) (*Chunk, error) {
	out := new(Chunk)
	err := c.cc.Invoke(ctx, "/BiblioService/UnChunks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *biblioServiceClient) UploadLibro(ctx context.Context, opts ...grpc.CallOption) (BiblioService_UploadLibroClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BiblioService_serviceDesc.Streams[0], "/BiblioService/UploadLibro", opts...)
	if err != nil {
		return nil, err
	}
	x := &biblioServiceUploadLibroClient{stream}
	return x, nil
}

type BiblioService_UploadLibroClient interface {
	Send(*UploadLibroRequest) error
	CloseAndRecv() (*UploadLibroResponse, error)
	grpc.ClientStream
}

type biblioServiceUploadLibroClient struct {
	grpc.ClientStream
}

func (x *biblioServiceUploadLibroClient) Send(m *UploadLibroRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *biblioServiceUploadLibroClient) CloseAndRecv() (*UploadLibroResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadLibroResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BiblioServiceServer is the server API for BiblioService service.
// All implementations must embed UnimplementedBiblioServiceServer
// for forward compatibility
type BiblioServiceServer interface {
	Chunks(context.Context, *Chunk) (*Mensaje, error)
	UnChunks(context.Context, *Mensaje) (*Chunk, error)
	UploadLibro(BiblioService_UploadLibroServer) error
	mustEmbedUnimplementedBiblioServiceServer()
}

// UnimplementedBiblioServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBiblioServiceServer struct {
}

func (UnimplementedBiblioServiceServer) Chunks(context.Context, *Chunk) (*Mensaje, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Chunks not implemented")
}
func (UnimplementedBiblioServiceServer) UnChunks(context.Context, *Mensaje) (*Chunk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnChunks not implemented")
}
func (UnimplementedBiblioServiceServer) UploadLibro(BiblioService_UploadLibroServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadLibro not implemented")
}
func (UnimplementedBiblioServiceServer) mustEmbedUnimplementedBiblioServiceServer() {}

// UnsafeBiblioServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BiblioServiceServer will
// result in compilation errors.
type UnsafeBiblioServiceServer interface {
	mustEmbedUnimplementedBiblioServiceServer()
}

func RegisterBiblioServiceServer(s grpc.ServiceRegistrar, srv BiblioServiceServer) {
	s.RegisterService(&_BiblioService_serviceDesc, srv)
}

func _BiblioService_Chunks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chunk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BiblioServiceServer).Chunks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BiblioService/Chunks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BiblioServiceServer).Chunks(ctx, req.(*Chunk))
	}
	return interceptor(ctx, in, info, handler)
}

func _BiblioService_UnChunks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mensaje)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BiblioServiceServer).UnChunks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BiblioService/UnChunks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BiblioServiceServer).UnChunks(ctx, req.(*Mensaje))
	}
	return interceptor(ctx, in, info, handler)
}

func _BiblioService_UploadLibro_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BiblioServiceServer).UploadLibro(&biblioServiceUploadLibroServer{stream})
}

type BiblioService_UploadLibroServer interface {
	SendAndClose(*UploadLibroResponse) error
	Recv() (*UploadLibroRequest, error)
	grpc.ServerStream
}

type biblioServiceUploadLibroServer struct {
	grpc.ServerStream
}

func (x *biblioServiceUploadLibroServer) SendAndClose(m *UploadLibroResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *biblioServiceUploadLibroServer) Recv() (*UploadLibroRequest, error) {
	m := new(UploadLibroRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _BiblioService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "BiblioService",
	HandlerType: (*BiblioServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Chunks",
			Handler:    _BiblioService_Chunks_Handler,
		},
		{
			MethodName: "UnChunks",
			Handler:    _BiblioService_UnChunks_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadLibro",
			Handler:       _BiblioService_UploadLibro_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "biblio/biblio.proto",
}