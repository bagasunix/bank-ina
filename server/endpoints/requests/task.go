package requests

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation"

	"github.com/bagasunix/bank-ina/pkg/errors"
)

type CreateTask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	UserID      string `json:"user_id"`
}

// Builder Object for CreateTask
type CreateTaskBuilder struct {
	title       string
	description string
	status      string
	userID      string
}

// Constructor for CreateTaskBuilder
func NewCreateTaskBuilder() *CreateTaskBuilder {
	o := new(CreateTaskBuilder)
	return o
}

// Build Method which creates CreateTask
func (b *CreateTaskBuilder) Build() *CreateTask {
	o := new(CreateTask)
	o.Title = b.title
	o.Description = b.description
	o.Status = b.status
	o.UserID = b.userID
	return o
}

// Setter method for the field title of type string in the object CreateTaskBuilder
func (c *CreateTaskBuilder) SetTitle(title string) {
	c.title = title
}

// Setter method for the field description of type string in the object CreateTaskBuilder
func (c *CreateTaskBuilder) SetDescription(description string) {
	c.description = description
}

// Setter method for the field status of type string in the object CreateTaskBuilder
func (c *CreateTaskBuilder) SetStatus(status string) {
	c.status = status
}

// Setter method for the field userID of type string in the object CreateTaskBuilder
func (c *CreateTaskBuilder) SetUserID(userID string) {
	c.userID = userID
}

func (s *CreateTask) ToJSON() []byte {
	j, err := json.Marshal(s)
	errors.HandlerReturnedVoid(err)
	return j
}

func (c CreateTask) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Title, validation.Required.Error("Nama harus diisi")),
		validation.Field(&c.UserID, validation.Required.Error("User ID harus diisi")),
		validation.Field(&c.Status, validation.Required.Error("Status harus diisi")),
	)
}

type UpdateTask struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// Builder Object for UpdateTask
type UpdateTaskBuilder struct {
	id          string
	title       string
	description string
	status      string
}

// Constructor for UpdateTaskBuilder
func NewUpdateTaskBuilder() *UpdateTaskBuilder {
	o := new(UpdateTaskBuilder)
	return o
}

// Build Method which creates UpdateTask
func (b *UpdateTaskBuilder) Build() *UpdateTask {
	o := new(UpdateTask)
	o.ID = b.id
	o.Title = b.title
	o.Description = b.description
	o.Status = b.status
	return o
}

// Setter method for the field title of type string in the object UpdateTaskBuilder
func (u *UpdateTaskBuilder) SetTitle(title string) {
	u.title = title
}

// Setter method for the field description of type string in the object UpdateTaskBuilder
func (u *UpdateTaskBuilder) SetDescription(description string) {
	u.description = description
}

// Setter method for the field status of type string in the object UpdateTaskBuilder
func (u *UpdateTaskBuilder) SetStatus(status string) {
	u.status = status
}

func (s *UpdateTask) ToJSON() []byte {
	j, err := json.Marshal(s)
	errors.HandlerReturnedVoid(err)
	return j
}

func (c UpdateTask) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Title, validation.Required.Error("Nama harus diisi")),
		validation.Field(&c.Status, validation.Required.Error("Status harus diisi")),
	)
}

// Setter method for the field id of type string in the object UpdateTaskBuilder
func (u *UpdateTaskBuilder) SetId(id string) {
	u.id = id
}
