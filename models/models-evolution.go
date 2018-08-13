package models

import (
	"github.com/jinzhu/gorm"
)

type Evolution struct {
	gorm.Model
	PokemonId	int		`json:"before_pokemon" gorm:"index"`
	AfterId		int		`json:"after_pokemon" gorm:"unique_index"`
	Condition	string 	`json:"condition"`
	Level		int		`json:"level"`
	IsMega		bool	`json:"is_mega"`
}

//// add Evolution
//func AddEvolution(Evolution *Evolution) (*Evolution, error) {
//	return Evolution, db.Save(&Evolution).Error
//}
//
//// fetch All Evolutions
//func FetchAllEvolution() (*gorm.DB){
//	var Evolutions []Evolution
//	return db.Find(&Evolutions)
//}
//
//// fetch one
//func FetchOneEvolution(id int) (*Evolution) {
//	var Evolution Evolution
//
//	db.First(&Evolution, id)
//
//	return &Evolution
//}
//
//// update Evolution
//func UpdateEvolution(id int, name string) (*Evolution) {
//	var Evolution Evolution
//	db.First(&Evolution,  id)
//	db.Model(&Evolution).Update("name", name)
//
//	return &Evolution
//}
//
//// delete Evolution
//func DeleteEvolution(id int) (bool) {
//	var Evolution Evolution
//	db.First(&Evolution, id)
//
//	if Evolution.ID == 0 {
//		return false
//	}
//
//	db.Delete(&Evolution)
//	return true
//}


