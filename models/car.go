package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Car struct {
	ID     uuid.UUID `gorm:"type:char(36);primaryKey;not null" json:"id"`
	Brand  string    `gorm:"type:varchar(100);not null" json:"brand"`
	Model  string    `gorm:"type:varchar(100);not null" json:"model"`
	Year   int       `gorm:"not null" json:"year"`
	Price  float64   `gorm:"type:decimal(12,2);not null" json:"price"`
	Status string    `gorm:"default:'available'" json:"status"`
}

func (c *Car) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}
