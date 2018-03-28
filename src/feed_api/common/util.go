package common

import (
	"crypto/md5"
	"fmt"
	api "git.inke.cn/video/panshi/panshi_api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// X is a convenient alias for a map[string]interface{} map
type X map[string]interface{}

// MD5 获取字符串md5值
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))

	return fmt.Sprintf("%x", h.Sum(nil))
}

// IsXhr 判断是否为Ajax请求
func IsXhr(c *gin.Context) bool {
	x := c.Request.Header.Get("X-Requested-With")

	if strings.ToLower(x) == "xmlhttprequest" {
		return true
	}

	return false
}

// Date 时间戳格式化日期
func Date(timestamp int64, format ...string) string {
	layout := "2017-10-19 15:04:05"

	if len(format) > 0 {
		layout = format[0]
	}

	date := time.Unix(timestamp, 0).Format(layout)

	return date
}

// OK API返回成功
func RespOK(c *gin.Context, data ...interface{}) {
	obj := gin.H{
		"code": api.ErrorCode_SUCCESS,
		"msg":  api.GetMessage(api.ErrorCode_SUCCESS),
	}

	if len(data) > 0 {
		obj["data"] = data[0]
	}

	c.JSON(http.StatusOK, obj)
}

// JSON API返回JSON数据
func RespJSON(c *gin.Context, code api.ErrorCode, data ...interface{}) {
	obj := gin.H{
		"code": code,
		"msg":  api.GetMessage(code),
	}

	if len(data) > 0 {
		obj["data"] = data[0]
	}

	c.JSON(http.StatusOK, obj)
}

//将Service RPC服务返回的String格式为Json
func RespJSONExtra(c *gin.Context, obj ...interface{}) {
	c.JSON(http.StatusOK, obj[0])
}
