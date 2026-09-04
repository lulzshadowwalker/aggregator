-- +goose Up
create table if not exists users (
	id uuid primary key default gen_random_uuid(),
	name text unique not null,
	created_at timestamp not null default now(),
	updated_at timestamp not null default now() 
);

-- +goose Down
drop table if exists users;
