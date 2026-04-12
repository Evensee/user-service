package user

import (
	"errors"
	"net/mail"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
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

	println("Looking for user with the same email!")
	if _, err = mail.ParseAddress(createUser.Email); err != nil {
		panic(status.Error(codes.InvalidArgument, "provided email is invalid"))
	}

	userWithSameEmail, _ := s.userRepo.GetOne(
		&FindUser{Email: &createUser.Email},
	)

	if userWithSameEmail != nil {
		panic(status.Error(codes.InvalidArgument, "provided email is used"))
	}
	println("User with email not has been found, creating new \n")
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(status.Error(codes.NotFound, "user does not exist"))
		}

		panic(err)
	}

	return user, err
}
