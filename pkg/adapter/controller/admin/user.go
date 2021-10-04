package admin

import (
	"net/http"

	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/registry"
)

type User struct {
	provider registry.Provider
}

func NewUser(p registry.Provider) *User {
	return &User{provider: p}
}

func (u *User) GetByID(w http.ResponseWriter, r *http.Request) {
	u.provider.UserUsecase().GetByID()
}
