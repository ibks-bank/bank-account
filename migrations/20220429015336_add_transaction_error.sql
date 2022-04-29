-- +goose Up
-- +goose StatementBegin
alter table if exists transactions
    add column if not exists "error" text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
