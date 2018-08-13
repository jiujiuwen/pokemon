package models

import (
	"github.com/jinzhu/gorm"
)

type Type struct {
	gorm.Model
	Name	string	`json:"name"`
}

// add Type
func AddType(Type *Type) (*Type, error) {
	return Type, db.Save(&Type).Error
}

// fetch All Types
func FetchAllType() (*gorm.DB){
	var Types []Type
	return db.Find(&Types)
}

// fetch one
func FetchOneType(id int) (*Type) {
	var Type Type

	db.First(&Type, id)

	return &Type
}

// update Type
func UpdateType(id int, name string) (*Type) {
	var Type Type
	db.First(&Type,  id)
	db.Model(&Type).Update("name", name)

	return &Type
}

// delete Type
func DeleteType(id int) (bool) {
	var Type Type
	db.First(&Type, id)

	if Type.ID == 0 {
		return false
	}

	db.Delete(&Type)
	return true
}


