package registry

import admin "github.com/hi-sasaki/clean-architecture-golang-sample/pkg/usecase/admin"

type Provider interface {
	UserUsecase() *admin.User
}
