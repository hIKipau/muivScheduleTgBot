-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id integer primary key,
    course text,
    groups text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
