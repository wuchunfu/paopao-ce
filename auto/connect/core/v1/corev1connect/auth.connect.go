// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: core/v1/auth.proto

package corev1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/rocboss/paopao-ce/auto/rpc/core/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// AuthenticateServiceName is the fully-qualified name of the AuthenticateService service.
	AuthenticateServiceName = "core.v1.AuthenticateService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AuthenticateServicePreLoginProcedure is the fully-qualified name of the AuthenticateService's
	// preLogin RPC.
	AuthenticateServicePreLoginProcedure = "/core.v1.AuthenticateService/preLogin"
	// AuthenticateServiceLoginProcedure is the fully-qualified name of the AuthenticateService's login
	// RPC.
	AuthenticateServiceLoginProcedure = "/core.v1.AuthenticateService/login"
	// AuthenticateServiceLogoutProcedure is the fully-qualified name of the AuthenticateService's
	// logout RPC.
	AuthenticateServiceLogoutProcedure = "/core.v1.AuthenticateService/logout"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	authenticateServiceServiceDescriptor        = v1.File_core_v1_auth_proto.Services().ByName("AuthenticateService")
	authenticateServicePreLoginMethodDescriptor = authenticateServiceServiceDescriptor.Methods().ByName("preLogin")
	authenticateServiceLoginMethodDescriptor    = authenticateServiceServiceDescriptor.Methods().ByName("login")
	authenticateServiceLogoutMethodDescriptor   = authenticateServiceServiceDescriptor.Methods().ByName("logout")
)

// AuthenticateServiceClient is a client for the core.v1.AuthenticateService service.
type AuthenticateServiceClient interface {
	PreLogin(context.Context, *connect.Request[v1.User]) (*connect.Response[v1.ActionReply], error)
	Login(context.Context, *connect.Request[v1.User]) (*connect.Response[v1.LoginReply], error)
	Logout(context.Context, *connect.Request[v1.User]) (*connect.Response[v1.ActionReply], error)
}

// NewAuthenticateServiceClient constructs a client for the core.v1.AuthenticateService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuthenticateServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AuthenticateServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &authenticateServiceClient{
		preLogin: connect.NewClient[v1.User, v1.ActionReply](
			httpClient,
			baseURL+AuthenticateServicePreLoginProcedure,
			connect.WithSchema(authenticateServicePreLoginMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		login: connect.NewClient[v1.User, v1.LoginReply](
			httpClient,
			baseURL+AuthenticateServiceLoginProcedure,
			connect.WithSchema(authenticateServiceLoginMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		logout: connect.NewClient[v1.User, v1.ActionReply](
			httpClient,
			baseURL+AuthenticateServiceLogoutProcedure,
			connect.WithSchema(authenticateServiceLogoutMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// authenticateServiceClient implements AuthenticateServiceClient.
type authenticateServiceClient struct {
	preLogin *connect.Client[v1.User, v1.ActionReply]
	login    *connect.Client[v1.User, v1.LoginReply]
	logout   *connect.Client[v1.User, v1.ActionReply]
}

// PreLogin calls core.v1.AuthenticateService.preLogin.
func (c *authenticateServiceClient) PreLogin(ctx context.Context, req *connect.Request[v1.User]) (*connect.Response[v1.ActionReply], error) {
	return c.preLogin.CallUnary(ctx, req)
}

// Login calls core.v1.AuthenticateService.login.
func (c *authenticateServiceClient) Login(ctx context.Context, req *connect.Request[v1.User]) (*connect.Response[v1.LoginReply], error) {
	return c.login.CallUnary(ctx, req)
}

// Logout calls core.v1.AuthenticateService.logout.
func (c *authenticateServiceClient) Logout(ctx context.Context, req *connect.Request[v1.User]) (*connect.Response[v1.ActionReply], error) {
	return c.logout.CallUnary(ctx, req)
}

// AuthenticateServiceHandler is an implementation of the core.v1.AuthenticateService service.
type AuthenticateServiceHandler interface {
	PreLogin(context.Context, *connect.Request[v1.User]) (*connect.Response[v1.ActionReply], error)
	Login(context.Context, *connect.Request[v1.User]) (*connect.Response[v1.LoginReply], error)
	Logout(context.Context, *connect.Request[v1.User]) (*connect.Response[v1.ActionReply], error)
}

// NewAuthenticateServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuthenticateServiceHandler(svc AuthenticateServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	authenticateServicePreLoginHandler := connect.NewUnaryHandler(
		AuthenticateServicePreLoginProcedure,
		svc.PreLogin,
		connect.WithSchema(authenticateServicePreLoginMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	authenticateServiceLoginHandler := connect.NewUnaryHandler(
		AuthenticateServiceLoginProcedure,
		svc.Login,
		connect.WithSchema(authenticateServiceLoginMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	authenticateServiceLogoutHandler := connect.NewUnaryHandler(
		AuthenticateServiceLogoutProcedure,
		svc.Logout,
		connect.WithSchema(authenticateServiceLogoutMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/core.v1.AuthenticateService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AuthenticateServicePreLoginProcedure:
			authenticateServicePreLoginHandler.ServeHTTP(w, r)
		case AuthenticateServiceLoginProcedure:
			authenticateServiceLoginHandler.ServeHTTP(w, r)
		case AuthenticateServiceLogoutProcedure:
			authenticateServiceLogoutHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAuthenticateServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAuthenticateServiceHandler struct{}

func (UnimplementedAuthenticateServiceHandler) PreLogin(context.Context, *connect.Request[v1.User]) (*connect.Response[v1.ActionReply], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.AuthenticateService.preLogin is not implemented"))
}

func (UnimplementedAuthenticateServiceHandler) Login(context.Context, *connect.Request[v1.User]) (*connect.Response[v1.LoginReply], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.AuthenticateService.login is not implemented"))
}

func (UnimplementedAuthenticateServiceHandler) Logout(context.Context, *connect.Request[v1.User]) (*connect.Response[v1.ActionReply], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.AuthenticateService.logout is not implemented"))
}
