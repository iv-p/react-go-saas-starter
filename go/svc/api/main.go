package main

import (
	userc "github.com/iv-p/react-go-saas-starter/svc/api/user"
	"github.com/iv-p/react-go-saas-starter/user"
	"github.com/kataras/iris/v12"
)

func main() {
	userRepository := user.NewMemoryRepository()
	userService := user.NewService(userRepository)
	userController := userc.NewController(userService)

	app := iris.New()

	app.Post("/user", userController.AddUser)
	app.Get("/user/:userID", userController.GetUser)
	app.Put("/user/:userID/name", userController.UpdateName)

	app.Listen(":8080")
}
