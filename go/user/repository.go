package user

import (
	"context"

	"github.com/iv-p/react-go-saas-starter/id"
)

// Repository is the interface for an user repository
type Repository interface {
	GetUser(ctx context.Context, id id.ID) (*User, error)
	AddUser(ctx context.Context, user User) error
	SetUser(ctx context.Context, user User) error
}
