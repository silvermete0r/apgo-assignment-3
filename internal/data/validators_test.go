package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestValidateEmail tests the ValidateEmail function
func TestValidateEmail(t *testing.T) {
	validEmail := "test@example.com"
	invalidEmail := "test@com"

	err := ValidateEmail(validEmail)
	assert.NoError(t, err, "Valid email validation failed")

	err = ValidateEmail(invalidEmail)
	assert.Error(t, err, "Invalid email validation passed")
}
