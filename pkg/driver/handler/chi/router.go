package chi

import (
	"context"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	firebaseAuth "firebase.google.com/go/auth"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	admin "github.com/hi-sasaki/clean-architecture-golang-sample/pkg/adapter/controller/admin"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/adapter/controller/security"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/driver/handler/middleware/auth"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/registry"
)

func NewRouter(p registry.Provider) *chi.Mux {
	firebaseApp, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	authClient, err := firebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	r.Post("/login", security.NewSecurity(p).Login)

	r.Group(func(r chi.Router) {
		r.Use(
			middleware.Logger,
			corsMiddleware(),
			authMiddleware(authClient),
			middleware.SetHeader("X-Content-Type-Options", "nosniff"),
			middleware.SetHeader("X-Frame-Options", "deny"),
		)
		u := admin.NewUser(p)
		r.Get("/user/{userID}", u.GetByID)
		r.Post("/user", u.CreateUser)
	})
	return r
}

func authMiddleware(authClient *firebaseAuth.Client) func(http.Handler) http.Handler {
	return auth.FirebaseHttpMiddleware{authClient}.Middleware

}

func corsMiddleware() func(http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler
}
