package models

import (
	"gin-rest/rest/m"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func init() {
	m.DB.AutoMigrate(&User{})
}
