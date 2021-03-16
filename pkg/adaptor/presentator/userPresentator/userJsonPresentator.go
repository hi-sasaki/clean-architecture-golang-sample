package userPresentator

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/domain/model/userModel"
	log "github.com/sirupsen/logrus"
)

type UserOutputData string
type UserJsonPresentator struct {
	
}

func (receiver *UserJsonPresentator) Serialize(user *userModel.UserData) (UserOutputData, error) {
	log.Debug("UserJsonPresentator Serialize start")

	userJson, err := json.Marshal(user)
	if err != nil {
		return "", errors.Wrap(err, "Wrap in serialize UserJsonPresentator: ")
	}

	log.WithFields(log.Fields{
		"userJson": userJson,
	}).Debug("UserJsonPresentator Serialize end")

	return UserOutputData(userJson), nil
}

func New() (*UserJsonPresentator, error) {
	presentator := UserJsonPresentator{}
	return &presentator, nil
}