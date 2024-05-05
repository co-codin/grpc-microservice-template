package orders

import (
	"grpc-microservice/services/common/genproto/orders"
	"grpc-microservice/services/orders/types"
	"net/http"
)

type OrdersHttpHandler struct {
	ordersService types.OrderService
}

func NewHttpOrdersService(ordersService types.OrderService) *OrdersHttpHandler {
	return &OrdersHttpHandler{
		ordersService: ordersService,
	}
}

func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := util.ParseJ
}