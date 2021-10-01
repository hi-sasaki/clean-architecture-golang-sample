package service

import "github.com/hi-sasaki/clean-architecture-golang-sample/pkg/entity/repository"

type User struct {
	User repository.User
}

func (u *User) GetByID() {
	//u.User.GetByID()
}
