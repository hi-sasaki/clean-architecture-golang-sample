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
	myBirthDay := time.Date(1991, 11, 4, 0, 0, 0, 0, jst)
	builder, err := userModel.NewUserBuilder("naoto", "oiso", myBirthDay)
	if err != nil {
		return nil, errors.Wrap(err, "Wrap in GetUser UserInMemoryRepository: ")
	}
	user := builder.Build()

	log.WithFields(log.Fields{
		"user": user,
	}).Debug("UserInMemoryRepository GetUser end")

	return user, nil

}

func (rcv *UserInMemoryRepository) GetUGetUserById(id string) (*userModel.UserEntity, error) {
	log.Debug("UserInMemoryRepository GetUser start")
	jst, _ := time.LoadLocation("Asia/Tokyo")
	myBirthDay1 := time.Date(1990, 11, 4, 0, 0, 0, 0, jst)
	builder1, err := userModel.NewUserBuilder("naoto", "oiso", myBirthDay1)
	if err != nil {
		return nil, errors.Wrap(err, "Wrap in GetUser UserInMemoryRepository: ")
	}
	user1 := builder1.AddId("0c51fe73-cbaf-432c-aa1a-1731f4b982c1").Build()

	myBirthDay2 := time.Date(1999, 1, 1, 0, 0, 0, 0, jst)
	builder2, err := userModel.NewUserBuilder("nao", "oiso", myBirthDay2)
	if err != nil {
		return nil, errors.Wrap(err, "Wrap in GetUser UserInMemoryRepository: ")
	}
	user2 := builder2.AddId("0c51fe73-cbaf-432c-aa1a-3731f4b982c1").Build()

	myBirthDay3 := time.Date(1991, 11, 4, 0, 0, 0, 0, jst)
	builder3, err := userModel.NewUserBuilder("naoto", "mizoi", myBirthDay3)
	if err != nil {
		return nil, errors.Wrap(err, "Wrap in GetUser UserInMemoryRepository: ")
	}
	user3 := builder3.AddId("0c51fe73-cbaf-432c-aa1a-3731f2b982c1").Build()

	users := []userModel.UserEntity {*user1,*user2,*user3}
	user := find(users, func(user userModel.UserEntity) bool {
		return user.Id() == id
	})

	log.WithFields(log.Fields{
		"user": user,
	}).Debug("UserInMemoryRepository GetUGetUserById end")

	return user, nil

}
func find(vs []userModel.UserEntity, f func(userModel.UserEntity) bool) *userModel.UserEntity {

	for _, v := range vs {
		log.Printf(v.FullName(),f(v))
		if f(v) {
			return &v
		}
	}
	return nil
}


// コンストラクタ関数
func New() (*UserInMemoryRepository, error) {
	repository := UserInMemoryRepository{}
	return &repository, nil
}
