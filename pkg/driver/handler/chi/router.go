package chi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/adapter/controller/admin"
)

func NewRouter(u *admin.User) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/user/{userID}", u.GetByID)
	return r
}
