package handler

import (
	"github.com/Evensee/user-service/internal/interface/resolver"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ServerApi struct {
	appResolver resolver.AppResolver
}

func New(appResolver resolver.AppResolver) ServerApi {
	return ServerApi{
		appResolver: appResolver,
	}
}

func (s *ServerApi) HealthCheck(ctx Ctx, req *emptypb.Empty) (*emptypb.Empty, error) {
	return new(emptypb.Empty), nil
}
