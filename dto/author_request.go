package dto

type ARequest struct{
	Name string `validate:"required" json:"name"`
	Birthdate string `validate:"required" json:"birthdate"`
}