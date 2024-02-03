package models

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time `gorm:"autoCreateTime;index;sort:desc;<-:create"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;<-:update"`
}

// Builder Object for BaseModel
type BaseModelBuilder struct {
	createdAt time.Time
	updatedAt time.Time
}

// Constructor for BaseModelBuilder
func NewBaseModelBuilder() *BaseModelBuilder {
	o := new(BaseModelBuilder)
	return o
}

// Build Method which creates BaseModel
func (b *BaseModelBuilder) Build() *BaseModel {
	o := new(BaseModel)
	o.CreatedAt = b.createdAt
	o.UpdatedAt = b.updatedAt
	return o
}

// Setter method for the field createdAt of type time.Time in the object BaseModelBuilder
func (b *BaseModelBuilder) SetCreatedAt(createdAt time.Time) {
	b.createdAt = createdAt
}

// Setter method for the field updatedAt of type time.Time in the object BaseModelBuilder
func (b *BaseModelBuilder) SetUpdatedAt(updatedAt time.Time) {
	b.updatedAt = updatedAt
}
