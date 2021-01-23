package dto

// BookUpdateDTO is used to catch body json from client
type BookUpdateDTO struct {
	ID          uint64 `json:"id" form:"id"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      string `json:"userID,omitempty" form:"userID,omitempty"`
}

// BookCreateDTO is used to catch body json from client
type BookCreateDTO struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      string `json:"userID,omitempty" form:"userID,omitempty"`
}
