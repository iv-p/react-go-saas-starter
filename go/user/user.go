package user

import (
	"github.com/iv-p/react-go-saas-starter/id"
	"gorm.io/gorm"
)

// User holds information for a single user
type User struct {
	gorm.Model
	UserID id.ID
	Name   string
}
