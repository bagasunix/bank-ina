package responses

import (
	"encoding/json"

	"github.com/bagasunix/bank-ina/pkg/errors"
)

type Empty struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (a *Empty) ToJSON() []byte {
	j, err := json.Marshal(a)
	errors.HandlerReturnedVoid(err)
	return j
}

// Builder Object for Empty
type EmptyBuilder struct {
	message string
	code    int
}

// Constructor for EmptyBuilder
func NewEmptyBuilder() *EmptyBuilder {
	o := new(EmptyBuilder)
	return o
}

// Build Method which creates Empty
func (b *EmptyBuilder) Build() *Empty {
	o := new(Empty)
	o.Message = b.message
	o.Code = b.code
	return o
}

// Setter method for the field msg of type string in the object EmptyBuilder
func (e *EmptyBuilder) SetMsg(message string) {
	e.message = message
}

// Setter method for the field code of type int in the object EmptyBuilder
func (e *EmptyBuilder) SetCode(code int) {
	e.code = code
}
