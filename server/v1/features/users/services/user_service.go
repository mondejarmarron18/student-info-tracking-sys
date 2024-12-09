package services

import (
	"errors"
	"server/v1/features/users/domains"
	"server/v1/features/users/repositories"
	"server/v1/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo     *repositories.UserRepo
	errorMessage utils.ErrorMessage
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repositories.NewUserRepo(),
	}
}
func (s *UserService) CreateUser(user domains.User) (domains.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, errors.New(s.errorMessage.BadRequest)
	}

	user.Password = string(hashedPassword)

	createdUser, err := s.userRepo.CreateUser(&user)
	if err != nil {
		return user, err
	}

	return *createdUser, nil
}

func (s *UserService) GetUsers(filter utils.Filter) ([]domains.User, error) {
	users, err := s.userRepo.GetUsers(filter)

	if err != nil {
		return nil, err
	}
	return *users, nil
}

func (s *UserService) GetUserById(id string) (domains.User, error) {
	user, err := s.userRepo.GetUserById(id)
	if err != nil {
		return domains.User{}, err
	}
	return *user, nil
}
