package request

type Task struct {
	Description string `json:"description" valid:"required~description is required"`
}

type TaskWithID struct {
	ID uint
	Description string `json:"description" valid:"required~description is required"`
}
