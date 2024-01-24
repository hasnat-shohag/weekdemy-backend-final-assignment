package types

import validation "github.com/go-ozzo/ozzo-validation"

// BookRequest AuthorRequest response struct | marshalled into json fromat from struct
type BookRequest struct {
	BookId      uint   `json:"book_id,omitempty"`
	BookName    string `json:"book_name"`
	AuthorId    uint   `json:"author_id"`
	Publication string `json:"publication,omitempty"`
}

type AuthorRequest struct {
	AuthorID          uint   `json:"author_id"`
	AuthorName        string `json:"author_name"`
	AuthorAddress     string `json:"author_address,omitempty"`
	AuthorPhoneNumber string `json:"author_phone_number,omitempty"`
}

func (book BookRequest) Validate() error {
	return validation.ValidateStruct(&book,
		validation.Field(&book.BookName,
			validation.Required.Error("Book name cannot be empty"),
			validation.Length(1, 50)),
		validation.Field(&book.AuthorId,
			validation.Required.Error("Author ID is required")))
}

func (author AuthorRequest) Validate() error {
	return validation.ValidateStruct(&author,
		validation.Field(&author.AuthorName,
			validation.Required.Error("Author name is required"),
			validation.Length(1, 50)))
}
