package v1

import (
	"learn_ginmall/pkg/util"
	"learn_ginmall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//创建商品
func CreateProduct(c *gin.Context) {
	var CreateProductService  service.ProductService
	form, _ := c.MultipartForm()
	files := form.File["file"]
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&CreateProductService); err == nil {
		res := CreateProductService.Create(c.Request.Context(), claim.ID, files)
		c.JSON(http.StatusOK, res)
	}else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logrusobj.Infoln("user createproduct api ",err)
	}
}