package user

import "context"

type DomainUserService struct {
	userRepo Repository
}

func NewUserDomainService(userRepo Repository) *DomainUserService {
	return &DomainUserService{userRepo: userRepo}
}

func (service *DomainUserService) Create(ctx context.Context, user User) (User, error) {
	return service.userRepo.Create(ctx, user)
}
