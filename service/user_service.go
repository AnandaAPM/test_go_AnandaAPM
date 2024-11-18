package service

import (
	"errors"
	"go_test/auth"
	"go_test/models"
	"go_test/repository"

	"gorm.io/gorm"
)

type UserService struct{
	Repos repository.UserRepository
}

func UserServices (repo repository.UserRepository) *UserService{
	return &UserService{Repos :repo}
}

func (s *UserService) RegisterUser(user *models.User) error {
	err := user.Hashed(user.Password)
	if err != nil {
		return err
	}

	
	existingUser, err := s.Repos.FindUser(user.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if existingUser != nil {
		return errors.New("username already exists")
	}
	return s.Repos.CreateUser(user)
	
}

func (s *UserService) LoginUser(user *models.User) (string,error) {
	

	
	exist, err := s.Repos.FindUser(user.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("user not found")
		}
		return "", err 
	}
	validate := user.ValidatePassword(exist.Password,user.Password)
	println(validate)
	if !validate {
		return "",errors.New("Password Invalid")
	}

	tokenStr,err:=auth.Generate(exist.Username)

	if err !=nil {
		return "", err
	}
	
	return tokenStr,nil
	
}