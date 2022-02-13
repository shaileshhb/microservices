package entity

import "github.com/satori/uuid"

type Comment struct {
	Base
	PostID  uuid.UUID `json:"postID" gorm:"type:varchar(36)"`
	Message string    `json:"message" gorm:"type:varchar(1000)"`
}
