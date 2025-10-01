package cleanmw

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CleanLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("日志已经清理完成")
		c.Next()
	}
}
