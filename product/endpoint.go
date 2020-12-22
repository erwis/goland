package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getProductByIDRequest struct {
	ProductID int
}

type getProductRequest struct {
	Limit  int
	Offset int
}

type getAddProductRequest struct {
	Category     string
	Description  string
	ListPrice    string
	StandardCost string
	ProductCode  string
	ProductName  string
}

type getUpdateProductRequest struct {
	ID           int
	Category     string
	Description  string
	ListPrice    float32
	StandardCost float32
	ProductCode  string
	ProductName  string
}

type getDeleteProductRequest struct {
	id string
}

func makeGetProductByIdEndPoint(s Service) endpoint.Endpoint {
	getProductByIdEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductByIDRequest)
		product, err := s.GetProductById(&req)
		if err != nil {
			panic(err)
		}
		return product, nil
	}
	return getProductByIdEndpoint
}

func makeGetProductsEndPoint(s Service) endpoint.Endpoint {
	getProductsEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductRequest)
		results, err := s.GetProducts(&req)
		if err != nil {
			panic(err)
		}
		return results, nil
	}
	return getProductsEndpoint
}

func makeAddProductEndPoint(s Service) endpoint.Endpoint {
	addProductEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddProductRequest)
		productId, err := s.InsertProduct(&req)
		if err != nil {
			panic(err)
		}
		return productId, nil
	}
	return addProductEndpoint
}

func makeUpdateProductEndPoint(s Service) endpoint.Endpoint {
	updateProductEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getUpdateProductRequest)
		result, err := s.UpdateProduct(&req)
		if err != nil {
			panic(err)
		}
		return result, nil
	}
	return updateProductEndpoint
}

func makeDeleteProductEndPoint(s Service) endpoint.Endpoint {
	deleteProductEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getDeleteProductRequest)
		result, err := s.DeleteProduct(&req)
		if err != nil {
			panic(err)
		}
		return result, nil
	}
	return deleteProductEndpoint
}
