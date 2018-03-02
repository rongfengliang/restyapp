// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: pb/echoservice.proto

/*
Package echoservice is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package echoservice

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_EchoService_Echo_0(ctx context.Context, marshaler runtime.Marshaler, client EchoServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq EchoMessage
	var metadata runtime.ServerMetadata

	if req.ContentLength > 0 {
		if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
			return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
		}
	}

	msg, err := client.Echo(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_UserServiceApp_Message_0(ctx context.Context, marshaler runtime.Marshaler, client UserServiceAppClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AppMessage
	var metadata runtime.ServerMetadata

	if req.ContentLength > 0 {
		if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
			return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
		}
	}

	msg, err := client.Message(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterEchoServiceHandlerFromEndpoint is same as RegisterEchoServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterEchoServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterEchoServiceHandler(ctx, mux, conn)
}

// RegisterEchoServiceHandler registers the http handlers for service EchoService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterEchoServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterEchoServiceHandlerClient(ctx, mux, NewEchoServiceClient(conn))
}

// RegisterEchoServiceHandler registers the http handlers for service EchoService to "mux".
// The handlers forward requests to the grpc endpoint over the given implementation of "EchoServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "EchoServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "EchoServiceClient" to call the correct interceptors.
func RegisterEchoServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client EchoServiceClient) error {

	mux.Handle("POST", pattern_EchoService_Echo_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_EchoService_Echo_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_EchoService_Echo_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_EchoService_Echo_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 1}, []string{"v1", "echo"}, ""))
)

var (
	forward_EchoService_Echo_0 = runtime.ForwardResponseMessage
)

// RegisterUserServiceAppHandlerFromEndpoint is same as RegisterUserServiceAppHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterUserServiceAppHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterUserServiceAppHandler(ctx, mux, conn)
}

// RegisterUserServiceAppHandler registers the http handlers for service UserServiceApp to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterUserServiceAppHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterUserServiceAppHandlerClient(ctx, mux, NewUserServiceAppClient(conn))
}

// RegisterUserServiceAppHandler registers the http handlers for service UserServiceApp to "mux".
// The handlers forward requests to the grpc endpoint over the given implementation of "UserServiceAppClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "UserServiceAppClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "UserServiceAppClient" to call the correct interceptors.
func RegisterUserServiceAppHandlerClient(ctx context.Context, mux *runtime.ServeMux, client UserServiceAppClient) error {

	mux.Handle("POST", pattern_UserServiceApp_Message_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_UserServiceApp_Message_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_UserServiceApp_Message_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_UserServiceApp_Message_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "user", "echo"}, ""))
)

var (
	forward_UserServiceApp_Message_0 = runtime.ForwardResponseMessage
)
