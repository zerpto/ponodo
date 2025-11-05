package response

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zerpto/ponodo/utils"
	"github.com/zerpto/ponodo/validation"
	"net/http"
	"time"
)

type MetaPagination struct {
}

type Meta struct {
	Timestamp         time.Time       `json:"timestamp"`
	Pagination        *MetaPagination `json:"pagination"`
	ExecutionDuration int64           `json:"execution_duration"`
}

type BaseSuccessResponse[T any] struct {
	Data T     `json:"data"`
	Meta *Meta `json:"meta"`
}

type BaseErrorResponse struct {
	Message string `json:"message"`
	Error   any    `json:"error"`
}

func Success[T any](ctx *gin.Context, statusCode int, data T) {
	if statusCode == 0 {
		statusCode = 200
	}
	timestamp := time.Now()
	startAt := time.Now()
	ctx.JSON(statusCode, &BaseSuccessResponse[T]{
		Data: data,
		Meta: &Meta{
			Timestamp:         timestamp,
			ExecutionDuration: time.Since(startAt).Nanoseconds(),
		},
	})
	ctx.Abort()

}

func Error(ctx *gin.Context, statusCode int, data error) {
	if statusCode == 0 {
		statusCode = 500
	}

	var errorContent any

	var validationErrors validator.ValidationErrors
	if errors.As(data, &validationErrors) {

		// handle validation error
		fieldsWithErrorValue := make(map[string][]string)
		for _, validationError := range validationErrors {
			fieldName := utils.ToSnakeCase(validationError.Field())
			message := validation.GetValidationMessage(validationError)
			fieldsWithErrorValue[fieldName] = append(fieldsWithErrorValue[validationError.Field()], message)
		}
		errorContent = fieldsWithErrorValue

	} else {

		// handle generic error
		errorContent = map[string]any{
			"generic": []string{data.Error()},
		}
	}

	ctx.JSON(statusCode, &BaseErrorResponse{
		Message: http.StatusText(statusCode),
		Error:   errorContent,
	})
	ctx.Abort()
}
