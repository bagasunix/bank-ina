package models

type Task struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	UserID      int
	User        *User  `gorm:"foreignKey:UserID"`
	Title       string `gorm:"size:255"`
	Description string
	Status      string `gorm:"size:50;default:pending"`
	BaseModel
}

// Builder Object for Task
type TaskBuilder struct {
	id          int
	userID      int
	user        *User
	title       string
	description string
	status      string
	BaseModelBuilder
}

// Constructor for TaskBuilder
func NewTaskBuilder() *TaskBuilder {
	o := new(TaskBuilder)
	return o
}

// Build Method which creates Task
func (b *TaskBuilder) Build() *Task {
	o := new(Task)
	o.ID = b.id
	o.UserID = b.userID
	o.User = b.user
	o.Title = b.title
	o.Description = b.description
	o.Status = b.status
	o.BaseModel = *b.BaseModelBuilder.Build()
	return o
}

// Setter method for the field id of type int in the object TaskBuilder
func (t *TaskBuilder) SetId(id int) {
	t.id = id
}

// Setter method for the field userID of type int in the object TaskBuilder
func (t *TaskBuilder) SetUserID(userID int) {
	t.userID = userID
}

// Setter method for the field user of type *User in the object TaskBuilder
func (t *TaskBuilder) SetUser(user *User) {
	t.user = user
}

// Setter method for the field title of type string in the object TaskBuilder
func (t *TaskBuilder) SetTitle(title string) {
	t.title = title
}

// Setter method for the field description of type string in the object TaskBuilder
func (t *TaskBuilder) SetDescription(description string) {
	t.description = description
}

// Setter method for the field status of type string in the object TaskBuilder
func (t *TaskBuilder) SetStatus(status string) {
	t.status = status
}
