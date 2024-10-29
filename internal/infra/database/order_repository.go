package database

import (
	"database/sql"

	"github.com/dpcamargo/fullcycle-clean-architecture/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *OrderRepository) GetOrder(orderId int) (*entity.Order, error) {
	var order entity.Order

	stmt, err := r.Db.Prepare("SELECT id, price, tax, final_price FROM orders WHERE id = ?")
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(orderId)
	err = row.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
