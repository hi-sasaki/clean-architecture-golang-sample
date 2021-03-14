package chi

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/rikodao/clean-architecture-golang-sample/src/lib/adaptor/controller/userController"
	"net/http"
)

func  newRouter(controller *userController.UserJsonController) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.SetLevel(log.Level(5))
		result, err := controller.GetUser()
		if err != nil {
			log.Fatal(err)
		}

		w.Write([]byte(result))

	})
	return r

}
