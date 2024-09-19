-- +goose Up
-- +goose StatementBegin
CREATE TABLE jwt_refresh_tokens
(
    token        text      NOT NULL,
    user_id      uuid      NOT NULL,
    ignore_after timestamp NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS jwt_refresh_tokens;
-- +goose StatementEnd
