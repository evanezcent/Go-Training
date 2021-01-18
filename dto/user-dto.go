package dto

// UserUpdateDTO is used to catch body json from client
type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" validate:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6"`
}

// UserCreateDTO is used to catch body json from client
type UserCreateDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" validate:"email" binding:"required,email"`
	Password string `json:"password" form:"password" validate:"min:6" binding:"required"`
}

// LoginDTO is used to catch body json from client
type LoginDTO struct {
	Email    string `json:"email" form:"email" validate:"email" binding:"required,email"`
	Password string `json:"password" form:"password" validate:"min:6" binding:"required"`
}
