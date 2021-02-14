package main

import (
	mw "github.com/iv-p/react-go-saas-starter/middleware"
	userc "github.com/iv-p/react-go-saas-starter/svc/api/user"
	"github.com/iv-p/react-go-saas-starter/user"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	auth := mw.NewAuth0Authenticator("", "")

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
