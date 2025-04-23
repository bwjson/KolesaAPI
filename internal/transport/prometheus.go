package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"strconv"
)

var httpRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	},
	[]string{"path", "method", "status"},
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
}

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}
		method := c.Request.Method

		defer func() {
			if r := recover(); r != nil {
				// Можно залогировать: log.Printf("panic recovered: %v", r)
				c.AbortWithStatus(http.StatusInternalServerError)
			}

			status := strconv.Itoa(c.Writer.Status())
			httpRequestsTotal.WithLabelValues(path, method, status).Inc()
		}()

		c.Next()
	}
}
