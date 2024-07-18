package products

// Principio Solid S Single responsibility principle
type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type CreateProductResponse struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
