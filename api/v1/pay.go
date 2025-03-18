package v1

import (
	"learn_ginmall/pkg/util"
	"learn_ginmall/service"

	"github.com/gin-gonic/gin"
)

func OrderPay(c *gin.Context) {
	var orderPay service.OrderPay
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&orderPay); err == nil {
		res := orderPay.PayDown(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		util.Logrusobj.Infoln(err)
		c.JSON(400, ErrorResponse(err))
	}
}
