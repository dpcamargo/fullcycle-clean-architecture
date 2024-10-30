# Fullcycle Clean Architecture Challenge

Full Cycle's post graduate course Go Expert challenge to develop a service with API RAPI REST, graphQL, gRPC and RabbitMQ for order registration and listing.

## Service Ports

web server on port :8001

gRPC server on port 50051

GraphQL server on port 8080

## Starting complimentary services

Run `docker compose up -d` to create the MySQL and RabbitMQ containers and do the necessary migration.

## API Rest

Test using order_create.http and order_get.http (vscode plugin needed)

## GraphQL

Open `http://localhost:8080/query` and use the following queries:

```
mutation createOrder{
  createOrder(
  	input: {id: 44, Price: 10, Tax: 100})
  	{id, Price, Tax, FinalPrice }
}

query getOrder {
  getOrder(id: 4) {
    id
    Price
    Tax
    FinalPrice
  }
}

query getList {
  getList {
    id
    Price
    Tax
    FinalPrice
  }
}
```

## gRPC

Use `evans -r repl` to connect to gRPC server.
