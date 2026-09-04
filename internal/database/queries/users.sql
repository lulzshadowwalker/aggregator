-- name: GetUser :one
select * from users
where name = $1;

-- name: CreateUser :one
insert into users (id, name, created_at, updated_at)
values ($1, $2, $3, $4)
returning *;

-- name: TruncateUsers :exec
truncate table users;
