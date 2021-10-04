package admin

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/registry"
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
	log.Println("id:" + id)
	res, err := u.provider.UserUsecase().GetByID(c, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}
