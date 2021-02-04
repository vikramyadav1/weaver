package api_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vikramyadav1/weaver/sampleApp/api"
	"github.com/vikramyadav1/weaver/sampleApp/models/product"
	"github.com/vikramyadav1/weaver/sampleApp/models/product/mocks"
)

var mockRepository *mocks.ProductRepository = new(mocks.ProductRepository)
var product1 product.Product = product.Product{Id: 1, Name: "product1",
	Description: "product 1", Price: 10, GroupId: 1}
var ps api.ProductServer = api.ProductServer{Pr: mockRepository}
var router *mux.Router = new(mux.Router)

func init() {
	ps.RegisterRoutes(router)
}

func TestGetAllProducts(t *testing.T) {
	assert := assert.New(t)
	products := make([]product.Product, 0)
	products = append(products, product1)

	mockRepository.On("All").Return(products, nil)
	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	res := w.Result()

	expectedJson := `[{"id":1,"name":"product1","description":"product 1","price":10,"references":{"group_id":1}}]`

	assert.Equal(200, res.StatusCode)
	assert.Equal(expectedJson, strings.TrimSpace(w.Body.String()))
}

func TestGetProduct(t *testing.T) {
	assert := assert.New(t)

	mockRepository.On("Get", 1).Return(&product1, nil)
	req, _ := http.NewRequest("GET", "/products/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	res := w.Result()

	expectedJson := `{"id":1,"name":"product1","description":"product 1","price":10,"references":{"group_id":1}}`
	assert.Equal(expectedJson, strings.TrimSpace(w.Body.String()))

	mockRepository.AssertExpectations(t)
	assert.Equal(200, res.StatusCode)

	mockRepository.On("Get", 2).Return(nil, nil)
	req, _ = http.NewRequest("GET", "/products/2", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	res = w.Result()

	mockRepository.AssertExpectations(t)
	assert.Equal(res.StatusCode, 404)
}

func TestCreateProduct(t *testing.T) {
	assert := assert.New(t)

	responseBody := []byte(`{"name":"product1","description":"product 1","price":10,"group_id":2}`)
	mockRepository.On("Create", mock.Anything).Return(new(product.Product), nil)
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(responseBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	res := w.Result()

	mockRepository.AssertExpectations(t)
	assert.Equal(201, res.StatusCode)
}

func TestUpdateProduct(t *testing.T) {
	assert := assert.New(t)

	requestBody := []byte(`{"name":"product3","description":"product 3","price":100,"group_id":3}`)
	mockRepository.On("Update", 1, mock.Anything).Return(new(product.Product), nil)
	req, _ := http.NewRequest("PUT", "/products/1", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	res := w.Result()

	mockRepository.AssertExpectations(t)
	assert.Equal(200, res.StatusCode)
}

func TestDeleteProduct(t *testing.T) {
	assert := assert.New(t)

	mockRepository.On("Delete", 1).Return(nil)
	req, _ := http.NewRequest("DELETE", "/products/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	res := w.Result()

	mockRepository.AssertExpectations(t)
	assert.Equal(200, res.StatusCode)
}
