package user

import (
	"context"

	"github.com/iv-p/react-go-saas-starter/id"
)

// AddUserRequest holds data needed to create a new User
type AddUserRequest struct {
	Name string
}

// Service is the interface for a service that handles User manipulation
type Service struct {
	repo Repository
}

// NewService instantiates a new user service with the provided repository
func NewService(repo Repository) Service {
	return Service{repo}
}

// AddUser appends a user record in the repository
func (s Service) AddUser(ctx context.Context, userRequest AddUserRequest) (User, error) {
	user := User{
		UserID: id.New(),
		Name:   userRequest.Name,
	}
	err := s.repo.AddUser(ctx, user)
	return user, err
}

// GetUser returns the user with corresponding `userID` or nil
func (s Service) GetUser(ctx context.Context, userID id.ID) (*User, error) {
	return s.repo.GetUser(ctx, userID)
}

// UpdateName of a user
func (s Service) UpdateName(ctx context.Context, userID id.ID, name string) error {
	user, err := s.repo.GetUser(ctx, userID)
	if err != nil {
		return err
	}
	if user == nil {
		return nil
	}
	user.Name = name
	err = s.repo.SetUser(ctx, *user)
	return err
}
