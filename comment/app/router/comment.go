package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satori/uuid"
	"github.com/shaileshhb/microservices/comment/app/comment"
	"github.com/shaileshhb/microservices/comment/app/entity"
	"github.com/shaileshhb/microservices/comment/app/web"
)

type CommentController struct {
	service *comment.CommentService
}

func NewCommentController(service *comment.CommentService) *CommentController {
	return &CommentController{
		service: service,
	}
}

// RegisterRoutes will register route for comments.
func (controller *CommentController) RegisterRoutes(r *gin.Engine) {

	r.POST("api/v1/post/:postID/comments", controller.AddComment)
	r.PUT("api/v1/post/:postID/comments/:commentID", controller.UpdateComment)
	r.PUT("api/v1/post/comments/:commentID", controller.UpdateComment)
	r.GET("api/v1/post/:postID/comments", controller.GetComments)
	r.GET("api/v1/post/comments", controller.GetComments)
}

// AddComment will add new comment for specified post.
func (controller *CommentController) AddComment(c *gin.Context) {
	fmt.Println(" ===================== Add Comment ===================== ")

	var comment entity.Comment

	err := web.UnmarshalJSON(c, &comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(" === comment ->", comment)

	comment.PostID, err = uuid.FromString(c.Param("postID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = controller.service.AddComment(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, nil)
}

// UpdateComment will update specified comment for specified post.
func (controller *CommentController) UpdateComment(c *gin.Context) {
	fmt.Println(" ===================== Update Comment ===================== ")

	var comment entity.Comment

	err := web.UnmarshalJSON(c, &comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	comment.PostID, err = uuid.FromString(c.Param("postID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	comment.ID, err = uuid.FromString(c.Param("commentID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = controller.service.UpdateComment(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, nil)
}

// DeleteComment will delete specified comment.
func (controller *CommentController) DeleteComment(c *gin.Context) {
	fmt.Println(" ===================== Delete Comment ===================== ")

	var comment entity.Comment
	var err error

	comment.ID, err = uuid.FromString(c.Param("commentID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = controller.service.DeleteComment(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, nil)
}

// GetComments will get all comments.
func (controller *CommentController) GetPostComments(c *gin.Context) {
	fmt.Println(" ===================== Get Post Comments ===================== ")

	var comments []entity.Comment

	postID, err := uuid.FromString(c.Param("commentID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = controller.service.GetPostComments(&comments, postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, comments)
}

// GetComments will get all comments.
func (controller *CommentController) GetComments(c *gin.Context) {
	fmt.Println(" ===================== Get Comments ===================== ")

	var comments []entity.Comment

	err := controller.service.GetComments(&comments)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, comments)
}
