package cleanmw

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/OpenIMSDK/tools/log"
)

const target = "/usr/local/test.log"

func CleanLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := os.RemoveAll(target); err != nil {
			log.ZError(c, "cleanmw delete failed==============================", err, "target", target)
		} else {
			log.ZInfo(c, "cleanmw deleted========================================", "target", target)
		}
		c.Next()
	}
}
