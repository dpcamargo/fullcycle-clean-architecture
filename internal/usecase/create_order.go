package usecase

import (
	"github.com/dpcamargo/fullcycle-clean-architecture/internal/entity"
	"github.com/dpcamargo/fullcycle-clean-architecture/pkg/events"
)

type OrderInputDTO struct {
	ID    string  `json:"id"`
	Price float32 `json:"price"`
	Tax   float32 `json:"tax"`
}

type OrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float32 `json:"price"`
	Tax        float32 `json:"tax"`
	FinalPrice float32 `json:"final_price"`
}

type CreateOrderUsecase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderCreated    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewCreateOrderUsecase(
	orderRepository entity.OrderRepositoryInterface,
	orderCreated events.EventInterface,
	eventDispatcher events.EventDispatcherInterface) *CreateOrderUsecase {
	return &CreateOrderUsecase{
		OrderRepository: orderRepository,
		OrderCreated:    orderCreated,
		EventDispatcher: eventDispatcher,
	}
}

func (c *CreateOrderUsecase) Execute(input OrderInputDTO) (OrderOutputDTO, error) {
	order := entity.Order{
		ID:    input.ID,
		Price: input.Price,
		Tax:   input.Tax,
	}
	err := order.CalculateFinalPrice()
	if err != nil {
		return OrderOutputDTO{}, err
	}

	err = c.OrderRepository.Save(&order)
	if err != nil {
		return OrderOutputDTO{}, err
	}

	dto := OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}

	c.OrderCreated.SetPayload(dto)
	c.EventDispatcher.Dispatch(c.OrderCreated)
	return dto, nil
}
