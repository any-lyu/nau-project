package middleware

import (
	"fmt"
	"net/http"

	"github.com/chanxuehong/log"
	"github.com/gin-gonic/gin"
)

// Recover 捕获 panic, 打印相关堆栈信息, 返回相关错误
func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func(ctx *gin.Context) {
			if r := recover(); r != nil {
				if r == http.ErrAbortHandler {
					panic(http.ErrAbortHandler) // net/http 会捕获 http.ErrAbortHandler
				}
				var err error
				switch v := r.(type) {
				case error:
					err = v
				default:
					err = fmt.Errorf("%v", v)
				}
				stack := stackString(callers(4))
				log.FatalContext(ctx.Request.Context(), "panic-recover", "panic", err.Error(), "stack", stack)
			}
		}(ctx)
		ctx.Next()
	}
}
