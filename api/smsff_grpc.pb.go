// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: api/smsff.proto

package api

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
	SmsFF_SendSms_FullMethodName = "/dnzzzsmsff.api.SmsFF/SendSms"
)

// SmsFFClient is the client API for SmsFF service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmsFFClient interface {
	SendSms(ctx context.Context, in *SendSmsReq, opts ...grpc.CallOption) (*SendSmsResp, error)
}

type smsFFClient struct {
	cc grpc.ClientConnInterface
}

func NewSmsFFClient(cc grpc.ClientConnInterface) SmsFFClient {
	return &smsFFClient{cc}
}

func (c *smsFFClient) SendSms(ctx context.Context, in *SendSmsReq, opts ...grpc.CallOption) (*SendSmsResp, error) {
	out := new(SendSmsResp)
	err := c.cc.Invoke(ctx, SmsFF_SendSms_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmsFFServer is the server API for SmsFF service.
// All implementations must embed UnimplementedSmsFFServer
// for forward compatibility
type SmsFFServer interface {
	SendSms(context.Context, *SendSmsReq) (*SendSmsResp, error)
	mustEmbedUnimplementedSmsFFServer()
}

// UnimplementedSmsFFServer must be embedded to have forward compatible implementations.
type UnimplementedSmsFFServer struct {
}

func (UnimplementedSmsFFServer) SendSms(context.Context, *SendSmsReq) (*SendSmsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSms not implemented")
}
func (UnimplementedSmsFFServer) mustEmbedUnimplementedSmsFFServer() {}

// UnsafeSmsFFServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmsFFServer will
// result in compilation errors.
type UnsafeSmsFFServer interface {
	mustEmbedUnimplementedSmsFFServer()
}

func RegisterSmsFFServer(s grpc.ServiceRegistrar, srv SmsFFServer) {
	s.RegisterService(&SmsFF_ServiceDesc, srv)
}

func _SmsFF_SendSms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSmsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsFFServer).SendSms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmsFF_SendSms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsFFServer).SendSms(ctx, req.(*SendSmsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SmsFF_ServiceDesc is the grpc.ServiceDesc for SmsFF service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SmsFF_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dnzzzsmsff.api.SmsFF",
	HandlerType: (*SmsFFServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSms",
			Handler:    _SmsFF_SendSms_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/smsff.proto",
}