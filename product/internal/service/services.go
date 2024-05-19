package service

import (
	"context"
	"product/internal/models"
	"product/internal/repository"
	"time"
)

type Service struct {
	Repository *repository.Repo
}

func NewService(repo *repository.Repo) *Service {
	return &Service{Repository: repo}
}

func (s *Service) CreateProduct(ctx context.Context, product models.Product) (int64, error) {
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	productID, err := s.Repository.CreateProduct(ctx, product)
	if err != nil {
		return 0, err
	}

	return productID, nil
}

func (s *Service) GetProductByID(ctx context.Context, productID int64) (*models.Product, error) {
	product, err := s.Repository.GetProductByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Service) UpdateProductByID(ctx context.Context, product models.Product) error {
	product.UpdatedAt = time.Now()

	err := s.Repository.UpdateProductByID(ctx, product)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteProductByID(ctx context.Context, productID int64) error {
	err := s.Repository.DeleteProductByID(ctx, productID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) CreateReview(ctx context.Context, review models.Review) (int64, error) {
	review.CreatedAt = time.Now()

	reviewID, err := s.Repository.CreateReview(ctx, review)
	if err != nil {
		return 0, err
	}

	return reviewID, nil
}

func (s *Service) GetReviewsByProductID(ctx context.Context, productID int64) ([]models.Review, error) {
	reviews, err := s.Repository.GetReviewsByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (s *Service) DeleteReviewByID(ctx context.Context, reviewID int64) error {
	err := s.Repository.DeleteReviewByID(ctx, reviewID)
	if err != nil {
		return err
	}

	return nil
}
