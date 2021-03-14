package userPresentator

import (
	"encoding/json"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/rikodao/clean-architecture-golang-sample/src/lib/domain/model/userModel"
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