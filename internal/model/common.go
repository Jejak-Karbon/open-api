package model

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	CreatedAt time.Time
	CreatedBy uint
	UpdatedAt time.Time
	UpdatedBy uint
	DeletedAt gorm.DeletedAt `gorm:"index"`
	DeletedBy uint
}
