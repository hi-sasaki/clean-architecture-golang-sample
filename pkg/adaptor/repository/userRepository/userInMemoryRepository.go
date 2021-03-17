package userRepository

import (
	"github.com/pkg/errors"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/domain/model/userModel"
	log "github.com/sirupsen/logrus"
	"time"
)

type UserInMemoryRepository struct {
}

func (rcv *UserInMemoryRepository) GetUser() (*userModel.UserEntity, error) {
	log.Debug("UserInMemoryRepository GetUser start")

	jst, _ := time.LoadLocation("Asia/Tokyo")
	myBirthDay := time.Date(1999, 1, 1, 0, 0, 0, 0, jst)
	user, err := userModel.New("naoto", "oiso", myBirthDay)
	if err != nil {
		return nil, errors.Wrap(err, "Wrap in GetUser UserInMemoryRepository: ")
	}


	log.WithFields(log.Fields{
		"user": user,
	}).Debug("UserInMemoryRepository GetUser end")

	return user, nil

}

// コンストラクタ関数
func New() (*UserInMemoryRepository, error) {
	repository := UserInMemoryRepository{}
	return &repository, nil
}
