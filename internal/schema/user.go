package schema

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Nickname string
	Phone    string
	Email    string
}
