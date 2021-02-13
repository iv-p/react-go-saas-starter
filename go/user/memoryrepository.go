package user

import (
	"context"
	"errors"

	"github.com/iv-p/react-go-saas-starter/id"
)

// MemoryRepository stores users in an in-memory array.
// Obviously not recommended to be used in production.
type MemoryRepository struct {
	users []User

	Repository
}

// NewMemoryRepository instantiates a new user repository
// that holds users in memory.
func NewMemoryRepository() Repository {
	return &MemoryRepository{}
}

// GetUser returns a user based on the provided `userID`, or nil if not found
func (mr MemoryRepository) GetUser(ctx context.Context, userID id.ID) (*User, error) {
	for _, user := range mr.users {
		if user.UserID == userID {
			return &user, nil
		}
	}
	return nil, nil
}

// AddUser appends a new user record. Returns an error if `user.UserID` already exists
func (mr *MemoryRepository) AddUser(ctx context.Context, user User) error {
	for _, user := range mr.users {
		if user.UserID == user.UserID {
			return errors.New("user already exists")
		}
	}

	mr.users = append(mr.users, user)
	return nil
}

// SetUser amends an user. Returns an error if `user.UserID` doesn't exists
func (mr *MemoryRepository) SetUser(ctx context.Context, user User) error {
	for idx, user := range mr.users {
		if user.UserID == user.UserID {
			mr.users[idx] = user
			return nil
		}
	}
	return errors.New("user doesn't exist")
}
