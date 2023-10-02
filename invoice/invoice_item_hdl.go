package invoice

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rickCrz7/acme/utils"
)

type InvoiceItemHandler struct {
	dao InvoiceItemDao
}

func NewInvoiceItemHandler(dao InvoiceItemDao, r *mux.Router) *InvoiceItemHandler {
	h := &InvoiceItemHandler{dao: dao}
	r.HandleFunc("/api/v1/invoice-items", h.getInvoiceItems).Methods("GET")
	r.HandleFunc("/api/v1/invoice-items/{id}", h.getInvoiceItemById).Methods("GET")
	r.HandleFunc("/api/v1/invoice-items", h.createInvoiceItem).Methods("POST")
	r.HandleFunc("/api/v1/invoice-items/{id}", h.updateInvoiceItem).Methods("PUT")
	r.HandleFunc("/api/v1/invoice-items/{id}", h.deleteInvoiceItem).Methods("DELETE")
	return h
}

func (h *InvoiceItemHandler) getInvoiceItems(w http.ResponseWriter, r *http.Request) {
	invoiceItems, err := h.dao.GetInvoiceItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoiceItems)
}

func (h *InvoiceItemHandler) getInvoiceItemById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	invoiceItem, err := h.dao.GetInvoiceItemById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoiceItem)
}

func (h *InvoiceItemHandler) createInvoiceItem(w http.ResponseWriter, r *http.Request) {
	invoiceItem := new(utils.InvoiceItem)
	err := json.NewDecoder(r.Body).Decode(invoiceItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.dao.CreateInvoiceItem(invoiceItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *InvoiceItemHandler) updateInvoiceItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	invoiceItem := new(utils.InvoiceItem)
	err := json.NewDecoder(r.Body).Decode(invoiceItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	invoiceItem.ID = id
	err = h.dao.UpdateInvoiceItem(invoiceItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *InvoiceItemHandler) deleteInvoiceItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := h.dao.DeleteInvoiceItem(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}