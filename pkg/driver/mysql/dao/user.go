package dao

import (
	"context"
	"database/sql"

	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/entity/model"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/entity/repository"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type User struct {
	db *sql.DB
}

func UserProvider(db *sql.DB) repository.User {
	return &User{db: db}
}

func (u *User) GetByID(ctx context.Context, id string) (*model.User, error) {
	return nil, nil
}

func (u *User) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	boil.SetDB(u.db)
	return nil, nil
}
