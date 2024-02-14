package repository

import (
    "context"
    "database/sql"
    "shopping/internal/models"
    "log"
)

type ProductRepository struct {
    db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
    return &ProductRepository{db: db}
}

func (r *ProductRepository) GetProduct(ctx context.Context, id int) (*models.Product, error) {
	product := &models.Product{1, "Product 1", 45.5, "https://googld,s,.com", 0.0, 10, "Description 1", 1}
	return product, nil
}

func (r *ProductRepository) AddProduct(ctx context.Context, product models.Product) (bool, error) {
    _, err := r.db.ExecContext(ctx, "INSERT INTO product (name, price, image_path, discount, quantity, description, category_id) VALUES (?, ?, ?, ?, ?, ?, ?)", product.Name, product.Price, product.Image, product.Discount, product.Quantity, product.Description, product.CategoryId)
    if err != nil {
        log.Fatal(err)
        return false, err
    }
    return true, nil
}

func (r *ProductRepository) GetProducts(ctx context.Context) ([]*models.Product, error) {
    rows, err := r.db.QueryContext(ctx, "SELECT id, name, price, image_path, discount, quantity, description, category_id FROM product")
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    defer rows.Close()

    products := make([]*models.Product, 0)
    for rows.Next() {
        var p models.Product
        if err := rows.Scan(&p.Id, &p.Name, &p.Price, &p.Image, &p.Discount, &p.Quantity, &p.Description, &p.CategoryId); err != nil {
            return nil, err
        }
        products = append(products, &p)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return products, nil
}