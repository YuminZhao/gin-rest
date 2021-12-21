package models

import (
	"gin-rest/rest/m"

	"gorm.io/gorm"
)

type order struct {
	gorm.Model
	UserId uint `gorm:"index"`
	Price  uint
}

type Order struct {
	gorm.Model `json:"-"`
	UserId     uint `json:"user_id"`
	Price      uint `json:"price"`
}

func init() {
	m.DB.AutoMigrate(&order{})
}
