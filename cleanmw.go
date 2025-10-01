package cleanmw

import (
	"github.com/gin-gonic/gin"
	"github.com/OpenIMSDK/tools/log"
)

func CleanLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.ZInfo(c, "clean log", "msg", "日志已经清理完成===================================")
		c.Next()
	}
}
