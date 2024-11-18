package repository

import (
	"go_test/models"

	"gorm.io/gorm"
)
type BookRepository interface{
	 GetAll()([]models.Book, error)
	 GetById(id uint)(models.Book, error)
	GetAuthorById(id uint)(models.Author, error)
	CreateBook(book *models.Book)error
	UpdateBook(book *models.Book)error
	DeleteBook(book *models.Book)error
}
type bookRepository struct{
	DB *gorm.DB
}

func BookRepo(db *gorm.DB) *bookRepository{
	return &bookRepository{DB:db}
}

func (b *bookRepository) GetAll()([]models.Book, error){
	var books []models.Book
	err := b.DB.Preload("Author").Find(&books).Error

	return books,err
}

func (b *bookRepository) GetById(id uint)(models.Book, error){
	var books models.Book
	err := b.DB.Preload("Author").Where("id = ?",id).First(&books).Error

	return books,err
}

func (b *bookRepository) GetAuthorById(id uint)(models.Author, error){
	var author models.Author
	err := b.DB.First(&author,id).Error

	return author,err
}

func (r *bookRepository) CreateBook(book *models.Book)error{
	return r.DB.Create(book).Error
}
func (r *bookRepository) UpdateBook(book *models.Book)error{
	return r.DB.Model(book).Updates(book).Error;
}
func (r *bookRepository) DeleteBook(book *models.Book)error{
	return r.DB.Delete(book).Error;
}