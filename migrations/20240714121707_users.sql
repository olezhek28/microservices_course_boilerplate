-- +goose Up
CREATE SCHEMA auth;
CREATE TABLE auth.users
(
    id bigserial primary key,
    email text not null unique,
    password text not null,
    name text,
    role int8 check ( role > 0 ),
    created_at timestamp(0) default CURRENT_TIMESTAMP,
    updated_at timestamp(0)
);

-- +goose Down
DROP TABLE auth.users;
DROP SCHEMA auth;