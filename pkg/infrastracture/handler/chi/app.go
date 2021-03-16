package chi

import (
	"github.com/go-chi/chi/v5"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/adaptor/controller/userController"
)

type Handler struct {

}

func (rcv Handler) Create(controller *userController.UserJsonController) *chi.Mux{
	return newRouter(controller)
}

func NewHandler() Handler {
	return Handler{}
}