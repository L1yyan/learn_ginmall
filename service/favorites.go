package service

import (
	"context"
	"learn_ginmall/dao"
	"learn_ginmall/model"
	"learn_ginmall/pkg/e"
	"learn_ginmall/serializer"

)

type FavoritesService struct {
	ProductId  uint `form:"product_id" json:"product_id"`
	BossId     uint `form:"boss_id" json:"boss_id"`
	FavoriteId uint `form:"favorite_id" json:"favorite_id"`
	PageNum    int  `form:"pageNum"`
	PageSize   int  `form:"pageSize"`
}

// Show 商品收藏夹
func (service *FavoritesService) Show(ctx context.Context, uId uint) serializer.Response {
	favoritesDao := dao.NewFavoritesDao(ctx)
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	favorites, total, err := favoritesDao.ListFavoriteByUserId(uId, service.PageSize, service.PageNum)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildFavorites(ctx, favorites), uint(total))
}

// Create 创建收藏夹
func (service *FavoritesService) Create(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	favoriteDao := dao.NewFavoritesDao(ctx)
	exist, _ := favoriteDao.FavoriteExistOrNot(service.ProductId, uId)
	if exist {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	bossDao := dao.NewUserDaoByDB(userDao.DB)
	boss, err := bossDao.GetUserById(service.BossId)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	favorite := &model.Favorite{
		UserId:    uId,
		User:      *user,
		ProductId: service.ProductId,
		Product:   *product,
		BossId:    service.BossId,
		Boss:      *boss,
	}
	favoriteDao = dao.NewFavoritesDaoByDB(favoriteDao.DB)
	err = favoriteDao.CreateFavorite(favorite)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Delete 删除收藏夹
func (service *FavoritesService) Delete(ctx context.Context) serializer.Response {
	code := e.Success

	favoriteDao := dao.NewFavoritesDao(ctx)
	err := favoriteDao.DeleteFavoriteById(service.FavoriteId)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   e.GetMsg(code),
	}
}
