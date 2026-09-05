-- +goose Up
create table if not exists feeds (
	id uuid primary key,
	user_id uuid not null,
	name text not null,
	url text unique not null,
	created_at timestamp not null,
	updated_at timestamp not null,

	foreign key (user_id) references users(id) on delete cascade
);

-- +goose Down
drop table if exists feeds;
