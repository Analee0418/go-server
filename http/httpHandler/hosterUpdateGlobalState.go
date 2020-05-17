package httpHandler

import (
	"log"
	"net/http"
	"runtime/debug"

	"com.lueey.shop/config"
	"com.lueey.shop/model"
	"com.lueey.shop/protocol"
	"com.lueey.shop/utils"
	"github.com/gin-gonic/gin"
)

// OnHosterUpdateState 主持人更新会议进度
func OnHosterUpdateState(c *gin.Context) {
	session := utils.HTTPSession(c)
	log.Printf("%s", session.Get("UserType"))
	userType, converted := session.Get("UserType").(int)
	if !converted {
		log.Printf("ERROR: login first pls.")
		c.String(403, "Invalid request.")
		return
	}

	userNname, converted := session.Get("UserID").(string)
	if !converted {
		log.Printf("ERROR: login first pls.")
		c.String(403, "Invalid request.")
		return
	}

	if userType != model.HTTPUserCategory_Hoster || userNname != config.HosterParams.Username {
		c.String(403, "Hoster cannot upload files.")
		return
	}

	if state, ok := c.GetPostForm("state"); ok {
		log.Printf("##### Will update global state later. #####\n%s", debug.Stack())

		if stateCode, error := protocol.NewGlobalStateValue(state); stateCode == -1 {
			c.String(http.StatusBadRequest, error.Error())
			return
		}

		utils.HTTPGlobalUpdateState(state)
		c.String(http.StatusOK, "OK")
		return
	}
	c.String(http.StatusForbidden, "Invalid parameters.")
}
