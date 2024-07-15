-- +goose Up
CREATE SCHEMA auth;
CREATE TABLE auth.users
(
    id bigserial primary key,
    email varchar(128) not null unique,
    password varchar(128) not null,
    name varchar(128),
    role int8 check ( role > 0 ),
    created_at timestamp(0) default CURRENT_TIMESTAMP,
    updated_at timestamp(0)
);

-- +goose Down
DROP TABLE auth.users;
DROP SCHEMA auth;