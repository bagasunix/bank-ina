package models

type User struct {
	ID       int `gorm:"primaryKey;autoIncrement"`
	Name     string
	Email    string
	Password string
	BaseModel
}

// Builder Object for User
type UserBuilder struct {
	id       int
	name     string
	email    string
	password string
	BaseModelBuilder
}

// Constructor for UserBuilder
func NewUserBuilder() *UserBuilder {
	o := new(UserBuilder)
	return o
}

// Build Method which creates User
func (b *UserBuilder) Build() *User {
	o := new(User)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.ID = b.id
	o.Name = b.name
	o.Email = b.email
	o.Password = b.password
	return o
}

// Setter method for the field id of type int in the object UserBuilder
func (u *UserBuilder) SetId(id int) {
	u.id = id
}

// Setter method for the field name of type string in the object UserBuilder
func (u *UserBuilder) SetName(name string) {
	u.name = name
}

// Setter method for the field email of type string in the object UserBuilder
func (u *UserBuilder) SetEmail(email string) {
	u.email = email
}

// Setter method for the field password of type string in the object UserBuilder
func (u *UserBuilder) SetPassword(password string) {
	u.password = password
}
func (User) TableName() string {
	return "users"
}
