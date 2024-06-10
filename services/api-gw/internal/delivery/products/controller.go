package productcontroller

// import (
// 	"fmt"
// 	"net/http"
// 	grpcutil "shopito/pkg/util/grpc"
// 	jsonutil "shopito/pkg/util/json"
// 	productservice "shopito/services/api-gw/internal/service/products"
// 	"shopito/services/api-gw/protobuf"

// 	"github.com/go-chi/chi"
// )

// type ProductController struct {
// 	service productservice.Service
// }

// func New(service *productservice.ProductService) *ProductController {
// 	return &ProductController{
// 		service: service,
// 	}
// }

// func (c *ProductController) ProductRoutes() chi.Router {
// 	r := chi.NewRouter()

// 	r.Get("/{id}", c.GetProductController)
// 	r.Get("/", c.ListProductController)
// 	r.Post("/", c.CreateProductController)
// 	r.Delete("/{id}", c.DeleteProductController)
// 	r.Put("/{id}", c.UpdateProductController)

// 	return r
// }

// func (c *ProductController) ListProductController(w http.ResponseWriter, r *http.Request) {
// }

// func (c *ProductController) GetProductController(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")
// 	product, err := c.service.GetProductService(id)
// 	if err != nil {
// 		status, msg := grpcutil.GRPCToHTTPError(err)
// 		jsonutil.EncodeJson(w, status, msg)
// 		return
// 	}
// 	jsonutil.EncodeJson(w, 200, product)
// }

// func (c *ProductController) DeleteProductController(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")
// 	err := c.service.DeleteProductService(id)
// 	if err != nil {
// 		status, msg := grpcutil.GRPCToHTTPError(err)
// 		jsonutil.EncodeJson(w, status, msg)
// 		return
// 	}
// 	jsonutil.EncodeJson(w, 200, fmt.Sprintf("deleted product with id=%v", id))
// }

// func (c *ProductController) UpdateProductController(w http.ResponseWriter, r *http.Request) {
// 	// id := chi.URLParam(r, "id")
// 	// if err != nil {
// 	// 	status, msg := grpcutil.GRPCToHTTPError(err)
// 	// 	jsonutil.EncodeJson(w, status, msg)
// 	// 	return
// 	// }
// }

// func (c *ProductController) CreateProductController(w http.ResponseWriter, r *http.Request) {
// 	product, err := jsonutil.DecodeJson[protobuf.CreateProductRequest](r)
// 	if err != nil {
// 		jsonutil.EncodeJson(w, 400, "Bad Gateway")
// 		return
// 	}
// 	id, err := c.service.CreateProductService(&product)
// 	if err != nil {
// 		status, msg := grpcutil.GRPCToHTTPError(err)
// 		jsonutil.EncodeJson(w, status, msg)
// 		return
// 	}
// 	jsonutil.EncodeJson(w, 201, fmt.Sprintf("Created product with id = %v", id))
// }
