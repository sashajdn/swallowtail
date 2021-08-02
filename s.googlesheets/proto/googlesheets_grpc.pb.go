// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package googlesheetsproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GooglesheetsClient is the client API for Googlesheets service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GooglesheetsClient interface {
	CreatePortfolioSheet(ctx context.Context, in *CreatePortfolioSheetRequest, opts ...grpc.CallOption) (*CreatePortfolioSheetResponse, error)
	RegisterNewPortfolioSheet(ctx context.Context, in *RegisterNewPortfolioSheetRequest, opts ...grpc.CallOption) (*RegisterNewPortfolioSheetResponse, error)
	TmpGetLatestPriceBySymbol(ctx context.Context, in *TmpGetLatestPriceBySymbolRequest, opts ...grpc.CallOption) (*TmpGetLatestPriceBySymbolResponse, error)
}

type googlesheetsClient struct {
	cc grpc.ClientConnInterface
}

func NewGooglesheetsClient(cc grpc.ClientConnInterface) GooglesheetsClient {
	return &googlesheetsClient{cc}
}

func (c *googlesheetsClient) CreatePortfolioSheet(ctx context.Context, in *CreatePortfolioSheetRequest, opts ...grpc.CallOption) (*CreatePortfolioSheetResponse, error) {
	out := new(CreatePortfolioSheetResponse)
	err := c.cc.Invoke(ctx, "/googlesheets/CreatePortfolioSheet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googlesheetsClient) RegisterNewPortfolioSheet(ctx context.Context, in *RegisterNewPortfolioSheetRequest, opts ...grpc.CallOption) (*RegisterNewPortfolioSheetResponse, error) {
	out := new(RegisterNewPortfolioSheetResponse)
	err := c.cc.Invoke(ctx, "/googlesheets/RegisterNewPortfolioSheet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googlesheetsClient) TmpGetLatestPriceBySymbol(ctx context.Context, in *TmpGetLatestPriceBySymbolRequest, opts ...grpc.CallOption) (*TmpGetLatestPriceBySymbolResponse, error) {
	out := new(TmpGetLatestPriceBySymbolResponse)
	err := c.cc.Invoke(ctx, "/googlesheets/TmpGetLatestPriceBySymbol", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GooglesheetsServer is the server API for Googlesheets service.
// All implementations must embed UnimplementedGooglesheetsServer
// for forward compatibility
type GooglesheetsServer interface {
	CreatePortfolioSheet(context.Context, *CreatePortfolioSheetRequest) (*CreatePortfolioSheetResponse, error)
	RegisterNewPortfolioSheet(context.Context, *RegisterNewPortfolioSheetRequest) (*RegisterNewPortfolioSheetResponse, error)
	TmpGetLatestPriceBySymbol(context.Context, *TmpGetLatestPriceBySymbolRequest) (*TmpGetLatestPriceBySymbolResponse, error)
	mustEmbedUnimplementedGooglesheetsServer()
}

// UnimplementedGooglesheetsServer must be embedded to have forward compatible implementations.
type UnimplementedGooglesheetsServer struct {
}

func (*UnimplementedGooglesheetsServer) CreatePortfolioSheet(context.Context, *CreatePortfolioSheetRequest) (*CreatePortfolioSheetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePortfolioSheet not implemented")
}
func (*UnimplementedGooglesheetsServer) RegisterNewPortfolioSheet(context.Context, *RegisterNewPortfolioSheetRequest) (*RegisterNewPortfolioSheetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNewPortfolioSheet not implemented")
}
func (*UnimplementedGooglesheetsServer) TmpGetLatestPriceBySymbol(context.Context, *TmpGetLatestPriceBySymbolRequest) (*TmpGetLatestPriceBySymbolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TmpGetLatestPriceBySymbol not implemented")
}
func (*UnimplementedGooglesheetsServer) mustEmbedUnimplementedGooglesheetsServer() {}

func RegisterGooglesheetsServer(s *grpc.Server, srv GooglesheetsServer) {
	s.RegisterService(&_Googlesheets_serviceDesc, srv)
}

func _Googlesheets_CreatePortfolioSheet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePortfolioSheetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GooglesheetsServer).CreatePortfolioSheet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/googlesheets/CreatePortfolioSheet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GooglesheetsServer).CreatePortfolioSheet(ctx, req.(*CreatePortfolioSheetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Googlesheets_RegisterNewPortfolioSheet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterNewPortfolioSheetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GooglesheetsServer).RegisterNewPortfolioSheet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/googlesheets/RegisterNewPortfolioSheet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GooglesheetsServer).RegisterNewPortfolioSheet(ctx, req.(*RegisterNewPortfolioSheetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Googlesheets_TmpGetLatestPriceBySymbol_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TmpGetLatestPriceBySymbolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GooglesheetsServer).TmpGetLatestPriceBySymbol(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/googlesheets/TmpGetLatestPriceBySymbol",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GooglesheetsServer).TmpGetLatestPriceBySymbol(ctx, req.(*TmpGetLatestPriceBySymbolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Googlesheets_serviceDesc = grpc.ServiceDesc{
	ServiceName: "googlesheets",
	HandlerType: (*GooglesheetsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePortfolioSheet",
			Handler:    _Googlesheets_CreatePortfolioSheet_Handler,
		},
		{
			MethodName: "RegisterNewPortfolioSheet",
			Handler:    _Googlesheets_RegisterNewPortfolioSheet_Handler,
		},
		{
			MethodName: "TmpGetLatestPriceBySymbol",
			Handler:    _Googlesheets_TmpGetLatestPriceBySymbol_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "s.googlesheets/proto/googlesheets.proto",
}
