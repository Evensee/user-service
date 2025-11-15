package user

import (
	"fmt"

	"github.com/google/uuid"
)

type DomainUserService struct {
	userRepo Repository
}

func NewUserDomainService(
	userRepo Repository,
) *DomainUserService {
	return &DomainUserService{userRepo: userRepo}
}

func (s *DomainUserService) Create(createUser *CreateUser) (*User, error) {
	createUserModel, err := NewUser(createUser)
	if err != nil {
		panic(err)
	}
	fmt.Printf("service create user: %v \n", createUserModel)

	fmt.Printf("repo in service %p \n", &s.userRepo)

	user, err := s.userRepo.CreateUser(createUserModel)
	if err != nil {
		panic(err)
	}

	return user, err
}

func (s *DomainUserService) Update(user_id uuid.UUID, updateUser *UpdateUser) (*User, error) {
	user, err := s.userRepo.Update(user_id, updateUser)
	if err != nil {
		panic(err)
	}
	return user, err
}

func (s *DomainUserService) GetAll(findUser *FindUser) (*[]User, error) {
	users, err := s.userRepo.GetAll(findUser)

	if err != nil {
		panic(err)
	}

	return users, err
}

func (s *DomainUserService) GetOne(findUser *FindUser) (*User, error) {
	user, err := s.userRepo.GetOne(findUser)

	if err != nil {
		panic(err)
	}

	return user, err
}
