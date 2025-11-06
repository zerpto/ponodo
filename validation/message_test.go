package validation

import (
	"reflect"
	"testing"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

// Ensure validator.FieldError is used in the test
var _ validator.FieldError = mockFieldError{}

// mockFieldError is a mock implementation of validator.FieldError
type mockFieldError struct {
	field string
	tag   string
	param string
}

func (m mockFieldError) Tag() string             { return m.tag }
func (m mockFieldError) ActualTag() string       { return m.tag }
func (m mockFieldError) Namespace() string       { return "" }
func (m mockFieldError) StructNamespace() string { return "" }
func (m mockFieldError) Field() string           { return m.field }
func (m mockFieldError) StructField() string     { return "" }
func (m mockFieldError) Value() interface{}      { return nil }
func (m mockFieldError) Param() string           { return m.param }
func (m mockFieldError) Kind() reflect.Kind      { return reflect.String }
func (m mockFieldError) Type() reflect.Type      { return reflect.TypeOf("") }
func (m mockFieldError) Translate(translator ut.Translator) string {
	return ""
}
func (m mockFieldError) Error() string { return "" }

func TestGetValidationMessage(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		tag      string
		param    string
		expected string
	}{
		{
			name:     "required field",
			field:    "Email",
			tag:      "required",
			param:    "",
			expected: "This field is required.",
		},
		{
			name:     "email field",
			field:    "EmailAddress",
			tag:      "email",
			param:    "",
			expected: "This field must be a valid email address.",
		},
		{
			name:     "min length with param",
			field:    "Password",
			tag:      "min",
			param:    "8",
			expected: "This field must be at least 8 characters.",
		},
		{
			name:     "max length with param",
			field:    "UserName",
			tag:      "max",
			param:    "50",
			expected: "This field must not be greater than 50 characters.",
		},
		{
			name:     "unknown tag",
			field:    "TestField",
			tag:      "unknown",
			param:    "",
			expected: "The test_field field is invalid.",
		},
		{
			name:     "numeric field",
			field:    "Age",
			tag:      "numeric",
			param:    "",
			expected: "This field must be numeric.",
		},
		{
			name:     "gte with param",
			field:    "Score",
			tag:      "gte",
			param:    "0",
			expected: "This field must be greater than or equal to 0.",
		},
		{
			name:     "oneof with param",
			field:    "Status",
			tag:      "oneof",
			param:    "active inactive",
			expected: "This field must be one of: active inactive.",
		},
		{
			name:     "camelCase field name conversion",
			field:    "UserName",
			tag:      "required",
			param:    "",
			expected: "This field is required.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := mockFieldError{
				field: tt.field,
				tag:   tt.tag,
				param: tt.param,
			}
			result := GetValidationMessage(err)
			assert.Equal(t, tt.expected, result)
		})
	}
}
