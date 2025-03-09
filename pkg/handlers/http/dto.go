package httpHandler

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/microcosm-cc/bluemonday"
)

type (
	CreateUserDto struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)

func (c *CreateUserDto) Sanitize() {
	policy := bluemonday.UGCPolicy()
	c.Name = policy.Sanitize(c.Name)
}

func (c CreateUserDto) Validate() error {
	// Validate the struct fields
	if err := validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required, validation.Length(5, 55)),
		validation.Field(&c.Email, validation.Required, is.Email),
	); err != nil {
		return err
	}

	return nil
}
