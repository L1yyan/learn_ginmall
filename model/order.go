package model

import "gorm.io/gorm"

//订单
type Order struct {
	gorm.Model
	UserId    uint   `gorm:"not null"`
	ProductId uint   `gorm:"not null"`
	BossId    uint   `gorm:"not null"`
	AddressId uint   `gorm:"not null"`
	Num       int    //数量
	OrderNum  uint64 //订单号
	Type      uint   //1.未支付 2.已经支付
	Money     float64
}
