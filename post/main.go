package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jinzhu/gorm"
	"github.com/shaileshhb/microservices/post/app/config"
	"github.com/shaileshhb/microservices/post/app/database"
	"github.com/shaileshhb/microservices/post/app/post"
	"github.com/shaileshhb/microservices/post/app/router"
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

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	fmt.Println(" === successfully connected to database == ")

	db.Debug().AutoMigrate(&post.Post{})
	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&post.Post{})

	return db, nil
}

func main() {

	db, err := initDatabase()
	if err != nil {
		fmt.Println("faild to connect to database -> ", err)
		return
	}
	database.DBConn = db
	defer database.DBConn.Close()

	app := fiber.New()
	app.Use(cors.New())
	router.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":4001"))
}
