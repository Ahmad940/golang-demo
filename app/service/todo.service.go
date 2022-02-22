package service

import (
	"context"
	"demo/platform/db"
	"demo/postgres"
	"github.com/google/uuid"
)

func InsertTodo(title string) (postgres.Todo, error) {
	return db.Query().AddTodo(context.Background(), title)
}

func GetTodos() ([]postgres.Todo, error)  {
	return db.Query().GetTodos(context.Background())
}

func GetATodo(id uuid.UUID) (postgres.Todo, error) {
	return db.Query().GetTodo(context.Background(), id)
}

func UpdateTodo(id uuid.UUID, todo postgres.UpdateTodoParams) (postgres.Todo, error)  {
	return db.Query().UpdateTodo(context.Background(), todo)
}


func DeleteTodo(id uuid.UUID) error {
	return db.Query().DeleteTodo(context.Background(), id)
}
