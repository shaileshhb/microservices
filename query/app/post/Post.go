package post

import (
	"github.com/satori/uuid"
	"github.com/shaileshhb/microservices/query/app/entity"
	"gorm.io/gorm"
)

type PostService struct {
	db *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{
		db: db,
	}
}

// GetPosts will get all posts.
func (service *PostService) GetPosts(posts *[]entity.Post) error {

	err := service.db.Debug().Preload("Comments").Find(posts).Error
	if err != nil {
		return err
	}

	return nil
}

// GetPost will get specified posts.
func (service *PostService) GetPost(post *[]entity.Post, postID uuid.UUID) error {

	err := service.db.Debug().Where("`id` = ?", postID).Preload("Comments").Find(post).Error
	if err != nil {
		return err
	}

	return nil
}
