package service

import (
	"errors"
	"go_test/models"
	"go_test/repository"
)

type BookService struct{
	Repos repository.BookRepository
}

func BookServices (repo repository.BookRepository) *BookService{
	return &BookService{Repos :repo}
}

func (s *BookService) GetAllBook()([]models.Book,error){
	return s.Repos.GetAll()
}

func (s *BookService) FindBookById(id uint)(models.Book,error){
	return s.Repos.GetById(id)
}

func (s *BookService) RegisterBook(book *models.Book) error {
	return s.Repos.CreateBook(book)
}

func (s *BookService) UpdateBook(id uint,book *models.Book) error {
	exist, err := s.Repos.GetById(id)
	if err != nil{
		return err
	}
	author, err := s.Repos.GetAuthorById(book.AuthorID)
	if err != nil {
		return errors.New("invalid AuthorID: author not found")
	}

	exist.Title = book.Title
	exist.ISBN = book.ISBN	
	exist.AuthorID = author.ID
	return s.Repos.UpdateBook(&exist)	
}

func (s *BookService) DeleteBook(id uint) error {
	exist, err := s.Repos.GetById(id)
	if err != nil{
		return err
	}
	return s.Repos.DeleteBook(&exist)
}