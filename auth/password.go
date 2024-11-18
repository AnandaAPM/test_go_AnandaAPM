package auth

import "golang.org/x/crypto/bcrypt"

func Hashed(passwordReq string)(string, error){

	hashedP,err := bcrypt.GenerateFromPassword([]byte(passwordReq),bcrypt.DefaultCost)

	return string(hashedP),err

}

func ValidatePassword(hashedPassword, plainPassword string)(bool){

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))

	return  err == nil
}