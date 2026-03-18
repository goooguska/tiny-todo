package task

type UpdateInput struct {
	Title       *string `json:"title" validate:"omitempty,max=255"`
	Description *string `json:"omitempty,description"`
	Completed   *bool   `json:"completed" validate:"omitempty,boolean"`
}
