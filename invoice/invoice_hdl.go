package invoice

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rickCrz7/acme/utils"
)

type InvoiceHandler struct {
	invoiceDao InvoiceDao
}

func NewInvoiceHandler(invoiceDao InvoiceDao, r *mux.Router) *InvoiceHandler {
	h := &InvoiceHandler{invoiceDao: invoiceDao}
	r.HandleFunc("/api/v1/invoices", h.getInvoices).Methods("GET")
	r.HandleFunc("/api/v1/invoices/{id}", h.getInvoiceById).Methods("GET")
	r.HandleFunc("/api/v1/invoices", h.createInvoice).Methods("POST")
	r.HandleFunc("/api/v1/invoices/{id}", h.updateInvoice).Methods("PUT")
	r.HandleFunc("/api/v1/invoices/paid/{id}", h.MarkAsPaid).Methods("PUT")
	r.HandleFunc("/api/v1/invoices/{id}", h.deleteInvoice).Methods("DELETE")
	return h
}

func (h *InvoiceHandler) getInvoices(w http.ResponseWriter, r *http.Request) {
	invoices, err := h.invoiceDao.GetInvoices()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoices)
}

func (h *InvoiceHandler) getInvoiceById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	invoice, err := h.invoiceDao.GetInvoiceById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoice)
}

func (h *InvoiceHandler) createInvoice(w http.ResponseWriter, r *http.Request) {
	invoice := new(utils.Invoice)
	err := json.NewDecoder(r.Body).Decode(invoice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.invoiceDao.CreateInvoice(invoice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *InvoiceHandler) updateInvoice(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	invoice := new(utils.Invoice)
	err := json.NewDecoder(r.Body).Decode(invoice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	invoice.ID = id
	err = h.invoiceDao.UpdateInvoice(invoice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *InvoiceHandler) MarkAsPaid(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := h.invoiceDao.MarkAsPaid(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *InvoiceHandler) deleteInvoice(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := h.invoiceDao.DeleteInvoice(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}