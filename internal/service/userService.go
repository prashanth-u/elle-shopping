package service

import (
    "context"
    "shopping/internal/repository"
    "shopping/internal/models"
)

type UserService struct {
    repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) ValidateRequest(ctx context.Context, sessionId string) (models.User, error) {
    return s.repo.ValidateUser(ctx, sessionId)
}

func (s *UserService) Register(ctx context.Context, username, password, firstName, lastName, token string) error {
    return s.repo.Register(ctx, username, password, firstName, lastName, token)
}