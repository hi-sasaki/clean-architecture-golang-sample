package security

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/registry"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/usecase/inout"
)

type Security struct {
	provider registry.Provider
}

func NewSecurity(p registry.Provider) *Security {
	return &Security{provider: p}
}

func (s *Security) Login(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	c, cancel := context.WithTimeout(ctx, time.Duration(time.Second)*30)
	defer cancel()

	user := new(inout.User)
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		return
	}
	log.Println("id:" + user.ID)
	log.Println("pass:" + user.Password)
	token, err := s.provider.LoginUsecase().Login(c, user.ID, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
	w.Write([]byte(token))
}
