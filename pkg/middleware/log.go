package middleware

import (
	"bytes"
	"io"
	"ji/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type responseWriter struct {
	gin.ResponseWriter
	b *bytes.Buffer
}

//重写 Write([]byte) (int, error) 方法
func (w responseWriter) Write(b []byte) (int, error) {
	//向一个bytes.buffer中写一份数据来为获取body使用
	w.b.Write(b)
	//完成gin.Context.Writer.Write()原有功能
	return w.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	log := logger.Logrus

	return func(c *gin.Context) {

		start := time.Now()
		requestTime := start.Format(time.RFC3339)
		requestMethod := c.Request.Method
		requestHeader := c.Request.Header
		path := c.Request.URL
		requestBody := ""
		b, err := c.GetRawData()
		if err != nil {
			requestBody = "failed to get request body"
		} else {
			requestBody = string(b)
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(b))
		host := c.Request.Host
		schema := c.Request.Proto
		client := c.ClientIP()

		writer := responseWriter{
			c.Writer,
			bytes.NewBuffer([]byte{}),
		}
		c.Writer = writer

		preEntry := log.WithFields(logrus.Fields{
			"requestTime":   requestTime,
			"client":        client,
			"requestMethod": requestMethod,
			"requestHeader": requestHeader,
			"requestBody":   requestBody,
			"host":          host,
			"schema":        schema,
			"path":          path,
		})
		if gin.Mode() == "debug" {
			preEntry.Debug()
		}else  {
			preEntry.Info()
		}
		
		c.Next()
		
		cost := time.Since(start).Milliseconds()
		responseStatus := c.Writer.Status()
		responseHeader := c.Writer.Header()
		responseBodySize := c.Writer.Size()
		responseBody := writer.b.String()

		postEntry := log.WithFields(logrus.Fields{
			"cost":             cost,
			"responseStatus":   responseStatus,
			"responseHeader":   responseHeader,
			"responseBodySize": responseBodySize,
			"responseBody":     responseBody,
		})
		postEntry.Info()

		if responseStatus >= 500 {
			postEntry.Error()
		} else if responseStatus >= 400 {
			postEntry.Warn()
		} else {
			postEntry.Info()
		}
	}
}
