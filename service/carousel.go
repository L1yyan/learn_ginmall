package service

import (
	"context"
	"learn_ginmall/dao"
	"learn_ginmall/pkg/e"
	"learn_ginmall/pkg/util"
	"learn_ginmall/serializer"
)

type CarouselService struct {

}

func (service *CarouselService) List (ctx context.Context) serializer.Response {
	carouselDao :=  dao.NewCarouselDao(ctx)
	code := e.Success
	carousels, err := carouselDao.ListCarousel()
	if err != nil {
		util.Logrusobj.Info("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCarousels(carousels), uint(len(carousels)))
	   
} 