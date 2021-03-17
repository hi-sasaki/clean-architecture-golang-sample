package userModel

import (
	"time"
)

type UserData struct {
	Id         string `json:"id"`
	FirstName  string `json:"firstName"`
	FamilyName string `json:"familyName"`
	Birthday   time.Time `json:"birthday"`
}

func Data(user UserEntity) *UserData {
	return &UserData{
		user.id.value(),
		user.firstName.value(),
		user.familyName.value(),
		user.birthday.value(),
	}

}