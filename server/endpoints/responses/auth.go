package responses

import (
	"encoding/json"

	"github.com/bagasunix/bank-ina/pkg/errors"
)

type SignIn struct {
	Token string `json:"token"`
}

// Builder Object for SignIn
type SignInBuilder struct {
	token string
}

// Constructor for SignInBuilder
func NewSignInBuilder() *SignInBuilder {
	o := new(SignInBuilder)
	return o
}

// Build Method which creates SignIn
func (b *SignInBuilder) Build() *SignIn {
	o := new(SignIn)
	o.Token = b.token
	return o
}

// Setter method for the field token of type string in the object SignInBuilder
func (s *SignInBuilder) SetToken(token string) {
	s.token = token
}

func (s *SignIn) ToJSON() []byte {
	j, errs := json.Marshal(s)
	errors.HandlerReturnedVoid(errs)
	return j
}
