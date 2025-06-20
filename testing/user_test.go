package testing

import (
	"context"
	"my-golang-service-pos/domain"
	"my-golang-service-pos/dto"
	"my-golang-service-pos/internal/config"
	"my-golang-service-pos/internal/service"
	"testing"
)

// Mock implementation struct
type mockUserRepository struct {
	existingEmail string
}

func (m *mockUserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	// Mock behavior - return empty user (email not found)
	if email == m.existingEmail {
		return domain.User{
			Email: email,
		}, nil
	}

	return domain.User{}, nil
}

func (m *mockUserRepository) Insert(ctx context.Context, user domain.User) error {
	// Mock behavior - always success
	return nil
}

func (m *mockUserRepository) FindUser(ctx context.Context) (domain.User, error) {
	// Mock behavior
	return domain.User{}, nil
}

func (m *mockUserRepository) FindUserId(ctx context.Context, id string) (domain.User, error) {
	// Mock behavior
	return domain.User{}, nil
}

func TestRegister(t *testing.T) {
	// Create mock repository
	mockRepo := &mockUserRepository{}

	// Create config
	conf := &config.Config{
		Jwt: config.Jwt{
			Key: "my_super_secret_jwt_key", // Fixed typo
			Exp: 60,
		},
	}

	//
	// Create user service with mock repo
	userService := service.NewUser(conf, mockRepo)

	t.Run("Valid registration", func(t *testing.T) { // Fixed typo
		// Test case 1: Valid registration
		ctx := context.Background()

		// Create test user data
		testUser := dto.UserData{
			Fullname: "azmi",
			Email:    "testexample.com",
			Password: "password123",
			// Add other required fields
		}

		// Call register method (adjust method name sesuai dengan service Anda)
		result, err := userService.Register(ctx, testUser)

		// Assertions
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result == "" { // atau sesuai return type dari Register method
			t.Error("Expected valid result, got empty")
		}
	})

	t.Run("Duplicate email registrasiton", func(t *testing.T) {
		ctx := context.Background()

		testUser := dto.UserData{
			Fullname: "azmi",
			Email:    "duplicate@gmail.com",
			Password: "password123",
		}

		mockRepo := &mockUserRepository{
			existingEmail: "duplicate@gmail.com",
		}

		userService := service.NewUser(conf, mockRepo)

		result, err := userService.Register(ctx, testUser)

		if err == nil {
			t.Errorf("Expected error for duplicate email, got nil")
		}

		if result != "" {
			t.Errorf("Expected empty result for duplicate email, got %v", result)
		}

	})

}
