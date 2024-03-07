package model

import (
	uuid "github.com/jackc/pgx/pgtype/ext/satori-uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	ID	uuid.UUID `gorm:"type:uuid;primary_key;"`
}