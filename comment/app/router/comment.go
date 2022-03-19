package router

import (
	"bytes"
	"encoding/json"
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

	// event-bus
	r.POST("api/v1/event-bus/events/listeners", controller.eventBus)

	r.POST("api/v1/post/:postID/comments", controller.addComment)
	r.PUT("api/v1/post/:postID/comments/:commentID", controller.updateComment)
	r.PUT("api/v1/post/comments/:commentID", controller.updateComment)
	r.GET("api/v1/post/:postID/comments", controller.getPostComments)
	r.GET("api/v1/post/comments", controller.getComments)
}

// eventBus wil listen to event-bus
func (controller *CommentController) eventBus(c *gin.Context) {
	fmt.Println(" ===================== Event Bus called ===================== ")

	var event entity.Event

	err := web.UnmarshalJSON(c, &event)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(" === event received for comment ->")

	err = controller.service.EventBus(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, nil)
}

// addComment will add new comment for specified post.
func (controller *CommentController) addComment(c *gin.Context) {
	fmt.Println(" ===================== Add Comment ===================== ")

	var comment entity.Comment

	err := web.UnmarshalJSON(c, &comment)
	if err != nil {
		fmt.Println(" ========== err in unmarshal ->", err)
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

	// emit an event
	event := entity.Event{
		Type: "CommentCreated",
		Data: comment,
	}

	body, err := json.Marshal(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:4005/event-bus/events",
		bytes.NewBuffer(body))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(" ================== req ->", req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	defer resp.Body.Close()

	c.JSON(http.StatusAccepted, nil)
}

// updateComment will update specified comment for specified post.
func (controller *CommentController) updateComment(c *gin.Context) {
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
func (controller *CommentController) getPostComments(c *gin.Context) {
	fmt.Println(" ===================== Get Post Comments ===================== ")

	var comments []entity.Comment

	postID, err := uuid.FromString(c.Param("postID"))
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

// getComments will get all comments.
func (controller *CommentController) getComments(c *gin.Context) {
	fmt.Println(" ===================== Get Comments ===================== ")

	var comments []entity.Comment

	err := controller.service.GetComments(&comments)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, comments)
}
