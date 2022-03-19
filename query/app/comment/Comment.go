package comment

import (
	"github.com/satori/uuid"
	"github.com/shaileshhb/microservices/query/app/entity"
	"gorm.io/gorm"
)

type CommentService struct {
	db *gorm.DB
}

func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{
		db: db,
	}
}

// GetPostComment will get all the comments for specified post.
func (service *CommentService) GetPostComment(comments *[]entity.Comment, postID uuid.UUID) error {

	err := service.doesPostExist(postID)
	if err != nil {
		return err
	}

	err = service.db.Debug().Where("`post_id` = ?", postID).Find(comments).Error
	if err != nil {
		return err
	}

	return nil
}

// GetComments will get all comments for specified comments.
func (service *CommentService) GetPostComments(comments *[]entity.Comment, postID uuid.UUID) error {

	err := service.db.Debug().Where("`post_id` = ?", postID).Find(comments).Error
	if err != nil {
		return err
	}

	return nil
}

// GetComments will get all comments for specified comments.
func (service *CommentService) GetComments(comments *[]entity.Comment) error {

	err := service.db.Debug().Find(comments).Error
	if err != nil {
		return err
	}

	return nil
}

func (service *CommentService) doesPostExist(postID uuid.UUID) error {

	return nil
}
