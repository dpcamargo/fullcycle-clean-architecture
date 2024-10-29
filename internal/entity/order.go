package entity

import "errors"

type Order struct {
	ID         int
	Price      float64
	Tax        float64
	FinalPrice float64
}

const (
	InvalidID    = "invalid id"
	InvalidPrice = "invalid price"
	InvalidTax   = "invalid tax"
)

func NewOrder(id int, price, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	err := order.IsValid()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) IsValid() error {
	if o.ID <= 0 {
		return errors.New(InvalidID)
	}
	if o.Price <= 0 {
		return errors.New(InvalidPrice)
	}
	if o.Tax <= 0 {
		return errors.New(InvalidTax)
	}
	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price * (1 + (o.Tax / 100))
	err := o.IsValid()
	if err != nil {
		return err
	}
	return nil
}
