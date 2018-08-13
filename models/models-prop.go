package models

import (
	"github.com/jinzhu/gorm"
)

type (
	ResProp struct {
		ID 		int 	`json:"id"`
		PropId	int		`json:"prop_id"`
		ResPropId	int	`json:"res_prop_id"`
	}


	Prop struct {
		gorm.Model
		Name	string	`json:"name"`
		//RestraintId []int 		`json:"restraint_id" gorm:"index"`// restraint prop id
		//Restraint	[]resProp	`json:"restraint" gorm:"ForeignKey:RestraintId"` // 1 - n

		//Restraint int `json:"restraint"`
	}

)

// add Prop
func AddProp(Prop *Prop) (*Prop, error) {
	return Prop, db.Save(&Prop).Error
}

// fetch All Props
func FetchAllProp() (*gorm.DB){
	var Props []Prop
	return db.Find(&Props)
}

// fetch one
func FetchOneProp(id int) (*Prop) {
	var Prop Prop

	db.Where("id = ?", id).First(&Prop)
	//db.Model(&Prop).Related(&Prop.Restraint)
	//db.First(&Prop, id)

	return &Prop
}

// update Prop
func UpdateProp(id int, prop *Prop) (*Prop) {
	var Prop Prop
	db.First(&Prop,  id)
	db.Model(&Prop).Update(prop)
	//db.Model(&Prop).Update("name", prop)
	//db.Model(&Prop).Update("restraint", restraint)

	return &Prop
}

// delete Prop
func DeleteProp(id int) (bool) {
	var Prop Prop
	db.First(&Prop, id)

	if Prop.ID == 0 {
		return false
	}

	db.Delete(&Prop)
	return true
}


