package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test User Validation
func TestValidateUser(t *testing.T) {
	user := &User{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "password123",
	}
	err := user.Validate()
	assert.NoError(t, err, "User validation failed")
}

// Test Password Hashing
func TestHashPassword(t *testing.T) {
	password := "password123"
	hashedPassword, err := HashPassword(password)
	assert.NoError(t, err, "Password hashing failed")
	assert.NotEqual(t, password, hashedPassword, "Hashed password should not match plain password")
}

// Test User Authentication
func TestAuthenticateUser(t *testing.T) {
	user := &User{
		Email:    "johndoe@example.com",
		Password: "password123",
	}
	hashedPassword, _ := HashPassword(user.Password)
	user.PasswordHash = hashedPassword

	err := user.Authenticate("password123")
	assert.NoError(t, err, "User authentication failed")

	err = user.Authenticate("wrongpassword")
	assert.Error(t, err, "User authentication should fail with wrong password")
}

// Test Password Matching
func TestPasswordMatches(t *testing.T) {
	password := "password123"
	hashedPassword, _ := HashPassword(password)

	match := PasswordMatches(password, hashedPassword)
	assert.True(t, match, "Passwords should match")

	match = PasswordMatches("wrongpassword", hashedPassword)
	assert.False(t, match, "Passwords should not match")
}
