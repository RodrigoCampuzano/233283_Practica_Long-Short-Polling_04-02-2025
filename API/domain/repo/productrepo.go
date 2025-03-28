package repo

import "ApiShortLong/domain/entities"

type ProductRepository interface {
    AddProduct(product entities.Product) error
    GetLastAddedProducts(limit int) ([]entities.Product, error) 
    CountProductsInDiscount() (int, error)
}

type ProductService interface {
    AddProduct(product entities.Product) error
    GetLastAddedProducts(limit int) ([]entities.Product, error) 
    CountProductsInDiscount() (int, error)
}