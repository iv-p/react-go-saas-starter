package controller

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/iv-p/react-go-saas-starter/user"
)

// Controller is the entrypoint of the user related endpoibts
type Controller struct {
	userService user.Service
}

// NewController instantiates a new user controller
func NewController(userService user.Service) chi.Router {
	c := Controller{userService}
	r := chi.NewRouter()

	r.Post("/", c.AddUser)
	r.Get("/{userID}", c.GetUser)
	r.Put("/{userID}/name", c.UpdateName)

	return r
}

// GetUser returns a user based on the provided `userID`
func (c Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	user, err := c.userService.GetUser(r.Context(), userID)
	if err != nil {
		http.Error(w, "could not fetch user", http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	render.JSON(w, r, user)
}

// AddUserRequest holds necessary data for new user creation
type AddUserRequest struct {
	Name string `json:"name"`
}

func (a *AddUserRequest) Bind(r *http.Request) error {
	return nil
}

// AddUser creates a new user
func (c Controller) AddUser(w http.ResponseWriter, r *http.Request) {
	request := &AddUserRequest{}
	if err := render.Bind(r, request); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user, err := c.userService.AddUser(r.Context(), user.AddUserRequest(*request))
	if err != nil {
		http.Error(w, "could not add user", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, user)
}

// UpdateNameRequest holds necessary data for updating the name of a user
type UpdateNameRequest struct {
	Name string `json:"name"`
}

func (a *UpdateNameRequest) Bind(r *http.Request) error {
	return nil
}

// UpdateName changes the name of a user if found
func (c Controller) UpdateName(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	request := &UpdateNameRequest{}
	if err := render.Bind(r, request); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err := c.userService.UpdateName(r.Context(), userID, request.Name)
	if err != nil {
		http.Error(w, "could not update user", http.StatusInternalServerError)
		return
	}
}
