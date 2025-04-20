package order

import (
	"errors"
	httppkg "grpc/pkg/http"
	"grpc/src/genproto/orders"
	"log"
	"net/http"
)

// NewCreateOrderHandler is an interface for handling order creation requests.
func NewCreateOrderHandler(grpcOrder orders.OrderServiceClient) httppkg.Http {
	return &create{
		grpcOrder: grpcOrder,
	}
}

type create struct {
	grpcOrder orders.OrderServiceClient
}

func (uc *create) Method() string { return http.MethodPost }

func (uc *create) Handle(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := httppkg.ParseJSON(r, &req)
	if err != nil {
		httppkg.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data, err := uc.grpcOrder.CreateOrder(r.Context(), &orders.CreateOrderRequest{
		CustomerID: req.CustomerID,
		ProductID:  req.ProductID,
		Quantity:   req.Quantity,
	})

	if err != nil {
		log.Printf("client error: %v", err)
		httppkg.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if data == nil {
		log.Println("client error: data is nil")
		httppkg.WriteError(w, http.StatusNotFound, errors.New("data is nil"))
		return
	}

	httppkg.WriteJSON(w, http.StatusOK, data)
	return
}
