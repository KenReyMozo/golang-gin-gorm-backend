package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID	uuid.UUID `json:"ID" gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (entry *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	entry.ID = uuid.New()
	return nil
}