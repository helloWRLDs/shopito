package productsdelivery

import "github.com/go-chi/chi"

func (d *ProductDeliveryImpl) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", d.InsertProduct)
	r.Get("/", d.GetProducts)
	r.Get("/{id}", d.GetProduct)
	r.Put("/{id}", d.UpdateProduct)
	r.Delete("/{id}", d.DeleteProduct)

	return r
}
