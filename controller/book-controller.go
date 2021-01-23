package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"../dto"
	"../helper"
	"../model"
	"../service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// BookController interface for login, register, read, and update user
type BookController interface {
	All(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type bookController struct {
	bookService service.BookService
	jwtService  service.JWTService
}

// NewBookController is like constructor of the model
func NewBookController(bookService service.BookService, jwtService service.JWTService) BookController {
	return &bookController{
		bookService,
		jwtService,
	}
}

func (controller *bookController) getUserIDFromToken(token string) string {
	authToken, err := controller.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}

	claims := authToken.Claims.(jwt.MapClaims)
	return fmt.Sprintf("%v", claims["userID"])
}

func (controller *bookController) All(ctx *gin.Context) {
	var books []model.Book = controller.bookService.GetAll()
	res := helper.ResponseSucces(true, "OK", books)
	ctx.JSON(http.StatusOK, res)
}

func (controller *bookController) FindByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)

	if err != nil {
		res := helper.ResponseFailed("ID not found", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	var book model.Book = controller.bookService.FindBookByID(id)
	if (book == model.Book{}) {
		res := helper.ResponseFailed("Data not found", "Wrong id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		res := helper.ResponseSucces(true, "success", book)
		ctx.JSON(http.StatusOK, res)
	}
}

func (controller *bookController) Insert(ctx *gin.Context) {
	var newBook dto.BookCreateDTO
	errDTO := ctx.ShouldBind(&newBook)

	if errDTO != nil {
		res := helper.ResponseFailed("Failed to process the request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := ctx.GetHeader("Authorization")
		userID := controller.getUserIDFromToken(authHeader)
		id, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			newBook.UserID = id
		}

		res := controller.bookService.InsertBook(newBook)
		response := helper.ResponseSucces(true, "success", res)
		ctx.JSON(http.StatusOK, response)
	}
}

func (controller *bookController) Update(ctx *gin.Context) {
	var newBook dto.BookUpdateDTO
	errDTO := ctx.ShouldBind(&newBook)

	if errDTO != nil {
		res := helper.ResponseFailed("Failed to process the request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := controller.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["userID"])

	if controller.bookService.AuthorizeForEdit(userID, newBook.ID) {
		res := controller.bookService.UpdateBook(newBook)
		response := helper.ResponseSucces(true, "success", res)
		ctx.JSON(http.StatusOK, response)
	} else {
		res := helper.ResponseFailed("You don't have any permission", "Invalid credential authorization", helper.EmptyObj{})
		ctx.JSON(http.StatusForbidden, res)
	}
}

func (controller *bookController) Delete(ctx *gin.Context) {
	var book model.Book
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helper.ResponseFailed("Failed to get ID", "Invalid ID", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	}

	book.ID = id
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := controller.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["userID"])

	if controller.bookService.AuthorizeForEdit(userID, book.ID) {
		controller.bookService.DeleteBook(book)
		response := helper.ResponseSucces(true, "Book Deleted", helper.EmptyObj{})
		ctx.JSON(http.StatusOK, response)
	} else {
		res := helper.ResponseFailed("You don't have any permission", "Invalid credential authorization", helper.EmptyObj{})
		ctx.JSON(http.StatusForbidden, res)
	}
}
