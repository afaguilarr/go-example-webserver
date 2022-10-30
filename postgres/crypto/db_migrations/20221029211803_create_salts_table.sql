-- +goose Up
-- +goose StatementBegin
CREATE TABLE salts (
    id         SERIAL PRIMARY KEY,
    context    text NOT NULL,
    salts      BYTEA NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    updated_at TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);
CREATE INDEX salts_id ON salts (id);
CREATE INDEX salts_context ON salts (context);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE salts;
-- +goose StatementEnd
