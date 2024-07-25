-- +goose Up
-- +goose StatementBegin
ALTER TABLE auth.users DROP constraint users_role_check;
ALTER TABLE auth.users ADD constraint users_role_check check (role >= 0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
