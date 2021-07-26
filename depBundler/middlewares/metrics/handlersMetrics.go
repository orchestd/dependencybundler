package metrics

import (
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/monitoring"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func AverageRequestDurationMetric(metrics monitoring.Metrics) gin.HandlerFunc{
	return func(c *gin.Context) {
		start := time.Now()
		statusMethod:= c.Request.Method
		statusCode := strconv.Itoa(c.Writer.Status())
		path:= c.FullPath()
		c.Next()
		defer func() {
			timer := metrics.WithTags(map[string]string{"path":path,"statusCode" : statusCode , "statusMethod" : statusMethod}).Timer("http_avg_request_duration_seconds" , "Avg duration of all HTTP requests")
			timer.Record(time.Since(start))
		}()
	}
}
