package reports

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ReportsHandler struct {
	dao ReportsDao
}

func NewReportsHandler(dao ReportsDao, r *mux.Router) *ReportsHandler {
	h := &ReportsHandler{dao: dao}
	r.HandleFunc("/api/v1/reports/most-sold-products", h.getMostSoldProducts).Methods("GET")
	r.HandleFunc("/api/v1/reports/total-sales-by-product", h.getTotalSalesByProduct).Methods("GET")
	r.HandleFunc("/api/v1/reports/total-sales-by-customers", h.getTotalSalesByCustomers).Methods("GET")
	return h
}

func (h *ReportsHandler) getMostSoldProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.dao.MostSoldProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *ReportsHandler) getTotalSalesByProduct(w http.ResponseWriter, r *http.Request) {
	products, err := h.dao.TotalSalesByProduct()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *ReportsHandler) getTotalSalesByCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.dao.TotalSalesByCustomers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}