-- +goose Up
-- +goose StatementBegin
create type currency as enum ('rub', 'eur', 'usd');

create table if not exists accounts
(
    "id"         bigserial primary key,
    "created_at" timestamp not null check ( "created_at" > '1970-01-01' ) default now(),
    "currency"   currency  not null,
    "balance"    bigint    not null check ( "balance" != 0 ),
    "limit"      bigint    not null check ( "limit" != 0 ),
    "user_id"    bigint    not null check ( "user_id" != 0 ),
    "name"       text      not null check ( "name" != '' )
);

create type transaction_type as enum ('transfer', 'payment');

create table if not exists transactions
(
    "id"              bigserial primary key,
    "created_at"      timestamp        not null check ( "created_at" > '1970-01-01' ) default now(),
    "account_to"      bigint           not null,
    "account_to_name" text             not null                                       default '',
    "account_from"    bigint           not null check ( "account_from" != 0 ),
    "amount"          bigint           not null check ( "amount" != 0 ),
    "type"            transaction_type not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists accounts;
drop table if exists transactions;
-- +goose StatementEnd
