// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// NexelliawalletdClient is the client API for Nexelliawalletd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NexelliawalletdClient interface {
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	GetExternalSpendableUTXOs(ctx context.Context, in *GetExternalSpendableUTXOsRequest, opts ...grpc.CallOption) (*GetExternalSpendableUTXOsResponse, error)
	CreateUnsignedTransactions(ctx context.Context, in *CreateUnsignedTransactionsRequest, opts ...grpc.CallOption) (*CreateUnsignedTransactionsResponse, error)
	ShowAddresses(ctx context.Context, in *ShowAddressesRequest, opts ...grpc.CallOption) (*ShowAddressesResponse, error)
	NewAddress(ctx context.Context, in *NewAddressRequest, opts ...grpc.CallOption) (*NewAddressResponse, error)
	Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error)
	Broadcast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (*BroadcastResponse, error)
	// Since SendRequest contains a password - this command should only be used on a trusted or secure connection
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
	// Since SignRequest contains a password - this command should only be used on a trusted or secure connection
	Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error)
}

type nexelliawalletdClient struct {
	cc grpc.ClientConnInterface
}

func NewNexelliawalletdClient(cc grpc.ClientConnInterface) NexelliawalletdClient {
	return &nexelliawalletdClient{cc}
}

func (c *nexelliawalletdClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, "/nexelliawalletd.nexelliawalletd/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nexelliawalletdClient) GetExternalSpendableUTXOs(ctx context.Context, in *GetExternalSpendableUTXOsRequest, opts ...grpc.CallOption) (*GetExternalSpendableUTXOsResponse, error) {
	out := new(GetExternalSpendableUTXOsResponse)
	err := c.cc.Invoke(ctx, "/nexelliawalletd.nexelliawalletd/GetExternalSpendableUTXOs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nexelliawalletdClient) CreateUnsignedTransactions(ctx context.Context, in *CreateUnsignedTransactionsRequest, opts ...grpc.CallOption) (*CreateUnsignedTransactionsResponse, error) {
	out := new(CreateUnsignedTransactionsResponse)
	err := c.cc.Invoke(ctx, "/nexelliawalletd.nexelliawalletd/CreateUnsignedTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nexelliawalletdClient) ShowAddresses(ctx context.Context, in *ShowAddressesRequest, opts ...grpc.CallOption) (*ShowAddressesResponse, error) {
	out := new(ShowAddressesResponse)
	err := c.cc.Invoke(ctx, "/nexelliawalletd.nexelliawalletd/ShowAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nexelliawalletdClient) NewAddress(ctx context.Context, in *NewAddressRequest, opts ...grpc.CallOption) (*NewAddressResponse, error) {
	out := new(NewAddressResponse)
	err := c.cc.Invoke(ctx, "/nexelliawalletd.nexelliawalletd/NewAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nexelliawalletdClient) Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error) {
	out := new(ShutdownResponse)
	err := c.cc.Invoke(ctx, "/nexelliawalletd.nexelliawalletd/Shutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nexelliawalletdClient) Broadcast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (*BroadcastResponse, error) {
	out := new(BroadcastResponse)
	err := c.cc.Invoke(ctx, "/nexelliawalletd.nexelliawalletd/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nexelliawalletdClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/nexelliawalletd.nexelliawalletd/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nexelliawalletdClient) Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error) {
	out := new(SignResponse)
	err := c.cc.Invoke(ctx, "/nexelliawalletd.nexelliawalletd/Sign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NexelliawalletdServer is the server API for Nexelliawalletd service.
// All implementations must embed UnimplementedNexelliawalletdServer
// for forward compatibility
type NexelliawalletdServer interface {
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	GetExternalSpendableUTXOs(context.Context, *GetExternalSpendableUTXOsRequest) (*GetExternalSpendableUTXOsResponse, error)
	CreateUnsignedTransactions(context.Context, *CreateUnsignedTransactionsRequest) (*CreateUnsignedTransactionsResponse, error)
	ShowAddresses(context.Context, *ShowAddressesRequest) (*ShowAddressesResponse, error)
	NewAddress(context.Context, *NewAddressRequest) (*NewAddressResponse, error)
	Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error)
	Broadcast(context.Context, *BroadcastRequest) (*BroadcastResponse, error)
	// Since SendRequest contains a password - this command should only be used on a trusted or secure connection
	Send(context.Context, *SendRequest) (*SendResponse, error)
	// Since SignRequest contains a password - this command should only be used on a trusted or secure connection
	Sign(context.Context, *SignRequest) (*SignResponse, error)
	mustEmbedUnimplementedNexelliawalletdServer()
}

// UnimplementedNexelliawalletdServer must be embedded to have forward compatible implementations.
type UnimplementedNexelliawalletdServer struct {
}

func (UnimplementedNexelliawalletdServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedNexelliawalletdServer) GetExternalSpendableUTXOs(context.Context, *GetExternalSpendableUTXOsRequest) (*GetExternalSpendableUTXOsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExternalSpendableUTXOs not implemented")
}
func (UnimplementedNexelliawalletdServer) CreateUnsignedTransactions(context.Context, *CreateUnsignedTransactionsRequest) (*CreateUnsignedTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUnsignedTransactions not implemented")
}
func (UnimplementedNexelliawalletdServer) ShowAddresses(context.Context, *ShowAddressesRequest) (*ShowAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowAddresses not implemented")
}
func (UnimplementedNexelliawalletdServer) NewAddress(context.Context, *NewAddressRequest) (*NewAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewAddress not implemented")
}
func (UnimplementedNexelliawalletdServer) Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}
func (UnimplementedNexelliawalletdServer) Broadcast(context.Context, *BroadcastRequest) (*BroadcastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedNexelliawalletdServer) Send(context.Context, *SendRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedNexelliawalletdServer) Sign(context.Context, *SignRequest) (*SignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}
func (UnimplementedNexelliawalletdServer) mustEmbedUnimplementedNexelliawalletdServer() {}

// UnsafeNexelliawalletdServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NexelliawalletdServer will
// result in compilation errors.
type UnsafeNexelliawalletdServer interface {
	mustEmbedUnimplementedNexelliawalletdServer()
}

func RegisterNexelliawalletdServer(s grpc.ServiceRegistrar, srv NexelliawalletdServer) {
	s.RegisterService(&Nexelliawalletd_ServiceDesc, srv)
}

func _Nexelliawalletd_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NexelliawalletdServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexelliawalletd.nexelliawalletd/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NexelliawalletdServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nexelliawalletd_GetExternalSpendableUTXOs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExternalSpendableUTXOsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NexelliawalletdServer).GetExternalSpendableUTXOs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexelliawalletd.nexelliawalletd/GetExternalSpendableUTXOs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NexelliawalletdServer).GetExternalSpendableUTXOs(ctx, req.(*GetExternalSpendableUTXOsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nexelliawalletd_CreateUnsignedTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUnsignedTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NexelliawalletdServer).CreateUnsignedTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexelliawalletd.nexelliawalletd/CreateUnsignedTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NexelliawalletdServer).CreateUnsignedTransactions(ctx, req.(*CreateUnsignedTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nexelliawalletd_ShowAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NexelliawalletdServer).ShowAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexelliawalletd.nexelliawalletd/ShowAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NexelliawalletdServer).ShowAddresses(ctx, req.(*ShowAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nexelliawalletd_NewAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NexelliawalletdServer).NewAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexelliawalletd.nexelliawalletd/NewAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NexelliawalletdServer).NewAddress(ctx, req.(*NewAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nexelliawalletd_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NexelliawalletdServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexelliawalletd.nexelliawalletd/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NexelliawalletdServer).Shutdown(ctx, req.(*ShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nexelliawalletd_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NexelliawalletdServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexelliawalletd.nexelliawalletd/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NexelliawalletdServer).Broadcast(ctx, req.(*BroadcastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nexelliawalletd_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NexelliawalletdServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexelliawalletd.nexelliawalletd/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NexelliawalletdServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nexelliawalletd_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NexelliawalletdServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexelliawalletd.nexelliawalletd/Sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NexelliawalletdServer).Sign(ctx, req.(*SignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Nexelliawalletd_ServiceDesc is the grpc.ServiceDesc for Nexelliawalletd service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Nexelliawalletd_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nexelliawalletd.nexelliawalletd",
	HandlerType: (*NexelliawalletdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBalance",
			Handler:    _Nexelliawalletd_GetBalance_Handler,
		},
		{
			MethodName: "GetExternalSpendableUTXOs",
			Handler:    _Nexelliawalletd_GetExternalSpendableUTXOs_Handler,
		},
		{
			MethodName: "CreateUnsignedTransactions",
			Handler:    _Nexelliawalletd_CreateUnsignedTransactions_Handler,
		},
		{
			MethodName: "ShowAddresses",
			Handler:    _Nexelliawalletd_ShowAddresses_Handler,
		},
		{
			MethodName: "NewAddress",
			Handler:    _Nexelliawalletd_NewAddress_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _Nexelliawalletd_Shutdown_Handler,
		},
		{
			MethodName: "Broadcast",
			Handler:    _Nexelliawalletd_Broadcast_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Nexelliawalletd_Send_Handler,
		},
		{
			MethodName: "Sign",
			Handler:    _Nexelliawalletd_Sign_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nexelliawalletd.proto",
}
