package gin

import (
	"family-joint-school/common/errors/apierror"
	"family-joint-school/common/errors/errorcode"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandlerFunc 是 http handler 函数模型.
type HandlerFunc func(c Context) (data interface{}, err error)

// HandlerFuncWrapper 返回的统一封装
func HandlerFuncWrapper(fn HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type response struct {
			Code    int         `json:"code"`
			Message string      `json:"msg"`
			Data    interface{} `json:"data"`
		}
		resp := response{
			Code: int(errorcode.Code_OK),
		}
		data, err := fn(WrapContext(ctx))
		if err != nil {
			if e, ok := err.(apierror.Error); ok {
				resp.Code = e.ErrorCode()
				resp.Message = e.ErrorMessage()
			} else {
				resp.Code = int(errorcode.Code_INTERNAL)
				resp.Message = err.Error()
			}
		}
		resp.Data = data
		ctx.Header("X-Corporation", "nau-693") // 标识是内部服务, api call 会根据这个标识来决定是否打印访问日志
		ctx.JSON(http.StatusOK, &resp)
	}
}

// MiddlewareFunc defines a function to process middleware.
type MiddlewareFunc func(HandlerFunc) HandlerFunc
