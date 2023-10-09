package request

type CreateTask struct{
	Description string `json:"description" valid:"required~description is required"`
}