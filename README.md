# Fullcycle Clean Architecture Challenge
Desafio da pós graduação Go Expert da Full Cycle para produzir um serviço com API REST, graphQL, gRPC e RabbitMQ para o cadastro e coleta de orders.




GraphQL queries:
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
```