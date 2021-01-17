package dto

// BookUpdateDTO is used to catch body json from client
type BookUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"userID,omitempty" form:"userID,omitempty" binding:"required"`
}

// BookCreateDTO is used to catch body json from client
type BookCreateDTO struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"userID,omitempty" form:"userID,omitempty" binding:"required"`
}
