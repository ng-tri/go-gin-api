-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS migration_batches (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    version_id BIGINT NOT NULL,
    batch INT NOT NULL,
    migration VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE migration_batches;
-- +goose StatementEnd
