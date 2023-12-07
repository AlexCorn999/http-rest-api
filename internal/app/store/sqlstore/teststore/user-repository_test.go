package teststore_test

import (
	"testing"

	"github.com/AlexCorn999/http-rest-api/internal/app/model"
	"github.com/AlexCorn999/http-rest-api/internal/app/store"
	"github.com/AlexCorn999/http-rest-api/internal/app/store/sqlstore/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryCreate(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepositoryFindByEmail(t *testing.T) {
	s := teststore.New()

	email := "alexcorn999@list.ru"

	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email

	s.User().Create(u)

	usr, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, usr)
}
