package task

type UpdateInput struct {
	Title       string `json:"title" validate:"required,max=255"`
	Description string `json:"description"`
	Completed   bool   `json:"completed" validate:"boolean"`
}
