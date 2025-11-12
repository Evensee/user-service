package user

import "github.com/google/uuid"

type DomainUserService struct {
	userRepo Repository
}

func NewUserDomainService(
	userRepo Repository,
) *DomainUserService {
	return &DomainUserService{userRepo: userRepo}
}

func (s *DomainUserService) Create(ctx Ctx, createUser *CreateUser) (*User, error) {
	createUserModel, err := NewUser(createUser)
	if err != nil {
		panic(err)
	}
	user, err := s.userRepo.Create(ctx, createUserModel)
	if err != nil {
		panic(err)
	}

	return user, err
}

func (s *DomainUserService) Update(ctx Ctx, user_id uuid.UUID, updateUser *UpdateUser) (*User, error) {
	user, err := s.userRepo.Update(ctx, user_id, updateUser)
	if err != nil {
		panic(err)
	}
	return user, err
}

func (s *DomainUserService) GetAll(ctx Ctx, findUser *FindUser) (*[]User, error) {
	users, err := s.userRepo.GetAll(ctx, findUser)

	if err != nil {
		panic(err)
	}

	return users, err
}

func (s *DomainUserService) GetOne(ctx Ctx, findUser *FindUser) (*User, error) {
	user, err := s.userRepo.GetOne(ctx, findUser)

	if err != nil {
		panic(err)
	}

	return user, err
}
