package product

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHeadler(s Service) http.Handler {
	r := chi.NewRouter()
	getProductByIdHandler := kithttp.NewServer(makeGetProductByIdEndPoint(s), getProductByIdRequestDecode, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getProductByIdHandler)

	getProductsHandler := kithttp.NewServer(makeGetProductsEndPoint(s), getProductsRequestDecode, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getProductsHandler)

	addProductsHandler := kithttp.NewServer(makeAddProductEndPoint(s), addProductRequestDecode, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", addProductsHandler)

	updateProductsHandler := kithttp.NewServer(makeUpdateProductEndPoint(s), updateProductRequestDecode, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPut, "/", updateProductsHandler)

	deleteProductsHandler := kithttp.NewServer(makeDeleteProductEndPoint(s), deleteProductRequestDecode, kithttp.EncodeJSONResponse)
	r.Method(http.MethodDelete, "/{id}", deleteProductsHandler)

	return r
}

func getProductByIdRequestDecode(context context.Context, r *http.Request) (interface{}, error) {
	productId, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return getProductByIDRequest{
		ProductID: productId,
	}, nil
}

func getProductsRequestDecode(context context.Context, r *http.Request) (interface{}, error) {
	request := getProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	return request, nil
}

func addProductRequestDecode(context context.Context, r *http.Request) (interface{}, error) {
	request := getAddProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	return request, nil
}

func updateProductRequestDecode(context context.Context, r *http.Request) (interface{}, error) {
	request := getUpdateProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	return request, nil
}

func deleteProductRequestDecode(context context.Context, r *http.Request) (interface{}, error) {
	return getDeleteProductRequest{
		id: chi.URLParam(r, "id"),
	}, nil
}
