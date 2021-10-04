package chi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	admin "github.com/hi-sasaki/clean-architecture-golang-sample/pkg/adapter/controller/admin"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/registry"
)

func NewRouter(p registry.Provider) *chi.Mux {
	r := chi.NewRouter()

	u := admin.NewUser(p)
	r.Use(middleware.Logger)
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/user/{userID}", u.GetByID)
	return r
}
