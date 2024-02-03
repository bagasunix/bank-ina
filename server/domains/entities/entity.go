package entities

import (
	"time"
)

type BaseEntity struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Builder Object for BaseEntity
type BaseEntityBuilder struct {
	createdAt time.Time
	updatedAt time.Time
}

// Constructor for BaseEntityBuilder
func NewBaseEntityBuilder() *BaseEntityBuilder {
	o := new(BaseEntityBuilder)
	return o
}

// Build Method which creates BaseEntity
func (b *BaseEntityBuilder) Build() *BaseEntity {
	o := new(BaseEntity)
	o.CreatedAt = b.createdAt
	o.UpdatedAt = b.updatedAt
	return o
}

// Setter method for the field createdAt of type time.Time in the object BaseEntityBuilder
func (b *BaseEntityBuilder) SetCreatedAt(createdAt time.Time) {
	b.createdAt = createdAt
}

// Setter method for the field updatedAt of type time.Time in the object BaseEntityBuilder
func (b *BaseEntityBuilder) SetUpdatedAt(updatedAt time.Time) {
	b.updatedAt = updatedAt
}
