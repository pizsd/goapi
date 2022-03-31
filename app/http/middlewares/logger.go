package middlewares

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/pizsd/goapi/pkg/logger"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"io/ioutil"
	"time"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		}
		start := time.Now()
		c.Next()
		cost := time.Since(start)
		responseStatus := c.Writer.Status()
		logFields := []zap.Field{
			zap.Int("status", responseStatus),
			zap.String("request", c.Request.Method+" "+c.Request.URL.String()),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time", cost.String()),
		}
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" || c.Request.Method == "PATCH" {
			logFields = append(logFields, zap.String("ReqBody", string(requestBody)))
			logFields = append(logFields, zap.String("RespBody", w.body.String()))
		}
		if responseStatus > 400 && responseStatus <= 499 {
			logger.Warn("HTTP Warning"+cast.ToString(responseStatus), logFields...)
		} else if responseStatus >= 500 && responseStatus <= 599 {
			logger.Error("HTTP Error "+cast.ToString(responseStatus), logFields...)
		} else {
			logger.Debug("HTTP Access", logFields...)
		}
	}
}
