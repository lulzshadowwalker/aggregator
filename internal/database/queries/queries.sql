-- name: ListUsers :many
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
