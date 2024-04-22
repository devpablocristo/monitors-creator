package presenter

import (
	"github.com/gin-gonic/gin"
)


type ApiErrorType struct {
	StatusCode int    `json:"status_code"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func ApiError(err error, c *gin.Context) ApiErrorType {
	return ApiErrorType{
		StatusCode: c.Request.Response.StatusCode,
		Code:       c.Request.Response.Status,
		Message:    err.Error(),
	}
}
