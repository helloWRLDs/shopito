package productsdelivery

import (
	"database/sql"
	"fmt"
	"net/http"
	productdomain "shopito/api/internal/domain/product"
	productusecase "shopito/api/internal/usecase/product"
	"shopito/api/pkg/types/errors"
	"shopito/api/pkg/types/response"
	jsonutil "shopito/api/pkg/util/json"
	"strconv"

	"github.com/go-chi/chi"
)

type ProductDeliveryImpl struct {
	uc productusecase.ProductUseCase
}

func New(db *sql.DB) *ProductDeliveryImpl {
	return &ProductDeliveryImpl{
		uc: productusecase.New(db),
	}
}

func (d *ProductDeliveryImpl) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", d.InsertProduct)
	r.Get("/", d.GetProducts)
	r.Get("/{id}", d.GetProduct)
	r.Put("/{id}", d.UpdateProduct)
	r.Delete("/{id}", d.DeleteProduct)

	return r
}

func (d *ProductDeliveryImpl) GetProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		jsonutil.EncodeJson(w, errors.ErrUnpocessableEntity.Status(), errors.ErrUnpocessableEntity.SetMessage("couldn't process product's id"))
		return
	}
	product, error := d.uc.GetProduct(r.Context(), id)
	if error != nil {
		jsonutil.EncodeJson(w, error.Status(), error)
		return
	}
	jsonutil.EncodeJson(w, 200, product)
}

func (d *ProductDeliveryImpl) GetProducts(w http.ResponseWriter, r *http.Request) {

}

func (d *ProductDeliveryImpl) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		jsonutil.EncodeJson(w, errors.ErrUnpocessableEntity.Status(), errors.ErrUnpocessableEntity.SetMessage("couldn't process product's id"))
		return
	}
	product, err := jsonutil.DecodeJson[productdomain.Product](r)
	if err != nil {
		jsonutil.EncodeJson(w, errors.ErrUnpocessableEntity.Status(), errors.ErrUnpocessableEntity.SetMessage("couldn't process product's data"))
		return
	}
	if err := d.uc.UpdateProduct(r.Context(), id, &product); err != nil {
		jsonutil.EncodeJson(w, err.Status(), err)
		return
	}
	jsonutil.EncodeJson(w, 200, response.NewJsonMessage(200, fmt.Sprintf("product updated with id=%v", id)))
}

func (d *ProductDeliveryImpl) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		jsonutil.EncodeJson(w, errors.ErrUnpocessableEntity.Status(), errors.ErrUnpocessableEntity.SetMessage("couldn't process user's id"))
		return
	}
	if err := d.uc.DeleteProduct(r.Context(), id); err != nil {
		jsonutil.EncodeJson(w, err.Status(), err)
		return
	}
	jsonutil.EncodeJson(w, 200, response.NewJsonMessage(200, fmt.Sprintf("product deleted with id=%v", id)))
}

func (d *ProductDeliveryImpl) InsertProduct(w http.ResponseWriter, r *http.Request) {
	product, err := jsonutil.DecodeJson[productdomain.Product](r)
	if err != nil {
		jsonutil.EncodeJson(w, errors.ErrUnpocessableEntity.Status(), errors.ErrUnpocessableEntity.SetMessage("couldn't process user's data"))
		return
	}
	insertedId, error := d.uc.InsertProduct(r.Context(), &product)
	if error != nil {
		jsonutil.EncodeJson(w, error.Status(), error)
		return
	}
	jsonutil.EncodeJson(w, 201, fmt.Sprintf("created product with id=%v", insertedId))
}
