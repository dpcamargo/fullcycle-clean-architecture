package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dpcamargo/fullcycle-clean-architecture/internal/entity"
	"github.com/dpcamargo/fullcycle-clean-architecture/internal/usecase"
	"github.com/dpcamargo/fullcycle-clean-architecture/pkg/events"
)

type WebOrderHandler struct {
	EventDispatcher   events.EventDispatcherInterface
	OrderRepository   entity.OrderRepositoryInterface
	OrderCreatedEvent events.EventInterface
}

func NewWebOrderHandler(
	eventDispatcher events.EventDispatcherInterface,
	orderRepository entity.OrderRepositoryInterface,
	orderCreatedEvent events.EventInterface,
) *WebOrderHandler {
	return &WebOrderHandler{
		EventDispatcher:   eventDispatcher,
		OrderRepository:   orderRepository,
		OrderCreatedEvent: orderCreatedEvent,
	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	orderUsecase := usecase.NewOrderUsecase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispatcher)
	output, err := orderUsecase.CreateOrder(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebOrderHandler) Get(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("id")
	orderID, err := strconv.Atoi(param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orderUsecase := usecase.NewOrderUsecase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispatcher)
	order, err := orderUsecase.GetOrder(orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebOrderHandler) List(w http.ResponseWriter, r *http.Request) {

	orderUsecase := usecase.NewOrderUsecase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispatcher)
	orders, err := orderUsecase.ListOrders()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
