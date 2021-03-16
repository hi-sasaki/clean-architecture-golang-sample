package userModel

import (
	"github.com/pkg/errors"
	"time"
)

type UserData struct{
	Id   string
	FirstName string
	FamilyName string
	Birthday time.Time
}


// 先頭を小文字にすることで外部パッケージから参照できないようにする
type UserEntity struct{
	id         idValue
	firstName  nameValue
	familyName nameValue
	birthday   birthdayValue
}

// コンストラクタ関数
func New(firstName string,familyName string,birthday time.Time) (*UserEntity, error) {

	createdID, err := newIdValue()
	if err != nil {
		return nil, err
	}

	createdFirstName, err := newNameValue(firstName)
	if err != nil {
		return nil, errors.Wrap(err, "Wrap in New userEntity: ")
	}

	createdFamilyName, err := newNameValue(familyName)
	if err != nil {
		return nil, errors.Wrap(err, "Wrap in New userEntity: ")
	}

	createdBirthday, err := newBirthdayValue(birthday)
	if err != nil {
		return nil, errors.Wrap(err, "Wrap in New userEntity: ")
	}

	user := UserEntity{
		id:   *createdID,
		firstName: *createdFirstName,
		familyName: *createdFamilyName,
		birthday: *createdBirthday,
	}
	return &user, nil
}

func (rcv *UserEntity) Age() int {
	return rcv.birthday.ConvertToAge()
}

func (rcv *UserEntity) FullName() string {
	return string(rcv.familyName + rcv.firstName)
}

func (rcv *UserEntity) Value() *UserData {
	return &UserData{
		rcv.id.value(),
		rcv.firstName.value(),
		rcv.familyName.value(),
		rcv.birthday.value(),
	}
}