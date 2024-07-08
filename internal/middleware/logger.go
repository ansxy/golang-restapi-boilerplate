package middleware

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LoggerBody struct {
	gin.ResponseWriter
	Body       bytes.Buffer
	StatusCode int
}

func (l *LoggerBody) Write(b []byte) (int, error) {
	size, err := l.Body.Write(b)
	if err != nil {
		return size, err
	}
	return l.ResponseWriter.Write(b)
}

func (l *LoggerBody) WriteHeader(statusCode int) {
	l.StatusCode = statusCode
	l.ResponseWriter.WriteHeader(statusCode)
}

func NewLogger(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ww := &LoggerBody{ResponseWriter: c.Writer}
		c.Writer = ww

		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		c.Next()

		var requestPayload map[string]interface{}
		json.Unmarshal(bodyBytes, &requestPayload)

		var responsePayload map[string]interface{}
		if err := json.Unmarshal(ww.Body.Bytes(), &responsePayload); err != nil {
			log.WithError(err).Error("Failed to unmarshal response body")
		}

		log.WithFields(logrus.Fields{
			"status_code": ww.StatusCode,
			"request":     requestPayload,
			"response":    responsePayload,
		}).Info("Request log")
	}
}
