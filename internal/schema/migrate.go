package schema

import "bic-gin/pkg/db"

func AutoMigrate() error {
	return db.SqlDB().AutoMigrate(&User{})
}
