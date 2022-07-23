package middleware

import (
	"github.com/chanxuehong/log"
	"github.com/chanxuehong/log/trace"
	"github.com/gin-gonic/gin"
)

// NewLogger 新建一个 log.Logger 对象, 注入到 request.Context 里
func NewLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceID, _ := trace.FromHeader(ctx.Request.Header)
		if traceID == "" {
			traceID = trace.NewTraceId()
			ctx.Header(trace.TraceIdHeaderKey, traceID)
		}
		ctx.Header(trace.TraceIdHeaderKey, traceID) // sets traceID to response header

		// adds log.Logger to request.Context
		l := log.New(log.WithTraceId(traceID))
		ctx.Request = ctx.Request.WithContext(log.NewContext(ctx.Request.Context(), l))
		ctx.Next()
	}
}
