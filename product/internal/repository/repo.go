package repository

import (
	"context"
	"database/sql"
	"product/internal/models"
)

type Repo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) CreateProduct(ctx context.Context, product models.Product) (int64, error) {
	query := `INSERT INTO products(name, description, price, category, created_at, updated_at)
				VALUES (?, ?, ?, ?, ?, ?)`

	res, err := r.db.ExecContext(ctx, query,
		product.Name,
		product.Description,
		product.Price,
		product.Category,
		product.CreatedAt,
		product.UpdatedAt,
	)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (r *Repo) GetProductByID(ctx context.Context, productID int64) (*models.Product, error) {
	query := `SELECT id, name, description, price, category, created_at, updated_at
			  FROM products
			  WHERE id = ?`

	product := models.Product{}
	err := r.db.QueryRowContext(ctx, query, productID).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Category,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *Repo) UpdateProductByID(ctx context.Context, product models.Product) error {
	query := `UPDATE products
			  SET name = ?, description = ?, price = ?, category = ?, updated_at = ?
			  WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query,
		product.Name,
		product.Description,
		product.Price,
		product.Category,
		product.UpdatedAt,
		product.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) DeleteProductByID(ctx context.Context, productID int64) error {
	query := `DELETE FROM products
			  WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, productID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) CreateReview(ctx context.Context, review models.Review) (int64, error) {
	query := `INSERT INTO reviews(product_id, user_id, rating, comment, created_at)
				VALUES (?, ?, ?, ?, ?)`

	res, err := r.db.ExecContext(ctx, query,
		review.ProductID,
		review.UserID,
		review.Rating,
		review.Comment,
		review.CreatedAt,
	)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (r *Repo) GetReviewsByProductID(ctx context.Context, productID int64) ([]models.Review, error) {
	query := `SELECT id, product_id, user_id, rating, comment, created_at
			  FROM reviews
			  WHERE product_id = ?`

	rows, err := r.db.QueryContext(ctx, query, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		review := models.Review{}
		err := rows.Scan(
			&review.ID,
			&review.ProductID,
			&review.UserID,
			&review.Rating,
			&review.Comment,
			&review.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *Repo) DeleteReviewByID(ctx context.Context, reviewID int64) error {
	query := `DELETE FROM reviews
			  WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, reviewID)
	if err != nil {
		return err
	}

	return nil
}
