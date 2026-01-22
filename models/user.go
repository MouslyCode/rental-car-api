package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:char(36);primaryKey;not null" json:"id"`
	Name     string    `gorm:"type:varchar(100);not null" json:"name"`
	Email    string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password string    `gorm:"type:varchar(120);not null" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
