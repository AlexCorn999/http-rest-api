package store_test

import (
	"testing"

	"github.com/AlexCorn999/http-rest-api/internal/app/model"
	"github.com/AlexCorn999/http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryCreate(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "alexcorn999@list.ru",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepositoryFindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "alexcorn999@list.ru"

	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(&model.User{
		Email: email,
	})

	u, err := s.User().FindByEmail("alexcorn999@list.ru")
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
