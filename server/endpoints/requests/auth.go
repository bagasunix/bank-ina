package requests

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"github.com/bagasunix/bank-ina/pkg/errors"
)

type Signin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c Signin) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Email, validation.Required.Error("Email harus diisi"), is.Email.Error("Format Email salah")),
		validation.Field(&c.Password, validation.Required.Error("Password harus diisi")),
	)
}

func (s *Signin) ToJSON() []byte {
	j, errs := json.Marshal(s)
	errors.HandlerReturnedVoid(errs)
	return j
}

// Builder Object for Signin
type SigninBuilder struct {
	email    string
	password string
}

// Constructor for SigninBuilder
func NewSigninBuilder() *SigninBuilder {
	o := new(SigninBuilder)
	return o
}

// Build Method which creates Signin
func (b *SigninBuilder) Build() *Signin {
	o := new(Signin)
	o.Email = b.email
	o.Password = b.password
	return o
}

// Setter method for the field email of type string in the object SigninBuilder
func (s *SigninBuilder) SetEmail(email string) {
	s.email = email
}

// Setter method for the field password of type string in the object SigninBuilder
func (s *SigninBuilder) SetPassword(password string) {
	s.password = password
}
