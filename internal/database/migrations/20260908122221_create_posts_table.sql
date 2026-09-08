-- +goose Up
create table if not exists posts (
	id uuid primary key,
	feed_id uuid not null,
	url text unique not null,
	title text not null,
	description text,
	published_at text,

	created_at timestamp not null,
	updated_at timestamp not null,

	foreign key (feed_id) references feeds(id) on delete cascade
);

-- +goose Down
drop table if exists posts;

