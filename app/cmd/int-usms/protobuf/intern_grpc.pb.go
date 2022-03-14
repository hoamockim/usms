// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package intern

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

// UpdateExtraClient is the client API for UpdateExtra service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UpdateExtraClient interface {
	Update(ctx context.Context, in *ExtraUserInfoReq, opts ...grpc.CallOption) (*ExtraUserInfoRes, error)
}

type updateExtraClient struct {
	cc grpc.ClientConnInterface
}

func NewUpdateExtraClient(cc grpc.ClientConnInterface) UpdateExtraClient {
	return &updateExtraClient{cc}
}

func (c *updateExtraClient) Update(ctx context.Context, in *ExtraUserInfoReq, opts ...grpc.CallOption) (*ExtraUserInfoRes, error) {
	out := new(ExtraUserInfoRes)
	err := c.cc.Invoke(ctx, "/intern.UpdateExtra/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateExtraServer is the server API for UpdateExtra service.
// All implementations must embed UnimplementedUpdateExtraServer
// for forward compatibility
type UpdateExtraServer interface {
	Update(context.Context, *ExtraUserInfoReq) (*ExtraUserInfoRes, error)
	mustEmbedUnimplementedUpdateExtraServer()
}

// UnimplementedUpdateExtraServer must be embedded to have forward compatible implementations.
type UnimplementedUpdateExtraServer struct {
}

func (UnimplementedUpdateExtraServer) Update(context.Context, *ExtraUserInfoReq) (*ExtraUserInfoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUpdateExtraServer) mustEmbedUnimplementedUpdateExtraServer() {}

// UnsafeUpdateExtraServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UpdateExtraServer will
// result in compilation errors.
type UnsafeUpdateExtraServer interface {
	mustEmbedUnimplementedUpdateExtraServer()
}

func RegisterUpdateExtraServer(s grpc.ServiceRegistrar, srv UpdateExtraServer) {
	s.RegisterService(&UpdateExtra_ServiceDesc, srv)
}

func _UpdateExtra_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtraUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateExtraServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/intern.UpdateExtra/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateExtraServer).Update(ctx, req.(*ExtraUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UpdateExtra_ServiceDesc is the grpc.ServiceDesc for UpdateExtra service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UpdateExtra_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "intern.UpdateExtra",
	HandlerType: (*UpdateExtraServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _UpdateExtra_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "intern.proto",
}