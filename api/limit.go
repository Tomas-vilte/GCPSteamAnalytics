package api

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	rateLimiter *rate.Limiter
	lock        sync.Mutex
	lastReset   time.Time
	resetPeriod = 1 * time.Minute
)

func initRateLimiter() {
	rateLimiter = rate.NewLimiter(rate.Every(resetPeriod), 100)
	lastReset = time.Now()
}

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		lock.Lock()
		defer lock.Unlock()

		// Verifica si ha pasado el tiempo de reinicio
		if time.Since(lastReset) >= resetPeriod {
			initRateLimiter() // Reinicia el límite
		}

		// Intenta adquirir un permiso del límite de velocidad
		if rateLimiter.Allow() {
			c.Next()
		} else {
			waitTime := int(resetPeriod.Seconds())
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message":      "Para loco me vas a matar! espera un toque " + strconv.Itoa(waitTime) + " segundos antes de intentar nuevamente.",
				"wait_seconds": waitTime,
			})
			c.Abort()
		}
	}
}
