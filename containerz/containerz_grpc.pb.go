//
// Copyright 2023 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// This file defines the gNOI APIs used to perform container operations on a
// network device. This specification is still under design and subject to
// change.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: containerz/containerz.proto

package containerz

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
	Containerz_Deploy_FullMethodName = "/gnoi.containerz.Containerz/Deploy"
	Containerz_Remove_FullMethodName = "/gnoi.containerz.Containerz/Remove"
	Containerz_List_FullMethodName   = "/gnoi.containerz.Containerz/List"
	Containerz_Start_FullMethodName  = "/gnoi.containerz.Containerz/Start"
	Containerz_Stop_FullMethodName   = "/gnoi.containerz.Containerz/Stop"
	Containerz_Log_FullMethodName    = "/gnoi.containerz.Containerz/Log"
)

// ContainerzClient is the client API for Containerz service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContainerzClient interface {
	// Deploy sets a container image on the target. The container is sent as
	// a sequential stream of messages containing up to 64KB of data. Upon
	// reception of a valid container, the target must load it into its registry.
	// Whether the registry is local or remote is target and deployment specific.
	// A valid container is one that has passed its checksum.
	Deploy(ctx context.Context, opts ...grpc.CallOption) (Containerz_DeployClient, error)
	// Remove deletes containers that match the spec defined in the request. If
	// the specified container does not exist, this operation is a no-op.
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
	// List returns all containers that match the spec defined in the request.
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (Containerz_ListClient, error)
	// Start starts a container. If the image does not exist on the target,
	// Start returns an error. A started container is identified by an instance
	// name, which  can optionally be supplied by the caller otherwise the target
	// should provide one. If the instance name already exists, the target should
	// return an error.
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	// Stop stops a container. If the container does not exist or is not running
	// this operation returns an error. This operation can, optionally, force
	// (i.e. kill) a container.
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
	// Log streams the logs of a running container. If the container if no longer
	// running this operation streams the latest logs and returns.
	Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (Containerz_LogClient, error)
}

type containerzClient struct {
	cc grpc.ClientConnInterface
}

func NewContainerzClient(cc grpc.ClientConnInterface) ContainerzClient {
	return &containerzClient{cc}
}

func (c *containerzClient) Deploy(ctx context.Context, opts ...grpc.CallOption) (Containerz_DeployClient, error) {
	stream, err := c.cc.NewStream(ctx, &Containerz_ServiceDesc.Streams[0], Containerz_Deploy_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &containerzDeployClient{stream}
	return x, nil
}

type Containerz_DeployClient interface {
	Send(*DeployRequest) error
	Recv() (*DeployResponse, error)
	grpc.ClientStream
}

type containerzDeployClient struct {
	grpc.ClientStream
}

func (x *containerzDeployClient) Send(m *DeployRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *containerzDeployClient) Recv() (*DeployResponse, error) {
	m := new(DeployResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *containerzClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	out := new(RemoveResponse)
	err := c.cc.Invoke(ctx, Containerz_Remove_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerzClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (Containerz_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &Containerz_ServiceDesc.Streams[1], Containerz_List_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &containerzListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Containerz_ListClient interface {
	Recv() (*ListResponse, error)
	grpc.ClientStream
}

type containerzListClient struct {
	grpc.ClientStream
}

func (x *containerzListClient) Recv() (*ListResponse, error) {
	m := new(ListResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *containerzClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, Containerz_Start_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerzClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := c.cc.Invoke(ctx, Containerz_Stop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerzClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (Containerz_LogClient, error) {
	stream, err := c.cc.NewStream(ctx, &Containerz_ServiceDesc.Streams[2], Containerz_Log_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &containerzLogClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Containerz_LogClient interface {
	Recv() (*LogResponse, error)
	grpc.ClientStream
}

type containerzLogClient struct {
	grpc.ClientStream
}

func (x *containerzLogClient) Recv() (*LogResponse, error) {
	m := new(LogResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ContainerzServer is the server API for Containerz service.
// All implementations must embed UnimplementedContainerzServer
// for forward compatibility
type ContainerzServer interface {
	// Deploy sets a container image on the target. The container is sent as
	// a sequential stream of messages containing up to 64KB of data. Upon
	// reception of a valid container, the target must load it into its registry.
	// Whether the registry is local or remote is target and deployment specific.
	// A valid container is one that has passed its checksum.
	Deploy(Containerz_DeployServer) error
	// Remove deletes containers that match the spec defined in the request. If
	// the specified container does not exist, this operation is a no-op.
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)
	// List returns all containers that match the spec defined in the request.
	List(*ListRequest, Containerz_ListServer) error
	// Start starts a container. If the image does not exist on the target,
	// Start returns an error. A started container is identified by an instance
	// name, which  can optionally be supplied by the caller otherwise the target
	// should provide one. If the instance name already exists, the target should
	// return an error.
	Start(context.Context, *StartRequest) (*StartResponse, error)
	// Stop stops a container. If the container does not exist or is not running
	// this operation returns an error. This operation can, optionally, force
	// (i.e. kill) a container.
	Stop(context.Context, *StopRequest) (*StopResponse, error)
	// Log streams the logs of a running container. If the container if no longer
	// running this operation streams the latest logs and returns.
	Log(*LogRequest, Containerz_LogServer) error
	mustEmbedUnimplementedContainerzServer()
}

// UnimplementedContainerzServer must be embedded to have forward compatible implementations.
type UnimplementedContainerzServer struct {
}

func (UnimplementedContainerzServer) Deploy(Containerz_DeployServer) error {
	return status.Errorf(codes.Unimplemented, "method Deploy not implemented")
}
func (UnimplementedContainerzServer) Remove(context.Context, *RemoveRequest) (*RemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedContainerzServer) List(*ListRequest, Containerz_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedContainerzServer) Start(context.Context, *StartRequest) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedContainerzServer) Stop(context.Context, *StopRequest) (*StopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedContainerzServer) Log(*LogRequest, Containerz_LogServer) error {
	return status.Errorf(codes.Unimplemented, "method Log not implemented")
}
func (UnimplementedContainerzServer) mustEmbedUnimplementedContainerzServer() {}

// UnsafeContainerzServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContainerzServer will
// result in compilation errors.
type UnsafeContainerzServer interface {
	mustEmbedUnimplementedContainerzServer()
}

func RegisterContainerzServer(s grpc.ServiceRegistrar, srv ContainerzServer) {
	s.RegisterService(&Containerz_ServiceDesc, srv)
}

func _Containerz_Deploy_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ContainerzServer).Deploy(&containerzDeployServer{stream})
}

type Containerz_DeployServer interface {
	Send(*DeployResponse) error
	Recv() (*DeployRequest, error)
	grpc.ServerStream
}

type containerzDeployServer struct {
	grpc.ServerStream
}

func (x *containerzDeployServer) Send(m *DeployResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *containerzDeployServer) Recv() (*DeployRequest, error) {
	m := new(DeployRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Containerz_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerzServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Containerz_Remove_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerzServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Containerz_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ContainerzServer).List(m, &containerzListServer{stream})
}

type Containerz_ListServer interface {
	Send(*ListResponse) error
	grpc.ServerStream
}

type containerzListServer struct {
	grpc.ServerStream
}

func (x *containerzListServer) Send(m *ListResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Containerz_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerzServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Containerz_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerzServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Containerz_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerzServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Containerz_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerzServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Containerz_Log_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ContainerzServer).Log(m, &containerzLogServer{stream})
}

type Containerz_LogServer interface {
	Send(*LogResponse) error
	grpc.ServerStream
}

type containerzLogServer struct {
	grpc.ServerStream
}

func (x *containerzLogServer) Send(m *LogResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Containerz_ServiceDesc is the grpc.ServiceDesc for Containerz service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Containerz_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnoi.containerz.Containerz",
	HandlerType: (*ContainerzServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Remove",
			Handler:    _Containerz_Remove_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Containerz_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Containerz_Stop_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Deploy",
			Handler:       _Containerz_Deploy_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "List",
			Handler:       _Containerz_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Log",
			Handler:       _Containerz_Log_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "containerz/containerz.proto",
}
