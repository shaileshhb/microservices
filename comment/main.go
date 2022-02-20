package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shaileshhb/microservices/comment/app/comment"
	"github.com/shaileshhb/microservices/comment/app/config"
	"github.com/shaileshhb/microservices/comment/app/entity"
	"github.com/shaileshhb/microservices/comment/app/router"
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
	db.Debug().AutoMigrate(&entity.Comment{})

	return db, nil
}

func main() {

	db, err := initDatabase()
	if err != nil {
		fmt.Println("faild to connect to database -> ", err)
		return
	}
	// defer db.Close()

	r := gin.Default()
	// r.Group("/api/v1").
	// .BasePath()

	service := comment.NewCommentService(db)
	controller := router.NewCommentController(service)
	controller.RegisterRoutes(r)

	r.Run(":4002")

}
