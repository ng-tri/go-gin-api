-- +goose Up
-- +goose StatementBegin
CREATE TYPE order_status AS ENUM ('pending', 'paid', 'shipping', 'completed', 'cancelled');
CREATE TYPE payment_method AS ENUM('credit_card', 'paypal', 'bank_transfer', 'momo', 'cod');
CREATE TYPE payment_status AS ENUM ('pending', 'paid', 'failed', 'refunded');

CREATE TABLE orders (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    status order_status NOT NULL DEFAULT 'pending',
    payment_method payment_method NOT NULL DEFAULT 'credit_card',
    payment_status payment_status NOT NULL DEFAULT 'pending',
    total_amount FLOAT NOT NULL,
    shipping_address TEXT,
    shipping_fee FLOAT DEFAULT 0,
    tracking_number VARCHAR(50),
    note TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
DROP TYPE order_status;
DROP TYPE payment_method;
DROP TYPE payment_status;
-- +goose StatementEnd
