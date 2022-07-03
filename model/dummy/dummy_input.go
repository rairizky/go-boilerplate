package dummy

type DummyInput struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}