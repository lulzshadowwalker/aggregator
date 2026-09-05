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

-- name: GetFeeds :many
select
    feeds.id,
    feeds.name,
    feeds.url,
    feeds.user_id,
    feeds.created_at,
    feeds.updated_at,
    users.name as user_name
from feeds
join users on feeds.user_id = users.id
order by feeds.created_at asc;

