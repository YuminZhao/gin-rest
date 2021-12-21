package services

import (
	"gin-rest/models"
	"gin-rest/rest/m"

	"gorm.io/gorm"
)

type orderService struct {
	GetOrder    func() ([]orderList, error)
	CreateOrder func(price uint) error
}

type orderList struct {
	Id     uint
	UserId uint
	Price  string
	User   struct {
		Id     uint
		Name   string
		Orders []struct {
			UserId uint
			Price  uint
		} `gorm:"foreignkey:UserId"`
	}
}

var OrderService = orderService{
	GetOrder: func() ([]orderList, error) {
		var order []orderList
		err := m.DB.Table("orders").Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Table("users")
		}).Preload("User.Orders", func(db *gorm.DB) *gorm.DB {
			return db.Table("orders")
		}).Find(&order).Error
		return order, err
	},

	CreateOrder: func(price uint) error {
		var order models.Order
		order.UserId = 1
		order.Price = price
		return m.DB.Save(&order).Error
	},
}
