-- +goose Up
-- +goose StatementBegin
CREATE TABLE hello_world (
    id SERIAL PRIMARY KEY,
    hw_text text NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE hello_world;
-- +goose StatementEnd
