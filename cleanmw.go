package cleanmw

import (
	"os"
	"github.com/gin-gonic/gin"
)

func CleanLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = os.RemoveAll("/usr/local/test.log")
		c.Next()
	}
}
