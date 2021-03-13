package userPresentator

import (
	"encoding/json"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"go-study-sample/src/lib/domain/model/userModel"
)

type UserOutputData string
type userJsonPresentator struct {
	
}

func (receiver *userJsonPresentator) Serialize(user *userModel.UserData) (UserOutputData, error) {
	log.Debug("userJsonPresentator Serialize start")

	userJson, err := json.Marshal(user)
	if err != nil {
		return "", errors.Wrap(err, "Wrap in serialize userJsonPresentator: ")
	}

	log.WithFields(log.Fields{
		"userJson": userJson,
	}).Debug("userJsonPresentator Serialize end")

	return UserOutputData(userJson), nil
}

func New() (*userJsonPresentator, error) {
	presentator := userJsonPresentator{}
	return &presentator, nil
}