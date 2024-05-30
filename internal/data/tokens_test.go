package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test Token Generation
func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(1, 24)
	assert.NoError(t, err, "Token generation failed")
	assert.NotEmpty(t, token, "Generated token is empty")
}
