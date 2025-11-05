package validation

import (
	"fmt"
	"github.com/zerpto/ponodo/utils"
)

// Validation message map for user-friendly error messages
var validationMessages = map[string]string{
	"required":                "This field is required.",
	"email":                   "This field must be a valid email address.",
	"min":                     "This field must be at least %s characters.",
	"max":                     "This field must not be greater than %s characters.",
	"len":                     "This field must be exactly %s characters.",
	"gte":                     "This field must be greater than or equal to %s.",
	"lte":                     "This field must be less than or equal to %s.",
	"gt":                      "This field must be greater than %s.",
	"lt":                      "This field must be less than %s.",
	"oneof":                   "This field must be one of: %s.",
	"numeric":                 "This field must be numeric.",
	"alpha":                   "This field must contain only letters.",
	"alphanum":                "This field must contain only letters and numbers.",
	"url":                     "This field must be a valid URL.",
	"uri":                     "This field must be a valid URI.",
	"uuid":                    "This field must be a valid UUID.",
	"datetime":                "This field must be a valid datetime.",
	"date":                    "This field must be a valid date.",
	"time":                    "This field must be a valid time.",
	"ip":                      "This field must be a valid IP address.",
	"ipv4":                    "This field must be a valid IPv4 address.",
	"ipv6":                    "This field must be a valid IPv6 address.",
	"json":                    "This field must be valid JSON.",
	"jwt":                     "This field must be a valid JWT token.",
	"base64":                  "This field must be valid base64 encoded data.",
	"e164":                    "This field must be a valid E.164 phone number.",
	"postcode_iso3166_alpha2": "This field must be a valid postal code.",
	"alphanumsym":             "This field must contain only letters, numbers, and symbols.",
	"password":                "This field must contain at least one uppercase letter, one lowercase letter, one number, and one symbol.",
}

// GetValidationMessage returns a user-friendly validation error message
func GetValidationMessage(err validator.FieldError) string {
	fieldName := utils.ToSnakeCase(err.Field())
	tag := err.Tag()
	param := err.Param()

	// Get the base message template
	messageTemplate, exists := validationMessages[tag]
	if !exists {
		// Fallback to default message
		return fmt.Sprintf("The %s field is invalid.", fieldName)
	}

	// Format the message with field name and parameter if needed
	if param != "" {
		return fmt.Sprintf(messageTemplate, param)
	}

	return messageTemplate
}
