package dao

import (
	"context"
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/driver/mysql/dto"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/entity/model"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/entity/repository"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type User struct {
	db *sql.DB
}

func UserProvider(db *sql.DB) repository.User {
	return &User{db: db}
}

func (u *User) GetByID(ctx context.Context, id string) (*model.User, error) {
	user := &model.User{}
	log.Println("id:", id)
	builder := sq.StatementBuilder.Where(
		sq.Eq{"id": id},
	).RunWith(u.db)

	err := builder.Select(
		"id",
		"first_name",
		"last_name",
	).From(dto.TableNames.Users).QueryRowContext(ctx).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	dt := &dto.User{}
	dt.ID = user.ID
	firstname := null.String{String: user.FirstName}
	dt.FirstName = firstname
	lastName := null.String{String: user.LastName}
	dt.LastName = lastName

	password := null.String{String: user.Password}
	dt.LastName = password

	if err := dt.Insert(ctx, u.db, boil.Infer()); err != nil {
		return nil, err
	}
	res, err := u.GetByID(ctx, dt.ID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
