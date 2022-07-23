package gin

import (
	"github.com/gin-gonic/gin"
)

// Context 支持 gin.Context.Request.Context().Value
type Context struct {
	*gin.Context
}

// Value Context().Value
func (ctx Context) Value(key interface{}) interface{} {
	value := ctx.Context.Value(key)
	if value == nil {
		return ctx.Request.Context().Value(key)
	}
	return value
}

// WrapContext 包装 gin.Context, 返回一个GinContext.
func WrapContext(ctx *gin.Context) Context {
	return Context{ctx}
}
