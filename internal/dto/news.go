package dto

type CreateNewsRequest struct {
	Title       string `json:"title" validate:"required"`
	Image       string `json:"image" validate:"required"`
	Description string `json:"description"`
	IsActive    int    `json:"is_active" validate:"required"`
}

type UpdateNewsRequest struct {
	Title       *string `json:"title"`
	Image       *string `json:"image"`
	Description *string `json:"description"`
	IsActive    int     `json:"is_active"`
}
