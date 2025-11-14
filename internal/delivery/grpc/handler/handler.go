package handler

import "github.com/Evensee/user-service/internal/interface/resolver"

type ServerApi struct {
	appResolver resolver.AppResolver
}

func New(appResolver resolver.AppResolver) ServerApi {
	return ServerApi{
		appResolver: appResolver,
	}
}
