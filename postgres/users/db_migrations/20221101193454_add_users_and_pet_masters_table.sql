-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id                 BIGSERIAL PRIMARY KEY,
    username           text NOT NULL UNIQUE,
    description        text,
    encrypted_password text NOT NULL,
    created_at         TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    updated_at         TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);
CREATE INDEX users_id ON users (id);
CREATE INDEX users_username ON users (username);

CREATE TABLE pet_masters (
    id             BIGSERIAL PRIMARY KEY,
    name           text NOT NULL,
    contact_number text,
    user_id        BIGSERIAL REFERENCES users(id) ON DELETE CASCADE,
    created_at     TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    updated_at     TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);
CREATE INDEX pet_masters_id ON pet_masters (id);
CREATE INDEX pet_masters_user_id ON pet_masters (user_id);

CREATE TABLE locations (
    id                   BIGSERIAL PRIMARY KEY,
    country              text NOT NULL,
    state_or_province    text,
    city_or_municipality text,
    neighborhood         text,
    zip_code             text,
    pet_master_id        BIGSERIAL REFERENCES pet_masters(id) ON DELETE CASCADE,
    created_at           TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    updated_at           TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);
CREATE INDEX locations_id ON locations (id);
CREATE INDEX locations_pet_master_id ON locations (pet_master_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE locations;
DROP TABLE pet_masters;
DROP TABLE users;
-- +goose StatementEnd
