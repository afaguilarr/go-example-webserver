-- +goose Up
-- +goose StatementBegin
CREATE TABLE encryption_data (
    id            SERIAL PRIMARY KEY,
    context       text NOT NULL UNIQUE,
    crypto_secret BYTEA NOT NULL,
    iv            BYTEA NOT NULL,
    salts         BYTEA NOT NULL,
    created_at    TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    updated_at    TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);
CREATE INDEX encryption_data_id ON encryption_data (id);
CREATE INDEX encryption_data_context ON encryption_data (context);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE encryption_data;
-- +goose StatementEnd
