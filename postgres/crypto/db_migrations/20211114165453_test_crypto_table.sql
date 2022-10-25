-- +goose Up
-- +goose StatementBegin
CREATE TABLE test_crypto_table (
    id SERIAL PRIMARY KEY,
    encrypted_pw BYTEA NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE test_crypto_table;
-- +goose StatementEnd
