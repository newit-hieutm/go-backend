package middlewares

import (
	"bytes"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/newit-hieutm/go-backend/pkg/loggers"
	"go.uber.org/zap"
)

func UserZapLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger := loggers.InitLogger()

		// Get request body
		var bodyBytes []byte
		if ctx.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(ctx.Request.Body)
			// Restore the body for later use
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		// Get query parameters
		queryParams := ctx.Request.URL.Query()

		// Get headers
		headers := make(map[string]string)
		for k, v := range ctx.Request.Header {
			headers[k] = strings.Join(v, ",")
		}

		// Log request details
		logger.Info("Request Info:",
			zap.Dict("Request",
				zap.String("client_ip", ctx.ClientIP()), zap.String("forwarded_for", ctx.GetHeader("X-Forwarded-For")),
				zap.String("real_ip", ctx.GetHeader("X-Real-IP")),
				zap.String("forwarded_for", ctx.GetHeader("X-Forwarded-For")),
				// Request details
				zap.String("method", ctx.Request.Method),
				zap.String("path", ctx.Request.URL.Path),
				zap.String("user_agent", ctx.Request.UserAgent()),
				// Headers
				zap.Any("headers", headers),
				// Query parameters
				zap.Any("query_params", queryParams),
				// Request body
				zap.String("body", string(bodyBytes))),
		)

		// Process request
		ctx.Next()
	}
}
