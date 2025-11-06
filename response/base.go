package response

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/zerpto/ponodo/utils"
	"github.com/zerpto/ponodo/validation"
)

// MetaPagination represents pagination metadata for API responses.
// It can be extended with fields like page, limit, total, etc.
// to provide pagination information in response metadata.
type MetaPagination struct {
}

// Meta represents response metadata that provides additional information
// about the API response, including timestamp, pagination details, and
// execution duration for performance monitoring.
type Meta struct {
	Timestamp         time.Time       `json:"timestamp"`
	Pagination        *MetaPagination `json:"pagination"`
	ExecutionDuration int64           `json:"execution_duration"`
}

// BaseSuccessResponse represents a standardized success response structure.
// It includes the response data and metadata, providing a consistent format
// for all successful API responses across the application.
type BaseSuccessResponse[T any] struct {
	Data T     `json:"data"`
	Meta *Meta `json:"meta"`
}

// BaseErrorResponse represents a standardized error response structure.
// It includes an error message and detailed error information, providing
// a consistent format for all error responses across the application.
type BaseErrorResponse struct {
	Message string `json:"message"`
	Error   any    `json:"error"`
}

// Success sends a standardized success response with the provided data.
// It includes metadata such as timestamp and execution duration, and
// uses the specified HTTP status code for the response.
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

// Error sends a standardized error response with the provided error.
// It handles validation errors by formatting them into a structured format,
// and uses the specified HTTP status code for the response.
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
