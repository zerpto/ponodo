package request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BaseRequest struct {
	Ctx       *gin.Context        `form:"-" json:"-"`
	Err       error               `form:"-" json:"-"`
	Validator *validator.Validate `form:"-" json:"-"`
}
