package service

import (
	"context"
	"learn_ginmall/conf"
	"learn_ginmall/dao"
	"learn_ginmall/model"
	"learn_ginmall/pkg/e"
	"learn_ginmall/pkg/util"
	"learn_ginmall/serializer"
	"mime/multipart"
	"strconv"
	"sync"
)

type ProductService struct {
	Id            uint   `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	CategoryId    uint   `json:"category_id" form:"category_id"`
	Title         string `json:"title" form:"title"`
	Info          string `json:"info" form:"info"`
	ImgPath       string `json:"img_path" form:"img_path"`
	Price         string `json:"price" form:"price"`
	DiscountPrice string `json:"discount_price" form:"discount_price"`
	OnSale        bool   `json:"on_sale" form:"on_sale"`
	Num           int    `json:"num" form:"num"`
	model.BasePage
}

func (service *ProductService) Create(ctx context.Context, uId uint, files []*multipart.FileHeader) serializer.Response {
	var boss *model.User
	var err error
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	boss, _ = userDao.GetUserById(uId)
	//以第一张作为封面图 这里写的是上传到本地
	//TODO:上传到七牛云 
	tmp, _ := files[0].Open()
	var path string
	if conf.UploadModel == "local" {
		path, err = UploadProductToLocalStatic(tmp, uId, service.Name)
	} else {
		path, err = util.UploadToQiNiu(tmp, file[0].Size)
	}
	
	if err != nil {
		code = e.ErrorProductImgUpload
		util.Logrusobj.Infoln("service create uploadProductTolocalStatic ",err)
		return serializer.Response{
			Status: code,
			Data: e.GetMsg(code),
			Error: err.Error(),
		}
	}
	product := &model.Product{
			Name:	service.Name,
			Category:	service.CategoryId,
			Title:	service.Title,
			Info:	service.Info,
			ImgPath:	path,
			Price:	service.Price,
			DiscountPrice:	service.Price,
			OnSale:	true,
			Num:	service.Num,
			BossId: uId,
			BossName: boss.UserName,
			BossAvatar: boss.Avatar,
	}
	productDao := dao.NewProductDao(ctx)
	err = productDao.CreateProduct(product)
	if err != nil {
		code = e.Error
		util.Logrusobj.Infoln("service create CreateProduct ",err)
		return serializer.Response{
			Status: code,
			Data: e.GetMsg(code),
			Error: err.Error(),
		}
	}
	//并发创建
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for index, file := range files {
		num := strconv.Itoa(index)
		productImgDao := dao.NewProductImgDaoByDB(productDao.DB)
		tmp, _ = file.Open()
		path, err = UploadProductToLocalStatic(tmp, uId, service.Name+num)
		if err != nil {
			code = e.ErrorProductImgUpload
			return serializer.Response{
				Status: code,
				Data: nil,
				Msg: e.GetMsg(code),
				Error: err.Error(),
			}
		}
		productImg := model.ProductImg {
			ProductId: product.ID,
			ImgPath: path,
		}
		err = productImgDao.CreateProductImg(&productImg)
		if err != nil {
			code = e.Error
			return serializer.Response{
				Status: code,
				Data: nil,
				Msg: e.GetMsg(code),
				Error: err.Error(),
			}
		}
		wg.Done()
	}
	wg.Wait()
	return serializer.Response{
		Status: code,
		Data: serializer.BuildProduct(product),
		Msg: e.GetMsg(code),
	}
}

func (service *ProductService) Delete(ctx context.Context, pId string) serializer.Response {
	code := e.Success
	productDao := dao.NewProductDao(ctx)
	productId, _ := strconv.Atoi(pId)
	err := productDao.DeleteProduct(uint(productId))
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *ProductService) Update(ctx context.Context, pId string) serializer.Response {
	code := e.Success
	ProductDao := dao.NewProductDao(ctx)
	productId, _ := strconv.Atoi(pId)
	product :=&model.Product{
		Name:       service.Name,
		Category: uint(service.CategoryId),
		Title:      service.Title,
		Info:       service.Info,
		ImgPath:       service.ImgPath,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice,
		OnSale:        service.OnSale,
	}
	err := ProductDao.UpdateProduct(uint(productId), product)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg: e.GetMsg(code),
	}
}
