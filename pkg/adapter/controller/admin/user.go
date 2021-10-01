package admin

import (
	"net/http"

	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/registry"
)

type User struct {
	provider registry.Provider
}

func (u *User) GetByID(w http.ResponseWriter, r *http.Request) {
	u.provider.UserUsecase().GetByID()
}
