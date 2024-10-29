package service

import (
	"context"

	"github.com/dpcamargo/fullcycle-clean-architecture/internal/infra/grpc/pb"
	"github.com/dpcamargo/fullcycle-clean-architecture/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUsecase usecase.OrderUsecase
}

func NewOrderService(createOrderUsecase usecase.OrderUsecase) *OrderService {
	return &OrderService{
		CreateOrderUsecase: createOrderUsecase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: in.Price,
		Tax:   in.Tax,
	}
	output, err := s.CreateOrderUsecase.CreateOrder(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      output.Price,
		Tax:        output.Tax,
		FinalPrice: output.FinalPrice,
	}, nil
}
