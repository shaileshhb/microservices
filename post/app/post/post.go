package post

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/satori/uuid"
	"github.com/shaileshhb/microservices/post/app/database"
	"github.com/shaileshhb/microservices/post/app/web"
)

type Base struct {
	ID        uuid.UUID  `json:"id" gorm:"type:varchar(36);primarykey"`
	CreatedAt time.Time  `json:"-" gorm:"type:datetime;DEFAULT"`
	UpdatedAt time.Time  `json:"-" gorm:"type:datetime;DEFAULT"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
}

type Post struct {
	Base
	Title string `json:"title" gorm:"type:varchar(100)"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.NewV4()
	return
}

func GetPosts(c *fiber.Ctx) error {
	fmt.Println(" ===================== GetPosts called ===================== ")
	db := database.DBConn
	var posts []Post

	err := db.Debug().Find(&posts).Error
	if err != nil {
		fmt.Println(" === err -> ", err)
		return fiber.ErrBadRequest
	}

	return c.JSON(&posts)
}

func GetPost(c *fiber.Ctx) error {
	fmt.Println(" ===================== GetPost called ===================== ")
	db := database.DBConn
	var post Post
	postID := c.Params("postID")

	err := db.Debug().Find(&post, postID).Error
	if err != nil {
		fmt.Println(" === err -> ", err)
		return fiber.ErrBadRequest
	}

	return c.JSON(&post)
}

func AddPosts(c *fiber.Ctx) error {
	fmt.Println(" ===================== AddPosts called ===================== ")
	var post Post
	err := web.UnmarshalJSON(c, &post)
	if err != nil {
		return fiber.ErrBadRequest
	}

	db := database.DBConn

	err = db.Debug().Create(&post).Error
	if err != nil {
		fmt.Println(" === err -> ", err)
		return fiber.ErrBadRequest
	}

	return c.SendString("Book added")
}

func UpdatePost(c *fiber.Ctx) error {
	fmt.Println(" ===================== UpdatePost called ===================== ")
	var post Post
	err := web.UnmarshalJSON(c, &post)
	if err != nil {
		return fiber.ErrBadRequest
	}

	post.ID, err = uuid.FromString(c.Params("postID"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	db := database.DBConn

	err = db.Debug().Model(&Post{}).Update(&post).Error
	if err != nil {
		fmt.Println(" === err -> ", err)
		return fiber.ErrBadRequest
	}
	return c.SendString("Book updated")
}

func DeletedPost(c *fiber.Ctx) error {
	fmt.Println(" ===================== DeletedPost called ===================== ")
	var post Post
	var err error

	post.ID, err = uuid.FromString(c.Params("postID"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	db := database.DBConn

	err = db.Debug().Delete(&post).Error
	if err != nil {
		fmt.Println(" === err -> ", err)
		return fiber.ErrBadRequest
	}
	return c.SendString("Book deleted")
}
