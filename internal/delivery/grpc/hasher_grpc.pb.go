// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: hasher.proto

package grpc

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
	Hasher_GetHashes_FullMethodName = "/hasher.Hasher/GetHashes"
)

// HasherClient is the client API for Hasher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HasherClient interface {
	GetHashes(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*HashResponse, error)
}

type hasherClient struct {
	cc grpc.ClientConnInterface
}

func NewHasherClient(cc grpc.ClientConnInterface) HasherClient {
	return &hasherClient{cc}
}

func (c *hasherClient) GetHashes(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*HashResponse, error) {
	out := new(HashResponse)
	err := c.cc.Invoke(ctx, Hasher_GetHashes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HasherServer is the server API for Hasher service.
// All implementations must embed UnimplementedHasherServer
// for forward compatibility
type HasherServer interface {
	GetHashes(context.Context, *HashRequest) (*HashResponse, error)
	mustEmbedUnimplementedHasherServer()
}

// UnimplementedHasherServer must be embedded to have forward compatible implementations.
type UnimplementedHasherServer struct {
}

func (UnimplementedHasherServer) GetHashes(context.Context, *HashRequest) (*HashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHashes not implemented")
}
func (UnimplementedHasherServer) mustEmbedUnimplementedHasherServer() {}

// UnsafeHasherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HasherServer will
// result in compilation errors.
type UnsafeHasherServer interface {
	mustEmbedUnimplementedHasherServer()
}

func RegisterHasherServer(s grpc.ServiceRegistrar, srv HasherServer) {
	s.RegisterService(&Hasher_ServiceDesc, srv)
}

func _Hasher_GetHashes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HasherServer).GetHashes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hasher_GetHashes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HasherServer).GetHashes(ctx, req.(*HashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Hasher_ServiceDesc is the grpc.ServiceDesc for Hasher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hasher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hasher.Hasher",
	HandlerType: (*HasherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHashes",
			Handler:    _Hasher_GetHashes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hasher.proto",
}
