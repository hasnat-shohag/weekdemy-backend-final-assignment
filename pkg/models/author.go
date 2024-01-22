package models

type AuthorDetail struct {
	AuthorId          uint   `gorm:"primaryKey; autoIncrement"`
	AuthorName        string `json:"author_name"`
	AuthorAddress     string `json:"author_address"`
	AuthorPhoneNumber string `json:"phone_number"`
}
