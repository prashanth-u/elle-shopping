package service

import (
    "context"
    "shopping/internal/repository"
    "shopping/internal/models"
)

type ProductService struct {
    repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
    return &ProductService{repo: repo}
}

func (s *ProductService) GetProduct(ctx context.Context, id int) (*models.Product, error) {
    return s.repo.GetProduct(ctx, id)
}

func (s *ProductService) GetProducts(ctx context.Context) ([]*models.Product, error) {
    return s.repo.GetProducts(ctx)
}

func (s *ProductService) AddProduct(ctx context.Context, product models.Product) (bool, error) {
    return s.repo.AddProduct(ctx, product)
}