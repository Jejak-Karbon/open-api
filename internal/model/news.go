package model

import (
	"time"

	"gorm.io/gorm"
)

func (News) TableName() string {
	return "news"
}

type News struct {
	ID          uint   `gorm:"primarykey;autoIncrement"`
	Title       string `json:"title" gorm:"size:200;not null"`
	Description string `json:"description" gorm:"size:1500;not null"`
	Image       string `json:"image" gorm:"size:200;not null"`
	IsActive    int    `json:"is_active" gorm:"size:1;not null"`
	Model
}

func (c *News) BeforeCreate(tx *gorm.DB) (err error) {
	c.CreatedAt = time.Now()
	return
}

func (c *News) BeforeUpdate(tx *gorm.DB) (err error) {
	c.UpdatedAt = time.Now()
	return
}
