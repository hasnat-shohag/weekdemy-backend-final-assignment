package models

type BookDetail struct {
	BookId      uint   `gorm:"primaryKey;autoIncrement"`
	BookName    string `json:"book_name"`
	AuthorId    uint   `json:"author_id"`
	Publication string `json:"publication;omitempty"`
}
