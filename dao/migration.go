package dao

import (
	"fmt"
	"learn_ginmall/model"
)

func Migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.User{},
		&model.Address{},
		&model.Admin{},
		&model.Category{},
		&model.Carousel{},
		&model.Cart{},
		&model.Notice{},
		&model.Product{},
		&model.ProductImg{},
		&model.Order{},
		&model.Favorite{},
	)
	if err != nil {
		fmt.Println("err: ", err)
	}
}
