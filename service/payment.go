package service

import (
	"context"
	"errors"
	"fmt"
	"learn_ginmall/cache"
	"learn_ginmall/dao"
	"learn_ginmall/model"
	"learn_ginmall/pkg/e"
	"learn_ginmall/pkg/util"
	"learn_ginmall/serializer"
	"strconv"
	"time"

	logging "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OrderPay struct {
	OrderId   uint    `form:"order_id" json:"order_id"`
	Money     float64 `form:"money" json:"money"`
	OrderNo   string  `form:"orderNo" json:"orderNo"`
	ProductId int     `form:"product_id" json:"product_id"`
	PayTime   string  `form:"payTime" json:"payTime" `
	Sign      string  `form:"sign" json:"sign" `
	BossId    int     `form:"boss_id" json:"boss_id"`
	BossName  string  `form:"boss_name" json:"boss_name"`
	Num       int     `form:"num" json:"num"`
	Key       string  `form:"key" json:"key"`
}

func (service *OrderPay) PayDown(ctx context.Context, uId uint) serializer.Response {
	code := e.Success

	err := dao.NewOrderDao(ctx).Transaction(func(tx *gorm.DB) error {

		util.Encrypt.SetKey(service.Key)
		orderDao := dao.NewOrderDaoByDB(tx)

		order, err := orderDao.GetOrderById(service.OrderId)
		if err != nil {
			logging.Info(err)
			return err
		}
		//先对该订单是否超时进行判断 从redis取出，如果超时则回滚，否则继续
		score, err := cache.RedisClient.ZScore(OrderTimeKey, strconv.FormatUint(order.OrderNum, 10)).Result()
		currentime := time.Now().Unix()
		if currentime > int64(score) {
			//订单超时了 后续操作取消
			code = e.ErrorRedis
			return errors.New("订单已超时")
		}

		money := order.Money
		num := order.Num
		money = money * float64(num)

		userDao := dao.NewUserDaoByDB(tx)
		user, err := userDao.GetUserById(uId)
		if err != nil {
			logging.Info(err)
			code = e.ErrorRedis
			return err
		}

		// 对钱进行解密。减去订单。再进行加密。
		moneyStr := util.Encrypt.AesDecoding(user.Money)
		moneyFloat, _ := strconv.ParseFloat(moneyStr, 64)
		if moneyFloat-money < 0.0 { // 金额不足进行回滚
			logging.Info(err)
			code = e.ErrorDatabase
			return errors.New("金币不足")
		}

		finMoney := fmt.Sprintf("%f", moneyFloat-money)
		user.Money = util.Encrypt.AesEncoding(finMoney)

		err = userDao.UpdateUserById(uId, user)
		if err != nil { // 更新用户金额失败，回滚
			logging.Info(err)
			code = e.ErrorDatabase
			return err
		}
		boss := new(model.User)
		boss, err = userDao.GetUserById(uint(service.BossId))
		moneyStr = util.Encrypt.AesDecoding(boss.Money)
		moneyFloat, _ = strconv.ParseFloat(moneyStr, 64)
		finMoney = fmt.Sprintf("%f", moneyFloat+money)
		boss.Money = util.Encrypt.AesEncoding(finMoney)

		err = userDao.UpdateUserById(uint(service.BossId), boss)
		if err != nil { // 更新boss金额失败，回滚
			logging.Info(err)
			code = e.ErrorDatabase
			return err
		}

		product := new(model.Product)
		productDao := dao.NewProductDaoByDB(tx)
		product, err = productDao.GetProductById(uint(service.ProductId))
		if err != nil {
			return err
		}
		product.Num -= num
		err = productDao.UpdateProduct(uint(service.ProductId), product)
		if err != nil { // 更新商品数量减少失败，回滚
			logging.Info(err)
			code = e.ErrorDatabase
			return err
		}

		// 更新订单状态
		order.Type = 2
		err = orderDao.UpdateOrderById(service.OrderId, order)
		if err != nil { // 更新订单失败，回滚
			logging.Info(err)
			code = e.ErrorDatabase
			return err
		}

		productUser := model.Product{
			Name:          product.Name,
			Category:      product.Category,
			Title:         product.Title,
			Info:          product.Info,
			ImgPath:       product.ImgPath,
			Price:         product.Price,
			DiscountPrice: product.DiscountPrice,
			Num:           num,
			OnSale:        false,
			BossId:        uId,
			BossName:      user.UserName,
			BossAvatar:    user.Avatar,
		}

		err = productDao.CreateProduct(&productUser)
		if err != nil { // 买完商品后创建成了自己的商品失败。订单失败，回滚
			logging.Info(err)
			code = e.ErrorDatabase
			return err
		}

		return nil

	})

	if err != nil {
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
