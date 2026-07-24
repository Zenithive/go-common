package response

import (
	"encoding/json"
	"net/http"

	"github.com/Zenithive/go-common/errorx"
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

func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}

func OK(w http.ResponseWriter, data any) {
	JSON(w, http.StatusOK, data)
}

func Created(w http.ResponseWriter, data any) {
	JSON(w, http.StatusCreated, data)
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func WriteErr(w http.ResponseWriter, err error, reqID string) {
	appErr := errorx.From(err)

	JSON(w, appErr.Status, Err{
		Error: ErrBody{
			Code:   appErr.Code,
			Msg:    appErr.Msg,
			ReqID:  reqID,
			Detail: appErr.Detail,
		},
	})
}
