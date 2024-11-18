package models

type Book struct {
	ID       uint   `gorm:"primaryKey"`
	Title    string `gorm:"not null" validate:"required"`
	ISBN     string `gorm:"not null;unique" validate:"required"`
	AuthorID  uint   `gorm:"index" json:"authorid"`
	Author    Author `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}