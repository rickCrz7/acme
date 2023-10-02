package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	dao ProductDao
}

func NewProductHandler(dao ProductDao, r *mux.Router) *ProductHandler {
	h := &ProductHandler{dao: dao}
	r.HandleFunc("/api/v1/products", h.getProducts).Methods("GET")
	r.HandleFunc("/api/v1/products/{id}", h.getProductById).Methods("GET")
	r.HandleFunc("/api/v1/products", h.createProduct).Methods("POST")
	r.HandleFunc("/api/v1/products/{id}", h.updateProduct).Methods("PUT")
	r.HandleFunc("/api/v1/products/{id}", h.deleteProduct).Methods("DELETE")
	return h
}

func (h *ProductHandler) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.dao.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) getProductById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	product, err := h.dao.GetProductById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) createProduct(w http.ResponseWriter, r *http.Request) {
	product := new(Product)
	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.dao.CreateProduct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ProductHandler) updateProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	product := new(Product)
	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	product.ID = id
	err = h.dao.UpdateProduct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) deleteProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := h.dao.DeleteProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

