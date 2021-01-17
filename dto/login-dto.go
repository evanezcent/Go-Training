package dto

// LoginDTO is used to catch body json from client
type LoginDTO struct {
	Email    string `json:"email" form:"email" validate:"email" binding:"required"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6" binding:"required"`
}

