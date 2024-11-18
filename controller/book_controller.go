package controller

import (
	"go_test/dto"
	"go_test/models"
	"go_test/service"
	"go_test/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	Service *service.BookService
	Validator *utils.Validator

}

func BookControllerDI(service *service.BookService) *BookController {
	
	validator := utils.Valid()

	return &BookController{
		Service:   service,
		Validator: validator,
	}
}

func (c *BookController) GetAllBooks (ctx *gin.Context){
	books, err := c.Service.GetAllBook()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, books)
}


func (c *BookController) FindBookById (ctx *gin.Context){
	idRaw := ctx.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	Book, err := c.Service.Repos.GetById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Book)
}

func (ctr *BookController) RegisterBook(ctx *gin.Context) {
	var book dto.BRequest

	
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	
	if err := ctr.Validator.Validate.Struct(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}
	
	fbook := models.Book{
		Title: book.Title,
		ISBN: book.ISBN,
		AuthorID: (book.AuthorId),
	}
	
	if err := ctr.Service.RegisterBook(&fbook); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register book", "details": err.Error()})
		return
	}


	ctx.JSON(http.StatusCreated, gin.H{"message": "Book registered successfully"})
}

func (c *BookController) UpdateBook(ctx *gin.Context) {
	var book dto.BRequest
	idRaw := ctx.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Bind the incoming JSON to the book struct
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	// Validate the request payload
	if err := c.Validator.Validate.Struct(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}

	// Check if the book exists in the database
	existingBook, err := c.Service.FindBookById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found", "details": err.Error()})
		return
	}

	// Prepare the book model for update
	fbook := models.Book{
		ID:       existingBook.ID, // Keep the existing book ID
		Title:    book.Title,
		ISBN:     book.ISBN,
		AuthorID: book.AuthorId,
	}

	// Update the book in the database
	if err := c.Service.UpdateBook(uint(id), &fbook); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book", "details": err.Error()})
		return
	}

	// Respond with success message
	ctx.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}


func (c *BookController) DeleteBook(ctx *gin.Context) {
	// var book dto.BRequest
	idRaw := ctx.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// if err := ctx.ShouldBindJSON(&book); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
	// 	return
	// }

	if err := c.Service.DeleteBook(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book", "details": err.Error()})
		return
	}


	ctx.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}