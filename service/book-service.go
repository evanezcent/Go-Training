package service

import (
	"fmt"
	"log"

	"../dto"
	"../model"
	"../repository"
	"github.com/mashingan/smapping"
)

// BookService interface for book service
type BookService interface {
	InsertBook(book dto.BookCreateDTO) model.Book
	UpdateBook(book dto.BookUpdateDTO) model.Book
	DeleteBook(book model.Book)
	GetAll() []model.Book
	FindBookByID(id uint64) model.Book
	AuthorizeForEdit(userID string, bookID uint64) bool
}

type bookService struct {
	bookRepository repository.BookRepository
}

// NewBookService is new Instance
func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepo,
	}
}

func (service *bookService) InsertBook(book dto.BookCreateDTO) model.Book {
	newBook := model.Book{}
	err := smapping.FillStruct(&newBook, smapping.MapFields(&book))

	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}

	updatedBook := service.bookRepository.InsertBook(newBook)
	return updatedBook
}

func (service *bookService) UpdateBook(book dto.BookUpdateDTO) model.Book {
	newBook := model.Book{}
	err := smapping.FillStruct(&newBook, smapping.MapFields(&book))

	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}

	updatedBook := service.bookRepository.UpdateBook(newBook)
	return updatedBook
}

func (service *bookService) DeleteBook(book model.Book) {
	service.bookRepository.DeleteBook(book)
}

func (service *bookService) GetAll() []model.Book {
	return service.bookRepository.GetAllBook()
}

func (service *bookService) FindBookByID(id uint64) model.Book {
	return service.bookRepository.FindBookByID(id)
}

func (service *bookService) AuthorizeForEdit(userID string, bookID uint64) bool {
	book := service.bookRepository.FindBookByID(bookID)
	id := fmt.Sprintf(book.UserID)

	return userID == id
}
