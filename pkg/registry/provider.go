package registry

import (
	"context"
	"database/sql"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	fi "github.com/hi-sasaki/clean-architecture-golang-sample/pkg/driver/firebase"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/driver/mysql"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/driver/mysql/dao"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/entity/service"
	admin "github.com/hi-sasaki/clean-architecture-golang-sample/pkg/usecase/admin"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/usecase/security"
)

type Provider interface {
	UserUsecase() *admin.User
	LoginUsecase() *security.Security
}

type Injection struct {
}

var db *sql.DB
var authClient *auth.Client

func init() {
	d, err := mysql.NewDB()
	if err != nil {
		panic(err)
	}
	db = d

	firebaseApp, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	client, err := firebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	authClient = client
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

func (p *Injection) LoginUsecase() *security.Security {
	s := &service.Security{}
	s.Security = fi.Provider(authClient)

	userService := &service.User{}
	userService.User = dao.UserProvider(db)
	return &security.Security{
		Service:     s,
		UserService: userService,
	}
}
