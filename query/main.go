package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/shaileshhb/microservices/query/app/comment"
	"github.com/shaileshhb/microservices/query/app/config"
	"github.com/shaileshhb/microservices/query/app/post"
	"github.com/shaileshhb/microservices/query/app/router"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() (*gorm.DB, error) {

	config, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// data source name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBDatabase)

	fmt.Println(" =================== dbURL -> ", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("successfully connected to database")
	// db.Debug().AutoMigrate(&entity.Comment{})
	// db.Debug().AutoMigrate(&entity.Comment{})

	return db, nil
}

func main() {

	db, err := initDatabase()
	if err != nil {
		fmt.Println("faild to connect to database -> ", err)
		return
	}

	app := mux.NewRouter().StrictSlash(true)
	routes := app.PathPrefix("/api/v1").Subrouter()
	routerRegister(db, routes)

	headers := handlers.AllowedHeaders([]string{"*"})
	methods := handlers.AllowedMethods([]string{http.MethodPost, http.MethodPut, http.MethodGet, http.MethodDelete, http.MethodOptions})
	origin := handlers.AllowedOrigins([]string{"*"})

	server := &http.Server{
		Addr:    "0.0.0.0:4003",
		Handler: handlers.CORS(headers, methods, origin)(routes),
	}

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(" ========= error occured ========= ", err.Error())
		return
	}
	fmt.Println(" ===== Server started ===== ")
}

func routerRegister(db *gorm.DB, routes *mux.Router) {

	postService := post.NewPostService(db)
	postController := router.NewPostController(postService)
	postController.RegisterRoutes(routes)

	commentService := comment.NewCommentService(db)
	// commentController := router.NewCommentController(commentService)
	// commentController.RegisterRoutes(routes)

	_ = router.NewEventBusController(postService, commentService)

}
