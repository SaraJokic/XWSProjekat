// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: accommodation_service/accommodation_service.proto

package accommodation_service

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

const (
	AccommodationService_Get_FullMethodName                          = "/AccommodationService/Get"
	AccommodationService_GetByUserId_FullMethodName                  = "/AccommodationService/GetByUserId"
	AccommodationService_GetAll_FullMethodName                       = "/AccommodationService/GetAll"
	AccommodationService_GetAllProminentAccommodation_FullMethodName = "/AccommodationService/GetAllProminentAccommodation"
	AccommodationService_CreateNewAccommodation_FullMethodName       = "/AccommodationService/CreateNewAccommodation"
	AccommodationService_UpdateAccommodation_FullMethodName          = "/AccommodationService/UpdateAccommodation"
	AccommodationService_Search_FullMethodName                       = "/AccommodationService/Search"
)

// AccommodationServiceClient is the client API for AccommodationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccommodationServiceClient interface {
	Get(ctx context.Context, in *AccGetRequest, opts ...grpc.CallOption) (*AccGetResponse, error)
	GetByUserId(ctx context.Context, in *AccGetByUserIdRequest, opts ...grpc.CallOption) (*AccGetByUserIdResponse, error)
	GetAll(ctx context.Context, in *AccGetAllRequest, opts ...grpc.CallOption) (*AccGetAllResponse, error)
	GetAllProminentAccommodation(ctx context.Context, in *AccGetAllRequest, opts ...grpc.CallOption) (*AccGetAllResponse, error)
	CreateNewAccommodation(ctx context.Context, in *AccCreateRequest, opts ...grpc.CallOption) (*AccCreateResponse, error)
	UpdateAccommodation(ctx context.Context, in *UpdateAccommodationRequest, opts ...grpc.CallOption) (*UpdateAccommodationResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*AccGetAllResponse, error)
}

type accommodationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccommodationServiceClient(cc grpc.ClientConnInterface) AccommodationServiceClient {
	return &accommodationServiceClient{cc}
}

func (c *accommodationServiceClient) Get(ctx context.Context, in *AccGetRequest, opts ...grpc.CallOption) (*AccGetResponse, error) {
	out := new(AccGetResponse)
	err := c.cc.Invoke(ctx, AccommodationService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetByUserId(ctx context.Context, in *AccGetByUserIdRequest, opts ...grpc.CallOption) (*AccGetByUserIdResponse, error) {
	out := new(AccGetByUserIdResponse)
	err := c.cc.Invoke(ctx, AccommodationService_GetByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAll(ctx context.Context, in *AccGetAllRequest, opts ...grpc.CallOption) (*AccGetAllResponse, error) {
	out := new(AccGetAllResponse)
	err := c.cc.Invoke(ctx, AccommodationService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAllProminentAccommodation(ctx context.Context, in *AccGetAllRequest, opts ...grpc.CallOption) (*AccGetAllResponse, error) {
	out := new(AccGetAllResponse)
	err := c.cc.Invoke(ctx, AccommodationService_GetAllProminentAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) CreateNewAccommodation(ctx context.Context, in *AccCreateRequest, opts ...grpc.CallOption) (*AccCreateResponse, error) {
	out := new(AccCreateResponse)
	err := c.cc.Invoke(ctx, AccommodationService_CreateNewAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) UpdateAccommodation(ctx context.Context, in *UpdateAccommodationRequest, opts ...grpc.CallOption) (*UpdateAccommodationResponse, error) {
	out := new(UpdateAccommodationResponse)
	err := c.cc.Invoke(ctx, AccommodationService_UpdateAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*AccGetAllResponse, error) {
	out := new(AccGetAllResponse)
	err := c.cc.Invoke(ctx, AccommodationService_Search_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccommodationServiceServer is the server API for AccommodationService service.
// All implementations must embed UnimplementedAccommodationServiceServer
// for forward compatibility
type AccommodationServiceServer interface {
	Get(context.Context, *AccGetRequest) (*AccGetResponse, error)
	GetByUserId(context.Context, *AccGetByUserIdRequest) (*AccGetByUserIdResponse, error)
	GetAll(context.Context, *AccGetAllRequest) (*AccGetAllResponse, error)
	GetAllProminentAccommodation(context.Context, *AccGetAllRequest) (*AccGetAllResponse, error)
	CreateNewAccommodation(context.Context, *AccCreateRequest) (*AccCreateResponse, error)
	UpdateAccommodation(context.Context, *UpdateAccommodationRequest) (*UpdateAccommodationResponse, error)
	Search(context.Context, *SearchRequest) (*AccGetAllResponse, error)
	mustEmbedUnimplementedAccommodationServiceServer()
}

// UnimplementedAccommodationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccommodationServiceServer struct {
}

func (UnimplementedAccommodationServiceServer) Get(context.Context, *AccGetRequest) (*AccGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAccommodationServiceServer) GetByUserId(context.Context, *AccGetByUserIdRequest) (*AccGetByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUserId not implemented")
}
func (UnimplementedAccommodationServiceServer) GetAll(context.Context, *AccGetAllRequest) (*AccGetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedAccommodationServiceServer) GetAllProminentAccommodation(context.Context, *AccGetAllRequest) (*AccGetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProminentAccommodation not implemented")
}
func (UnimplementedAccommodationServiceServer) CreateNewAccommodation(context.Context, *AccCreateRequest) (*AccCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewAccommodation not implemented")
}
func (UnimplementedAccommodationServiceServer) UpdateAccommodation(context.Context, *UpdateAccommodationRequest) (*UpdateAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccommodation not implemented")
}
func (UnimplementedAccommodationServiceServer) Search(context.Context, *SearchRequest) (*AccGetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedAccommodationServiceServer) mustEmbedUnimplementedAccommodationServiceServer() {}

// UnsafeAccommodationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccommodationServiceServer will
// result in compilation errors.
type UnsafeAccommodationServiceServer interface {
	mustEmbedUnimplementedAccommodationServiceServer()
}

func RegisterAccommodationServiceServer(s grpc.ServiceRegistrar, srv AccommodationServiceServer) {
	s.RegisterService(&AccommodationService_ServiceDesc, srv)
}

func _AccommodationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).Get(ctx, req.(*AccGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccGetByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_GetByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetByUserId(ctx, req.(*AccGetByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccGetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAll(ctx, req.(*AccGetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAllProminentAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccGetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAllProminentAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_GetAllProminentAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAllProminentAccommodation(ctx, req.(*AccGetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_CreateNewAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).CreateNewAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_CreateNewAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).CreateNewAccommodation(ctx, req.(*AccCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_UpdateAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).UpdateAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_UpdateAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).UpdateAccommodation(ctx, req.(*UpdateAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_Search_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccommodationService_ServiceDesc is the grpc.ServiceDesc for AccommodationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccommodationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AccommodationService",
	HandlerType: (*AccommodationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _AccommodationService_Get_Handler,
		},
		{
			MethodName: "GetByUserId",
			Handler:    _AccommodationService_GetByUserId_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _AccommodationService_GetAll_Handler,
		},
		{
			MethodName: "GetAllProminentAccommodation",
			Handler:    _AccommodationService_GetAllProminentAccommodation_Handler,
		},
		{
			MethodName: "CreateNewAccommodation",
			Handler:    _AccommodationService_CreateNewAccommodation_Handler,
		},
		{
			MethodName: "UpdateAccommodation",
			Handler:    _AccommodationService_UpdateAccommodation_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _AccommodationService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accommodation_service/accommodation_service.proto",
}
