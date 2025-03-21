package dao

import (
	"context"
	"learn_ginmall/model"

	"gorm.io/gorm"
)

type ProductImgDao struct {
	*gorm.DB
}

func NewProductImgDao(ctx context.Context) *ProductImgDao {
	return &ProductImgDao{NewDBclient(ctx)}
}

func NewProductImgDaoByDB(db *gorm.DB) *ProductImgDao {
	return &ProductImgDao {db}
}

func (dao *ProductImgDao)CreateProductImg(productImg *model.ProductImg) error {
	return dao.DB.Model(&model.ProductImg{}).Create(&productImg).Error
}

