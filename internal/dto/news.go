package dto

type CreateNewsRequest struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Image       string `json:"image" form:"image"`
	Description string `json:"description" form:"description"`
	IsActive    int    `json:"is_active" form:"is_active" validate:"required"`
}

type UpdateNewsRequest struct {
	Title       *string `json:"title"`
	Image       *string `json:"image"`
	Description *string `json:"description"`
	IsActive    int     `json:"is_active"`
}
