package repo

import (
	"github.com/andresxlp/api-mvc-arch/dto"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user dto.User) error
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) UserRepository {
	return &repo{db}
}

func (r *repo) CreateUser(ctx context.Context, user dto.User) error {
	return r.db.WithContext(ctx).Create(&user).Error
}
