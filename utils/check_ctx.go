package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckContext[T any](ctx *gin.Context, key string) (T, error) {
	var zero T
	value, exist := ctx.Get(key)

	if !exist {
		return zero, errors.New(key + "tidak ditemukan di ctx")
	}

	casted, ok := value.(T)

	if !ok {
		return zero, errors.New("type assertion gagal untuk" + key)
	}

	return casted, nil
}
