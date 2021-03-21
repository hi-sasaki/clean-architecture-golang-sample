package model

import (
	"github.com/rikodao/clean-architecture-golang-sample/pkg/domain/model/userModel"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"time"
)

func TestUserEntitySuccess(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	myBirthDay := time.Date(1991, 11, 4, 0, 0, 0, 0, jst)
	builder, err := userModel.NewUserBuilder("naoto", "oiso", myBirthDay)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	result := builder.Build()

	log.Print(result)
	today := time.Now()
	myAge := int(math.Floor(today.Sub(myBirthDay).Hours() / 24 / 365))
	assert.Equal(t, result.Age(), myAge, "they should be equal")
	assert.Equal(t, result.FullName(), "oisonaoto", "they should be equal")
}

func TestUserEntityWithIdSuccess(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	myBirthDay := time.Date(1991, 11, 4, 0, 0, 0, 0, jst)
	builder, err := userModel.NewUserBuilder("naoto", "oiso", myBirthDay)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	id := "0c51fe73-cbaf-432c-aa1a-3731f4b982c1"
	result := builder.AddId(id).Build()

	log.Print(result)
	today := time.Now()
	myAge := int(math.Floor(today.Sub(myBirthDay).Hours() / 24 / 365))
	assert.Equal(t, result.Id(), id, "they should be equal")
	assert.Equal(t, result.Age(), myAge, "they should be equal")
	assert.Equal(t, result.FullName(), "oisonaoto", "they should be equal")
}

func TestUserEntityAgeCheckFailed(t *testing.T) {
	tooYoungBirthday := time.Now().Add(-time.Hour * 340)

	_, err := userModel.NewUserBuilder("naoto", "oiso", tooYoungBirthday)
	assert.Error(t, err, "TestUserEntityAgeCheckFailed")
}

func TestUserEntityNameEmptyCheckFailed(t *testing.T) {
	tooYoungBirthday := time.Now().Add(-time.Hour * 340)
	_, err := userModel.NewUserBuilder("", "oiso", tooYoungBirthday)
	assert.Error(t, err, "TestUserEntityNameEmptyCheckFailed")
}

func TestUserEntityNameOverCheckFailed(t *testing.T) {
	tooYoungBirthday := time.Now().Add(-time.Hour * 340)
	_, err := userModel.NewUserBuilder("", "oisosssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss", tooYoungBirthday)
	assert.Error(t, err, "TestUserEntityNameOverCheckFailed")
}
