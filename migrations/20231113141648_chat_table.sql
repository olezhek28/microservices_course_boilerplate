-- +goose Up
create table chat (
    id serial primary key,
    user_id bigint,
    from_user text,
    message text
);

-- +goose Down
drop table chat;

