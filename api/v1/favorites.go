package v1

import (
	"learn_ginmall/pkg/util"
	"learn_ginmall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建收藏
func CreateFavorite(c *gin.Context) {
	var service service.FavoritesService
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// 收藏夹详情接口
func ShowFavorites(c *gin.Context) {
	var	service  service.FavoritesService
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))

	}
}

func DeleteFavorite(c *gin.Context) {
	var	service  service.FavoritesService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
