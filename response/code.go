package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ok[T any](ctx *gin.Context, data T) {
	Success(ctx, http.StatusOK, data)
}

func Created[T any](ctx *gin.Context, data T) {
	Success(ctx, http.StatusCreated, data)
}

func NoContent(ctx *gin.Context) {
	ctx.String(http.StatusNoContent, "")
}

func BadRequest(ctx *gin.Context, data error) {
	Error(ctx, http.StatusBadRequest, data)
}

func Unauthorized(ctx *gin.Context, data error) {
	Error(ctx, http.StatusUnauthorized, data)
}

func Forbidden(ctx *gin.Context, data error) {
	Error(ctx, http.StatusForbidden, data)
}

func NotFound(ctx *gin.Context, data error) {
	Error(ctx, http.StatusNotFound, data)
}

func MethodNotAllowed(ctx *gin.Context, data error) {
	Error(ctx, http.StatusMethodNotAllowed, data)
}

func InternalServerError(ctx *gin.Context, data error) {
	Error(ctx, http.StatusInternalServerError, data)
}

func BadGateway(ctx *gin.Context, data error) {
	Error(ctx, http.StatusBadGateway, data)
}

func ServiceUnavailable(ctx *gin.Context, data error) {
	Error(ctx, http.StatusServiceUnavailable, data)
}
