package dao

import (
	"context"
	"learn_ginmall/model"

	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBclient(ctx)}
}

func (dao *ProductDao) CreateProduct(product *model.Product) (err error) {
	return dao.DB.Model(&model.Product{}).Create(&product).Error
}

func (dao *ProductDao) DeleteProduct(pId uint) error {
	return dao.DB.Where("id = ?", pId).Delete(&model.Product{}).Error
}

func (dao *ProductDao) UpdateProduct(pId uint, product *model.Product) error {
	return dao.DB.Model(&model.Product{}).Where("id=?", pId).
		Updates(&product).Error
}

