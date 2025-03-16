package serializer

import (
	"learn_ginmall/conf"
	"learn_ginmall/model"
)

type Product struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	CategoryId    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	View          uint   `json:"view"`
	CreateAt       int64  `json:"creat_at"`
	Num           int    `json:"num"`
	OnSale        bool    `json:"on_sale"`
	BossId        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
	BossAvatar    string `json:"boss_avatar"`
}

func BuildProduct(item *model.Product) *Product {
	return &Product{
		Id:	item.ID,
		Name: item.Name,
		CategoryId: item.Category,
		Title: item.Ttile,
		Info: item.Info,
		ImgPath: conf.Host + conf.HttpPort + conf.ProductPath + item.ImgPath,
		Price: item.Price,
		DiscountPrice: item.DiscountPrice,
		View: uint(item.View()),//点击数
		CreateAt: item.CreatedAt.Unix(),
		Num: item.Num,
		OnSale: item.OnSale,
		BossId: item.BossId,
		BossName: item.BossName,
		BossAvatar: conf.Host + conf.HttpPort + conf.AvatarPath + item.BossAvatar,
	}
}