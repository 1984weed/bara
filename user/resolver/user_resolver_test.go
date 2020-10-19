package resolver_test

import (
	"bara/auth"
	"bara/graphql_model"
	"bara/model"
	"bara/user/mocks"
	"bara/user/resolver"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetMe(t *testing.T) {
	mockUserUC := new(mocks.UserUsecase)

	u := resolver.NewUserResolver(mockUserUC)

	userID := int64(1)
	t.Run("success", func(t *testing.T) {
		expected := &model.Users{
			ID:       userID,
			UserName: "userName",
		}

		mockUserUC.On("GetUserByID", mock.Anything, userID).Return(expected, nil).Once()

		user, err := u.GetMe(withUserIDContext(userID))

		assert.NotNil(t, user)
		assert.NoError(t, err)
	})

	t.Run("fail because there is an error", func(t *testing.T) {
		mockUserUC.On("GetUserByID", mock.Anything, userID).Return(nil, errors.New("not found")).Once()

		user, err := u.GetMe(withUserIDContext(userID))

		assert.Nil(t, user)
		assert.Error(t, err)
	})
}

func TestGetUser(t *testing.T) {
	mockUserUC := new(mocks.UserUsecase)
	u := resolver.NewUserResolver(mockUserUC)

	userID := int64(1)

	t.Run("success", func(t *testing.T) {
		expected := &model.Users{
			ID:       userID,
			UserName: "userName",
		}
		mockUserUC.On("GetUserByUserName", mock.Anything, "userName").Return(expected, nil).Once()
		user, err := u.GetUser(withUserIDContext(1), "userName")

		assert.Equal(t, &graphql_model.User{ID: "1", DisplayName: "", UserName: "userName", Email: "", Image: "", Role: (*graphql_model.UserRole)(nil), Bio: ""}, user)
		assert.NoError(t, err)
	})

	t.Run("failure because GetUserByUserName causes an error", func(t *testing.T) {
		expected := &model.Users{
			ID:       userID,
			UserName: "userName",
		}
		mockUserUC.On("GetUserByUserName", mock.Anything, "userName").Return(expected, errors.New("Something happened")).Once()
		user, err := u.GetUser(withUserIDContext(1), "userName")

		assert.Nil(t, user)
		assert.Error(t, err)
	})

}

func withUserIDContext(userID int64) context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, auth.UserCtxKey, &auth.CurrentUser{Sub: userID})

	return ctx
}
