package sqlstore_test

import (
	"testing"

	"github.com/AlexCorn999/http-rest-api/internal/app/model"
	"github.com/AlexCorn999/http-rest-api/internal/app/store"
	"github.com/AlexCorn999/http-rest-api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryCreate(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepositoryFindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	email := "alexcorn999@list.ru"

	s := sqlstore.New(db)
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email

	s.User().Create(u)

	usr, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, usr)
}
