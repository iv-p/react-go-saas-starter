package user

import (
	"context"
	"errors"

	"github.com/iv-p/react-go-saas-starter/id"
	"gorm.io/gorm"
)

// PostgresRepository stores users in an in-memory array.
// Obviously not recommended to be used in production.
type PostgresRepository struct {
	db *gorm.DB

	Repository
}

// NewPostgresRepository instantiates a new user repository
// that holds users in memory.
func NewPostgresRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&User{})
	return &PostgresRepository{
		db: db,
	}
}

// GetUser returns a user based on the provided `userID`, or nil if not found
func (mr PostgresRepository) GetUser(ctx context.Context, userID id.ID) (*User, error) {
	var user User
	res := mr.db.First(&user, userID)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return &user, nil
}

// AddUser appends a new user record. Returns an error if `user.UserID` already exists
func (mr *PostgresRepository) AddUser(ctx context.Context, user User) error {
	res := mr.db.Create(user)
	return res.Error
}

// SetUser amends an user. Returns an error if `user.UserID` doesn't exists
func (mr *PostgresRepository) SetUser(ctx context.Context, user User) error {
	res := mr.db.Save(user)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("user doesn't exist")
	}
	return nil
}
