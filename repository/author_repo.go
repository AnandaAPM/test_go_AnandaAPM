package repository

import (
	"go_test/models"

	"gorm.io/gorm"
)
type AuthorRepository interface{
	 GetAll()([]models.Author, error)
	 GetById(id uint)(models.Author, error)
	 CreateAuthor (book *models.Author)error
	 UpdateAuthor (book *models.Author)error
	 DeleteAuthor (book *models.Author)error
}
type authorRepository struct{
	DB *gorm.DB
}

func AuthorRepo(db *gorm.DB) *authorRepository{
	return &authorRepository{DB:db}
}

func (b *authorRepository) GetAll()([]models.Author, error){
	var Authors []models.Author
	err := b.DB.Find(&Authors).Error

	return Authors,err
}

func (b *authorRepository) GetById(id uint)(models.Author, error){
	var author models.Author
	err := b.DB.Where("id = ?",id).First(&author).Error

	return author,err
}

func (r *authorRepository) CreateAuthor (author *models.Author)error{
	return r.DB.Create(author).Error
}
func (r *authorRepository) UpdateAuthor (author *models.Author)error{
	return r.DB.Updates(author).Error;
}
func (r *authorRepository) DeleteAuthor (author *models.Author)error{
	return r.DB.Delete(author).Error;
}