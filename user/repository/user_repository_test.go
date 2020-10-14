package repository_test

import (
	"bara/model"
	"bara/repository_suite"
	"bara/user/repository"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type userRepositoryTest struct {
	repository_suite.RepositoryTestSuite
}

func TestUserSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip repository test")
	}

	userSuite := &userRepositoryTest{
		repository_suite.RepositoryTestSuite{},
	}

	suite.Run(t, userSuite)
}

func (u *userRepositoryTest) SetupTest() {
	db := u.DB

	users := mockUserData()

	for _, user := range users {
		err := db.Insert(&user)
		require.NoError(u.T(), err)
	}
}

func mockUserData() []model.Users {
	return []model.Users{
		{
			UserName:    "user-1",
			DisplayName: "James Smith",
			Password:    "user-1-password",
			Email:       "user-1-email@testtest.com",
			Bio:         "user-1-bio",
			ImageURL:    "user-1-image",
			UpdatedAt:   time.Now().UTC(),
			CreatedAt:   time.Now().UTC(),
		},
		{
			UserName:    "user-2",
			DisplayName: "Maria Garcia",
			Password:    "user-2-password",
			Email:       "user-2-email@testtest.com",
			Bio:         "user-2-bio",
			ImageURL:    "user-2-image",
			UpdatedAt:   time.Now().UTC(),
			CreatedAt:   time.Now().UTC(),
		},
	}
}

func (u *userRepositoryTest) TearDownTest() {
	u.RepositoryTestSuite.ClearDatabase()
}

// TestRegister...
func (u *userRepositoryTest) TestRegister() {
	runner := repository.NewUserRepositoryRunner(u.DB)
	repo := runner.GetRepository()

	user := &model.Users{
		UserName:  "super man",
		Password:  "password",
		Email:     "testetst@testtest.com",
		UpdatedAt: time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
	}

	me, err := repo.Register(context.Background(), user)

	require.NoError(u.T(), err)
	assert.Equal(u.T(), me, &model.Users{
		ID:        me.ID,
		UserName:  user.UserName,
		Password:  user.Password,
		Email:     user.Email,
		Bio:       "",
		ImageURL:  "",
		UpdatedAt: user.UpdatedAt,
		CreatedAt: user.CreatedAt,
	})
}

// TestGetUserByEmail...
func (a *userRepositoryTest) TestGetUserByEmail() {
	runner := repository.NewUserRepositoryRunner(a.DB)
	repo := runner.GetRepository()

	targetUser := mockUserData()[1]

	a.T().Run("success", func(t *testing.T) {
		res, err := repo.GetUserByEmail(context.Background(), targetUser.Email)

		require.NoError(a.T(), err)
		assert.Equal(a.T(), &model.Users{
			ID:          res.ID,
			UserName:    targetUser.UserName,
			DisplayName: targetUser.DisplayName,
			Password:    targetUser.Password,
			Email:       targetUser.Email,
			Bio:         targetUser.Bio,
			ImageURL:    targetUser.ImageURL,
			UpdatedAt:   res.UpdatedAt,
			CreatedAt:   res.CreatedAt,
		}, res)
	})

	a.T().Run("fail", func(t *testing.T) {
		res, err := repo.GetUserByEmail(context.Background(), "invalidemail")

		assert.Empty(t, res)
		require.Error(a.T(), err)
	})
}

// TestGetUserByEmail...
func (a *userRepositoryTest) TestGetUserByUserName() {
	runner := repository.NewUserRepositoryRunner(a.DB)
	repo := runner.GetRepository()

	targetUser := mockUserData()[0]

	a.T().Run("success", func(t *testing.T) {
		res, err := repo.GetUserByUserName(context.Background(), targetUser.UserName)

		require.NoError(a.T(), err)
		assert.Equal(a.T(), &model.Users{
			ID:          res.ID,
			UserName:    targetUser.UserName,
			DisplayName: targetUser.DisplayName,
			Password:    targetUser.Password,
			Email:       targetUser.Email,
			Bio:         targetUser.Bio,
			ImageURL:    targetUser.ImageURL,
			UpdatedAt:   res.UpdatedAt,
			CreatedAt:   res.CreatedAt,
		}, res)
	})

	a.T().Run("fail", func(t *testing.T) {
		res, err := repo.GetUserByEmail(context.Background(), "invalid-username")

		assert.Empty(t, res)
		require.Error(a.T(), err)
	})
}
