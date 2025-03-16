package model

import "gorm.io/gorm"
//购物车
type Cart struct {
	gorm.Model
	UserId uint `gorm:"not null"`//用户id
	ProductId uint `gorm:"not null"`//产品id
	BossId uint `gorm:"not null"`//商家
	Num uint `gorm:"not null"`//数量
	MaxNum uint `gorm:"not null"`//购买限额
	Check bool//是否支付检验
}