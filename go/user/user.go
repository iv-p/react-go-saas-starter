package user

import "github.com/iv-p/react-go-saas-starter/id"

// User holds information for a single user
type User struct {
	UserID id.ID
	Name string
}