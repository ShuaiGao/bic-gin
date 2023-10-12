package schema

import "gin-bic/pkg/db"

func AutoMigrate() error {
	return db.SqlDB().AutoMigrate(&User{})
}
