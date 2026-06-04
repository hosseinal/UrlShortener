package service

import (
	"context"
	"strconv"

	"github.com/hosseinal/UrlShortner/internal/model"
	"github.com/hosseinal/UrlShortner/internal/repository/user"
)

type AuthService interface {
	Register(ctx context.Context, username string, email string, password string) error
	Login(ctx context.Context, email string, password string) (string, error)
}

type authService struct {
	userRepository user.UserRepository
}

func NewAuthService(userRepository user.UserRepository) AuthService {
	return &authService{userRepository: userRepository}
}

func (s *authService) Register(ctx context.Context, username string, email string, password string) error {
	return s.userRepository.CreateUser(ctx, &model.User{
		Username: username,
		Email:    email,
		Password: password,
	})
}

func (s *authService) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	return strconv.FormatUint(uint64(user.ID), 10), nil
}
