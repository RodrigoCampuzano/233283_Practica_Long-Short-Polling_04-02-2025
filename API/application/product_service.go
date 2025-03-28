package service

import (
	"ApiShortLong/domain/entities"
	"ApiShortLong/domain/repo"
)

type productService struct {
	repo repo.ProductRepository
}

func NewProductService(repo repo.ProductRepository) repo.ProductService {
	return &productService{repo: repo}
}

func (s *productService) AddProduct(product entities.Product) error {
	return s.repo.AddProduct(product)
}

func (s *productService) GetLastAddedProducts(limit int) ([]entities.Product, error) {
    return s.repo.GetLastAddedProducts(limit) 
}

func (s *productService) CountProductsInDiscount() (int, error) {
	return s.repo.CountProductsInDiscount()
}