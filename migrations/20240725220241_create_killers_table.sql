-- +goose Up
-- +goose StatementBegin
CREATE TABLE killers (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    city TEXT NOT NULL,
    kills INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE killers;
-- +goose StatementEnd
