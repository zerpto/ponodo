package request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// BaseRequest provides a base structure for HTTP request handlers.
// It embeds common fields like context, error handling, and validator
// that are typically needed for request processing and validation.
type BaseRequest struct {
	Ctx       *gin.Context        `form:"-" json:"-"`
	Err       error               `form:"-" json:"-"`
	Validator *validator.Validate `form:"-" json:"-"`
}
