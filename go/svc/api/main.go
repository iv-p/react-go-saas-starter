package main

import (
	"github.com/caarlos0/env"
	"github.com/iv-p/react-go-saas-starter/config"
	mw "github.com/iv-p/react-go-saas-starter/middleware"
	userc "github.com/iv-p/react-go-saas-starter/svc/api/user"
	"github.com/iv-p/react-go-saas-starter/user"

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

	userRepository := user.NewMemoryRepository()
	userService := user.NewService(userRepository)
	userRouter := userc.NewController(userService)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(auth.Protect)
	r.Mount("/user", userRouter)

	http.ListenAndServe(":8080", r)
}
