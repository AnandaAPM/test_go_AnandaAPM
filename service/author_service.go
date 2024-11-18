package service

import (
	"go_test/models"
	"go_test/repository"
)

type AuthorService struct{
	Repos repository.AuthorRepository
}

func AuthorServices (repo repository.AuthorRepository) *AuthorService{
	return &AuthorService{Repos :repo}
}

func (s *AuthorService) GetAllAuthor()([]models.Author,error){
	return s.Repos.GetAll()
}
func (s *AuthorService) FindAuthorById(id uint)(models.Author,error){
	return s.Repos.GetById(id)
}

func (s *AuthorService) RegisterAuthor(author *models.Author) error {
	return s.Repos.CreateAuthor(author)
}

func (s *AuthorService) UpdateAuthor(id uint,author *models.Author) error {
	exist, err := s.Repos.GetById(id)
	if err != nil{
		return err
	}
	exist.Name = author.Name
	exist.Birthdate = author.Birthdate	
	return s.Repos.UpdateAuthor(&exist)	
}

func (s *AuthorService) DeleteAuthor(id uint) error {
	exist, err := s.Repos.GetById(id)
	if err != nil{
		return err
	}
	return s.Repos.DeleteAuthor(&exist)
}
