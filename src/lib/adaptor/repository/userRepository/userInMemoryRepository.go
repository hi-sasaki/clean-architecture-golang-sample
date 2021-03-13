package userRepository

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"go-study-sample/src/lib/domain/model/userModel"
	"time"
)

type UserInMemoryRepository struct {

}

func (rcv *UserInMemoryRepository) GetUser() (*userModel.UserData, error) {
	log.Debug("UserInMemoryRepository GetUser start")

	jst, _ := time.LoadLocation("Asia/Tokyo")
	myBirthDay := time.Date(1999,1,1,0,0,0,0, jst)
	userInstance, err := userModel.New("naoto","oiso",myBirthDay)
	if err != nil {
		return nil, errors.Wrap(err, "Wrap in GetUser UserInMemoryRepository: ")
	}
	userData := userInstance.Value()

	log.WithFields(log.Fields{
		"userData": userData,
	}).Debug("UserInMemoryRepository GetUser end")

	return userData,nil

}

// コンストラクタ関数
func New() (*UserInMemoryRepository, error) {
	repository := UserInMemoryRepository{}
	return &repository, nil
}