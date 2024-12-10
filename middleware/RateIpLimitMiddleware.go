package middleware

import (
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

// requestCounts 用于存储每个 IP 的请求计数
var requestCounts = make(map[string]int)

// mu 是一个互斥锁，用于保护 requestCounts 的并发访问
var mu sync.Mutex

// RateIpLimitMiddleware 限制每个 IP 每分钟最多 100 次请求
func RateIpLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		mu.Lock()
		defer mu.Unlock()

		ip := c.ClientIP()  // 获取请求的 IP 地址
		requestCounts[ip]++ // 增加该 IP 的请求计数

		// 每分钟后重置请求计数
		go func() {
			time.Sleep(1 * time.Minute)
			mu.Lock()
			requestCounts[ip] = 0 // 重置该 IP 的请求计数
			mu.Unlock()
		}()

		// 检查请求计数是否超过限制
		if requestCounts[ip] > 100 {
			c.JSON(429, gin.H{"error": "Too Many Requests"}) // 返回 429 错误
			c.Abort()                                        // 中止请求
			return
		}

		c.Next() // 继续处理请求
	}
}
