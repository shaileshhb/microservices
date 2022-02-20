package entity

import (
	"time"

	"github.com/satori/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID  `json:"id" gorm:"type:varchar(36);primarykey"`
	CreatedAt time.Time  `json:"-" gorm:"type:datetime"`
	UpdatedAt time.Time  `json:"-" gorm:"type:datetime"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.NewV4()
	return
}
