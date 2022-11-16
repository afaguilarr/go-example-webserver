-- +goose Up
-- +goose StatementBegin

ALTER TABLE users
  DROP COLUMN encrypted_password;

ALTER TABLE users
  ADD COLUMN encrypted_password BYTEA NOT NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE users
  DROP COLUMN encrypted_password;

ALTER TABLE users
  ADD COLUMN encrypted_password text NOT NULL;

-- +goose StatementEnd
