-- name: GetTask :one
SELECT * FROM tasks 
WHERE id = ? LIMIT 1;

-- name: ListTasks :many
SELECT * FROM tasks;

-- name: ListTasksIncomplete :many
SELECT * FROM tasks WHERE complete = 0;

-- name: CreateTask :one
INSERT INTO tasks (
  description, complete
) VALUES (
  ?, ?
)
RETURNING *;

-- name: UpdateTask :exec
UPDATE tasks 
set description = ?,
complete = ?
WHERE id = ?
RETURNING *;

-- name: UpdateTaskComplete :exec
UPDATE tasks 
set complete = ?
WHERE id = ?
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = ?;


-- name: GetMigration :one
SELECT * FROM migrations 
WHERE id = ? LIMIT 1;

-- name: ListMigrations :many
SELECT * FROM migrations;

-- name: GetLatestMigration :one
SELECT count FROM migrations;

