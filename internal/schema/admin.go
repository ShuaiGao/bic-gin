package schema

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;size:128"`
	Name     string `gorm:"size:255"`
	Nickname string `gorm:"size:255"`
	Phone    string `gorm:"size:255"`
	Email    string `gorm:"size:255"`
	Ban      bool
}

type Api struct {
	Key    string `gorm:"primarykey;size:128;"`
	Url    string `gorm:"uniqueIndex:api-unique;size:120"`
	Method string `gorm:"uniqueIndex:api-unique;size:7"`
	Label  string `gorm:"size:128"`
}

type Menu struct {
	Key       string `gorm:"primarykey;size:32;"`
	Label     string `gorm:"size:128"`
	Rank      int
	FatherKey string `gorm:"default:null;size:32"`
	Father    *Menu  `gorm:"default:galeone;foreignKey:FatherKey;References:key;"`
}

type PageAction struct {
	Key     string `gorm:"uniqueIndex;size:64"`
	Label   string `gorm:"size:128"`
	MenuKey string `gorm:"size:32"`
	Menu    Menu   `gorm:"foreignKey:MenuKey;References:Key"`
	Apis    []Api  `gorm:"many2many:page_action_api;foreignKey:Key;"`
}

type Role struct {
	gorm.Model
	Name        string `gorm:"size:128"`
	UserID      uint
	User        User
	PageActions []PageAction `gorm:"many2many:role_page_action;foreignKey:ID;References:Key"`
	Users       []User       `gorm:"many2many:role_user;foreignKey:ID;References:ID"`
}
