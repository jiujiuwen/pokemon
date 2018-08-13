package models

import (
	"github.com/jinzhu/gorm"
)

type Special struct {
	gorm.Model
	Name	string	`json:"name"`
}

// add Special
func AddSpecial(Special *Special) (*Special, error) {
	return Special, db.Save(&Special).Error
}

// fetch All Specials
func FetchAllSpecial() (*gorm.DB){
	var Specials []Special
	return db.Find(&Specials)
}

// fetch one
func FetchOneSpecial(id int) (*Special) {
	var Special Special

	db.First(&Special, id)

	return &Special
}

// update Special
func UpdateSpecial(id int, name string) (*Special) {
	var Special Special
	db.First(&Special,  id)
	db.Model(&Special).Update("name", name)

	return &Special
}

// delete Special
func DeleteSpecial(id int) (bool) {
	var Special Special
	db.First(&Special, id)

	if Special.ID == 0 {
		return false
	}

	db.Delete(&Special)
	return true
}


