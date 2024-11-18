package models

import "golang.org/x/crypto/bcrypt"
type User struct{
	Id uint `gorm:"primaryKey"`
	Username string `gorm:"not null;unique" validate:"required"`
	Password string `gorm:"not null;unique" validate:"required"`
}




func(u *User) Hashed(password string) error{

	hashedP,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	
	if err !=nil{
		return err
	}
	u.Password = string(hashedP)
	return nil

}

func (u *User)ValidatePassword(dbpassword,reqpassword string)(bool){

	err := bcrypt.CompareHashAndPassword([]byte(dbpassword), []byte(reqpassword))

	return  err == nil
}