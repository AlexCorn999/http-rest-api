package store

import "github.com/AlexCorn999/http-rest-api/internal/app/model"

// UserRepository ...
type UserRepository interface {
	Create(u *model.User) error
	FindByEmail(email string) (*model.User, error)
}
