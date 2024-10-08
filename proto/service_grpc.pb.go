// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: proto/service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	WeatherFetcher_FetchWeather_FullMethodName = "/WeatherFetcher/FetchWeather"
)

// WeatherFetcherClient is the client API for WeatherFetcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WeatherFetcherClient interface {
	FetchWeather(ctx context.Context, in *CityRequest, opts ...grpc.CallOption) (*WeatherResponse, error)
}

type weatherFetcherClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherFetcherClient(cc grpc.ClientConnInterface) WeatherFetcherClient {
	return &weatherFetcherClient{cc}
}

func (c *weatherFetcherClient) FetchWeather(ctx context.Context, in *CityRequest, opts ...grpc.CallOption) (*WeatherResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WeatherResponse)
	err := c.cc.Invoke(ctx, WeatherFetcher_FetchWeather_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeatherFetcherServer is the server API for WeatherFetcher service.
// All implementations must embed UnimplementedWeatherFetcherServer
// for forward compatibility.
type WeatherFetcherServer interface {
	FetchWeather(context.Context, *CityRequest) (*WeatherResponse, error)
	mustEmbedUnimplementedWeatherFetcherServer()
}

// UnimplementedWeatherFetcherServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWeatherFetcherServer struct{}

func (UnimplementedWeatherFetcherServer) FetchWeather(context.Context, *CityRequest) (*WeatherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchWeather not implemented")
}
func (UnimplementedWeatherFetcherServer) mustEmbedUnimplementedWeatherFetcherServer() {}
func (UnimplementedWeatherFetcherServer) testEmbeddedByValue()                        {}

// UnsafeWeatherFetcherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeatherFetcherServer will
// result in compilation errors.
type UnsafeWeatherFetcherServer interface {
	mustEmbedUnimplementedWeatherFetcherServer()
}

func RegisterWeatherFetcherServer(s grpc.ServiceRegistrar, srv WeatherFetcherServer) {
	// If the following call pancis, it indicates UnimplementedWeatherFetcherServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WeatherFetcher_ServiceDesc, srv)
}

func _WeatherFetcher_FetchWeather_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherFetcherServer).FetchWeather(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WeatherFetcher_FetchWeather_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherFetcherServer).FetchWeather(ctx, req.(*CityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WeatherFetcher_ServiceDesc is the grpc.ServiceDesc for WeatherFetcher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WeatherFetcher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "WeatherFetcher",
	HandlerType: (*WeatherFetcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchWeather",
			Handler:    _WeatherFetcher_FetchWeather_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}
