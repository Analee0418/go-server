package httpHandler

import (
	"log"
	"net/http"
	"strings"

	"com.lueey.shop/config"
	"com.lueey.shop/model"
	"com.lueey.shop/utils"
	"github.com/gin-gonic/gin"
)

// OnHosterSignin 主持人登录
func OnHosterSignin(c *gin.Context) {
	username, ok := c.GetPostForm("username")
	if config.DEBUG {
		log.Printf("[DEBUG] host signin username: %s", username)
	}
	if !ok {
		c.String(http.StatusForbidden, "Invalid request.")
		return
	}
	password, ok := c.GetPostForm("password")
	if config.DEBUG {
		log.Printf("[DEBUG] host signin password: %s", password)
	}

	if 0 != strings.Compare(username, config.HosterParams.Username) ||
		0 != strings.Compare(password, config.HosterParams.Password) {
		c.String(http.StatusForbidden, "Invalid request.")
		return
	}
	session := utils.HTTPSession(c)
	session.Set("UserType", model.HTTPUserCategory_Hoster)
	session.Set("UserID", config.HosterParams.Username)
	session.Save()
	c.String(http.StatusOK, "Login success")
}
