package api

import (
	"github.com/Evensee/user-service/internal/interface/resolver"
	"github.com/Evensee/user-service/protobuf_generated/user"
	"google.golang.org/grpc"
)

type serverApi struct {
	user.UnimplementedUserServiceServer
	appResolver resolver.AppResolver
}

func New(appResolver resolver.AppResolver) *serverApi {
	return &serverApi{
		appResolver: appResolver,
	}
}

func Register(gRPC *grpc.Server, serverApiInstance *serverApi) {
	user.RegisterUserServiceServer(
		gRPC,
		serverApiInstance,
	)
}
