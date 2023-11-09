package schema

import "bic-gin/pkg/db"

func AutoMigrate() error {
	if err := db.SqlDB().AutoMigrate(&User{}, &Menu{}, &Api{}); err != nil {
		return err
	}
	return db.SqlDB().AutoMigrate(
		&PageAction{},
		&Role{},
	)
}
