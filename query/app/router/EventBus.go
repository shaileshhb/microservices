package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shaileshhb/microservices/query/app/comment"
	"github.com/shaileshhb/microservices/query/app/entity"
	"github.com/shaileshhb/microservices/query/app/post"
	"github.com/shaileshhb/microservices/query/app/web"
)

type EventBusController struct {
	postService    *post.PostService
	commentService *comment.CommentService
}

func NewEventBusController(postService *post.PostService, commentService *comment.CommentService) *EventBusController {
	return &EventBusController{
		postService:    postService,
		commentService: commentService,
	}
}

// RegisterRoutes will register route for event-bus.
func (controller *EventBusController) RegisterRoutes(router *mux.Router) {

	// event-bus
	// r.POST("api/v1/event-bus/events/listeners", controller.eventBus)

	router.HandleFunc("/event-bus/events/listeners", controller.ListenEvent).Methods(http.MethodPost)

	fmt.Println(" =============== EventBus Routes Registered =============== ")
}

func (controller *EventBusController) ListenEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" ===================== ListenEvent ===================== ")

	var event entity.Event

	err := web.UnmarshalJSON(r, &event)
	if err != nil {
		web.RespondErrorMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	// err = controller.service.ListenEvent(&event)
	// if err != nil {
	// 	web.RespondErrorMessage(w, http.StatusBadRequest, err.Error())
	// 	return
	// }
	var posts []entity.Post

	if event.Type == "PostCreated" {
		// fetch the added post and its comments
		err = controller.postService.GetPosts(&posts)
		if err != nil {
			web.RespondErrorMessage(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	if event.Type == "CommentCreated" {
		// fetch all comments for specified post
		comment, ok := (event.Data).(entity.Comment)
		fmt.Println(" =============== ok value ->", ok)
		fmt.Println(" ============== comment ->", comment)
		err = controller.postService.GetPost(&posts, comment.PostID)
		if err != nil {
			web.RespondErrorMessage(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	web.RespondJSON(w, http.StatusOK, "event received")
}
