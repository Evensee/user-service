package api

import (
	"context"
	"fmt"
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

	fmt.Printf("resp: %v err: %v \n", resp, err)
	if err != nil {
		log.Printf("[gRPC] Error in %s: %v", info.FullMethod, err)
	}
	return resp, err
}
