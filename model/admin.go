package model

import "gorm.io/gorm"
//admin
type Admin struct {
	gorm.Model
	UserName string
	PasswordDigest string
	Avatar string
}