-- name: CreateUser :one
 insert into users (id, name, email, password_hash, role)
values ($1, $2, $3, $4, $5)
returning *;

-- name: FetchUser :one
select * from users where email = $1;
