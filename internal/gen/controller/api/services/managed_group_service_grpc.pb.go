// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

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

// ManagedGroupServiceClient is the client API for ManagedGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagedGroupServiceClient interface {
	// GetManagedGroup returns a stored ManagedGroup if present. The provided request must
	// include the id for the ManagedGroup be retrieved. If missing, malformed or
	// referencing a non existing ManagedGroup an error is returned.
	GetManagedGroup(ctx context.Context, in *GetManagedGroupRequest, opts ...grpc.CallOption) (*GetManagedGroupResponse, error)
	// ListManagedGroups returns a list of stored ManagedGroups which exist inside the
	// provided Auth Method. The request must include the Auth Method id which
	// contains the ManagedGroups being listed. If missing or malformed an error
	// is returned.
	ListManagedGroups(ctx context.Context, in *ListManagedGroupsRequest, opts ...grpc.CallOption) (*ListManagedGroupsResponse, error)
	// CreateManagedGroup creates and stores a ManagedGroup. The provided request
	// must include the Auth Method ID in which the ManagedGroup will be created.
	// If the Auth Method ID is missing, malformed, or references a non existing
	// resource an error is returned. If a name or login_name is provided that is
	// in use in another ManagedGroup in the same Auth Method an error is
	// returned.
	CreateManagedGroup(ctx context.Context, in *CreateManagedGroupRequest, opts ...grpc.CallOption) (*CreateManagedGroupResponse, error)
	// UpdateManagedGroup updates an existing ManagedGroup. The provided
	// ManagedGroup must not have any read only fields set. The update mask must
	// be included in the request and contain at least 1 mutable field. To unset a
	// field's value, include the field in the update mask and don't set it in the
	// provided ManagedGroup. An error is returned if the ManagedGroup id is
	// missing or references a non-existing resource. An error is also returned if
	// the request attempts to update the name to one that is already in use in
	// the containing Auth Method.
	UpdateManagedGroup(ctx context.Context, in *UpdateManagedGroupRequest, opts ...grpc.CallOption) (*UpdateManagedGroupResponse, error)
	// DeleteManagedGroup removes a ManagedGroup. If the provided ManagedGroup Id
	// is malformed or not provided an error is returned.
	DeleteManagedGroup(ctx context.Context, in *DeleteManagedGroupRequest, opts ...grpc.CallOption) (*DeleteManagedGroupResponse, error)
}

type managedGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewManagedGroupServiceClient(cc grpc.ClientConnInterface) ManagedGroupServiceClient {
	return &managedGroupServiceClient{cc}
}

func (c *managedGroupServiceClient) GetManagedGroup(ctx context.Context, in *GetManagedGroupRequest, opts ...grpc.CallOption) (*GetManagedGroupResponse, error) {
	out := new(GetManagedGroupResponse)
	err := c.cc.Invoke(ctx, "/controller.api.services.v1.ManagedGroupService/GetManagedGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managedGroupServiceClient) ListManagedGroups(ctx context.Context, in *ListManagedGroupsRequest, opts ...grpc.CallOption) (*ListManagedGroupsResponse, error) {
	out := new(ListManagedGroupsResponse)
	err := c.cc.Invoke(ctx, "/controller.api.services.v1.ManagedGroupService/ListManagedGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managedGroupServiceClient) CreateManagedGroup(ctx context.Context, in *CreateManagedGroupRequest, opts ...grpc.CallOption) (*CreateManagedGroupResponse, error) {
	out := new(CreateManagedGroupResponse)
	err := c.cc.Invoke(ctx, "/controller.api.services.v1.ManagedGroupService/CreateManagedGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managedGroupServiceClient) UpdateManagedGroup(ctx context.Context, in *UpdateManagedGroupRequest, opts ...grpc.CallOption) (*UpdateManagedGroupResponse, error) {
	out := new(UpdateManagedGroupResponse)
	err := c.cc.Invoke(ctx, "/controller.api.services.v1.ManagedGroupService/UpdateManagedGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managedGroupServiceClient) DeleteManagedGroup(ctx context.Context, in *DeleteManagedGroupRequest, opts ...grpc.CallOption) (*DeleteManagedGroupResponse, error) {
	out := new(DeleteManagedGroupResponse)
	err := c.cc.Invoke(ctx, "/controller.api.services.v1.ManagedGroupService/DeleteManagedGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagedGroupServiceServer is the server API for ManagedGroupService service.
// All implementations must embed UnimplementedManagedGroupServiceServer
// for forward compatibility
type ManagedGroupServiceServer interface {
	// GetManagedGroup returns a stored ManagedGroup if present. The provided request must
	// include the id for the ManagedGroup be retrieved. If missing, malformed or
	// referencing a non existing ManagedGroup an error is returned.
	GetManagedGroup(context.Context, *GetManagedGroupRequest) (*GetManagedGroupResponse, error)
	// ListManagedGroups returns a list of stored ManagedGroups which exist inside the
	// provided Auth Method. The request must include the Auth Method id which
	// contains the ManagedGroups being listed. If missing or malformed an error
	// is returned.
	ListManagedGroups(context.Context, *ListManagedGroupsRequest) (*ListManagedGroupsResponse, error)
	// CreateManagedGroup creates and stores a ManagedGroup. The provided request
	// must include the Auth Method ID in which the ManagedGroup will be created.
	// If the Auth Method ID is missing, malformed, or references a non existing
	// resource an error is returned. If a name or login_name is provided that is
	// in use in another ManagedGroup in the same Auth Method an error is
	// returned.
	CreateManagedGroup(context.Context, *CreateManagedGroupRequest) (*CreateManagedGroupResponse, error)
	// UpdateManagedGroup updates an existing ManagedGroup. The provided
	// ManagedGroup must not have any read only fields set. The update mask must
	// be included in the request and contain at least 1 mutable field. To unset a
	// field's value, include the field in the update mask and don't set it in the
	// provided ManagedGroup. An error is returned if the ManagedGroup id is
	// missing or references a non-existing resource. An error is also returned if
	// the request attempts to update the name to one that is already in use in
	// the containing Auth Method.
	UpdateManagedGroup(context.Context, *UpdateManagedGroupRequest) (*UpdateManagedGroupResponse, error)
	// DeleteManagedGroup removes a ManagedGroup. If the provided ManagedGroup Id
	// is malformed or not provided an error is returned.
	DeleteManagedGroup(context.Context, *DeleteManagedGroupRequest) (*DeleteManagedGroupResponse, error)
	mustEmbedUnimplementedManagedGroupServiceServer()
}

// UnimplementedManagedGroupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedManagedGroupServiceServer struct {
}

func (UnimplementedManagedGroupServiceServer) GetManagedGroup(context.Context, *GetManagedGroupRequest) (*GetManagedGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManagedGroup not implemented")
}
func (UnimplementedManagedGroupServiceServer) ListManagedGroups(context.Context, *ListManagedGroupsRequest) (*ListManagedGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListManagedGroups not implemented")
}
func (UnimplementedManagedGroupServiceServer) CreateManagedGroup(context.Context, *CreateManagedGroupRequest) (*CreateManagedGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateManagedGroup not implemented")
}
func (UnimplementedManagedGroupServiceServer) UpdateManagedGroup(context.Context, *UpdateManagedGroupRequest) (*UpdateManagedGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateManagedGroup not implemented")
}
func (UnimplementedManagedGroupServiceServer) DeleteManagedGroup(context.Context, *DeleteManagedGroupRequest) (*DeleteManagedGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteManagedGroup not implemented")
}
func (UnimplementedManagedGroupServiceServer) mustEmbedUnimplementedManagedGroupServiceServer() {}

// UnsafeManagedGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagedGroupServiceServer will
// result in compilation errors.
type UnsafeManagedGroupServiceServer interface {
	mustEmbedUnimplementedManagedGroupServiceServer()
}

func RegisterManagedGroupServiceServer(s grpc.ServiceRegistrar, srv ManagedGroupServiceServer) {
	s.RegisterService(&ManagedGroupService_ServiceDesc, srv)
}

func _ManagedGroupService_GetManagedGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManagedGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagedGroupServiceServer).GetManagedGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.api.services.v1.ManagedGroupService/GetManagedGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagedGroupServiceServer).GetManagedGroup(ctx, req.(*GetManagedGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagedGroupService_ListManagedGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListManagedGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagedGroupServiceServer).ListManagedGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.api.services.v1.ManagedGroupService/ListManagedGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagedGroupServiceServer).ListManagedGroups(ctx, req.(*ListManagedGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagedGroupService_CreateManagedGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateManagedGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagedGroupServiceServer).CreateManagedGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.api.services.v1.ManagedGroupService/CreateManagedGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagedGroupServiceServer).CreateManagedGroup(ctx, req.(*CreateManagedGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagedGroupService_UpdateManagedGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateManagedGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagedGroupServiceServer).UpdateManagedGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.api.services.v1.ManagedGroupService/UpdateManagedGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagedGroupServiceServer).UpdateManagedGroup(ctx, req.(*UpdateManagedGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagedGroupService_DeleteManagedGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteManagedGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagedGroupServiceServer).DeleteManagedGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.api.services.v1.ManagedGroupService/DeleteManagedGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagedGroupServiceServer).DeleteManagedGroup(ctx, req.(*DeleteManagedGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ManagedGroupService_ServiceDesc is the grpc.ServiceDesc for ManagedGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManagedGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "controller.api.services.v1.ManagedGroupService",
	HandlerType: (*ManagedGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetManagedGroup",
			Handler:    _ManagedGroupService_GetManagedGroup_Handler,
		},
		{
			MethodName: "ListManagedGroups",
			Handler:    _ManagedGroupService_ListManagedGroups_Handler,
		},
		{
			MethodName: "CreateManagedGroup",
			Handler:    _ManagedGroupService_CreateManagedGroup_Handler,
		},
		{
			MethodName: "UpdateManagedGroup",
			Handler:    _ManagedGroupService_UpdateManagedGroup_Handler,
		},
		{
			MethodName: "DeleteManagedGroup",
			Handler:    _ManagedGroupService_DeleteManagedGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controller/api/services/v1/managed_group_service.proto",
}