package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`
	Massage string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusOK, Response{
		Status:  "success",
		Massage: message,
		Data:    data,
	})
}
func Created(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusCreated, Response{
		Status:  "success",
		Massage: message,
		Data:    data,
	})
}

func Error(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, Response{
		Status:  "success",
		Massage: message,
	})
}
