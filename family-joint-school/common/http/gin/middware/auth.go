package middleware

import (
	"family-joint-school/common/errors/apierror"
	"family-joint-school/common/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// JwtAuth 认证 token 是否有效, 如果无效则返回错误.
func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenBytes := ctx.Request.Header.Get("Authorization")
		token, err := jwt.TokenParse(tokenBytes)
		if err != nil {
			// Request Basic Authentication otherwise
			ctx.Header("WWW-Authenticate", "Basic realm=Restricted")
			//ctx.JSON(http.StatusOK, apierror.ErrUnauthenticated)
			ctx.AbortWithStatusJSON(http.StatusOK, apierror.ErrUnauthenticated)
			return
		}
		if token.Subject != "student" {
			ctx.AbortWithStatusJSON(http.StatusOK, apierror.ErrPermissionDenied)
			return
		}
		ctx.Set("uid", token.UID)
		ctx.Next()
	}
}

func JwtAuthCMS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenBytes := ctx.Request.Header.Get("Authorization")
		token, err := jwt.TokenParse(tokenBytes)
		if err != nil {
			// Request Basic Authentication otherwise
			ctx.Header("WWW-Authenticate", "Basic realm=Restricted")
			ctx.AbortWithStatusJSON(http.StatusOK, apierror.ErrUnauthenticated)
			return
		}
		if token.Subject != "teacher" {
			ctx.AbortWithStatusJSON(http.StatusOK, apierror.ErrPermissionDenied)
			return
		}
		ctx.Set("uid", token.UID)
		ctx.Next()
	}
}
