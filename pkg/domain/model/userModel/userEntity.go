package userModel

import (
	"github.com/pkg/errors"
	"time"
)

// 先頭を小文字にすることで外部パッケージから参照できないようにする
type UserEntity struct {
	id         idValue
	firstName  nameValue
	familyName nameValue
	birthday   birthdayValue
}
type UserBuilder struct {
	userEntity *UserEntity
}

// コンストラクタ関数
func NewUserBuilder(firstName string, familyName string, birthday time.Time) (*UserBuilder, error) {
	firstNameInstance, err := newNameValue(firstName)
	if err != nil {
		return nil, errors.Wrap(err, "Wrap in New userEntity: ")
	}

	familyNameInstance, err := newNameValue(familyName)
	if err != nil {
		return nil, errors.Wrap(err, "Wrap in New userEntity: ")
	}

	birthdayInstance, err := newBirthdayValue(birthday)
	if err != nil {
		return nil, errors.Wrap(err, "Wrap in New userEntity: ")
	}

	user := UserEntity{
		firstName:  *firstNameInstance,
		familyName: *familyNameInstance,
		birthday:   *birthdayInstance,
	}

	return &UserBuilder{&user}, nil
}

func (rcv *UserBuilder) AddId(id string) *UserBuilder {
	rcv.userEntity.id = idValue(id)
	return rcv
}
func (rcv *UserBuilder) Build() *UserEntity {
	if rcv.userEntity.id.value() == "" {
		createdID := newIdValue()
		rcv.userEntity.id = *createdID
	}
	return rcv.userEntity

}

func (rcv *UserEntity) Id() string {
	return rcv.id.value()
}

func (rcv *UserEntity) Age() int {
	return rcv.birthday.ConvertToAge()
}

func (rcv *UserEntity) FullName() string {
	return string(rcv.familyName + rcv.firstName)
}
