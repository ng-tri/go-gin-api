-- +goose Up
-- +goose StatementBegin

CREATE TABLE orders (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,

    order_code VARCHAR(50) NOT NULL UNIQUE,

    status ENUM('pending', 'paid', 'shipping', 'completed', 'cancelled')
        NOT NULL DEFAULT 'pending',

    payment_method ENUM('credit_card', 'paypal', 'bank_transfer', 'momo', 'cod')
        NOT NULL DEFAULT 'credit_card',

    payment_status ENUM('pending', 'paid', 'failed', 'refunded')
        NOT NULL DEFAULT 'pending',

    total_amount DECIMAL(10,2) NOT NULL,
    shipping_address TEXT,
    shipping_fee DECIMAL(10,2) DEFAULT 0,
    tracking_number VARCHAR(50),
    note TEXT,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd
