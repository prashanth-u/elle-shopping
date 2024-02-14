package service

import (
    "context"
    "shopping/internal/repository"
    "shopping/internal/models"
)

type CategoryService struct {
    repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
    return &CategoryService{repo: repo}
}

func (s *CategoryService) GetCategory(ctx context.Context, id int) (*models.Category, error) {
    return s.repo.GetCategory(ctx, id)
}