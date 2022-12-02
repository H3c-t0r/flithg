// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: flyteidl/service/dataproxy.proto

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DataProxyServiceClient is the client API for DataProxyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataProxyServiceClient interface {
	// CreateUploadLocation creates a signed url to upload artifacts to for a given project/domain.
	CreateUploadLocation(ctx context.Context, in *CreateUploadLocationRequest, opts ...grpc.CallOption) (*CreateUploadLocationResponse, error)
	// Deprecated: Do not use.
	// CreateDownloadLocation creates a signed url to download artifacts.
	CreateDownloadLocation(ctx context.Context, in *CreateDownloadLocationRequest, opts ...grpc.CallOption) (*CreateDownloadLocationResponse, error)
	// CreateDownloadLocation creates a signed url to download artifacts.
	CreateDownloadLink(ctx context.Context, in *CreateDownloadLinkRequest, opts ...grpc.CallOption) (*CreateDownloadLinkResponse, error)
}

type dataProxyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataProxyServiceClient(cc grpc.ClientConnInterface) DataProxyServiceClient {
	return &dataProxyServiceClient{cc}
}

func (c *dataProxyServiceClient) CreateUploadLocation(ctx context.Context, in *CreateUploadLocationRequest, opts ...grpc.CallOption) (*CreateUploadLocationResponse, error) {
	out := new(CreateUploadLocationResponse)
	err := c.cc.Invoke(ctx, "/flyteidl.service.DataProxyService/CreateUploadLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *dataProxyServiceClient) CreateDownloadLocation(ctx context.Context, in *CreateDownloadLocationRequest, opts ...grpc.CallOption) (*CreateDownloadLocationResponse, error) {
	out := new(CreateDownloadLocationResponse)
	err := c.cc.Invoke(ctx, "/flyteidl.service.DataProxyService/CreateDownloadLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataProxyServiceClient) CreateDownloadLink(ctx context.Context, in *CreateDownloadLinkRequest, opts ...grpc.CallOption) (*CreateDownloadLinkResponse, error) {
	out := new(CreateDownloadLinkResponse)
	err := c.cc.Invoke(ctx, "/flyteidl.service.DataProxyService/CreateDownloadLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataProxyServiceServer is the server API for DataProxyService service.
// All implementations should embed UnimplementedDataProxyServiceServer
// for forward compatibility
type DataProxyServiceServer interface {
	// CreateUploadLocation creates a signed url to upload artifacts to for a given project/domain.
	CreateUploadLocation(context.Context, *CreateUploadLocationRequest) (*CreateUploadLocationResponse, error)
	// Deprecated: Do not use.
	// CreateDownloadLocation creates a signed url to download artifacts.
	CreateDownloadLocation(context.Context, *CreateDownloadLocationRequest) (*CreateDownloadLocationResponse, error)
	// CreateDownloadLocation creates a signed url to download artifacts.
	CreateDownloadLink(context.Context, *CreateDownloadLinkRequest) (*CreateDownloadLinkResponse, error)
}

// UnimplementedDataProxyServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDataProxyServiceServer struct {
}

func (UnimplementedDataProxyServiceServer) CreateUploadLocation(context.Context, *CreateUploadLocationRequest) (*CreateUploadLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUploadLocation not implemented")
}
func (UnimplementedDataProxyServiceServer) CreateDownloadLocation(context.Context, *CreateDownloadLocationRequest) (*CreateDownloadLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDownloadLocation not implemented")
}
func (UnimplementedDataProxyServiceServer) CreateDownloadLink(context.Context, *CreateDownloadLinkRequest) (*CreateDownloadLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDownloadLink not implemented")
}

// UnsafeDataProxyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataProxyServiceServer will
// result in compilation errors.
type UnsafeDataProxyServiceServer interface {
	mustEmbedUnimplementedDataProxyServiceServer()
}

func RegisterDataProxyServiceServer(s grpc.ServiceRegistrar, srv DataProxyServiceServer) {
	s.RegisterService(&DataProxyService_ServiceDesc, srv)
}

func _DataProxyService_CreateUploadLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUploadLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataProxyServiceServer).CreateUploadLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.DataProxyService/CreateUploadLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataProxyServiceServer).CreateUploadLocation(ctx, req.(*CreateUploadLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataProxyService_CreateDownloadLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDownloadLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataProxyServiceServer).CreateDownloadLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.DataProxyService/CreateDownloadLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataProxyServiceServer).CreateDownloadLocation(ctx, req.(*CreateDownloadLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataProxyService_CreateDownloadLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDownloadLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataProxyServiceServer).CreateDownloadLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.DataProxyService/CreateDownloadLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataProxyServiceServer).CreateDownloadLink(ctx, req.(*CreateDownloadLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataProxyService_ServiceDesc is the grpc.ServiceDesc for DataProxyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataProxyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flyteidl.service.DataProxyService",
	HandlerType: (*DataProxyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUploadLocation",
			Handler:    _DataProxyService_CreateUploadLocation_Handler,
		},
		{
			MethodName: "CreateDownloadLocation",
			Handler:    _DataProxyService_CreateDownloadLocation_Handler,
		},
		{
			MethodName: "CreateDownloadLink",
			Handler:    _DataProxyService_CreateDownloadLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flyteidl/service/dataproxy.proto",
}
