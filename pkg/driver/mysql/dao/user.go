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

	var lastName *string
	var firstName *string
	var password *string
	err := builder.Select(
		"id",
		"first_name",
		"last_name",
		"password",
	).From(dto.TableNames.Users).QueryRowContext(ctx).Scan(
		&user.ID,
		&firstName,
		&lastName,
		&password,
	)
	if firstName != nil {
		user.FirstName = *firstName
	}
	if lastName != nil {
		user.LastName = *lastName
	}

	if password != nil {
		//サンプルコードのためパスワードも返すようにしているが、本来望ましくない
		user.Password = *password
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	dt := &dto.User{}
	dt.ID = user.ID
	firstname := null.String{String: user.FirstName, Valid: true}
	dt.FirstName = firstname
	lastName := null.String{String: user.LastName, Valid: true}
	dt.LastName = lastName

	password := null.String{String: user.Password, Valid: true}
	dt.Password = password

	if err := dt.Insert(ctx, u.db, boil.Infer()); err != nil {
		return nil, err
	}
	res, err := u.GetByID(ctx, dt.ID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
