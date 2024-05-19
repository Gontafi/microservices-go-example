package grpcapi

import (
	"context"
	"log"
	productpb "product/gen/go/product"
	"product/internal/models"
	"product/internal/service"
)

type ServerAPI struct {
	productpb.UnimplementedProductServiceServer
	productService *service.Service
}

func NewServerAPI(productService *service.Service) *ServerAPI {
	return &ServerAPI{
		productService: productService,
	}
}

func (s *ServerAPI) CreateProduct(ctx context.Context, req *productpb.CreateProductRequest) (*productpb.CreateProductResponse, error) {
	product := models.Product{
		Name:        req.Product.Name,
		Description: req.Product.Description,
		Price:       float64(req.Product.Price),
		Category:    req.Product.Category,
	}

	productID, err := s.productService.CreateProduct(ctx, product)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := &productpb.CreateProductResponse{
		ProductId: productID,
	}

	return res, nil
}

func (s *ServerAPI) GetProduct(ctx context.Context, req *productpb.GetProductRequest) (*productpb.GetProductResponse, error) {
	product, err := s.productService.GetProductByID(ctx, req.ProductId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := &productpb.GetProductResponse{
		Product: &productpb.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       float32(product.Price),
			Category:    product.Category,
			CreatedAt:   product.CreatedAt.String(),
			UpdatedAt:   product.UpdatedAt.String(),
		},
	}

	return res, nil
}

func (s *ServerAPI) UpdateProduct(ctx context.Context, req *productpb.UpdateProductRequest) (*productpb.UpdateProductResponse, error) {
	product := models.Product{
		ID:          req.Product.Id,
		Name:        req.Product.Name,
		Description: req.Product.Description,
		Price:       float64(req.Product.Price),
		Category:    req.Product.Category,
	}

	err := s.productService.UpdateProductByID(ctx, product)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := &productpb.UpdateProductResponse{}

	return res, nil
}

func (s *ServerAPI) DeleteProduct(ctx context.Context, req *productpb.DeleteProductRequest) (*productpb.DeleteProductResponse, error) {
	err := s.productService.DeleteProductByID(ctx, req.ProductId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := &productpb.DeleteProductResponse{}

	return res, nil
}
