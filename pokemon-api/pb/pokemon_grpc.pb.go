// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: pokemon.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Pokedex_Create_FullMethodName  = "/pokemon.Pokedex/Create"
	Pokedex_Read_FullMethodName    = "/pokemon.Pokedex/Read"
	Pokedex_ReadOne_FullMethodName = "/pokemon.Pokedex/ReadOne"
	Pokedex_Update_FullMethodName  = "/pokemon.Pokedex/Update"
	Pokedex_Delete_FullMethodName  = "/pokemon.Pokedex/Delete"
)

// PokedexClient is the client API for Pokedex service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PokedexClient interface {
	Create(ctx context.Context, in *PokemonRequest, opts ...grpc.CallOption) (*Pokemon, error)
	Read(ctx context.Context, in *PokemonFilter, opts ...grpc.CallOption) (*PokemonListResponse, error)
	ReadOne(ctx context.Context, in *PokemonID, opts ...grpc.CallOption) (*Pokemon, error)
	Update(ctx context.Context, in *PokemonUpdateRequest, opts ...grpc.CallOption) (*Pokemon, error)
	Delete(ctx context.Context, in *PokemonID, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type pokedexClient struct {
	cc grpc.ClientConnInterface
}

func NewPokedexClient(cc grpc.ClientConnInterface) PokedexClient {
	return &pokedexClient{cc}
}

func (c *pokedexClient) Create(ctx context.Context, in *PokemonRequest, opts ...grpc.CallOption) (*Pokemon, error) {
	out := new(Pokemon)
	err := c.cc.Invoke(ctx, Pokedex_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokedexClient) Read(ctx context.Context, in *PokemonFilter, opts ...grpc.CallOption) (*PokemonListResponse, error) {
	out := new(PokemonListResponse)
	err := c.cc.Invoke(ctx, Pokedex_Read_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokedexClient) ReadOne(ctx context.Context, in *PokemonID, opts ...grpc.CallOption) (*Pokemon, error) {
	out := new(Pokemon)
	err := c.cc.Invoke(ctx, Pokedex_ReadOne_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokedexClient) Update(ctx context.Context, in *PokemonUpdateRequest, opts ...grpc.CallOption) (*Pokemon, error) {
	out := new(Pokemon)
	err := c.cc.Invoke(ctx, Pokedex_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokedexClient) Delete(ctx context.Context, in *PokemonID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Pokedex_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PokedexServer is the server API for Pokedex service.
// All implementations must embed UnimplementedPokedexServer
// for forward compatibility
type PokedexServer interface {
	Create(context.Context, *PokemonRequest) (*Pokemon, error)
	Read(context.Context, *PokemonFilter) (*PokemonListResponse, error)
	ReadOne(context.Context, *PokemonID) (*Pokemon, error)
	Update(context.Context, *PokemonUpdateRequest) (*Pokemon, error)
	Delete(context.Context, *PokemonID) (*emptypb.Empty, error)
	mustEmbedUnimplementedPokedexServer()
}

// UnimplementedPokedexServer must be embedded to have forward compatible implementations.
type UnimplementedPokedexServer struct {
}

func (UnimplementedPokedexServer) Create(context.Context, *PokemonRequest) (*Pokemon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPokedexServer) Read(context.Context, *PokemonFilter) (*PokemonListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedPokedexServer) ReadOne(context.Context, *PokemonID) (*Pokemon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadOne not implemented")
}
func (UnimplementedPokedexServer) Update(context.Context, *PokemonUpdateRequest) (*Pokemon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPokedexServer) Delete(context.Context, *PokemonID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPokedexServer) mustEmbedUnimplementedPokedexServer() {}

// UnsafePokedexServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PokedexServer will
// result in compilation errors.
type UnsafePokedexServer interface {
	mustEmbedUnimplementedPokedexServer()
}

func RegisterPokedexServer(s grpc.ServiceRegistrar, srv PokedexServer) {
	s.RegisterService(&Pokedex_ServiceDesc, srv)
}

func _Pokedex_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PokemonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokedexServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pokedex_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokedexServer).Create(ctx, req.(*PokemonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pokedex_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PokemonFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokedexServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pokedex_Read_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokedexServer).Read(ctx, req.(*PokemonFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pokedex_ReadOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PokemonID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokedexServer).ReadOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pokedex_ReadOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokedexServer).ReadOne(ctx, req.(*PokemonID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pokedex_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PokemonUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokedexServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pokedex_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokedexServer).Update(ctx, req.(*PokemonUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pokedex_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PokemonID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokedexServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pokedex_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokedexServer).Delete(ctx, req.(*PokemonID))
	}
	return interceptor(ctx, in, info, handler)
}

// Pokedex_ServiceDesc is the grpc.ServiceDesc for Pokedex service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pokedex_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pokemon.Pokedex",
	HandlerType: (*PokedexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Pokedex_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Pokedex_Read_Handler,
		},
		{
			MethodName: "ReadOne",
			Handler:    _Pokedex_ReadOne_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Pokedex_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Pokedex_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pokemon.proto",
}
