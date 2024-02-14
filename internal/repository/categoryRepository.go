package repository

import (
    "context"
    "database/sql"
    "shopping/internal/models"
)

type CategoryRepository struct {
    db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
    return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetCategory(ctx context.Context, id int) (*models.Category, error) {
	category := &models.Category{1, "Category 1"}
	return category, nil
}