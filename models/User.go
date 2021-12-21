package models

import (
	"gin-rest/rest/m"

	"gorm.io/gorm"
)

type user struct {
	gorm.Model
	Name string `gorm:"size:255"`
}

type User struct {
	gorm.Model `json:"-"`
	Name       string `json:"name"`
}

func init() {
	m.DB.AutoMigrate(&user{})
}
