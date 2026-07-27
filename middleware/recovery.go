package middleware

import (
	"log/slog"
	"runtime/debug"

	"github.com/Zenithive/go-common/errorx"
	"github.com/Zenithive/go-common/response"
	"github.com/gin-gonic/gin"
)

func Recovery(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.ErrorContext(
					c.Request.Context(),
					"panic recovered",
					slog.Any("panic", err),
					slog.String("request_id", c.GetString("request_id")),
					slog.String("method", c.Request.Method),
					slog.String("path", c.Request.URL.Path),
					slog.String("stack", string(debug.Stack())),
				)

				response.WriteErr(
					c,
					errorx.Internal(nil),
					c.GetString("request_id"), // omit if WriteErr reads it from context
				)

				c.Abort()
			}
		}()

		c.Next()
	}
}
