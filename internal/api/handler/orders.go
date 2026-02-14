package handler

import (
	"net/http"
	"time"

	"github.com/francisco3ferraz/go-nuts/internal/orders"
)

type OrdersHandler struct {
	service orders.Service
}

type orderResponse struct {
	ID          string    `json:"id"`
	Customer    string    `json:"customer"`
	AmountCents int64     `json:"amount_cents"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewOrdersHandler(service orders.Service) OrdersHandler {
	return OrdersHandler{service: service}
}

func (h OrdersHandler) List(w http.ResponseWriter, r *http.Request) {
	orderList, err := h.service.List(r.Context())
	if err != nil {
		http.Error(w, "failed to list orders", http.StatusInternalServerError)
		return
	}

	response := make([]orderResponse, 0, len(orderList))
	for _, item := range orderList {
		response = append(response, orderResponse{
			ID:          item.ID,
			Customer:    item.Customer,
			AmountCents: item.AmountCents,
			CreatedAt:   item.CreatedAt,
		})
	}

	respondJSON(w, http.StatusOK, response)
}
