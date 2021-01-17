package dto

// UserUpdateDTO is used to catch body json from client
type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" validate:"email" binding:"required"`
	Password string `json:"password,omiempty" form:"password,omiempty" validate:"min:6" binding:"required"`
}

// UserCreateDTO is used to catch body json from client
type UserCreateDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" validate:"email" binding:"required"`
	Password string `json:"password,omiempty" form:"password,omiempty" validate:"min:6" binding:"required"`
}
