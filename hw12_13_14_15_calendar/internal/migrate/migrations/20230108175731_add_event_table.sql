-- +goose Up
-- +goose StatementBegin
create table events
(
    id         UUID primary key,
    owner      text      not null,
    title      text      not null,
    text       text,
    start_time timestamp not null,
    end_time   timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table events;
-- +goose StatementEnd
