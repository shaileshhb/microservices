package comment

import (
	"github.com/satori/uuid"
	"github.com/shaileshhb/microservices/comment/app/entity"
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

// EventBus wil listen to event-bus
func (service *CommentService) EventBus(comment *entity.Event) error {
	return nil
}

// AddComment will add comment for specified post.
func (service *CommentService) AddComment(comment *entity.Comment) error {

	err := service.doesPostExist(comment.PostID)
	if err != nil {
		return err
	}

	// uow := database.NewUnitOfWork(service.db, false)

	err = service.db.Debug().Create(comment).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateComment will update specified comment for specified post.
func (service *CommentService) UpdateComment(comment *entity.Comment) error {

	err := service.doesPostExist(comment.PostID)
	if err != nil {
		return err
	}

	err = service.doesCommentExist(comment.ID)
	if err != nil {
		return err
	}

	// uow := database.NewUnitOfWork(service.db, false)

	var tempComment entity.Comment
	tempComment.ID = comment.ID

	err = service.db.Debug().Where("`id` = ?", comment.ID).First(&tempComment).Error
	if err != nil {
		return err
	}

	comment.CreatedAt = tempComment.CreatedAt

	err = service.db.Debug().Model(&entity.Comment{}).
		Where("`id` = ?", comment.ID).Save(comment).Error
	if err != nil {
		return err
	}

	return nil
}

// DeleteComment will delete specified comment.
func (service *CommentService) DeleteComment(comment *entity.Comment) error {

	err := service.doesCommentExist(comment.ID)
	if err != nil {
		return err
	}

	// uow := database.NewUnitOfWork(service.db, false)

	err = service.db.Debug().Delete(comment).Error
	if err != nil {
		return err
	}

	return nil
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

func (service *CommentService) doesCommentExist(commentID uuid.UUID) error {

	return nil
}
