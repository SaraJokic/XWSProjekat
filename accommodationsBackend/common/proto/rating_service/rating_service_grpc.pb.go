// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: rating_service/rating_service.proto

package rating_service

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
	RatingService_GetRateAccommodation_FullMethodName                      = "/RatingService/GetRateAccommodation"
	RatingService_GetRateHost_FullMethodName                               = "/RatingService/GetRateHost"
	RatingService_GetHostRatingsByGuestId_FullMethodName                   = "/RatingService/GetHostRatingsByGuestId"
	RatingService_GetAccommodationsRatingsByGuestId_FullMethodName         = "/RatingService/GetAccommodationsRatingsByGuestId"
	RatingService_GetAccommodationsRatingsByAccommodationId_FullMethodName = "/RatingService/GetAccommodationsRatingsByAccommodationId"
	RatingService_GetHostRatingsByHostId_FullMethodName                    = "/RatingService/GetHostRatingsByHostId"
	RatingService_UpdateAccommodationRating_FullMethodName                 = "/RatingService/UpdateAccommodationRating"
	RatingService_UpdateHostRating_FullMethodName                          = "/RatingService/UpdateHostRating"
	RatingService_CreateNewHostRating_FullMethodName                       = "/RatingService/CreateNewHostRating"
	RatingService_CreateNewAccommodationRating_FullMethodName              = "/RatingService/CreateNewAccommodationRating"
	RatingService_GetAvgRatingHost_FullMethodName                          = "/RatingService/GetAvgRatingHost"
	RatingService_GetAvgAccommodationRating_FullMethodName                 = "/RatingService/GetAvgAccommodationRating"
)

// RatingServiceClient is the client API for RatingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RatingServiceClient interface {
	GetRateAccommodation(ctx context.Context, in *GetRateAccommodationRequest, opts ...grpc.CallOption) (*GetRateAccommodationResponse, error)
	GetRateHost(ctx context.Context, in *GetRateHostRequest, opts ...grpc.CallOption) (*GetRateHostResponse, error)
	GetHostRatingsByGuestId(ctx context.Context, in *GetRateHostByGuestRequest, opts ...grpc.CallOption) (*GetRateHostByGuestResponse, error)
	GetAccommodationsRatingsByGuestId(ctx context.Context, in *GetRateAccommodationByGuestRequest, opts ...grpc.CallOption) (*GetRateAccommodationByGuestResponse, error)
	GetAccommodationsRatingsByAccommodationId(ctx context.Context, in *GetRateAccommodationByAccommodationRequest, opts ...grpc.CallOption) (*GetRateAccommodationByAccommodationResponse, error)
	GetHostRatingsByHostId(ctx context.Context, in *GetRateHostByHostRequest, opts ...grpc.CallOption) (*GetRateHostByHostResponse, error)
	UpdateAccommodationRating(ctx context.Context, in *UpdateAccommodationRatingRequest, opts ...grpc.CallOption) (*UpdateAccommodationRatingResponse, error)
	UpdateHostRating(ctx context.Context, in *UpdateHostRatingRequest, opts ...grpc.CallOption) (*UpdateHostRatingResponse, error)
	CreateNewHostRating(ctx context.Context, in *CreateNewHostRatingRequest, opts ...grpc.CallOption) (*CreateNewHostRatingResponse, error)
	CreateNewAccommodationRating(ctx context.Context, in *CreateNewAccommodationRatingRequest, opts ...grpc.CallOption) (*CreateNewAccommodationRatingResponse, error)
	GetAvgRatingHost(ctx context.Context, in *GetAvgHostRatingRequest, opts ...grpc.CallOption) (*GetAvgHostRatingResponse, error)
	GetAvgAccommodationRating(ctx context.Context, in *GetAvgAccommodationRatingRequest, opts ...grpc.CallOption) (*GetAvgAccommodationRatingResponse, error)
}

type ratingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRatingServiceClient(cc grpc.ClientConnInterface) RatingServiceClient {
	return &ratingServiceClient{cc}
}

func (c *ratingServiceClient) GetRateAccommodation(ctx context.Context, in *GetRateAccommodationRequest, opts ...grpc.CallOption) (*GetRateAccommodationResponse, error) {
	out := new(GetRateAccommodationResponse)
	err := c.cc.Invoke(ctx, RatingService_GetRateAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetRateHost(ctx context.Context, in *GetRateHostRequest, opts ...grpc.CallOption) (*GetRateHostResponse, error) {
	out := new(GetRateHostResponse)
	err := c.cc.Invoke(ctx, RatingService_GetRateHost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetHostRatingsByGuestId(ctx context.Context, in *GetRateHostByGuestRequest, opts ...grpc.CallOption) (*GetRateHostByGuestResponse, error) {
	out := new(GetRateHostByGuestResponse)
	err := c.cc.Invoke(ctx, RatingService_GetHostRatingsByGuestId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetAccommodationsRatingsByGuestId(ctx context.Context, in *GetRateAccommodationByGuestRequest, opts ...grpc.CallOption) (*GetRateAccommodationByGuestResponse, error) {
	out := new(GetRateAccommodationByGuestResponse)
	err := c.cc.Invoke(ctx, RatingService_GetAccommodationsRatingsByGuestId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetAccommodationsRatingsByAccommodationId(ctx context.Context, in *GetRateAccommodationByAccommodationRequest, opts ...grpc.CallOption) (*GetRateAccommodationByAccommodationResponse, error) {
	out := new(GetRateAccommodationByAccommodationResponse)
	err := c.cc.Invoke(ctx, RatingService_GetAccommodationsRatingsByAccommodationId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetHostRatingsByHostId(ctx context.Context, in *GetRateHostByHostRequest, opts ...grpc.CallOption) (*GetRateHostByHostResponse, error) {
	out := new(GetRateHostByHostResponse)
	err := c.cc.Invoke(ctx, RatingService_GetHostRatingsByHostId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) UpdateAccommodationRating(ctx context.Context, in *UpdateAccommodationRatingRequest, opts ...grpc.CallOption) (*UpdateAccommodationRatingResponse, error) {
	out := new(UpdateAccommodationRatingResponse)
	err := c.cc.Invoke(ctx, RatingService_UpdateAccommodationRating_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) UpdateHostRating(ctx context.Context, in *UpdateHostRatingRequest, opts ...grpc.CallOption) (*UpdateHostRatingResponse, error) {
	out := new(UpdateHostRatingResponse)
	err := c.cc.Invoke(ctx, RatingService_UpdateHostRating_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) CreateNewHostRating(ctx context.Context, in *CreateNewHostRatingRequest, opts ...grpc.CallOption) (*CreateNewHostRatingResponse, error) {
	out := new(CreateNewHostRatingResponse)
	err := c.cc.Invoke(ctx, RatingService_CreateNewHostRating_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) CreateNewAccommodationRating(ctx context.Context, in *CreateNewAccommodationRatingRequest, opts ...grpc.CallOption) (*CreateNewAccommodationRatingResponse, error) {
	out := new(CreateNewAccommodationRatingResponse)
	err := c.cc.Invoke(ctx, RatingService_CreateNewAccommodationRating_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetAvgRatingHost(ctx context.Context, in *GetAvgHostRatingRequest, opts ...grpc.CallOption) (*GetAvgHostRatingResponse, error) {
	out := new(GetAvgHostRatingResponse)
	err := c.cc.Invoke(ctx, RatingService_GetAvgRatingHost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetAvgAccommodationRating(ctx context.Context, in *GetAvgAccommodationRatingRequest, opts ...grpc.CallOption) (*GetAvgAccommodationRatingResponse, error) {
	out := new(GetAvgAccommodationRatingResponse)
	err := c.cc.Invoke(ctx, RatingService_GetAvgAccommodationRating_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingServiceServer is the server API for RatingService service.
// All implementations must embed UnimplementedRatingServiceServer
// for forward compatibility
type RatingServiceServer interface {
	GetRateAccommodation(context.Context, *GetRateAccommodationRequest) (*GetRateAccommodationResponse, error)
	GetRateHost(context.Context, *GetRateHostRequest) (*GetRateHostResponse, error)
	GetHostRatingsByGuestId(context.Context, *GetRateHostByGuestRequest) (*GetRateHostByGuestResponse, error)
	GetAccommodationsRatingsByGuestId(context.Context, *GetRateAccommodationByGuestRequest) (*GetRateAccommodationByGuestResponse, error)
	GetAccommodationsRatingsByAccommodationId(context.Context, *GetRateAccommodationByAccommodationRequest) (*GetRateAccommodationByAccommodationResponse, error)
	GetHostRatingsByHostId(context.Context, *GetRateHostByHostRequest) (*GetRateHostByHostResponse, error)
	UpdateAccommodationRating(context.Context, *UpdateAccommodationRatingRequest) (*UpdateAccommodationRatingResponse, error)
	UpdateHostRating(context.Context, *UpdateHostRatingRequest) (*UpdateHostRatingResponse, error)
	CreateNewHostRating(context.Context, *CreateNewHostRatingRequest) (*CreateNewHostRatingResponse, error)
	CreateNewAccommodationRating(context.Context, *CreateNewAccommodationRatingRequest) (*CreateNewAccommodationRatingResponse, error)
	GetAvgRatingHost(context.Context, *GetAvgHostRatingRequest) (*GetAvgHostRatingResponse, error)
	GetAvgAccommodationRating(context.Context, *GetAvgAccommodationRatingRequest) (*GetAvgAccommodationRatingResponse, error)
	mustEmbedUnimplementedRatingServiceServer()
}

// UnimplementedRatingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRatingServiceServer struct {
}

func (UnimplementedRatingServiceServer) GetRateAccommodation(context.Context, *GetRateAccommodationRequest) (*GetRateAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRateAccommodation not implemented")
}
func (UnimplementedRatingServiceServer) GetRateHost(context.Context, *GetRateHostRequest) (*GetRateHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRateHost not implemented")
}
func (UnimplementedRatingServiceServer) GetHostRatingsByGuestId(context.Context, *GetRateHostByGuestRequest) (*GetRateHostByGuestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHostRatingsByGuestId not implemented")
}
func (UnimplementedRatingServiceServer) GetAccommodationsRatingsByGuestId(context.Context, *GetRateAccommodationByGuestRequest) (*GetRateAccommodationByGuestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccommodationsRatingsByGuestId not implemented")
}
func (UnimplementedRatingServiceServer) GetAccommodationsRatingsByAccommodationId(context.Context, *GetRateAccommodationByAccommodationRequest) (*GetRateAccommodationByAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccommodationsRatingsByAccommodationId not implemented")
}
func (UnimplementedRatingServiceServer) GetHostRatingsByHostId(context.Context, *GetRateHostByHostRequest) (*GetRateHostByHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHostRatingsByHostId not implemented")
}
func (UnimplementedRatingServiceServer) UpdateAccommodationRating(context.Context, *UpdateAccommodationRatingRequest) (*UpdateAccommodationRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccommodationRating not implemented")
}
func (UnimplementedRatingServiceServer) UpdateHostRating(context.Context, *UpdateHostRatingRequest) (*UpdateHostRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHostRating not implemented")
}
func (UnimplementedRatingServiceServer) CreateNewHostRating(context.Context, *CreateNewHostRatingRequest) (*CreateNewHostRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewHostRating not implemented")
}
func (UnimplementedRatingServiceServer) CreateNewAccommodationRating(context.Context, *CreateNewAccommodationRatingRequest) (*CreateNewAccommodationRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewAccommodationRating not implemented")
}
func (UnimplementedRatingServiceServer) GetAvgRatingHost(context.Context, *GetAvgHostRatingRequest) (*GetAvgHostRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvgRatingHost not implemented")
}
func (UnimplementedRatingServiceServer) GetAvgAccommodationRating(context.Context, *GetAvgAccommodationRatingRequest) (*GetAvgAccommodationRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvgAccommodationRating not implemented")
}
func (UnimplementedRatingServiceServer) mustEmbedUnimplementedRatingServiceServer() {}

// UnsafeRatingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RatingServiceServer will
// result in compilation errors.
type UnsafeRatingServiceServer interface {
	mustEmbedUnimplementedRatingServiceServer()
}

func RegisterRatingServiceServer(s grpc.ServiceRegistrar, srv RatingServiceServer) {
	s.RegisterService(&RatingService_ServiceDesc, srv)
}

func _RatingService_GetRateAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRateAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetRateAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_GetRateAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetRateAccommodation(ctx, req.(*GetRateAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetRateHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRateHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetRateHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_GetRateHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetRateHost(ctx, req.(*GetRateHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetHostRatingsByGuestId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRateHostByGuestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetHostRatingsByGuestId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_GetHostRatingsByGuestId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetHostRatingsByGuestId(ctx, req.(*GetRateHostByGuestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetAccommodationsRatingsByGuestId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRateAccommodationByGuestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetAccommodationsRatingsByGuestId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_GetAccommodationsRatingsByGuestId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetAccommodationsRatingsByGuestId(ctx, req.(*GetRateAccommodationByGuestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetAccommodationsRatingsByAccommodationId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRateAccommodationByAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetAccommodationsRatingsByAccommodationId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_GetAccommodationsRatingsByAccommodationId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetAccommodationsRatingsByAccommodationId(ctx, req.(*GetRateAccommodationByAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetHostRatingsByHostId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRateHostByHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetHostRatingsByHostId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_GetHostRatingsByHostId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetHostRatingsByHostId(ctx, req.(*GetRateHostByHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_UpdateAccommodationRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccommodationRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).UpdateAccommodationRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_UpdateAccommodationRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).UpdateAccommodationRating(ctx, req.(*UpdateAccommodationRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_UpdateHostRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHostRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).UpdateHostRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_UpdateHostRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).UpdateHostRating(ctx, req.(*UpdateHostRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_CreateNewHostRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNewHostRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).CreateNewHostRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_CreateNewHostRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).CreateNewHostRating(ctx, req.(*CreateNewHostRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_CreateNewAccommodationRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNewAccommodationRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).CreateNewAccommodationRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_CreateNewAccommodationRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).CreateNewAccommodationRating(ctx, req.(*CreateNewAccommodationRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetAvgRatingHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvgHostRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetAvgRatingHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_GetAvgRatingHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetAvgRatingHost(ctx, req.(*GetAvgHostRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetAvgAccommodationRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvgAccommodationRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetAvgAccommodationRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_GetAvgAccommodationRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetAvgAccommodationRating(ctx, req.(*GetAvgAccommodationRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RatingService_ServiceDesc is the grpc.ServiceDesc for RatingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RatingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RatingService",
	HandlerType: (*RatingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRateAccommodation",
			Handler:    _RatingService_GetRateAccommodation_Handler,
		},
		{
			MethodName: "GetRateHost",
			Handler:    _RatingService_GetRateHost_Handler,
		},
		{
			MethodName: "GetHostRatingsByGuestId",
			Handler:    _RatingService_GetHostRatingsByGuestId_Handler,
		},
		{
			MethodName: "GetAccommodationsRatingsByGuestId",
			Handler:    _RatingService_GetAccommodationsRatingsByGuestId_Handler,
		},
		{
			MethodName: "GetAccommodationsRatingsByAccommodationId",
			Handler:    _RatingService_GetAccommodationsRatingsByAccommodationId_Handler,
		},
		{
			MethodName: "GetHostRatingsByHostId",
			Handler:    _RatingService_GetHostRatingsByHostId_Handler,
		},
		{
			MethodName: "UpdateAccommodationRating",
			Handler:    _RatingService_UpdateAccommodationRating_Handler,
		},
		{
			MethodName: "UpdateHostRating",
			Handler:    _RatingService_UpdateHostRating_Handler,
		},
		{
			MethodName: "CreateNewHostRating",
			Handler:    _RatingService_CreateNewHostRating_Handler,
		},
		{
			MethodName: "CreateNewAccommodationRating",
			Handler:    _RatingService_CreateNewAccommodationRating_Handler,
		},
		{
			MethodName: "GetAvgRatingHost",
			Handler:    _RatingService_GetAvgRatingHost_Handler,
		},
		{
			MethodName: "GetAvgAccommodationRating",
			Handler:    _RatingService_GetAvgAccommodationRating_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rating_service/rating_service.proto",
}
