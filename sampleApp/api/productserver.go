package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vikramyadav1/weaver/sampleApp/models/product"
	"net/http"
	"strconv"
)

type ProductJson struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`

	References ProductReference `json:"references"`
}

type ProductReference struct {
	GroupId int `json:"group_id"`
}

func toJsonProduct(p product.Product) ProductJson {
	return ProductJson{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,

		References: ProductReference{
			GroupId: p.GroupId,
		},
	}
}

func toDBProduct(p ProductJson) product.Product {
	return product.Product{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		GroupId:     p.References.GroupId,
	}
}

type ProductServer struct {
	Pr product.ProductRepository
}

func (ps *ProductServer) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", ps.GetAllProducts).Methods("GET")
	router.HandleFunc("/products", ps.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", ps.GetProduct).Methods("GET")
	router.HandleFunc("/products/{id}", ps.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", ps.DeleteProduct).Methods("DELETE")
}

func (ps *ProductServer) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	jsonProducts := make([]ProductJson, 0)
	products, _ := ps.Pr.All()

	for _, product := range products {
		jsonProducts = append(jsonProducts, toJsonProduct(product))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonProducts)
}

func (ps *ProductServer) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	productId, _ := strconv.Atoi(id)
	p, err := ps.Pr.Get(productId)

	if p == nil && err == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if p == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toJsonProduct(*p))
}

func (ps *ProductServer) CreateProduct(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var pj ProductJson
	decoder.Decode(&pj)

	p, err := ps.Pr.Create(toDBProduct(pj))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toJsonProduct(*p))
}

func (ps *ProductServer) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	productId, _ := strconv.Atoi(id)

	decoder := json.NewDecoder(r.Body)
	var pj ProductJson
	decoder.Decode(&pj)

	p, err := ps.Pr.Update(productId, toDBProduct(pj))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toJsonProduct(*p))
}

func (ps *ProductServer) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	productId, _ := strconv.Atoi(id)

	err := ps.Pr.Delete(productId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
