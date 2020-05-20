package httpHandler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"path"
	"strings"

	"com.lueey.shop/config"
	"com.lueey.shop/model"
	"com.lueey.shop/utils"
	"github.com/gin-gonic/gin"
)

// OnUploadFile 上传文件处理器
func OnUploadFile(c *gin.Context) {
	session := utils.HTTPSession(c)
	log.Printf("[INFO] %s", session.Get("user"))
	user, converted := session.Get("user").(model.HTTPUser)
	if !converted {
		log.Printf("\033[1;31m[ERROR] \033[0mlogin first pls.")
		c.String(403, "Invalid request.")
		return
	}

	if user.UserType == model.HTTPUserCategory_Hoster {
		c.String(403, "Hoster cannot upload files.")
		return
	}

	session.Set("", "")
	// single file
	file, _ := c.FormFile("file")
	fileName := file.Filename
	log.Printf("[INFO] Upload file, filename: %s", fileName)

	extname := strings.ToLower(path.Ext(fileName))
	switch extname {
	case ".jpg", ".jpeg", ".png", ".gif":
		break
	default:
		c.String(http.StatusForbidden, "Invalid file type.")
		return
	}

	var fileBuffer []byte = []byte{}
	if src, err := file.Open(); err != nil {
		src.Read(fileBuffer)
		src.Seek(io.SeekStart, io.SeekStart)
	}

	fileMD5 := utils.CalcMD5(fileBuffer)

	fileName = fmt.Sprintf("%s?%s=%s",
		fileName, user.UserID, fileMD5)
	// Upload the file to specific dst.
	c.SaveUploadedFile(file, utils.ExpandUser(fmt.Sprintf("~/contract.tmp/%s", fileName)))

	c.String(http.StatusOK, fmt.Sprintf("%s:%s/assets/%s", config.HTTPParams.IP, config.HTTPParams.Port, fileName))
}
