-- name: GetTodos :many
select * from todo;

-- name: GetTodo :one
select * from todo where id = $1 LIMIT 1;

-- name: UpdateTodo :one
update todo set title = $1, completed = $2 where id = $3 RETURNING *;

-- name: AddTodo :one
insert into todo (title) values ($1) RETURNING *;

-- name: DeleteTodo :exec
delete from todo where id = $1;
