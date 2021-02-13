package controller

import (
	"github.com/iv-p/react-go-saas-starter/user"
	"github.com/kataras/iris/v12"
)

// Controller is the entrypoint of the user related endpoibts
type Controller struct {
	userService user.Service
}

// NewController instantiates a new user controller
func NewController(userService user.Service) Controller {
	return Controller{userService}
}

// GetUser returns a user based on the provided `userID`
func (c Controller) GetUser(ctx iris.Context) {
	userID := ctx.Params().Get("userID")
	user, err := c.userService.GetUser(ctx.Request().Context(), userID)
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Could not fetch user").DetailErr(err))
		return
	}

	if user == nil {
		ctx.StatusCode(404)
		return
	}

	ctx.JSON(user)
}

// AddUserRequest holds necessary data for new user creation
type AddUserRequest struct {
	Name string `json:"name"`
}

// AddUser creates a new user
func (c Controller) AddUser(ctx iris.Context) {
	var request UpdateNameRequest
	err := ctx.ReadJSON(&request)

	user, err := c.userService.AddUser(ctx.Request().Context(), user.AddUserRequest(request))
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Could not add user").DetailErr(err))
		return
	}

	ctx.JSON(user)
}

// UpdateNameRequest holds necessary data for updating the name of a user
type UpdateNameRequest struct {
	Name string `json:"name"`
}

// UpdateName changes the name of a user if found
func (c Controller) UpdateName(ctx iris.Context) {
	userID := ctx.Params().Get("userID")
	var request UpdateNameRequest
	err := ctx.ReadJSON(&request)
	if err != nil || len(request.Name) < 2 {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}

	err = c.userService.UpdateName(ctx.Request().Context(), userID, request.Name)
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Could not fetch user").DetailErr(err))
		return
	}
}
