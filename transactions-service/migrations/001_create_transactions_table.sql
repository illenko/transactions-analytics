-- +goose Up
CREATE TABLE IF NOT EXISTS transactions
(
    id       uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    datetime timestamp      not null,
    amount   decimal(20, 2) not null,
    category varchar(256)   not null,
    merchant varchar(256)   not null
);