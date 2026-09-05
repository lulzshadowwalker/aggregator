-- name: GetUsers :many
select * from users;

-- name: GetUser :one
select * from users
where name = $1;

-- name: CreateUser :one
insert into users (id, name, created_at, updated_at)
values ($1, $2, $3, $4)
returning *;

-- name: DeleteUsers :exec
delete from users;

-- name: CreateFeed :one
insert into feeds (id, user_id, name, url, created_at, updated_at)
values ($1, $2, $3, $4, $5, $6)
returning *;

-- name: CreateFeedFollow :one
with inserted as (
	insert into feed_follows (id, user_id, feed_id, created_at, updated_at)
	values ($1, $2, $3, $4, $5)
	returning *
)
select 
	inserted.*,
	users.name as user_name,
	feeds.name as feed_name
from 
	inserted
	join users on users.id = inserted.user_id
	join feeds on feeds.id = inserted.feed_id;


-- name: GetFeeds :many
select
	feeds.*,
	users.name as user_name
from feeds
join users on feeds.user_id = users.id
order by feeds.created_at asc;

-- name: GetFeedsByUserID :many
select 
	users.*,
	feeds.name as feed_name,
	feeds.url as feed_url
from users
join feeds on feeds.user_id = users.id
where users.id = $1;


-- name: GetFeedByURL :one
select * from feeds where url = $1;
