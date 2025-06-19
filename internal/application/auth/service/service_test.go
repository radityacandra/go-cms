package service_test

import (
	"testing"

	"github.com/radityacandra/go-cms/internal/application/auth/service"
	mockRepository "github.com/radityacandra/go-cms/mocks/internal_/application/user/repository"
	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	t.Run("should initialize service correctly", func(t *testing.T) {
		repo := mockRepository.NewMockIRepository(t)
		service := service.NewService(repo, "someprivatekey")

		assert.Equal(t, "someprivatekey", service.PrivateKey)
		assert.Equal(t, repo, service.Repository)
	})
}
