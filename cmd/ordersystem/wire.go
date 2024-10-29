//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/dpcamargo/fullcycle-clean-architecture/internal/entity"
	"github.com/dpcamargo/fullcycle-clean-architecture/internal/event"
	"github.com/dpcamargo/fullcycle-clean-architecture/internal/infra/database"
	"github.com/dpcamargo/fullcycle-clean-architecture/internal/infra/web"
	"github.com/dpcamargo/fullcycle-clean-architecture/internal/usecase"
	"github.com/dpcamargo/fullcycle-clean-architecture/pkg/events"
	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.OrderUsecase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewOrderUsecase,
	)
	return &usecase.OrderUsecase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
