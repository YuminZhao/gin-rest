package services

import (
	"gin-rest/models"
	"gin-rest/rest/m"

	"gorm.io/gorm"
)

type userService struct {
	GetUserById func(id uint) (models.User, error)
	CreateUser  func(name string) error
	GetUser     func() ([]userList, error)
}

type userList struct {
	Id     uint
	Name   string
	Orders []struct {
		Id     uint
		UserId uint
		Price  uint
	} `gorm:"foreignkey:UserId"`
}

var UserService = userService{
	GetUserById: func(id uint) (models.User, error) {
		var user models.User
		err := m.DB.Model(&models.User{}).First(&user, id).Error
		return user, err
	},
	CreateUser: func(name string) error {
		var user models.User
		user.Name = name
		return m.DB.Save(&user).Error
	},
	GetUser: func() ([]userList, error) {
		var data []userList
		err := m.DB.Table("users").Preload("Orders", func(db *gorm.DB) *gorm.DB {
			return db.Table("orders")
		}).Find(&data).Error
		return data, err
	},
}
