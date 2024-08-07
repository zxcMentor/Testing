// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: rates.proto

package proto

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
	GetRates_Get_FullMethodName = "/proto.GetRates/Get"
)

// GetRatesClient is the client API for GetRates service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetRatesClient interface {
	Get(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Response, error)
}

type getRatesClient struct {
	cc grpc.ClientConnInterface
}

func NewGetRatesClient(cc grpc.ClientConnInterface) GetRatesClient {
	return &getRatesClient{cc}
}

func (c *getRatesClient) Get(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, GetRates_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetRatesServer is the server API for GetRates service.
// All implementations must embed UnimplementedGetRatesServer
// for forward compatibility
type GetRatesServer interface {
	Get(context.Context, *Req) (*Response, error)
	mustEmbedUnimplementedGetRatesServer()
}

// UnimplementedGetRatesServer must be embedded to have forward compatible implementations.
type UnimplementedGetRatesServer struct {
}

func (UnimplementedGetRatesServer) Get(context.Context, *Req) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedGetRatesServer) mustEmbedUnimplementedGetRatesServer() {}

// UnsafeGetRatesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetRatesServer will
// result in compilation errors.
type UnsafeGetRatesServer interface {
	mustEmbedUnimplementedGetRatesServer()
}

func RegisterGetRatesServer(s grpc.ServiceRegistrar, srv GetRatesServer) {
	s.RegisterService(&GetRates_ServiceDesc, srv)
}

func _GetRates_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetRatesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GetRates_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetRatesServer).Get(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

// GetRates_ServiceDesc is the grpc.ServiceDesc for GetRates service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetRates_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GetRates",
	HandlerType: (*GetRatesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _GetRates_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rates.proto",
}
