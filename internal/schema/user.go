package schema

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Name     string
	Nickname string
	Phone    string
	Email    string
	Ban      bool
}
