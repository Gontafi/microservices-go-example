package service

import (
	"context"
	"log"
	pb "orders/gen/go/product"
	"orders/internal/models"
	"orders/internal/repository"
	"time"
)

type Service struct {
	Repository    *repository.Repo
	ProductClient pb.ProductServiceClient
}

func NewService(repo *repository.Repo, productClient pb.ProductServiceClient) *Service {
	return &Service{Repository: repo, ProductClient: productClient}
}

func (s *Service) CreateOrder(ctx context.Context, order models.Order, items []models.OrderItem) (int64, error) {
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	for i := range items {
		items[i].CreatedAt = time.Now()
		items[i].UpdatedAt = time.Now()

		// Check if the product exists
		req := &pb.GetProductRequest{ProductId: items[i].ProductID}
		_, err := s.ProductClient.GetProduct(ctx, req)
		if err != nil {
			log.Printf("Product with ID %d does not exist: %v", items[i].ProductID, err)
			return 0, err
		}
	}

	orderID, err := s.Repository.CreateOrder(ctx, order, items)
	if err != nil {
		return 0, err
	}

	return orderID, nil
}

func (s *Service) GetOrderByID(ctx context.Context, orderID int64) (*models.Order, error) {
	order, err := s.Repository.GetOrderByID(ctx, orderID)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *Service) GetOrderItemsByOrderID(ctx context.Context, orderID int64) ([]models.OrderItem, error) {
	items, err := s.Repository.GetOrderItemsByOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (s *Service) UpdateOrderStatus(ctx context.Context, orderID int64, status string) error {
	err := s.Repository.UpdateOrderStatus(ctx, orderID, status)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteOrderByID(ctx context.Context, orderID int64) error {
	err := s.Repository.DeleteOrderByID(ctx, orderID)
	if err != nil {
		return err
	}

	return nil
}
