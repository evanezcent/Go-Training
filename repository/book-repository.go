package repository

import (
	"../model"
	"gorm.io/gorm"
)

// BookRepository as interface that cover all function
type BookRepository interface {
	InsertBook(book model.Book) model.Book
	UpdateBook(book model.Book) model.Book
	DeleteBook(book model.Book)
	GetAllBook() []model.Book
	FindBookByID(id uint64) model.Book
}

type bookConnection struct {
	connection *gorm.DB
}

// NewBookRepo used to create new Instance of user repository
func NewBookRepo(db *gorm.DB) BookRepository {
	return &bookConnection{
		connection: db,
	}
}

func (db *bookConnection) InsertBook(book model.Book) model.Book {
	db.connection.Save(&book)

	// handle action to get the owner of the book (User)
	db.connection.Preload("User").Find(&book)
	return book
}

func (db *bookConnection) UpdateBook(book model.Book) model.Book {
	db.connection.Updates(&book)

	// handle action to get the owner of the book (User)
	db.connection.Preload("User").Find(&book)
	return book
}

func (db *bookConnection) DeleteBook(book model.Book) {
	db.connection.Delete(&book)
}

func (db *bookConnection) FindBookByID(bookID uint64) model.Book {
	var book model.Book
	db.connection.Preload("User").Find(&book, bookID)

	return book
}

func (db *bookConnection) GetAllBook() []model.Book {
	var books []model.Book
	db.connection.Preload("User").Find(&books)

	return books
}
