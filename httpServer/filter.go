package httpServer

import (
	"cloud/common/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

//过滤器
func Filter(c *gin.Context) {
	url := c.Request.URL.String()
	if strings.Index(url, "?") != -1 {
		url = strings.Split(url, "?")[0]
	}
	value, ok := HttpRoute[url]
	if !ok {
		log := fmt.Sprintf("Not found this action,ok:%v,url:%s", ok, url)
		logger.Error(log)
		c.String(200, log)
		return
	}
	method := strings.ToUpper(c.Request.Method)
	if method != value.Method {
		logger.Error(fmt.Sprintf("Request method is wrong,method:%s", method))
		return
	}
	value.Action(c) //执行方法
}
