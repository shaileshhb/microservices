package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

	r.POST("/:postID/comments", controller.AddComment)
	r.GET("/comments", controller.GetComments)
}

func (controller *CommentController) AddComment(c *gin.Context) {
	fmt.Println(" ===================== Add Comment ===================== ")

	var comment entity.Comment

	err := web.UnmarshalJSON(c, &comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = controller.service.AddComment(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, comment)
}

// GetComments will get all comments.
func (controller *CommentController) GetComments(c *gin.Context) {
	fmt.Println(" ===================== Get Comments ===================== ")

	var comments []entity.Comment

	err := controller.service.GetComments(&comments)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, comments)
}
