package entities

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	BaseEntity
}

// Builder Object for User
type UserBuilder struct {
	id       string
	name     string
	email    string
	password string
	BaseEntityBuilder
}

// Constructor for UserBuilder
func NewUserBuilder() *UserBuilder {
	o := new(UserBuilder)
	return o
}

// Build Method which creates User
func (b *UserBuilder) Build() *User {
	o := new(User)
	o.ID = b.id
	o.Name = b.name
	o.Email = b.email
	o.Password = b.password
	o.BaseEntity = *b.BaseEntityBuilder.Build()
	return o
}

// Setter method for the field id of type string in the object UserBuilder
func (u *UserBuilder) SetID(id string) {
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
