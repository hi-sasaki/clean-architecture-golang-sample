package registry

import (
	"database/sql"

	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/driver/mysql"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/driver/mysql/dao"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/entity/service"
	admin "github.com/hi-sasaki/clean-architecture-golang-sample/pkg/usecase/admin"
)

type Provider interface {
	UserUsecase() *admin.User
}

type Injection struct {
	Config
}

var db *sql.DB

func init() {
	d, err := mysql.NewDB()
	if err != nil {
		panic(err)
	}
	db = d
}

func NewProvider() Provider {
	return &Injection{}
}

func (p *Injection) UserUsecase() *admin.User {
	service := &service.User{}
	service.User = dao.UserProvider(db)
	return &admin.User{
		Service: service,
	}
}
