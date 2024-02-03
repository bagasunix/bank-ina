package requests

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"github.com/bagasunix/bank-ina/pkg/errors"
)

type CreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Builder Object for CreateUser
type CreateUserBuilder struct {
	name     string
	email    string
	password string
}

// Constructor for CreateUserBuilder
func NewCreateUserBuilder() *CreateUserBuilder {
	o := new(CreateUserBuilder)
	return o
}

// Build Method which creates CreateUser
func (b *CreateUserBuilder) Build() *CreateUser {
	o := new(CreateUser)
	o.Email = b.email
	o.Name = b.name
	o.Password = b.password
	return o
}

// Setter method for the field name of type string in the object CreateUserBuilder
func (c *CreateUserBuilder) SetName(name string) {
	c.name = name
}

// Setter method for the field email of type string in the object CreateUserBuilder
func (c *CreateUserBuilder) SetEmail(email string) {
	c.email = email
}

// Setter method for the field password of type string in the object CreateUserBuilder
func (c *CreateUserBuilder) SetPassword(password string) {
	c.password = password
}

func (s *CreateUser) ToJSON() []byte {
	j, err := json.Marshal(s)
	errors.HandlerReturnedVoid(err)
	return j
}

func (c CreateUser) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required.Error("Nama harus diisi")),
		validation.Field(&c.Email, validation.Required.Error("Email harus diisi"), is.Email.Error("Format email salah")),
		validation.Field(&c.Password, validation.Required.Error("Nama harus diisi")),
	)
}

type UpdateUser struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Builder Object for UpdateUser
type UpdateUserBuilder struct {
	id       string
	name     string
	email    string
	password string
}

// Constructor for UpdateUserBuilder
func NewUpdateUserBuilder() *UpdateUserBuilder {
	o := new(UpdateUserBuilder)
	return o
}

// Build Method which creates UpdateUser
func (b *UpdateUserBuilder) Build() *UpdateUser {
	o := new(UpdateUser)
	o.ID = b.id
	o.Name = b.name
	o.Email = b.email
	o.Password = b.password
	return o
}

// Setter method for the field id of type string in the object UpdateUserBuilder
func (u *UpdateUserBuilder) SetId(id string) {
	u.id = id
}

// Setter method for the field name of type string in the object UpdateUserBuilder
func (u *UpdateUserBuilder) SetName(name string) {
	u.name = name
}

// Setter method for the field password of type string in the object UpdateUserBuilder
func (u *UpdateUserBuilder) SetPassword(password string) {
	u.password = password
}

// Setter method for the field email of type string in the object UpdateUserBuilder
func (u *UpdateUserBuilder) SetEmail(email string) {
	u.email = email
}

func (c UpdateUser) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required.Error("Nama harus diisi")),
		validation.Field(&c.Password, validation.Required.Error("Nama harus diisi")),
	)
}
