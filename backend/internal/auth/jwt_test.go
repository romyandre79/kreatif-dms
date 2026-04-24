package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestJWT(t *testing.T) {
	secret := "test-secret"
	payload := TokenPayload{
		UserID:       uuid.New(),
		Role:         "admin",
		DepartmentID: uuid.New(),
	}

	t.Run("Generate and Validate Valid Token", func(t *testing.T) {
		token, err := GenerateToken(payload, secret, 15*time.Minute)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		claims, err := ValidateToken(token, secret)
		assert.NoError(t, err)
		assert.Equal(t, payload.UserID, claims.UserID)
		assert.Equal(t, payload.Role, claims.Role)
		assert.Equal(t, payload.DepartmentID, claims.DepartmentID)
	})

	t.Run("Invalid Secret", func(t *testing.T) {
		token, _ := GenerateToken(payload, secret, 15*time.Minute)
		_, err := ValidateToken(token, "wrong-secret")
		assert.Error(t, err)
	})

	t.Run("Expired Token", func(t *testing.T) {
		token, _ := GenerateToken(payload, secret, -1*time.Minute)
		_, err := ValidateToken(token, secret)
		assert.Error(t, err)
	})
}
