package user

import (
	"context"

	"github.com/Evensee/user-service/internal"
)

type DomainUserService struct {
	userRepo  Repository
	appConfig internal.AppConfig
}

func NewUserDomainService(
	userRepo Repository,
	appConfig *internal.AppConfig,
) *DomainUserService {
	return &DomainUserService{userRepo: userRepo}
}

func (service *DomainUserService) Create(ctx context.Context, user User) (User, error) {
	return service.userRepo.Create(ctx, user)
}
