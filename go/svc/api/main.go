package main

import (
	"fmt"

	"github.com/caarlos0/env"
	mw "github.com/iv-p/react-go-saas-starter/middleware"
	"github.com/iv-p/react-go-saas-starter/svc/api/config"
	userc "github.com/iv-p/react-go-saas-starter/svc/api/user"
	"github.com/iv-p/react-go-saas-starter/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	var c config.Config
	err := env.Parse(&c)
	if err != nil {
		panic(err)
	}
	auth := mw.NewAuth0Authenticator(c.Auth0Domain)

	db, _ := gorm.Open(
		postgres.Open(
			fmt.Sprintf("host=postgres-dev user=postgresadmin password=admin123 dbname=postgresdb port=5432 sslmode=disable"),
		),
		&gorm.Config{},
	)

	userRepository := user.NewPostgresRepository(db)
	userService := user.NewService(userRepository)
	userRouter := userc.NewController(userService)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(auth.Protect)
	r.Mount("/user", userRouter)

	fmt.Println("Starting api")
	http.ListenAndServe(fmt.Sprintf(":%d", c.Port), r)
}
