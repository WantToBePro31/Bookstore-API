package models

type Book struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Title        string `json:"title"`
	Author       string `json:"author"`
	Release_Year int    `json:"release_year"`
}
