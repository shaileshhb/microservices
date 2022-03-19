package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/satori/uuid"
	"github.com/shaileshhb/microservices/query/app/entity"
	"github.com/shaileshhb/microservices/query/app/post"
	"github.com/shaileshhb/microservices/query/app/web"
)

type PostController struct {
	service *post.PostService
}

func NewPostController(service *post.PostService) *PostController {
	return &PostController{
		service: service,
	}
}

// RegisterRoutes will register route for posts.
func (controller *PostController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/posts", controller.getPosts).Methods(http.MethodGet)
	router.HandleFunc("/posts/{postID}", controller.getPost).Methods(http.MethodGet)

	fmt.Println(" =============== Post Routes Registered =============== ")
}

func (controller *PostController) getPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" ===================== Get Posts ===================== ")

	var posts []entity.Post

	err := controller.service.GetPosts(&posts)
	if err != nil {
		web.RespondErrorMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	web.RespondJSON(w, http.StatusOK, posts)
}

func (controller *PostController) getPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" ===================== Get Post ===================== ")

	var pt []entity.Post

	ptID, err := uuid.FromString(mux.Vars(r)["postID"])
	if err != nil {
		return
	}

	err = controller.service.GetPost(&pt, ptID)
	if err != nil {
		web.RespondErrorMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	web.RespondJSON(w, http.StatusOK, pt)
}
