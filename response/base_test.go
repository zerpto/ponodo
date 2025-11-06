package response

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		statusCode int
		data       interface{}
	}{
		{
			name:       "success with 200",
			statusCode: http.StatusOK,
			data:       map[string]string{"message": "success"},
		},
		{
			name:       "success with 201",
			statusCode: http.StatusCreated,
			data:       map[string]int{"id": 1},
		},
		{
			name:       "success with default status code",
			statusCode: 0,
			data:       "test data",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			Success(c, tt.statusCode, tt.data)

			expectedStatus := tt.statusCode
			if expectedStatus == 0 {
				expectedStatus = http.StatusOK
			}

			assert.Equal(t, expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), "data")
			assert.Contains(t, w.Body.String(), "meta")
		})
	}
}

func TestError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		statusCode int
		err        error
	}{
		{
			name:       "error with 400",
			statusCode: http.StatusBadRequest,
			err:        errors.New("bad request"),
		},
		{
			name:       "error with 500",
			statusCode: http.StatusInternalServerError,
			err:        errors.New("internal error"),
		},
		{
			name:       "error with default status code",
			statusCode: 0,
			err:        errors.New("error"),
		},
		{
			name:       "validation error",
			statusCode: http.StatusBadRequest,
			err:        validator.ValidationErrors{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			Error(c, tt.statusCode, tt.err)

			expectedStatus := tt.statusCode
			if expectedStatus == 0 {
				expectedStatus = http.StatusInternalServerError
			}

			assert.Equal(t, expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), "message")
			assert.Contains(t, w.Body.String(), "error")
		})
	}
}

func TestError_ValidationErrors(t *testing.T) {
	gin.SetMode(gin.TestMode)

	validatorInstance := validator.New()
	type TestStruct struct {
		Email string `validate:"required,email"`
		Name  string `validate:"required"`
	}

	testStruct := TestStruct{
		Email: "invalid-email",
		Name:  "",
	}

	err := validatorInstance.Struct(testStruct)
	assert.Error(t, err)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	Error(c, http.StatusBadRequest, err)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "message")
	assert.Contains(t, w.Body.String(), "error")
}

