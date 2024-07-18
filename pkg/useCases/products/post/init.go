package post

import (
	products "e-commer/infrastructure/data/repositoryProducts"
	productsDomain "e-commer/pkg/domain/products"
)

type Ipost interface {
	CreateProduct(product productsDomain.CreateProductRequest) productsDomain.CreateProductResponse
}

type Post struct {
	RepositoryProducts products.Repository
}
