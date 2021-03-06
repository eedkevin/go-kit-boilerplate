// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 7040e72f5f
// Version Date: 2020-09-19T18:42:02Z

package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/eedkevin/go-kit-boilerplate/geo"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC GeoServer.
func MakeGRPCServer(endpoints Endpoints, options ...grpctransport.ServerOption) pb.GeoServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	serverOptions = append(serverOptions, options...)
	return &grpcServer{
		// geo

		status: grpctransport.NewServer(
			endpoints.StatusEndpoint,
			DecodeGRPCStatusRequest,
			EncodeGRPCStatusResponse,
			serverOptions...,
		),
		distance: grpctransport.NewServer(
			endpoints.DistanceEndpoint,
			DecodeGRPCDistanceRequest,
			EncodeGRPCDistanceResponse,
			serverOptions...,
		),
	}
}

// grpcServer implements the GeoServer interface
type grpcServer struct {
	status   grpctransport.Handler
	distance grpctransport.Handler
}

// Methods for grpcServer to implement GeoServer interface

func (s *grpcServer) Status(ctx context.Context, req *pb.StatusRequest) (*pb.StatusResponse, error) {
	_, rep, err := s.status.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.StatusResponse), nil
}

func (s *grpcServer) Distance(ctx context.Context, req *pb.DistanceRequest) (*pb.DistanceResponse, error) {
	_, rep, err := s.distance.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DistanceResponse), nil
}

// Server Decode

// DecodeGRPCStatusRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC status request to a user-domain status request. Primarily useful in a server.
func DecodeGRPCStatusRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.StatusRequest)
	return req, nil
}

// DecodeGRPCDistanceRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC distance request to a user-domain distance request. Primarily useful in a server.
func DecodeGRPCDistanceRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.DistanceRequest)
	return req, nil
}

// Server Encode

// EncodeGRPCStatusResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain status response to a gRPC status reply. Primarily useful in a server.
func EncodeGRPCStatusResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.StatusResponse)
	return resp, nil
}

// EncodeGRPCDistanceResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain distance response to a gRPC distance reply. Primarily useful in a server.
func EncodeGRPCDistanceResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.DistanceResponse)
	return resp, nil
}

// Helpers

func metadataToContext(ctx context.Context, md metadata.MD) context.Context {
	for k, v := range md {
		if v != nil {
			// The key is added both in metadata format (k) which is all lower
			// and the http.CanonicalHeaderKey of the key so that it can be
			// accessed in either format
			ctx = context.WithValue(ctx, k, v[0])
			ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
		}
	}

	return ctx
}
