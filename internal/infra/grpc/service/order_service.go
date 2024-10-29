package service

import (
	"context"

	"github.com/dpcamargo/fullcycle-clean-architecture/internal/infra/grpc/pb"
	"github.com/dpcamargo/fullcycle-clean-architecture/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	OrderUsecase usecase.OrderUsecase
}

func NewOrderService(orderUsecase usecase.OrderUsecase) *OrderService {
	return &OrderService{
		OrderUsecase: orderUsecase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    int(in.GetId()),
		Price: in.GetPrice(),
		Tax:   in.GetTax(),
	}
	output, err := s.OrderUsecase.CreateOrder(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         int32(output.ID),
		Price:      output.Price,
		Tax:        output.Tax,
		FinalPrice: output.FinalPrice,
	}, nil
}

