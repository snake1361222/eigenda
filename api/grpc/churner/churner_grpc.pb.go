// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: churner/churner.proto

package churner

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
	Churner_Churn_FullMethodName = "/churner.Churner/Churn"
)

// ChurnerClient is the client API for Churner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChurnerClient interface {
	Churn(ctx context.Context, in *ChurnRequest, opts ...grpc.CallOption) (*ChurnReply, error)
}

type churnerClient struct {
	cc grpc.ClientConnInterface
}

func NewChurnerClient(cc grpc.ClientConnInterface) ChurnerClient {
	return &churnerClient{cc}
}

func (c *churnerClient) Churn(ctx context.Context, in *ChurnRequest, opts ...grpc.CallOption) (*ChurnReply, error) {
	out := new(ChurnReply)
	err := c.cc.Invoke(ctx, Churner_Churn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChurnerServer is the server API for Churner service.
// All implementations must embed UnimplementedChurnerServer
// for forward compatibility
type ChurnerServer interface {
	Churn(context.Context, *ChurnRequest) (*ChurnReply, error)
	mustEmbedUnimplementedChurnerServer()
}

// UnimplementedChurnerServer must be embedded to have forward compatible implementations.
type UnimplementedChurnerServer struct {
}

func (UnimplementedChurnerServer) Churn(context.Context, *ChurnRequest) (*ChurnReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Churn not implemented")
}
func (UnimplementedChurnerServer) mustEmbedUnimplementedChurnerServer() {}

// UnsafeChurnerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChurnerServer will
// result in compilation errors.
type UnsafeChurnerServer interface {
	mustEmbedUnimplementedChurnerServer()
}

func RegisterChurnerServer(s grpc.ServiceRegistrar, srv ChurnerServer) {
	s.RegisterService(&Churner_ServiceDesc, srv)
}

func _Churner_Churn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChurnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChurnerServer).Churn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Churner_Churn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChurnerServer).Churn(ctx, req.(*ChurnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Churner_ServiceDesc is the grpc.ServiceDesc for Churner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Churner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "churner.Churner",
	HandlerType: (*ChurnerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Churn",
			Handler:    _Churner_Churn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "churner/churner.proto",
}