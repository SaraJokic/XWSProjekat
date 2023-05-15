// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: availability-service/availability_service.proto

package availability_service

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
	AvailabilityService_Get_FullMethodName                   = "/AvailabilityService/Get"
	AvailabilityService_GetAll_FullMethodName                = "/AvailabilityService/GetAll"
	AvailabilityService_CreatePriceChange_FullMethodName     = "/AvailabilityService/CreatePriceChange"
	AvailabilityService_AddAvailableSlot_FullMethodName      = "/AvailabilityService/AddAvailableSlot"
	AvailabilityService_GetByAccommodationId_FullMethodName  = "/AvailabilityService/GetByAccommodationId"
	AvailabilityService_UpdateAvailableSlot_FullMethodName   = "/AvailabilityService/UpdateAvailableSlot"
	AvailabilityService_CreateNewAvailability_FullMethodName = "/AvailabilityService/CreateNewAvailability"
)

// AvailabilityServiceClient is the client API for AvailabilityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AvailabilityServiceClient interface {
	Get(ctx context.Context, in *GetAvailableRequest, opts ...grpc.CallOption) (*GetAvailableResponse, error)
	GetAll(ctx context.Context, in *GetAllAvailableRequest, opts ...grpc.CallOption) (*GetAllAvailableResponse, error)
	CreatePriceChange(ctx context.Context, in *CreatePriceChangeRequest, opts ...grpc.CallOption) (*CreatePriceChangeResponse, error)
	AddAvailableSlot(ctx context.Context, in *AddAvailableSlotRequest, opts ...grpc.CallOption) (*AddAvailableSlotResponse, error)
	GetByAccommodationId(ctx context.Context, in *GetByAccIdRequest, opts ...grpc.CallOption) (*GetByAccIdResponse, error)
	UpdateAvailableSlot(ctx context.Context, in *UpdateAvailableSlotRequest, opts ...grpc.CallOption) (*UpdateAvailableSlotResponse, error)
	CreateNewAvailability(ctx context.Context, in *CreateNewAvailabilityRequest, opts ...grpc.CallOption) (*CreateNewAvailabilityResponse, error)
}

type availabilityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAvailabilityServiceClient(cc grpc.ClientConnInterface) AvailabilityServiceClient {
	return &availabilityServiceClient{cc}
}

func (c *availabilityServiceClient) Get(ctx context.Context, in *GetAvailableRequest, opts ...grpc.CallOption) (*GetAvailableResponse, error) {
	out := new(GetAvailableResponse)
	err := c.cc.Invoke(ctx, AvailabilityService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *availabilityServiceClient) GetAll(ctx context.Context, in *GetAllAvailableRequest, opts ...grpc.CallOption) (*GetAllAvailableResponse, error) {
	out := new(GetAllAvailableResponse)
	err := c.cc.Invoke(ctx, AvailabilityService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *availabilityServiceClient) CreatePriceChange(ctx context.Context, in *CreatePriceChangeRequest, opts ...grpc.CallOption) (*CreatePriceChangeResponse, error) {
	out := new(CreatePriceChangeResponse)
	err := c.cc.Invoke(ctx, AvailabilityService_CreatePriceChange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *availabilityServiceClient) AddAvailableSlot(ctx context.Context, in *AddAvailableSlotRequest, opts ...grpc.CallOption) (*AddAvailableSlotResponse, error) {
	out := new(AddAvailableSlotResponse)
	err := c.cc.Invoke(ctx, AvailabilityService_AddAvailableSlot_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *availabilityServiceClient) GetByAccommodationId(ctx context.Context, in *GetByAccIdRequest, opts ...grpc.CallOption) (*GetByAccIdResponse, error) {
	out := new(GetByAccIdResponse)
	err := c.cc.Invoke(ctx, AvailabilityService_GetByAccommodationId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *availabilityServiceClient) UpdateAvailableSlot(ctx context.Context, in *UpdateAvailableSlotRequest, opts ...grpc.CallOption) (*UpdateAvailableSlotResponse, error) {
	out := new(UpdateAvailableSlotResponse)
	err := c.cc.Invoke(ctx, AvailabilityService_UpdateAvailableSlot_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *availabilityServiceClient) CreateNewAvailability(ctx context.Context, in *CreateNewAvailabilityRequest, opts ...grpc.CallOption) (*CreateNewAvailabilityResponse, error) {
	out := new(CreateNewAvailabilityResponse)
	err := c.cc.Invoke(ctx, AvailabilityService_CreateNewAvailability_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AvailabilityServiceServer is the server API for AvailabilityService service.
// All implementations must embed UnimplementedAvailabilityServiceServer
// for forward compatibility
type AvailabilityServiceServer interface {
	Get(context.Context, *GetAvailableRequest) (*GetAvailableResponse, error)
	GetAll(context.Context, *GetAllAvailableRequest) (*GetAllAvailableResponse, error)
	CreatePriceChange(context.Context, *CreatePriceChangeRequest) (*CreatePriceChangeResponse, error)
	AddAvailableSlot(context.Context, *AddAvailableSlotRequest) (*AddAvailableSlotResponse, error)
	GetByAccommodationId(context.Context, *GetByAccIdRequest) (*GetByAccIdResponse, error)
	UpdateAvailableSlot(context.Context, *UpdateAvailableSlotRequest) (*UpdateAvailableSlotResponse, error)
	CreateNewAvailability(context.Context, *CreateNewAvailabilityRequest) (*CreateNewAvailabilityResponse, error)
	mustEmbedUnimplementedAvailabilityServiceServer()
}

// UnimplementedAvailabilityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAvailabilityServiceServer struct {
}

func (UnimplementedAvailabilityServiceServer) Get(context.Context, *GetAvailableRequest) (*GetAvailableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAvailabilityServiceServer) GetAll(context.Context, *GetAllAvailableRequest) (*GetAllAvailableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedAvailabilityServiceServer) CreatePriceChange(context.Context, *CreatePriceChangeRequest) (*CreatePriceChangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePriceChange not implemented")
}
func (UnimplementedAvailabilityServiceServer) AddAvailableSlot(context.Context, *AddAvailableSlotRequest) (*AddAvailableSlotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAvailableSlot not implemented")
}
func (UnimplementedAvailabilityServiceServer) GetByAccommodationId(context.Context, *GetByAccIdRequest) (*GetByAccIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByAccommodationId not implemented")
}
func (UnimplementedAvailabilityServiceServer) UpdateAvailableSlot(context.Context, *UpdateAvailableSlotRequest) (*UpdateAvailableSlotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAvailableSlot not implemented")
}
func (UnimplementedAvailabilityServiceServer) CreateNewAvailability(context.Context, *CreateNewAvailabilityRequest) (*CreateNewAvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewAvailability not implemented")
}
func (UnimplementedAvailabilityServiceServer) mustEmbedUnimplementedAvailabilityServiceServer() {}

// UnsafeAvailabilityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AvailabilityServiceServer will
// result in compilation errors.
type UnsafeAvailabilityServiceServer interface {
	mustEmbedUnimplementedAvailabilityServiceServer()
}

func RegisterAvailabilityServiceServer(s grpc.ServiceRegistrar, srv AvailabilityServiceServer) {
	s.RegisterService(&AvailabilityService_ServiceDesc, srv)
}

func _AvailabilityService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvailableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvailabilityServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AvailabilityService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvailabilityServiceServer).Get(ctx, req.(*GetAvailableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AvailabilityService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllAvailableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvailabilityServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AvailabilityService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvailabilityServiceServer).GetAll(ctx, req.(*GetAllAvailableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AvailabilityService_CreatePriceChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePriceChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvailabilityServiceServer).CreatePriceChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AvailabilityService_CreatePriceChange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvailabilityServiceServer).CreatePriceChange(ctx, req.(*CreatePriceChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AvailabilityService_AddAvailableSlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAvailableSlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvailabilityServiceServer).AddAvailableSlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AvailabilityService_AddAvailableSlot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvailabilityServiceServer).AddAvailableSlot(ctx, req.(*AddAvailableSlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AvailabilityService_GetByAccommodationId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByAccIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvailabilityServiceServer).GetByAccommodationId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AvailabilityService_GetByAccommodationId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvailabilityServiceServer).GetByAccommodationId(ctx, req.(*GetByAccIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AvailabilityService_UpdateAvailableSlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAvailableSlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvailabilityServiceServer).UpdateAvailableSlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AvailabilityService_UpdateAvailableSlot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvailabilityServiceServer).UpdateAvailableSlot(ctx, req.(*UpdateAvailableSlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AvailabilityService_CreateNewAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNewAvailabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvailabilityServiceServer).CreateNewAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AvailabilityService_CreateNewAvailability_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvailabilityServiceServer).CreateNewAvailability(ctx, req.(*CreateNewAvailabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AvailabilityService_ServiceDesc is the grpc.ServiceDesc for AvailabilityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AvailabilityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AvailabilityService",
	HandlerType: (*AvailabilityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _AvailabilityService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _AvailabilityService_GetAll_Handler,
		},
		{
			MethodName: "CreatePriceChange",
			Handler:    _AvailabilityService_CreatePriceChange_Handler,
		},
		{
			MethodName: "AddAvailableSlot",
			Handler:    _AvailabilityService_AddAvailableSlot_Handler,
		},
		{
			MethodName: "GetByAccommodationId",
			Handler:    _AvailabilityService_GetByAccommodationId_Handler,
		},
		{
			MethodName: "UpdateAvailableSlot",
			Handler:    _AvailabilityService_UpdateAvailableSlot_Handler,
		},
		{
			MethodName: "CreateNewAvailability",
			Handler:    _AvailabilityService_CreateNewAvailability_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "availability-service/availability_service.proto",
}