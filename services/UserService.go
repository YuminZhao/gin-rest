package services

import (
	"gin-rest/models"
	"gin-rest/rest/m"

	"gorm.io/gorm"
)

type UserService struct {
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

type UserCreate struct {
	Name     string `form:"name" label:"姓名" validate:"required,exists=users%id" message:"usercreate_name"`
	UserName string `form:"username" label:"用户名" validate:"required,username"`
}

func (s *UserService) GetUserById(id uint) (models.User, error) {
	var user models.User
	err := m.DB.Model(&models.User{}).First(&user, id).Error
	return user, err
}
func (s *UserService) CreateUser(name string) error {
	var user models.User
	user.Name = name
	return m.DB.Save(&user).Error
}
func (s *UserService) GetUser() ([]userList, error) {
	var data []userList
	err := m.DB.Table("users").Preload("Orders", func(db *gorm.DB) *gorm.DB {
		return db.Table("orders")
	}).Find(&data).Error
	return data, err
}
