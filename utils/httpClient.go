package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const HTTPSessionAge int = 1200

// ProwlNotify 发送提醒
func ProwlNotify(msg string) {
	v := url.Values{}
	v.Set("apikey", "455220800d0b687cd58d8328fde6164721478329")
	v.Set("application", "Turtle-chen ERROR")
	v.Set("event", msg)
	v.Set("priority", "1")
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://api.prowlapp.com/publicapi/add", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时
	log.Printf("[INFO] %+v\n", req)                                                  //看下发送的结构

	resp, err := client.Do(req) //发送
	if err != nil {
		log.Printf("\033[1;31m[ERROR] \033[0mProwl notify error, %v", err)
	}
	defer resp.Body.Close() //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	log.Println("[INFO] ", string(data), err)
}

// HTTPSession HTTP 服务器 session
func HTTPSession(c *gin.Context) sessions.Session {
	s := sessions.Default(c)
	s.Options(sessions.Options{MaxAge: HTTPSessionAge})
	return s
}
