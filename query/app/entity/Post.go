package entity

type Post struct {
	Base
	Title    string    `json:"title" gorm:"type:varchar(100)"`
	Comments []Comment `json:"comments"`
}
