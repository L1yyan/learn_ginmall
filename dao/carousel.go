package dao

import (
	"context"
	"learn_ginmall/model"

	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

func NewCarouselDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{NewDBclient(ctx)}
}

func NewCarouselDaoByDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

// GetNOticeById 根据id获取carousel
func (dao *NoticeDao) ListCarousel() (carousel []model.Carousel, err error) {
	err = dao.DB.Model(&model.Carousel{}).Find(&carousel).Error
	return carousel, err
}
