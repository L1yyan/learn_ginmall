package model

import (
	"learn_ginmall/cache"
	"strconv"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name string
	Category uint//类别
	Title string
	Info string
	ImgPath string
	Price string
	DiscountPrice string
	OnSale bool `gorm:"default:false"`
	Num int
	BossId uint
	BossName string
	BossAvatar string//头像

}

func (product *Product) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.ProductViewKey(product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count 
}
func (product *Product) AddView() {
	//增加商品点击数
	cache.RedisClient.Incr(cache.ProductViewKey(product.ID))
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(product.ID)))
	
}