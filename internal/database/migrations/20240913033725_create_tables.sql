-- +goose Up
-- +goose StatementBegin
SET TIME ZONE 'Asia/Jakarta';

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users
(
    id         uuid PRIMARY KEY             DEFAULT gen_random_uuid(),
    email      VARCHAR(255) UNIQUE NOT NULL,
    password   TEXT                NOT NULL,
    role_id    INT                 NOT NULL DEFAULT 0,
    created_at TIMESTAMP           NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE jwt_refresh_tokens
(
    token        TEXT      NOT NULL,
    user_id      uuid      NOT NULL,
    ignore_after TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS jwt_refresh_tokens;

DROP TABLE IF EXISTS users;

DROP EXTENSION IF EXISTS pgcrypto;
-- +goose StatementEnd
