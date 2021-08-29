package userController

import (
	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/application/interface/controller"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/application/usecase"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type UserJsonController struct {
	getUserInteractor     usecase.GetUserInteractor
	getUserByIdInteractor usecase.GetUserByIdInteractor
}

func (rcv UserJsonController) GetUser(w http.ResponseWriter, r *http.Request) {
	log.SetLevel(log.Level(5))
	result, err := rcv.getUserInteractor.Handle()
	if err != nil {
		log.Error(err)
	}

	w.Write([]byte(result))

}

func (rcv UserJsonController) GetUserById(w http.ResponseWriter, r *http.Request) {
	log.SetLevel(log.Level(5))
	id := chi.URLParam(r, "userID")
	if id == "" {
		log.Error("userIDが空です")
	}
	command := controller.GetUserByIdCommand{Id: id}
	result, err := rcv.getUserByIdInteractor.Handle(command)
	if err != nil {
		log.Error(errors.Wrap(err, "Wrap in GetUser UserJsonController: "))
	}

	if err != nil {
		log.Error(err)
	}

	w.Write([]byte(result))

}

func New(getUserInteractor *usecase.GetUserInteractor, getUserByIdInteractor *usecase.GetUserByIdInteractor) *UserJsonController {
	controller := UserJsonController{*getUserInteractor, *getUserByIdInteractor}
	return &controller
}
