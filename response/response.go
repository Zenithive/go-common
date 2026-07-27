package response

import (
	"net/http"

	"github.com/Zenithive/go-common/errorx"
	"github.com/gin-gonic/gin"
)

type Err struct {
	Error ErrBody `json:"error"`
}

type ErrBody struct {
	Code   string `json:"code"`
	Msg    string `json:"message"`
	ReqID  string `json:"request_id,omitempty"`
	Detail any    `json:"details,omitempty"`
}

func JSON(c *gin.Context, status int, data any) {
	c.JSON(status, data)
}

func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

func Created(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, data)
}

func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func WriteErr(c *gin.Context, err error, reqID string) {
	appErr := errorx.From(err)

	c.JSON(appErr.Status, Err{
		Error: ErrBody{
			Code:   appErr.Code,
			Msg:    appErr.Msg,
			ReqID:  reqID,
			Detail: appErr.Detail,
		},
	})
}
