-- +goose Up
CREATE TABLE orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    price DECIMAL(10, 2) NOT NULL,
    tax DECIMAL(10, 2) NOT NULL,
    final_price DECIMAL(10, 2)
);

-- +goose Down
DROP TABLE orders;