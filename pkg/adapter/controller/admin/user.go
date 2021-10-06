package admin

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/registry"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/usecase/inout"
)

type User struct {
	provider registry.Provider
}

func NewUser(p registry.Provider) *User {
	return &User{provider: p}
}

func (u *User) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	c, cancel := context.WithTimeout(ctx, time.Duration(time.Second)*30)
	defer cancel()

	id := chi.URLParam(r, "userID")
	res, err := u.provider.UserUsecase().GetByID(c, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	w.WriteHeader(http.StatusOK)
}

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	c, cancel := context.WithTimeout(ctx, time.Duration(time.Second)*30)
	defer cancel()

	user := new(inout.User)
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		return
	}

	res, err := u.provider.UserUsecase().CreateUser(c, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	w.WriteHeader(http.StatusCreated)
}
