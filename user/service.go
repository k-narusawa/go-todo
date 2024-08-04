package user

import (
	"go-todo/domain"
	"go-todo/domain/repository"
	"go-todo/domain/value"
)

type Service struct {
	userRepo repository.UserRepository
}

func NewService(userRepo repository.UserRepository) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

type RegisterUserInput struct {
	Name string `json:"name"`
}

func (s *Service) Register(in RegisterUserInput) (*domain.User, error) {
	user := domain.NewUser(in.Name)

	if err := s.userRepo.Store(user); err != nil {
		return nil, err
	}

	return user, nil
}

type GetUserInput struct {
	UserID value.UserID `json:"user_id"`
}

func (s *Service) Get(in GetUserInput) (*domain.User, error) {
	userID := in.UserID
	user, err := s.userRepo.FindByID(userID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetAll() ([]*domain.User, error) {
	return s.userRepo.FindAll()
}

type UpdateUserInput struct {
	UserID value.UserID `json:"user_id"`
	Name   string       `json:"name"`
}

func (s *Service) Update(in UpdateUserInput) (*domain.User, error) {
	user, err := s.userRepo.FindByID(in.UserID)
	if err != nil {
		return nil, err
	}

	user.ChangeName(in.Name)
	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}
