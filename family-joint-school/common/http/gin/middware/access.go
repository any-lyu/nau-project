package middleware

import (
	"bytes"
	"errors"
	"family-joint-school/common/errors/apierror"
	"family-joint-school/common/errors/errorcode"
	"family-joint-school/common/pool"
	"github.com/chanxuehong/log"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// ErrRequestBodyTooLarge 是 MaxBytesReader 返回的 io.ReadCloser 读取超过限制时返回的错误.
var ErrRequestBodyTooLarge = errors.New("http-request_body_too_large")

// AccessHandler  log + cross
func AccessHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Info("header", "header", log.JSON(ctx.Request.Header))
		origin := ctx.GetHeader("Origin")
		if origin != "" {
			ctx.Header("Access-Control-Allow-Origin", ctx.GetHeader("Origin"))
		} else {
			ctx.Header("Access-Control-Allow-Origin", "*")
		}
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		buf := pool.BytesBufferPool16KB.GetBytesBufferFromPool()
		buf.Reset()
		defer pool.BytesBufferPool16KB.PutBytesBufferToPool(buf)

		blw := &bodyLogWriter{body: buf, ResponseWriter: ctx.Writer}
		ctx.Writer = blw
		request := ctx.Request

		l, ok := log.FromRequest(request)
		if !ok {
			l = log.New()
		}
		var body = ""
		if request.Method == http.MethodPost ||
			request.Method == http.MethodPut {
			bodyBytes, err := io.ReadAll(request.Body)
			request.Body.Close()
			switch err {
			case nil:
				request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
				body = string(bodyBytes)
			case ErrRequestBodyTooLarge:
				l.Error("failed to read body", "error", err.Error())
				ctx.JSON(http.StatusRequestEntityTooLarge, apierror.New(errorcode.Code_REQUEST_ENTITY_TOO_LARGE, err.Error()))
				return
			default:
				l.Error("failed to read body", "error", err.Error())
				ctx.JSON(http.StatusInternalServerError, apierror.New(errorcode.Code_INTERNAL, err.Error()))
				return
			}
		}

		t := time.Now()
		ctx.Next()

		since := time.Since(t).Milliseconds()

		var method = ctx.Request.Method
		var statusCode = blw.Status()
		log.InfoContext(ctx.Request.Context(), "access-log",
			"method", method,
			"path", request.RequestURI,
			"since", since,
			"request", body,
			"statusCode", statusCode,
			"reply", string(blw.body.Bytes()))
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
