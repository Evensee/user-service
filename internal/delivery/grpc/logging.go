package api

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func LoggingInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	log.Printf("[gRPC] Method called: %s", info.FullMethod)
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("[gRPC] Error in %s: %v", info.FullMethod, err)
	}
	return resp, err
}
