package entities

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	UserID      string `json:"user_id"`
	BaseEntity
}

// Builder Object for Task
type TaskBuilder struct {
	id          string
	title       string
	description string
	status      string
	userID      string
	BaseEntityBuilder
}

// Constructor for TaskBuilder
func NewTaskBuilder() *TaskBuilder {
	o := new(TaskBuilder)
	return o
}

// Build Method which creates Task
func (b *TaskBuilder) Build() *Task {
	o := new(Task)
	o.BaseEntity = *b.BaseEntityBuilder.Build()
	o.ID = b.id
	o.Title = b.title
	o.Description = b.description
	o.UserID = b.userID
	o.Status = b.status
	return o
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

// Setter method for the field userID of type string in the object TaskBuilder
func (t *TaskBuilder) SetUserID(userID string) {
	t.userID = userID
}

// Setter method for the field id of type string in the object TaskBuilder
func (t *TaskBuilder) SetId(id string) {
	t.id = id
}
