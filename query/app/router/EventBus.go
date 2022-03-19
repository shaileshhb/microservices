package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	router.HandleFunc("/event-bus/events/listeners", controller.ListenEvent).Methods(http.MethodPost)

	// calling getEvents to sync data
	controller.getEvents()

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
	err = controller.handleEvents(event, &posts)
	if err != nil {
		web.RespondErrorMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	web.RespondJSON(w, http.StatusOK, "event received")
}

func (controller *EventBusController) getEvents() {
	fmt.Println(" ===================== getEvents ===================== ")

	req, err := http.NewRequest(http.MethodGet, "http://localhost:4005/event-bus/events", nil)
	if err != nil {
		return
	}

	fmt.Println(" ================== req ->", req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// panic(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(" ========= resp ->", resp)
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var events []entity.Event

	err = json.Unmarshal(respBody, &events)
	if err != nil {
		return
	}

	var allPosts []entity.Post

	for _, event := range events {
		fmt.Println(" ================= event ->", event.Type)
		var posts []entity.Post
		err = controller.handleEvents(event, &posts)
		if err != nil {
			return
		}

		allPosts = append(allPosts, posts...)
	}

}

func (controller *EventBusController) handleEvents(event entity.Event, posts *[]entity.Post) error {

	if event.Type == "PostCreated" {
		// fetch the added post and its comments
		err := controller.postService.GetPosts(posts)
		if err != nil {
			return err
		}
	}

	if event.Type == "CommentCreated" {
		// fetch all comments for specified post
		comment, ok := (event.Data).(entity.Comment)
		fmt.Println(" =============== ok value ->", ok)
		fmt.Println(" ============== comment ->", comment)
		err := controller.postService.GetPost(posts, comment.PostID)
		if err != nil {
			return err
		}
	}

	return nil
}
