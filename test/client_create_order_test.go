package test

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"grpc/src/genproto/orders"
	"grpc/src/ucase/order"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// mockOrderServiceClient is a mock implementation of an OrderService client used for testing.
// createOrderResponse is the mock response returned by CreateOrder calls.
// createOrderError is the mock error returned by CreateOrder calls.
type mockOrderServiceClient struct {
	createOrderResponse *orders.CreateOrderResponse
	createOrderError    error
}

// CreateOrder sends a request to create a new order and returns a response or an error.
func (m *mockOrderServiceClient) CreateOrder(ctx context.Context, in *orders.CreateOrderRequest, opts ...grpc.CallOption) (*orders.CreateOrderResponse, error) {
	return m.createOrderResponse, m.createOrderError
}

// GetOrders retrieves a list of orders based on the given request parameters and returns a response or an error.
func (m *mockOrderServiceClient) GetOrders(ctx context.Context, in *orders.GetOrdersRequest, opts ...grpc.CallOption) (*orders.GetOrderResponse, error) {
	return nil, errors.New("not implemented")
}

// TestNewCreateOrderHandler validates the behavior of the create order handler with various test cases and mock responses.
func TestNewCreateOrderHandler(t *testing.T) {
	tests := []struct {
		name               string
		mockResponse       *orders.CreateOrderResponse
		mockError          error
		expectedStatusCode int
		requestBody        string
	}{
		{
			name: "success",
			mockResponse: &orders.CreateOrderResponse{
				Status: "success",
				Order: &orders.Order{
					OrderID:    1,
					CustomerID: 3,
					ProductID:  4,
					Quantity:   6,
				},
			},
			requestBody:        `{"CustomerID":42,"ProductID": 32,"Quantity": 3}`,
			mockError:          nil,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "validation error",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "grpc error",
			mockResponse:       nil,
			mockError:          errors.New("grpc error"),
			requestBody:        `{"CustomerID":42,"ProductID": 32,"Quantity": 3}`,
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:               "invalid grpc response",
			mockResponse:       nil,
			mockError:          nil,
			requestBody:        `{"CustomerID":42,"ProductID": 32,"Quantity": 3}`,
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &mockOrderServiceClient{
				createOrderResponse: tt.mockResponse,
				createOrderError:    tt.mockError,
			}

			handler := order.NewCreateOrderHandler(mockClient)

			req := httptest.NewRequest(handler.Method(), "/orders", strings.NewReader(tt.requestBody))
			rec := httptest.NewRecorder()

			handler.Handle(rec, req)

			assert.Equal(t, tt.expectedStatusCode, rec.Code)
		})
	}
}
