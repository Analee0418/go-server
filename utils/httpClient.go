package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

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
	log.Printf("%+v\n", req)                                                         //看下发送的结构

	resp, err := client.Do(req) //发送
	if err != nil {
		log.Printf("ERROR: Prowl notify error, %v", err)
	}
	defer resp.Body.Close() //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(data), err)
}
