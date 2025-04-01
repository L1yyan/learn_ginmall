package routes

import (
	api "learn_ginmall/api/v1"
	"learn_ginmall/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		//用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		
		//轮播图
		v1.GET("carousels", api.ListCarousel)
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			//用户操作
			authed.PUT("user", api.UserUpdate)
			authed.PUT("avatar", api.UploadAvatar)
			authed.POST("user/sending-email", api.SendEmail)
			authed.POST("user/valid-email", api.ValidEmail)


			
			//商品操作
			authed.POST("product", api.CreateProduct)
			authed.PUT("product/:id", api.UpdateProduct)
			authed.DELETE("product/:id", api.DeleteProduct)
			// 收藏夹
			authed.GET("favorites", api.ShowFavorites)
			authed.POST("favorites", api.CreateFavorite)
			authed.DELETE("favorites/:id", api.DeleteFavorite)
			// 订单操作
			authed.POST("orders", api.CreateOrder)
			authed.GET("orders", api.ListOrders)
			authed.GET("orders/:id", api.ShowOrder)
			authed.DELETE("orders/:id", api.DeleteOrder)

			// 购物车
			authed.POST("carts", api.CreateCart)
			authed.GET("carts", api.ShowCarts)
			authed.PUT("carts/:id", api.UpdateCart) // 购物车id
			authed.DELETE("carts/:id", api.DeleteCart)

			// 收获地址操作
			authed.POST("addresses", api.CreateAddress)
			authed.GET("addresses/:id", api.GetAddress)
			authed.GET("addresses", api.ListAddress)
			authed.PUT("addresses/:id", api.UpdateAddress)
			authed.DELETE("addresses/:id", api.DeleteAddress)

			// 支付功能
			authed.POST("paydown", api.OrderPay)

			// 显示金额
			authed.POST("money", api.ShowMoney)
		}

	}

	return r

}
