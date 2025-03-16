package v1

import (
	"learn_ginmall/pkg/util"
	"learn_ginmall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var UserRegister service.UserService
	if err := c.ShouldBind(&UserRegister); err == nil {
		res := UserRegister.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logrusobj.Infoln("user register api ",err)
	}
}

func UserLogin(c *gin.Context) {
	var UserLogin service.UserService
	if err := c.ShouldBind(&UserLogin); err == nil {
		res := UserLogin.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logrusobj.Infoln("user login api ",err)
	}
}

func UserUpdate(c *gin.Context) {
	var UserUpdate service.UserService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&UserUpdate); err == nil {
		res := UserUpdate.Update(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logrusobj.Infoln("user update api ",err)
	}
}

func UploadAvatar(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	var uploadAvatar service.UserService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&uploadAvatar); err == nil {
		res := uploadAvatar.Post(c.Request.Context(), claims.ID, file, fileSize)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logrusobj.Infoln("user uploadavatar api ",err)
	}
}

func SendEmail (c *gin.Context) {
	var SendEmail service.SendEmailService

		claims, _ := util.ParseToken(c.GetHeader("Authorization"))
		if err := c.ShouldBind(&SendEmail); err == nil {
			res := SendEmail.Send(c.Request.Context(), claims.ID)
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.Logrusobj.Infoln("user sendemail api ",err)
		}
	
	
}

func ValidEmail (c *gin.Context) {
	var ValidEmail service.ValidEmailService
		if err := c.ShouldBind(&ValidEmail); err == nil {
			res := ValidEmail.Valid(c.Request.Context(), c.GetHeader("Authorization"))
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.Logrusobj.Infoln("user validemail api ",err)
		}
}

func ShowMoney (c *gin.Context) {
	var ShowMoney service.ShowMoneyService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
		if err := c.ShouldBind(&ShowMoney); err == nil {
			res := ShowMoney.Show(c.Request.Context(), claims.ID)
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			util.Logrusobj.Infoln("user showmoney api ",err)
		}
}


