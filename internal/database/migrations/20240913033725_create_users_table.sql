-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users
(
    id         uuid PRIMARY KEY             DEFAULT gen_random_uuid(),
    email      varchar(255) UNIQUE NOT NULL,
    password   text                NOT NULL,
    role_id    int                 NOT NULL DEFAULT 1,
    is_active  boolean             NOT NULL DEFAULT FALSE,
    created_at timestamp           NOT NULL DEFAULT NOW(),
    updated_at timestamp                    DEFAULT NULL,
    deleted_at timestamp                    DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;

DROP EXTENSION IF EXISTS pgcrypto;
-- +goose StatementEnd
