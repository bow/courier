// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: courier.proto

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

// CourierClient is the client API for Courier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourierClient interface {
	// AddFeeds adds a new feed source.
	AddFeed(ctx context.Context, in *AddFeedRequest, opts ...grpc.CallOption) (*AddFeedResponse, error)
	// EditFeeds sets one or more fields of feeds.
	EditFeeds(ctx context.Context, in *EditFeedsRequest, opts ...grpc.CallOption) (*EditFeedsResponse, error)
	// ListFeeds lists all added feed sources.
	ListFeeds(ctx context.Context, in *ListFeedsRequest, opts ...grpc.CallOption) (*ListFeedsResponse, error)
	// DeleteFeeds removes one or more feed sources.
	DeleteFeeds(ctx context.Context, in *DeleteFeedsRequest, opts ...grpc.CallOption) (*DeleteFeedsResponse, error)
	// UpdateFeeds updates all feeds and returns ones with new entries.
	UpdateFeeds(ctx context.Context, in *UpdateFeedsRequest, opts ...grpc.CallOption) (Courier_UpdateFeedsClient, error)
	// EditEntries sets one or more fields of an entry.
	EditEntries(ctx context.Context, in *EditEntriesRequest, opts ...grpc.CallOption) (*EditEntriesResponse, error)
	// ExportOPML exports feed subscriptions as an OPML document.
	ExportOPML(ctx context.Context, in *ExportOPMLRequest, opts ...grpc.CallOption) (*ExportOPMLResponse, error)
	// ImportOPML imports an OPML document.
	ImportOPML(ctx context.Context, in *ImportOPMLRequest, opts ...grpc.CallOption) (*ImportOPMLResponse, error)
	// GetInfo returns the version info of the running server.
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error)
}

type courierClient struct {
	cc grpc.ClientConnInterface
}

func NewCourierClient(cc grpc.ClientConnInterface) CourierClient {
	return &courierClient{cc}
}

func (c *courierClient) AddFeed(ctx context.Context, in *AddFeedRequest, opts ...grpc.CallOption) (*AddFeedResponse, error) {
	out := new(AddFeedResponse)
	err := c.cc.Invoke(ctx, "/courier.Courier/AddFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierClient) EditFeeds(ctx context.Context, in *EditFeedsRequest, opts ...grpc.CallOption) (*EditFeedsResponse, error) {
	out := new(EditFeedsResponse)
	err := c.cc.Invoke(ctx, "/courier.Courier/EditFeeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierClient) ListFeeds(ctx context.Context, in *ListFeedsRequest, opts ...grpc.CallOption) (*ListFeedsResponse, error) {
	out := new(ListFeedsResponse)
	err := c.cc.Invoke(ctx, "/courier.Courier/ListFeeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierClient) DeleteFeeds(ctx context.Context, in *DeleteFeedsRequest, opts ...grpc.CallOption) (*DeleteFeedsResponse, error) {
	out := new(DeleteFeedsResponse)
	err := c.cc.Invoke(ctx, "/courier.Courier/DeleteFeeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierClient) UpdateFeeds(ctx context.Context, in *UpdateFeedsRequest, opts ...grpc.CallOption) (Courier_UpdateFeedsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Courier_ServiceDesc.Streams[0], "/courier.Courier/UpdateFeeds", opts...)
	if err != nil {
		return nil, err
	}
	x := &courierUpdateFeedsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Courier_UpdateFeedsClient interface {
	Recv() (*UpdateFeedsResponse, error)
	grpc.ClientStream
}

type courierUpdateFeedsClient struct {
	grpc.ClientStream
}

func (x *courierUpdateFeedsClient) Recv() (*UpdateFeedsResponse, error) {
	m := new(UpdateFeedsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *courierClient) EditEntries(ctx context.Context, in *EditEntriesRequest, opts ...grpc.CallOption) (*EditEntriesResponse, error) {
	out := new(EditEntriesResponse)
	err := c.cc.Invoke(ctx, "/courier.Courier/EditEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierClient) ExportOPML(ctx context.Context, in *ExportOPMLRequest, opts ...grpc.CallOption) (*ExportOPMLResponse, error) {
	out := new(ExportOPMLResponse)
	err := c.cc.Invoke(ctx, "/courier.Courier/ExportOPML", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierClient) ImportOPML(ctx context.Context, in *ImportOPMLRequest, opts ...grpc.CallOption) (*ImportOPMLResponse, error) {
	out := new(ImportOPMLResponse)
	err := c.cc.Invoke(ctx, "/courier.Courier/ImportOPML", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error) {
	out := new(GetInfoResponse)
	err := c.cc.Invoke(ctx, "/courier.Courier/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourierServer is the server API for Courier service.
// All implementations must embed UnimplementedCourierServer
// for forward compatibility
type CourierServer interface {
	// AddFeeds adds a new feed source.
	AddFeed(context.Context, *AddFeedRequest) (*AddFeedResponse, error)
	// EditFeeds sets one or more fields of feeds.
	EditFeeds(context.Context, *EditFeedsRequest) (*EditFeedsResponse, error)
	// ListFeeds lists all added feed sources.
	ListFeeds(context.Context, *ListFeedsRequest) (*ListFeedsResponse, error)
	// DeleteFeeds removes one or more feed sources.
	DeleteFeeds(context.Context, *DeleteFeedsRequest) (*DeleteFeedsResponse, error)
	// UpdateFeeds updates all feeds and returns ones with new entries.
	UpdateFeeds(*UpdateFeedsRequest, Courier_UpdateFeedsServer) error
	// EditEntries sets one or more fields of an entry.
	EditEntries(context.Context, *EditEntriesRequest) (*EditEntriesResponse, error)
	// ExportOPML exports feed subscriptions as an OPML document.
	ExportOPML(context.Context, *ExportOPMLRequest) (*ExportOPMLResponse, error)
	// ImportOPML imports an OPML document.
	ImportOPML(context.Context, *ImportOPMLRequest) (*ImportOPMLResponse, error)
	// GetInfo returns the version info of the running server.
	GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error)
	mustEmbedUnimplementedCourierServer()
}

// UnimplementedCourierServer must be embedded to have forward compatible implementations.
type UnimplementedCourierServer struct {
}

func (UnimplementedCourierServer) AddFeed(context.Context, *AddFeedRequest) (*AddFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFeed not implemented")
}
func (UnimplementedCourierServer) EditFeeds(context.Context, *EditFeedsRequest) (*EditFeedsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditFeeds not implemented")
}
func (UnimplementedCourierServer) ListFeeds(context.Context, *ListFeedsRequest) (*ListFeedsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFeeds not implemented")
}
func (UnimplementedCourierServer) DeleteFeeds(context.Context, *DeleteFeedsRequest) (*DeleteFeedsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFeeds not implemented")
}
func (UnimplementedCourierServer) UpdateFeeds(*UpdateFeedsRequest, Courier_UpdateFeedsServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateFeeds not implemented")
}
func (UnimplementedCourierServer) EditEntries(context.Context, *EditEntriesRequest) (*EditEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditEntries not implemented")
}
func (UnimplementedCourierServer) ExportOPML(context.Context, *ExportOPMLRequest) (*ExportOPMLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportOPML not implemented")
}
func (UnimplementedCourierServer) ImportOPML(context.Context, *ImportOPMLRequest) (*ImportOPMLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportOPML not implemented")
}
func (UnimplementedCourierServer) GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (UnimplementedCourierServer) mustEmbedUnimplementedCourierServer() {}

// UnsafeCourierServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourierServer will
// result in compilation errors.
type UnsafeCourierServer interface {
	mustEmbedUnimplementedCourierServer()
}

func RegisterCourierServer(s grpc.ServiceRegistrar, srv CourierServer) {
	s.RegisterService(&Courier_ServiceDesc, srv)
}

func _Courier_AddFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServer).AddFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/courier.Courier/AddFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServer).AddFeed(ctx, req.(*AddFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courier_EditFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServer).EditFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/courier.Courier/EditFeeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServer).EditFeeds(ctx, req.(*EditFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courier_ListFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServer).ListFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/courier.Courier/ListFeeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServer).ListFeeds(ctx, req.(*ListFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courier_DeleteFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServer).DeleteFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/courier.Courier/DeleteFeeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServer).DeleteFeeds(ctx, req.(*DeleteFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courier_UpdateFeeds_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UpdateFeedsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CourierServer).UpdateFeeds(m, &courierUpdateFeedsServer{stream})
}

type Courier_UpdateFeedsServer interface {
	Send(*UpdateFeedsResponse) error
	grpc.ServerStream
}

type courierUpdateFeedsServer struct {
	grpc.ServerStream
}

func (x *courierUpdateFeedsServer) Send(m *UpdateFeedsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Courier_EditEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServer).EditEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/courier.Courier/EditEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServer).EditEntries(ctx, req.(*EditEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courier_ExportOPML_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportOPMLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServer).ExportOPML(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/courier.Courier/ExportOPML",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServer).ExportOPML(ctx, req.(*ExportOPMLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courier_ImportOPML_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportOPMLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServer).ImportOPML(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/courier.Courier/ImportOPML",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServer).ImportOPML(ctx, req.(*ImportOPMLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courier_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/courier.Courier/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServer).GetInfo(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Courier_ServiceDesc is the grpc.ServiceDesc for Courier service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Courier_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "courier.Courier",
	HandlerType: (*CourierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFeed",
			Handler:    _Courier_AddFeed_Handler,
		},
		{
			MethodName: "EditFeeds",
			Handler:    _Courier_EditFeeds_Handler,
		},
		{
			MethodName: "ListFeeds",
			Handler:    _Courier_ListFeeds_Handler,
		},
		{
			MethodName: "DeleteFeeds",
			Handler:    _Courier_DeleteFeeds_Handler,
		},
		{
			MethodName: "EditEntries",
			Handler:    _Courier_EditEntries_Handler,
		},
		{
			MethodName: "ExportOPML",
			Handler:    _Courier_ExportOPML_Handler,
		},
		{
			MethodName: "ImportOPML",
			Handler:    _Courier_ImportOPML_Handler,
		},
		{
			MethodName: "GetInfo",
			Handler:    _Courier_GetInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UpdateFeeds",
			Handler:       _Courier_UpdateFeeds_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "courier.proto",
}
