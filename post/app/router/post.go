package router

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/shaileshhb/microservices/post/app/post"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello world!!!") // => ✋ register
}

func SetupRoutes(app *fiber.App) {
	app.Post("/api/v1/post", post.AddPosts)
	app.Put("/api/v1/post/:postID", post.UpdatePost)
	app.Delete("/api/v1/post/:postID", post.DeletedPost)
	app.Get("/api/v1/post", post.GetPosts)
	app.Get("/api/v1/post/:postID", post.GetPost)
}
