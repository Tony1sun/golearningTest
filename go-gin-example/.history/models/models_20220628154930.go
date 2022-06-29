package models

import "github.com/jinzhu/gorm"

var db *gorm.DB

type Model struct {
	ID        int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on`
}
