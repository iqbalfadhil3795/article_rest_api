package structs

import "github.com/jinzhu/gorm"

type Article struct{
	gorm.Model
	Title string
	Content string
	Category string
}