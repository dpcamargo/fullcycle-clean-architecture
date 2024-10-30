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

func (s *OrderService) GetOrder(ctx context.Context, in *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	output, err := s.OrderUsecase.GetOrder(int(in.GetId()))
	if err != nil {
		return nil, err
	}
	return &pb.GetOrderResponse{
		Id:         int32(output.ID),
		Price:      output.Price,
		Tax:        output.Tax,
		FinalPrice: output.FinalPrice,
	}, nil
}

func (s *OrderService) GetList(ctx context.Context, in *pb.Empty) (*pb.GetOrderListResponse, error) {
	output, err := s.OrderUsecase.ListOrders()
	if err != nil {
		return nil, err
	}
	var res *pb.GetOrderListResponse
	var orders []*pb.GetOrderResponse
	for _, order := range output {
		temp := pb.GetOrderResponse{
			Id:         int32(order.ID),
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		orders = append(orders, &temp)
	}
	res = &pb.GetOrderListResponse{
		Orders: orders,
	}
	return res, nil
}
