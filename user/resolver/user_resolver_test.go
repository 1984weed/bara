package resolver_test

import (
	"bara/user/mocks"
	"bara/user/resolver"
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	mockUserUC := new(mocks.UserUsecase)

	u := resolver.NewUserResolver(mockUserUC)

	t.Run("success", func(t *testing.T) {
		mockUserUC.On("Register", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil).Once()
		err := u.Register(context.Background(), "user-name", "test@test.com", "password")

		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		mockUserUC.On("Register", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(errors.New("Error test")).Once()
		err := u.Register(context.Background(), "user-name", "test@test.com", "password")

		assert.NoError(t, err)
		fmt.Println(err.Error())
	})
}
