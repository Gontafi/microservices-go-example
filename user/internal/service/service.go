package service

import (
	"context"
	"fmt"
	"log"
	"user/internal/models"
	"user/internal/repository"
	"user/pkg/utils"
)

type Service struct {
	Repository *repository.Repo
}

func NewService(repo *repository.Repo) *Service {
	return &Service{Repository: repo}
}

func (s *Service) RegisterUser(ctx context.Context, user models.User, profile models.UserProfile) error {
	if !utils.IsEmailValid(user.Email) {
		return fmt.Errorf("invalid email format")
	}

	passwordHash, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = passwordHash

	err = s.Repository.CreateUser(ctx, user, profile)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) LoginUser(ctx context.Context, username string, password string) (*models.User, error) {
	user, err := s.Repository.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, fmt.Errorf("wrong password or username")
	}

	return user, nil
}

func (s *Service) ResetPassword(ctx context.Context, username string, email string, password string, code string) error {
	user, err := s.Repository.GetUserByUsername(ctx, username)
	if err != nil {
		return err
	}

	if email != "" {
		if !utils.IsEmailValid(email) {
			return fmt.Errorf("invalid email format")
		}

		user.Email = email
	}

	if code != "9999" {
		return fmt.Errorf("wrong code(try 9999)")
	}

	log.Println(password)

	hashPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user.Password = hashPassword

	err = s.Repository.UpdateUserByID(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ChangeUserProfile(ctx context.Context, profile models.UserProfile, userID int64) error {
	err := s.Repository.UpdateUserProfileByUserID(ctx, userID, profile)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetUserProfile(ctx context.Context, userID int64) (*models.UserProfile, error) {
	profile, err := s.Repository.GetUserProfileByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return profile, nil
}
