package ginx

import (
	"github.com/Zenithive/go-common/errorx"
	"github.com/Zenithive/go-common/response"
	"github.com/Zenithive/go-common/validation"
	"github.com/gin-gonic/gin"
)

func Bind(c *gin.Context, v any) bool {
	if err := c.ShouldBindJSON(v); err != nil {
		detail := validation.Parse(err)

		response.WriteErr(
			c.Writer,
			errorx.Validation(detail),
			c.GetString("request_id"),
		)

		c.Abort()

		return false
	}

	return true
}

func ReqID(c *gin.Context) string {
	return c.GetString("request_id")
}
