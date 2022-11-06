-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
  ADD COLUMN encrypted_refresh_token_secret BYTEA;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
  DROP COLUMN encrypted_refresh_token_secret;
-- +goose StatementEnd
