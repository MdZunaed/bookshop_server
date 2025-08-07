package dto

import (
	"time"
)

type CreateBookDto struct {
	BookName     string    `bson:"book_name" json:"book_name"`
	BookCode     string    `bson:"book_code" json:"book_code"`
	Author       string    `bson:"author" json:"author"`
	PageCount    int       `bson:"page_count" json:"page_count"`
	PublishYear  int       `bson:"publish_year" json:"publish_year"`
	Price        float32   `bson:"price" json:"price"`
	Status       bool      `bson:"status" json:"status"`
	Description  string    `bson:"description,omitempty" json:"description,omitempty"`
	CreatedBy    string    `bson:"created_by" json:"created_by"`
	ModifiedBy   string    `bson:"modified_by" json:"modified_by"`
	CreatedDate  time.Time `bson:"created_date" json:"created_date"`
	ModifiedDate time.Time `bson:"modified_date" json:"modified_date"`
}

type UpdateBookDto struct {
	//Id           string    `bson:"_id" json:"id"`
	BookName     string    `bson:"book_name" json:"book_name"`
	BookCode     string    `bson:"book_code" json:"book_code"`
	PageCount    int       `bson:"page_count" json:"page_count"`
	Price        float32   `bson:"price" json:"price"`
	Status       bool      `bson:"status" json:"status"`
	Description  string    `bson:"description" json:"description"`
	ModifiedBy   string    `bson:"modified_by" json:"modified_by"`
	ModifiedDate time.Time `bson:"modified_date" json:"modified_date"`
}
