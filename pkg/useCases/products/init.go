package products

import "e-commer/pkg/useCases/products/post"

type Handler struct {
	Post post.Ipost
}

func (h *Handler) NewRecordsHandler() {
	h.Post = &post.Post{}
}
