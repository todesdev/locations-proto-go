// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: locationtypes.proto

package location_types

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LocationTypesServiceClient is the client API for LocationTypesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocationTypesServiceClient interface {
	GetAllLocationTypes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (LocationTypesService_GetAllLocationTypesClient, error)
	GetLocationTypeByPublicId(ctx context.Context, in *PublicIdRequest, opts ...grpc.CallOption) (*LocationTypeResponse, error)
	CreateLocationType(ctx context.Context, in *LocationTypeRequest, opts ...grpc.CallOption) (*LocationTypeResponse, error)
	UpdateLocationType(ctx context.Context, in *UpdateLocationTypeRequest, opts ...grpc.CallOption) (*LocationTypeResponse, error)
	DeleteLocationType(ctx context.Context, in *PublicIdRequest, opts ...grpc.CallOption) (*LocationTypeResponse, error)
}

type locationTypesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocationTypesServiceClient(cc grpc.ClientConnInterface) LocationTypesServiceClient {
	return &locationTypesServiceClient{cc}
}

func (c *locationTypesServiceClient) GetAllLocationTypes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (LocationTypesService_GetAllLocationTypesClient, error) {
	stream, err := c.cc.NewStream(ctx, &LocationTypesService_ServiceDesc.Streams[0], "/locationtypes.LocationTypesService/GetAllLocationTypes", opts...)
	if err != nil {
		return nil, err
	}
	x := &locationTypesServiceGetAllLocationTypesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LocationTypesService_GetAllLocationTypesClient interface {
	Recv() (*LocationTypeResponse, error)
	grpc.ClientStream
}

type locationTypesServiceGetAllLocationTypesClient struct {
	grpc.ClientStream
}

func (x *locationTypesServiceGetAllLocationTypesClient) Recv() (*LocationTypeResponse, error) {
	m := new(LocationTypeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *locationTypesServiceClient) GetLocationTypeByPublicId(ctx context.Context, in *PublicIdRequest, opts ...grpc.CallOption) (*LocationTypeResponse, error) {
	out := new(LocationTypeResponse)
	err := c.cc.Invoke(ctx, "/locationtypes.LocationTypesService/GetLocationTypeByPublicId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationTypesServiceClient) CreateLocationType(ctx context.Context, in *LocationTypeRequest, opts ...grpc.CallOption) (*LocationTypeResponse, error) {
	out := new(LocationTypeResponse)
	err := c.cc.Invoke(ctx, "/locationtypes.LocationTypesService/CreateLocationType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationTypesServiceClient) UpdateLocationType(ctx context.Context, in *UpdateLocationTypeRequest, opts ...grpc.CallOption) (*LocationTypeResponse, error) {
	out := new(LocationTypeResponse)
	err := c.cc.Invoke(ctx, "/locationtypes.LocationTypesService/UpdateLocationType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationTypesServiceClient) DeleteLocationType(ctx context.Context, in *PublicIdRequest, opts ...grpc.CallOption) (*LocationTypeResponse, error) {
	out := new(LocationTypeResponse)
	err := c.cc.Invoke(ctx, "/locationtypes.LocationTypesService/DeleteLocationType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocationTypesServiceServer is the server API for LocationTypesService service.
// All implementations must embed UnimplementedLocationTypesServiceServer
// for forward compatibility
type LocationTypesServiceServer interface {
	GetAllLocationTypes(*emptypb.Empty, LocationTypesService_GetAllLocationTypesServer) error
	GetLocationTypeByPublicId(context.Context, *PublicIdRequest) (*LocationTypeResponse, error)
	CreateLocationType(context.Context, *LocationTypeRequest) (*LocationTypeResponse, error)
	UpdateLocationType(context.Context, *UpdateLocationTypeRequest) (*LocationTypeResponse, error)
	DeleteLocationType(context.Context, *PublicIdRequest) (*LocationTypeResponse, error)
	mustEmbedUnimplementedLocationTypesServiceServer()
}

// UnimplementedLocationTypesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLocationTypesServiceServer struct {
}

func (UnimplementedLocationTypesServiceServer) GetAllLocationTypes(*emptypb.Empty, LocationTypesService_GetAllLocationTypesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllLocationTypes not implemented")
}
func (UnimplementedLocationTypesServiceServer) GetLocationTypeByPublicId(context.Context, *PublicIdRequest) (*LocationTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocationTypeByPublicId not implemented")
}
func (UnimplementedLocationTypesServiceServer) CreateLocationType(context.Context, *LocationTypeRequest) (*LocationTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLocationType not implemented")
}
func (UnimplementedLocationTypesServiceServer) UpdateLocationType(context.Context, *UpdateLocationTypeRequest) (*LocationTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLocationType not implemented")
}
func (UnimplementedLocationTypesServiceServer) DeleteLocationType(context.Context, *PublicIdRequest) (*LocationTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLocationType not implemented")
}
func (UnimplementedLocationTypesServiceServer) mustEmbedUnimplementedLocationTypesServiceServer() {}

// UnsafeLocationTypesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocationTypesServiceServer will
// result in compilation errors.
type UnsafeLocationTypesServiceServer interface {
	mustEmbedUnimplementedLocationTypesServiceServer()
}

func RegisterLocationTypesServiceServer(s grpc.ServiceRegistrar, srv LocationTypesServiceServer) {
	s.RegisterService(&LocationTypesService_ServiceDesc, srv)
}

func _LocationTypesService_GetAllLocationTypes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LocationTypesServiceServer).GetAllLocationTypes(m, &locationTypesServiceGetAllLocationTypesServer{stream})
}

type LocationTypesService_GetAllLocationTypesServer interface {
	Send(*LocationTypeResponse) error
	grpc.ServerStream
}

type locationTypesServiceGetAllLocationTypesServer struct {
	grpc.ServerStream
}

func (x *locationTypesServiceGetAllLocationTypesServer) Send(m *LocationTypeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _LocationTypesService_GetLocationTypeByPublicId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationTypesServiceServer).GetLocationTypeByPublicId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/locationtypes.LocationTypesService/GetLocationTypeByPublicId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationTypesServiceServer).GetLocationTypeByPublicId(ctx, req.(*PublicIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationTypesService_CreateLocationType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocationTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationTypesServiceServer).CreateLocationType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/locationtypes.LocationTypesService/CreateLocationType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationTypesServiceServer).CreateLocationType(ctx, req.(*LocationTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationTypesService_UpdateLocationType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLocationTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationTypesServiceServer).UpdateLocationType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/locationtypes.LocationTypesService/UpdateLocationType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationTypesServiceServer).UpdateLocationType(ctx, req.(*UpdateLocationTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationTypesService_DeleteLocationType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationTypesServiceServer).DeleteLocationType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/locationtypes.LocationTypesService/DeleteLocationType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationTypesServiceServer).DeleteLocationType(ctx, req.(*PublicIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LocationTypesService_ServiceDesc is the grpc.ServiceDesc for LocationTypesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocationTypesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "locationtypes.LocationTypesService",
	HandlerType: (*LocationTypesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLocationTypeByPublicId",
			Handler:    _LocationTypesService_GetLocationTypeByPublicId_Handler,
		},
		{
			MethodName: "CreateLocationType",
			Handler:    _LocationTypesService_CreateLocationType_Handler,
		},
		{
			MethodName: "UpdateLocationType",
			Handler:    _LocationTypesService_UpdateLocationType_Handler,
		},
		{
			MethodName: "DeleteLocationType",
			Handler:    _LocationTypesService_DeleteLocationType_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllLocationTypes",
			Handler:       _LocationTypesService_GetAllLocationTypes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "locationtypes.proto",
}
