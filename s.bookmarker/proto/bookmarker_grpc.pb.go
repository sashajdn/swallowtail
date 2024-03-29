// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: s.bookmarker/proto/bookmarker.proto

package bookmarkerproto

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BookmarkerClient is the client API for Bookmarker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookmarkerClient interface {
}

type bookmarkerClient struct {
	cc grpc.ClientConnInterface
}

func NewBookmarkerClient(cc grpc.ClientConnInterface) BookmarkerClient {
	return &bookmarkerClient{cc}
}

// BookmarkerServer is the server API for Bookmarker service.
// All implementations must embed UnimplementedBookmarkerServer
// for forward compatibility
type BookmarkerServer interface {
	mustEmbedUnimplementedBookmarkerServer()
}

// UnimplementedBookmarkerServer must be embedded to have forward compatible implementations.
type UnimplementedBookmarkerServer struct {
}

func (UnimplementedBookmarkerServer) mustEmbedUnimplementedBookmarkerServer() {}

// UnsafeBookmarkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookmarkerServer will
// result in compilation errors.
type UnsafeBookmarkerServer interface {
	mustEmbedUnimplementedBookmarkerServer()
}

func RegisterBookmarkerServer(s grpc.ServiceRegistrar, srv BookmarkerServer) {
	s.RegisterService(&Bookmarker_ServiceDesc, srv)
}

// Bookmarker_ServiceDesc is the grpc.ServiceDesc for Bookmarker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bookmarker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bookmarker",
	HandlerType: (*BookmarkerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "s.bookmarker/proto/bookmarker.proto",
}
