package service

import (
	"context"
	"demo/pkg/util"
	"demo/platform/db"
	"demo/postgres"
	"github.com/google/uuid"
)

func GetUserById(id uuid.UUID) (postgres.User, error) {
	return db.Query().GetAUserById(context.Background(), id)
}

func GetUserByEmail(email string) (postgres.User, error) {
	return db.Query().GetAUserByEmail(context.Background(), email)
}

func GetAllUsers() ([]postgres.User, error) {
	return db.Query().GetAllUsers(context.Background())
}

func InsertUser(user postgres.InsertUserParams) (postgres.User, error) {
	var err error
	user.Password, err = util.HashPassword(user.Password)
	if err != nil {
		return postgres.User{}, err
	}
	return db.Query().InsertUser(context.Background(), user)
}

func UpdateUser(user postgres.UpdateUserParams) (postgres.User, error) {
	return db.Query().UpdateUser(context.Background(), user)
}

func DeleteUser(id uuid.UUID) error {
	return db.Query().DeleteUser(context.Background(), id)
}
