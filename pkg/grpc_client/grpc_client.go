package grpc_client

import (
	"fmt"
	grpcopentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCClientServiceConn(host string, port int) (*grpc.ClientConn, error) {
	clientGRPCConn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpcopentracing.UnaryClientInterceptor()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to auth service: %w", err)
	}

	return clientGRPCConn, nil
}
