package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const reqIDKey = "request_id"

type ctxKey string

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader("X-Request-ID")

		if id == "" {
			id = uuid.NewString()
		}

		c.Set(reqIDKey, id)
		c.Header("X-Request-ID", id)

		ctx := context.WithValue(
			c.Request.Context(),
			ctxKey(reqIDKey),
			id,
		)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func ReqID(ctx context.Context) string {
	id, _ := ctx.Value(ctxKey(reqIDKey)).(string)
	return id
}
