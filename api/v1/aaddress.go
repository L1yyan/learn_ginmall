package v1

import (
	"learn_ginmall/service"
	"learn_ginmall/pkg/util"
	"github.com/gin-gonic/gin"
)

// CreateAddress 新增收货地址
func CreateAddress(c *gin.Context) {
	var addressService service.AddressService
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&addressService); err == nil {
		res := addressService.Create(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.Logrusobj.Infoln(err)
	}
}

// GetAddress 展示某个收货地址
func GetAddress(c *gin.Context) {
	var addressService service.AddressService
	res := addressService.Show(c.Request.Context(), c.Param("id"))
	c.JSON(200, res)
}

// ListAddress 展示收货地址
func ListAddress(c *gin.Context) {
	var addressService service.AddressService
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&addressService); err == nil {
		res := addressService.List(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.Logrusobj.Infoln(err)
	}
}

// UpdateAddress 修改收货地址
func UpdateAddress(c *gin.Context) {
	var addressService service.AddressService
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&addressService); err == nil {
		res := addressService.Update(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.Logrusobj.Infoln(err)
	}
}

// DeleteAddress 删除收获地址
func DeleteAddress(c *gin.Context) {
	var addressService service.AddressService
	if err := c.ShouldBind(&addressService); err == nil {
		res := addressService.Delete(c.Request.Context(), c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.Logrusobj.Infoln(err)
	}
}