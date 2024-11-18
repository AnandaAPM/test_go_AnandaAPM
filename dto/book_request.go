package dto

type BRequest struct{
	
	Title string ` validate:"required"`
	ISBN string `validate:"required"`
	AuthorId uint `json:"authorid"`
}