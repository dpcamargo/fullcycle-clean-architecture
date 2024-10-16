package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"

	graphHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	config "github.com/dpcamargo/fullcycle-clean-architecture/configs"
	"github.com/dpcamargo/fullcycle-clean-architecture/internal/event/handler"
	"github.com/dpcamargo/fullcycle-clean-architecture/internal/infra/graph"
	"github.com/dpcamargo/fullcycle-clean-architecture/internal/infra/grpc/pb"
	"github.com/dpcamargo/fullcycle-clean-architecture/internal/infra/grpc/service"
	"github.com/dpcamargo/fullcycle-clean-architecture/internal/infra/web/webserver"
	"github.com/dpcamargo/fullcycle-clean-architecture/pkg/events"
	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	configs, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName)
	db, err := sql.Open(configs.DBDriver, connString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createOrderUsecase := NewCreateOrderUseCase(db, eventDispatcher)

	webServer := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	webServer.AddHandler("/order", webOrderHandler.Create, webserver.POST)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webServer.Start()

	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*createOrderUsecase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)
	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	srv := graphHandler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	log.Fatal(http.ListenAndServe(":"+configs.GraphQLServerPort, nil))
}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
