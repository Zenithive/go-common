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
					slog.String(
						"request_id",
						ReqID(c.Request.Context()),
					),
					slog.String("method", c.Request.Method),
					slog.String("path", c.Request.URL.Path),
					slog.String("stack", string(debug.Stack())),
				)

				response.WriteErr(
					c.Writer,
					errorx.Internal(nil),
					ReqID(c.Request.Context()),
				)

				c.Abort()
			}
		}()

		c.Next()
	}
}
