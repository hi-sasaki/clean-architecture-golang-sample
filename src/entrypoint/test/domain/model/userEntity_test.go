package model

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/rikodao/clean-architecture-golang-sample/src/lib/domain/model/userModel"
	"math"
	"testing"
	"time"
)

func TestUserEntitySuccess(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	myBirthDay := time.Date(1991,11,4,0,0,0,0, jst)
	result, err := userModel.New("naoto","oiso", myBirthDay)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	log.Print(result)
	today := time.Now()
	myAge := int(math.Floor(today.Sub(myBirthDay).Hours() / 24 / 365))
	assert.Equal(t, result.Age(), myAge, "they should be equal")
	assert.Equal(t, result.FullName(), "oisonaoto", "they should be equal")
}

func TestUserEntityAgeCheckFailed(t *testing.T) {
	tooYoungBirthday := time.Now().Add(-time.Hour * 340)
	_, err := userModel.New("naoto","oiso", tooYoungBirthday)
	assert.Error(t,err, "TestUserEntityAgeCheckFailed")
}

func TestUserEntityNameEmptyCheckFailed(t *testing.T) {
	tooYoungBirthday := time.Now().Add(-time.Hour * 340)
	_, err := userModel.New("","oiso", tooYoungBirthday)
	assert.Error(t,err, "TestUserEntityNameEmptyCheckFailed")
}

func TestUserEntityNameOverCheckFailed(t *testing.T) {
	tooYoungBirthday := time.Now().Add(-time.Hour * 340)
	_, err := userModel.New("naoto","oisosssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss", tooYoungBirthday)
	assert.Error(t,err, "TestUserEntityNameOverCheckFailed")
}