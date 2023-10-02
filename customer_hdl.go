package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	dao CustomerDao
}

func NewCustomerHandler(dao CustomerDao, r *mux.Router) *CustomerHandler {
	h := &CustomerHandler{dao: dao}
	r.HandleFunc("/api/v1/customers", h.getCustomers).Methods("GET")
	r.HandleFunc("/api/v1/customers/{id}", h.getCustomerById).Methods("GET")
	r.HandleFunc("/api/v1/customers", h.createCustomer).Methods("POST")
	r.HandleFunc("/api/v1/customers/{id}", h.updateCustomer).Methods("PUT")
	r.HandleFunc("/api/v1/customers/{id}", h.deleteCustomer).Methods("DELETE")
	return h
}

func (h *CustomerHandler) getCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.dao.GetCustomers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h *CustomerHandler) getCustomerById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	customer, err := h.dao.GetCustomerById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func (h *CustomerHandler) createCustomer(w http.ResponseWriter, r *http.Request) {
	customer := new(Customer)
	err := json.NewDecoder(r.Body).Decode(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.dao.CreateCustomer(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *CustomerHandler) updateCustomer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	customer := new(Customer)
	err := json.NewDecoder(r.Body).Decode(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	customer.ID = id
	err = h.dao.UpdateCustomer(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *CustomerHandler) deleteCustomer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := h.dao.DeleteCustomer(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
