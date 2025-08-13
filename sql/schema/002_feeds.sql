-- +goose Up
CREATE TABLE feeds (
    id UUID primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    name text not null,
    url text not null unique,
    user_id UUID not null references users on delete cascade,
    foreign key(user_id) references users(id)
);



-- +goose Down
DROP TABLE feeds;