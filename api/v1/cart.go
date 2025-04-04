package v1

import (
	"learn_ginmall/pkg/util"
	"learn_ginmall/service"

	"github.com/gin-gonic/gin"
)

func CreateCart(c *gin.Context) {
	var createCartService service.CartService
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createCartService); err == nil {
		res := createCartService.Create(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.Logrusobj.Infoln(err)
	}
}

// 购物车详细信息
func ShowCarts(c *gin.Context) {
	var showCartService service.CartService
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	res := showCartService.Show(c.Request.Context(), claim.ID)
	c.JSON(200, res)
}

// 修改购物车信息
func UpdateCart(c *gin.Context) {
	var updateCartService service.CartService
	if err := c.ShouldBind(&updateCartService); err == nil {
		res := updateCartService.Update(c.Request.Context(), c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.Logrusobj.Infoln(err)
	}
}

// 删除购物车
func DeleteCart(c *gin.Context) {
	var deleteCartService service.CartService
	if err := c.ShouldBind(&deleteCartService); err == nil {
		res := deleteCartService.Delete(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.Logrusobj.Infoln(err)
	}
}
