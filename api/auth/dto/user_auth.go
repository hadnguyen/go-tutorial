package dto

import (
	"fmt"

	"go-tutorial/api/user/dto"
	"go-tutorial/api/user/model"

	"github.com/go-playground/validator/v10"
)

type UserAuth struct {
	User   *dto.InfoPrivateUser `json:"user" validate:"required"`
	Tokens *UserTokens          `json:"tokens" validate:"required"`
}

func NewUserAuth(user *model.User, tokens *UserTokens) *UserAuth {
	return &UserAuth{
		User:   dto.NewInfoPrivateUser(user),
		Tokens: tokens,
	}
}

func (d *UserAuth) GetValue() *UserAuth {
	return d
}

func (d *UserAuth) ValidateErrors(errs validator.ValidationErrors) ([]string, error) {
	var msgs []string
	for _, err := range errs {
		switch err.Tag() {
		case "required":
			msgs = append(msgs, fmt.Sprintf("%s is required", err.Field()))
		default:
			msgs = append(msgs, fmt.Sprintf("%s is invalid", err.Field()))
		}
	}
	return msgs, nil
}
