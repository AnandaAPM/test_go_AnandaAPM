package controller

import (
	"go_test/dto"
	"go_test/models"
	"go_test/service"
	"go_test/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	Service *service.AuthorService
	Validator *utils.Validator

}

func AuthorControllerDI(service *service.AuthorService) *AuthorController {
	validator := utils.Valid()

	return &AuthorController{
		Service:   service,
		Validator: validator,
	}
}

func (c *AuthorController) GetAllAuthors (ctx *gin.Context){
	Authors, err := c.Service.GetAllAuthor()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Authors)
}

func (c *AuthorController) FindAuthorById (ctx *gin.Context){
	idRaw := ctx.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	Author, err := c.Service.Repos.GetById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Author)
}

func (ctr *AuthorController) RegisterAuthor(ctx *gin.Context) {
	var author dto.ARequest

	if err := ctx.ShouldBindJSON(&author); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	if err := ctr.Validator.Validate.Struct(&author); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}

	parsedDate, err := time.Parse("2006-01-02", author.Birthdate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD.", "details": err.Error()})
		return
	}

	fAuthor := models.Author{
		Name:      author.Name,
		Birthdate: parsedDate,
	}

	if err := ctr.Service.RegisterAuthor(&fAuthor); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register author", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Author registered successfully"})
}

func (c *AuthorController) UpdateAuthor(ctx *gin.Context) {
	var author dto.ARequest
	idRaw := ctx.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := ctx.ShouldBindJSON(&author); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	if err := c.Validator.Validate.Struct(&author); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}

	parsedDate, err := time.Parse("2006-01-02", author.Birthdate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD.", "details": err.Error()})
		return
	}

	fAuthor := models.Author{
		Name:      author.Name,
		Birthdate: parsedDate,
	}

	if err := c.Service.UpdateAuthor(uint(id), &fAuthor); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update author", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Author updated successfully"})
}

func (c *AuthorController) DeleteAuthor(ctx *gin.Context) {
	// var author dto.ARequest
	idRaw := ctx.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// if err := ctx.ShouldBindJSON(&author); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
	// 	return
	// }

	

	
	if err := c.Service.DeleteAuthor(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update author", "details": err.Error()})
		return
	}


	ctx.JSON(http.StatusOK, gin.H{"message": "Author deleted successfully"})
}

