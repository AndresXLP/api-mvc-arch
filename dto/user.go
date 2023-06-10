package dto

import "github.com/go-playground/validator/v10"

var validate = validator.New()

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

func (u *User) Validate() error {
	return validate.Struct(u)
}
