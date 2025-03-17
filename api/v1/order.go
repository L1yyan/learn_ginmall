package v1

import (
	"learn_ginmall/pkg/util"
	"learn_ginmall/service"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var createOrderService service.OrderService
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createOrderService); err == nil {
		res := createOrderService.Create(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

func ListOrders(c *gin.Context) {
	var listOrdersService service.OrderService
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listOrdersService); err == nil {
		res := listOrdersService.List(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

// 订单详情
func ShowOrder(c *gin.Context) {
	var showOrderService service.OrderService
	
	if err := c.ShouldBind(&showOrderService); err == nil {
		res := showOrderService.Show(c.Request.Context(), c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

func DeleteOrder(c *gin.Context) {
	var deleteOrderService service.OrderService
	
	if err := c.ShouldBind(&deleteOrderService); err == nil {
		res := deleteOrderService.Delete(c.Request.Context(), c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}
