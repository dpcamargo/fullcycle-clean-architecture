
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