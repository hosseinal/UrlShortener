package service

import (
	"context"

	"github.com/hosseinal/UrlShortner/internal/model"
	"github.com/hosseinal/UrlShortner/internal/repository/helper"
	"github.com/hosseinal/UrlShortner/internal/repository/user"
)

type AuthService interface {
	Register(ctx context.Context, username string, email string, password string) (model.User, error)
	Login(ctx context.Context, email string, password string) (model.User, error)
	GetUserByID(ctx context.Context, id string) (model.User, error)
}

type authService struct {
	userRepository user.UserRepository
}

func NewAuthService(userRepository user.UserRepository) AuthService {
	return &authService{userRepository: userRepository}
}

func (s *authService) Register(ctx context.Context, username string, email string, password string) (model.User, error) {
	hashedPassword, err := helper.HashPassword(password)
	if err != nil {
		return model.User{}, err
	}
	user := model.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}
	err = s.userRepository.CreateUser(ctx, &user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (s *authService) Login(ctx context.Context, email string, password string) (model.User, error) {
	user, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return model.User{}, err
	}
	return *user, nil
}

func (s *authService) GetUserByID(ctx context.Context, id string) (model.User, error) {
	user, err := s.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return model.User{}, err
	}
	return *user, nil
}
