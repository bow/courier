// Protobuf interface of the neon service.
//
// Copyright (c) 2022 Wibowo Arindrarto <contact@arindrarto.dev>
// SPDX-License-Identifier: BSD-3-Clause

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: neon.proto

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
	Neon_AddFeed_FullMethodName       = "/neon.Neon/AddFeed"
	Neon_EditFeeds_FullMethodName     = "/neon.Neon/EditFeeds"
	Neon_ListFeeds_FullMethodName     = "/neon.Neon/ListFeeds"
	Neon_PullFeeds_FullMethodName     = "/neon.Neon/PullFeeds"
	Neon_DeleteFeeds_FullMethodName   = "/neon.Neon/DeleteFeeds"
	Neon_StreamEntries_FullMethodName = "/neon.Neon/StreamEntries"
	Neon_ListEntries_FullMethodName   = "/neon.Neon/ListEntries"
	Neon_EditEntries_FullMethodName   = "/neon.Neon/EditEntries"
	Neon_GetEntry_FullMethodName      = "/neon.Neon/GetEntry"
	Neon_ExportOPML_FullMethodName    = "/neon.Neon/ExportOPML"
	Neon_ImportOPML_FullMethodName    = "/neon.Neon/ImportOPML"
	Neon_GetStats_FullMethodName      = "/neon.Neon/GetStats"
	Neon_GetInfo_FullMethodName       = "/neon.Neon/GetInfo"
)

// NeonClient is the client API for Neon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NeonClient interface {
	// AddFeeds adds a new feed source.
	AddFeed(ctx context.Context, in *AddFeedRequest, opts ...grpc.CallOption) (*AddFeedResponse, error)
	// EditFeeds sets one or more fields of feeds.
	EditFeeds(ctx context.Context, in *EditFeedsRequest, opts ...grpc.CallOption) (*EditFeedsResponse, error)
	// ListFeeds lists all added feed sources.
	ListFeeds(ctx context.Context, in *ListFeedsRequest, opts ...grpc.CallOption) (*ListFeedsResponse, error)
	// PullFeeds checks feeds for updates and returns them.
	PullFeeds(ctx context.Context, in *PullFeedsRequest, opts ...grpc.CallOption) (Neon_PullFeedsClient, error)
	// DeleteFeeds removes one or more feed sources.
	DeleteFeeds(ctx context.Context, in *DeleteFeedsRequest, opts ...grpc.CallOption) (*DeleteFeedsResponse, error)
	// StreamEntries streams entries of a specific feed.
	StreamEntries(ctx context.Context, in *StreamEntriesRequest, opts ...grpc.CallOption) (Neon_StreamEntriesClient, error)
	// ListEntries lists entries of a specific feed.
	ListEntries(ctx context.Context, in *ListEntriesRequest, opts ...grpc.CallOption) (*ListEntriesResponse, error)
	// EditEntries sets one or more fields of an entry.
	EditEntries(ctx context.Context, in *EditEntriesRequest, opts ...grpc.CallOption) (*EditEntriesResponse, error)
	// GetEntry returns the content of an entry.
	GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (*GetEntryResponse, error)
	// ExportOPML exports feed subscriptions as an OPML document.
	ExportOPML(ctx context.Context, in *ExportOPMLRequest, opts ...grpc.CallOption) (*ExportOPMLResponse, error)
	// ImportOPML imports an OPML document.
	ImportOPML(ctx context.Context, in *ImportOPMLRequest, opts ...grpc.CallOption) (*ImportOPMLResponse, error)
	// GetStats returns various statistics of the feed subscriptions.
	GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error)
	// GetInfo returns the version info of the running server.
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error)
}

type neonClient struct {
	cc grpc.ClientConnInterface
}

func NewNeonClient(cc grpc.ClientConnInterface) NeonClient {
	return &neonClient{cc}
}

func (c *neonClient) AddFeed(ctx context.Context, in *AddFeedRequest, opts ...grpc.CallOption) (*AddFeedResponse, error) {
	out := new(AddFeedResponse)
	err := c.cc.Invoke(ctx, Neon_AddFeed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neonClient) EditFeeds(ctx context.Context, in *EditFeedsRequest, opts ...grpc.CallOption) (*EditFeedsResponse, error) {
	out := new(EditFeedsResponse)
	err := c.cc.Invoke(ctx, Neon_EditFeeds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neonClient) ListFeeds(ctx context.Context, in *ListFeedsRequest, opts ...grpc.CallOption) (*ListFeedsResponse, error) {
	out := new(ListFeedsResponse)
	err := c.cc.Invoke(ctx, Neon_ListFeeds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neonClient) PullFeeds(ctx context.Context, in *PullFeedsRequest, opts ...grpc.CallOption) (Neon_PullFeedsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Neon_ServiceDesc.Streams[0], Neon_PullFeeds_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &neonPullFeedsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Neon_PullFeedsClient interface {
	Recv() (*PullFeedsResponse, error)
	grpc.ClientStream
}

type neonPullFeedsClient struct {
	grpc.ClientStream
}

func (x *neonPullFeedsClient) Recv() (*PullFeedsResponse, error) {
	m := new(PullFeedsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *neonClient) DeleteFeeds(ctx context.Context, in *DeleteFeedsRequest, opts ...grpc.CallOption) (*DeleteFeedsResponse, error) {
	out := new(DeleteFeedsResponse)
	err := c.cc.Invoke(ctx, Neon_DeleteFeeds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neonClient) StreamEntries(ctx context.Context, in *StreamEntriesRequest, opts ...grpc.CallOption) (Neon_StreamEntriesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Neon_ServiceDesc.Streams[1], Neon_StreamEntries_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &neonStreamEntriesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Neon_StreamEntriesClient interface {
	Recv() (*StreamEntriesResponse, error)
	grpc.ClientStream
}

type neonStreamEntriesClient struct {
	grpc.ClientStream
}

func (x *neonStreamEntriesClient) Recv() (*StreamEntriesResponse, error) {
	m := new(StreamEntriesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *neonClient) ListEntries(ctx context.Context, in *ListEntriesRequest, opts ...grpc.CallOption) (*ListEntriesResponse, error) {
	out := new(ListEntriesResponse)
	err := c.cc.Invoke(ctx, Neon_ListEntries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neonClient) EditEntries(ctx context.Context, in *EditEntriesRequest, opts ...grpc.CallOption) (*EditEntriesResponse, error) {
	out := new(EditEntriesResponse)
	err := c.cc.Invoke(ctx, Neon_EditEntries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neonClient) GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (*GetEntryResponse, error) {
	out := new(GetEntryResponse)
	err := c.cc.Invoke(ctx, Neon_GetEntry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neonClient) ExportOPML(ctx context.Context, in *ExportOPMLRequest, opts ...grpc.CallOption) (*ExportOPMLResponse, error) {
	out := new(ExportOPMLResponse)
	err := c.cc.Invoke(ctx, Neon_ExportOPML_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neonClient) ImportOPML(ctx context.Context, in *ImportOPMLRequest, opts ...grpc.CallOption) (*ImportOPMLResponse, error) {
	out := new(ImportOPMLResponse)
	err := c.cc.Invoke(ctx, Neon_ImportOPML_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neonClient) GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error) {
	out := new(GetStatsResponse)
	err := c.cc.Invoke(ctx, Neon_GetStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neonClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error) {
	out := new(GetInfoResponse)
	err := c.cc.Invoke(ctx, Neon_GetInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NeonServer is the server API for Neon service.
// All implementations must embed UnimplementedNeonServer
// for forward compatibility
type NeonServer interface {
	// AddFeeds adds a new feed source.
	AddFeed(context.Context, *AddFeedRequest) (*AddFeedResponse, error)
	// EditFeeds sets one or more fields of feeds.
	EditFeeds(context.Context, *EditFeedsRequest) (*EditFeedsResponse, error)
	// ListFeeds lists all added feed sources.
	ListFeeds(context.Context, *ListFeedsRequest) (*ListFeedsResponse, error)
	// PullFeeds checks feeds for updates and returns them.
	PullFeeds(*PullFeedsRequest, Neon_PullFeedsServer) error
	// DeleteFeeds removes one or more feed sources.
	DeleteFeeds(context.Context, *DeleteFeedsRequest) (*DeleteFeedsResponse, error)
	// StreamEntries streams entries of a specific feed.
	StreamEntries(*StreamEntriesRequest, Neon_StreamEntriesServer) error
	// ListEntries lists entries of a specific feed.
	ListEntries(context.Context, *ListEntriesRequest) (*ListEntriesResponse, error)
	// EditEntries sets one or more fields of an entry.
	EditEntries(context.Context, *EditEntriesRequest) (*EditEntriesResponse, error)
	// GetEntry returns the content of an entry.
	GetEntry(context.Context, *GetEntryRequest) (*GetEntryResponse, error)
	// ExportOPML exports feed subscriptions as an OPML document.
	ExportOPML(context.Context, *ExportOPMLRequest) (*ExportOPMLResponse, error)
	// ImportOPML imports an OPML document.
	ImportOPML(context.Context, *ImportOPMLRequest) (*ImportOPMLResponse, error)
	// GetStats returns various statistics of the feed subscriptions.
	GetStats(context.Context, *GetStatsRequest) (*GetStatsResponse, error)
	// GetInfo returns the version info of the running server.
	GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error)
	mustEmbedUnimplementedNeonServer()
}

// UnimplementedNeonServer must be embedded to have forward compatible implementations.
type UnimplementedNeonServer struct {
}

func (UnimplementedNeonServer) AddFeed(context.Context, *AddFeedRequest) (*AddFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFeed not implemented")
}
func (UnimplementedNeonServer) EditFeeds(context.Context, *EditFeedsRequest) (*EditFeedsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditFeeds not implemented")
}
func (UnimplementedNeonServer) ListFeeds(context.Context, *ListFeedsRequest) (*ListFeedsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFeeds not implemented")
}
func (UnimplementedNeonServer) PullFeeds(*PullFeedsRequest, Neon_PullFeedsServer) error {
	return status.Errorf(codes.Unimplemented, "method PullFeeds not implemented")
}
func (UnimplementedNeonServer) DeleteFeeds(context.Context, *DeleteFeedsRequest) (*DeleteFeedsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFeeds not implemented")
}
func (UnimplementedNeonServer) StreamEntries(*StreamEntriesRequest, Neon_StreamEntriesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamEntries not implemented")
}
func (UnimplementedNeonServer) ListEntries(context.Context, *ListEntriesRequest) (*ListEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEntries not implemented")
}
func (UnimplementedNeonServer) EditEntries(context.Context, *EditEntriesRequest) (*EditEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditEntries not implemented")
}
func (UnimplementedNeonServer) GetEntry(context.Context, *GetEntryRequest) (*GetEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntry not implemented")
}
func (UnimplementedNeonServer) ExportOPML(context.Context, *ExportOPMLRequest) (*ExportOPMLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportOPML not implemented")
}
func (UnimplementedNeonServer) ImportOPML(context.Context, *ImportOPMLRequest) (*ImportOPMLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportOPML not implemented")
}
func (UnimplementedNeonServer) GetStats(context.Context, *GetStatsRequest) (*GetStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}
func (UnimplementedNeonServer) GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (UnimplementedNeonServer) mustEmbedUnimplementedNeonServer() {}

// UnsafeNeonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NeonServer will
// result in compilation errors.
type UnsafeNeonServer interface {
	mustEmbedUnimplementedNeonServer()
}

func RegisterNeonServer(s grpc.ServiceRegistrar, srv NeonServer) {
	s.RegisterService(&Neon_ServiceDesc, srv)
}

func _Neon_AddFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NeonServer).AddFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Neon_AddFeed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NeonServer).AddFeed(ctx, req.(*AddFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Neon_EditFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NeonServer).EditFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Neon_EditFeeds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NeonServer).EditFeeds(ctx, req.(*EditFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Neon_ListFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NeonServer).ListFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Neon_ListFeeds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NeonServer).ListFeeds(ctx, req.(*ListFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Neon_PullFeeds_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullFeedsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NeonServer).PullFeeds(m, &neonPullFeedsServer{stream})
}

type Neon_PullFeedsServer interface {
	Send(*PullFeedsResponse) error
	grpc.ServerStream
}

type neonPullFeedsServer struct {
	grpc.ServerStream
}

func (x *neonPullFeedsServer) Send(m *PullFeedsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Neon_DeleteFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NeonServer).DeleteFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Neon_DeleteFeeds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NeonServer).DeleteFeeds(ctx, req.(*DeleteFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Neon_StreamEntries_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamEntriesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NeonServer).StreamEntries(m, &neonStreamEntriesServer{stream})
}

type Neon_StreamEntriesServer interface {
	Send(*StreamEntriesResponse) error
	grpc.ServerStream
}

type neonStreamEntriesServer struct {
	grpc.ServerStream
}

func (x *neonStreamEntriesServer) Send(m *StreamEntriesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Neon_ListEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NeonServer).ListEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Neon_ListEntries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NeonServer).ListEntries(ctx, req.(*ListEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Neon_EditEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NeonServer).EditEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Neon_EditEntries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NeonServer).EditEntries(ctx, req.(*EditEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Neon_GetEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NeonServer).GetEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Neon_GetEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NeonServer).GetEntry(ctx, req.(*GetEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Neon_ExportOPML_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportOPMLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NeonServer).ExportOPML(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Neon_ExportOPML_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NeonServer).ExportOPML(ctx, req.(*ExportOPMLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Neon_ImportOPML_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportOPMLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NeonServer).ImportOPML(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Neon_ImportOPML_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NeonServer).ImportOPML(ctx, req.(*ImportOPMLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Neon_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NeonServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Neon_GetStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NeonServer).GetStats(ctx, req.(*GetStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Neon_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NeonServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Neon_GetInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NeonServer).GetInfo(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Neon_ServiceDesc is the grpc.ServiceDesc for Neon service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Neon_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "neon.Neon",
	HandlerType: (*NeonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFeed",
			Handler:    _Neon_AddFeed_Handler,
		},
		{
			MethodName: "EditFeeds",
			Handler:    _Neon_EditFeeds_Handler,
		},
		{
			MethodName: "ListFeeds",
			Handler:    _Neon_ListFeeds_Handler,
		},
		{
			MethodName: "DeleteFeeds",
			Handler:    _Neon_DeleteFeeds_Handler,
		},
		{
			MethodName: "ListEntries",
			Handler:    _Neon_ListEntries_Handler,
		},
		{
			MethodName: "EditEntries",
			Handler:    _Neon_EditEntries_Handler,
		},
		{
			MethodName: "GetEntry",
			Handler:    _Neon_GetEntry_Handler,
		},
		{
			MethodName: "ExportOPML",
			Handler:    _Neon_ExportOPML_Handler,
		},
		{
			MethodName: "ImportOPML",
			Handler:    _Neon_ImportOPML_Handler,
		},
		{
			MethodName: "GetStats",
			Handler:    _Neon_GetStats_Handler,
		},
		{
			MethodName: "GetInfo",
			Handler:    _Neon_GetInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullFeeds",
			Handler:       _Neon_PullFeeds_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamEntries",
			Handler:       _Neon_StreamEntries_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "neon.proto",
}
