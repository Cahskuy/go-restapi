package models

type Post struct {
	ID    uint64 `gorm:"column:id;primaryKey" json:"id"`
	Title string `gorm:"column:title;size:255;not null" json:"title"`
	Body  string `gorm:"column:body;size:255" json:"body"`
}
