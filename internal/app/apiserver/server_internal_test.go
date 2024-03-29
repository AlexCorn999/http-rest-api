package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AlexCorn999/http-rest-api/internal/app/store/sqlstore/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	s := newServer(teststore.New())
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}
