package validation

import (
	"github.com/Zenithive/go-common/errorx"
	"github.com/Zenithive/go-common/response"
	"github.com/gin-gonic/gin"
)

func Bind(c *gin.Context, v any) bool {
	if err := c.ShouldBindJSON(v); err != nil {
		detail := Parse(err)

		response.WriteErr(
			c,
			errorx.Validation(detail),
			c.GetString("request_id"), // omit if WriteErr fetches it itself
		)

		c.Abort()
		return false
	}

	return true
}

func ReqID(c *gin.Context) string {
	return c.GetString("request_id")
}
