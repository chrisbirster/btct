// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package app

import (
	"context"
)

const createTask = `-- name: CreateTask :one
INSERT INTO tasks (
  description, complete
) VALUES (
  ?, ?
)
RETURNING id, description, complete
`

type CreateTaskParams struct {
	Description string
	Complete    bool
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask, arg.Description, arg.Complete)
	var i Task
	err := row.Scan(&i.ID, &i.Description, &i.Complete)
	return i, err
}

const deleteTask = `-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = ?
`

func (q *Queries) DeleteTask(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTask, id)
	return err
}

const getTask = `-- name: GetTask :one
SELECT id, description, complete FROM tasks 
WHERE id = ? LIMIT 1
`

func (q *Queries) GetTask(ctx context.Context, id int64) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTask, id)
	var i Task
	err := row.Scan(&i.ID, &i.Description, &i.Complete)
	return i, err
}

const listTasks = `-- name: ListTasks :many
SELECT id, description, complete FROM tasks
`

func (q *Queries) ListTasks(ctx context.Context) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, listTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(&i.ID, &i.Description, &i.Complete); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTasksIncomplete = `-- name: ListTasksIncomplete :many
SELECT id, description, complete FROM tasks WHERE complete = 0
`

func (q *Queries) ListTasksIncomplete(ctx context.Context) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, listTasksIncomplete)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(&i.ID, &i.Description, &i.Complete); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTask = `-- name: UpdateTask :exec
UPDATE tasks 
set description = ?,
complete = ?
WHERE id = ?
RETURNING id, description, complete
`

type UpdateTaskParams struct {
	Description string
	Complete    bool
	ID          int64
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) error {
	_, err := q.db.ExecContext(ctx, updateTask, arg.Description, arg.Complete, arg.ID)
	return err
}

const updateTaskComplete = `-- name: UpdateTaskComplete :exec
UPDATE tasks 
set complete = ?
WHERE id = ?
RETURNING id, description, complete
`

type UpdateTaskCompleteParams struct {
	Complete bool
	ID       int64
}

func (q *Queries) UpdateTaskComplete(ctx context.Context, arg UpdateTaskCompleteParams) error {
	_, err := q.db.ExecContext(ctx, updateTaskComplete, arg.Complete, arg.ID)
	return err
}
