package userModel

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"math"
	"time"
)

type birthdayValue time.Time
const (
	minimumAge = 18
	maxAge = 128
)
func newBirthdayValue(birthday time.Time)(*birthdayValue, error) {
	log.WithFields(log.Fields{
		"birthday": birthday,
	}).Debug("newBirthdayValue start")
	ageValueObject := birthdayValue(birthday)

	age := ageValueObject.ConvertToAge()
	if age < minimumAge {
		return nil, errors.New("18歳未満はユーザー登録出来ません。")
	}
	if age > maxAge {
		return nil, errors.New("年齢上限を超えています。")
	}


	defer log.WithFields(log.Fields{
		"ageValueObject": ageValueObject,
	}).Debug("newBirthdayValue end")

	return &ageValueObject, nil

}


func (rcv *birthdayValue)  ConvertToAge() int {
	log.WithFields(log.Fields{
		"birthday": rcv,
	}).Debug("convertToAge start")

	today := time.Now()
	age := int(math.Floor(today.Sub(time.Time(*rcv)).Hours() / 24 / 365))

	defer log.WithFields(log.Fields{
		"age": age,
	}).Debug("convertToAge end")

	return age

}

func (rcv birthdayValue) value() time.Time {
	return time.Time(rcv)
}