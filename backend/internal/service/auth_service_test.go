package service

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kreatif/dms-backend/internal/config"
	"github.com/kreatif/dms-backend/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// Mock repository
type MockRepo struct {
	mock.Mock
	repository.Querier
}

func (m *MockRepo) GetUserByEmail(ctx context.Context, email string) (repository.GetUserByEmailRow, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(repository.GetUserByEmailRow), args.Error(1)
}

func TestAuthService_Login(t *testing.T) {
	mockRepo := new(MockRepo)
	cfg := config.Config{
		JWTSecret:      "test-secret",
		AccessTokenTTL: 15 * time.Minute,
		RefreshTokenTTL: 1 * time.Hour,
	}
	svc := NewAuthService(mockRepo, cfg)

	email := "test@example.com"
	password := "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	
	mockUser := repository.GetUserByEmailRow{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: string(hashedPassword),
		FullName:     "Test User",
		RoleName:     "user",
	}

	t.Run("Successful Login", func(t *testing.T) {
		mockRepo.On("GetUserByEmail", mock.Anything, email).Return(mockUser, nil).Once()

		res, err := svc.Login(context.Background(), email, password)
		
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, mockUser.ID, res.UserID)
		assert.Equal(t, mockUser.RoleName, res.Role)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid Password", func(t *testing.T) {
		mockRepo.On("GetUserByEmail", mock.Anything, email).Return(mockUser, nil).Once()

		res, err := svc.Login(context.Background(), email, "wrong-password")
		
		assert.Error(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "invalid email or password", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
