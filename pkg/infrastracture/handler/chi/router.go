package chi

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/adaptor/controller/userController"
	"net/http"
)

func NewRouter(controller *userController.UserJsonController) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/", controller.GetUser)
	r.Get("/user/{userID}", controller.GetUserById)
	return r

}
func NewHandler(controller *userController.UserJsonController) http.Handler {
	return http.Handler(NewHandler(controller))
}
