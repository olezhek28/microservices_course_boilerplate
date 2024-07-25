-- +goose Up
-- +goose StatementBegin
CREATE TABLE auth.user_actions
(
    user_id bigint references auth.users(id),
    name text not null,
    old_value text not null,
    new_value text,
    created_at timestamp(0) default CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE auth.user_actions;
-- +goose StatementEnd
