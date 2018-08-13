package models

import (
	"github.com/jinzhu/gorm"
)

type Skill struct {
	gorm.Model
	Name	string	`json:"name"`

	PropId int 		`json:"prop_id" gorm:"index"`// prop id
	Prop	Prop	`json:"prop" gorm:"ForeignKey:PropID"` // one to one

	Describe string	`json:"describe"`
}

// add Skill
func AddSkill(skill *Skill) (*Skill, error) {
	return skill, db.Save(&skill).Error
}

// get total num

func GetTotalNum() (count int) {
	db.Model(&Skill{}).Count(&count)

	return
}
// fetch All Skills
func FetchAllSkill(pageNum int, pageSize int,name string) (*gorm.DB){
	var skills []Skill
	var startPos = (pageNum -1) * pageSize

	if len(name) == 0 {
		return db.Preload("Prop").Offset(startPos).Limit(pageSize).Find(&skills)
	} else {
		return db.Preload("Prop").Offset(startPos).Limit(pageSize).Find(&skills, "name LIKE ?", "%" + name + "%")
	}
}

// fetch one
func FetchOneSkill(id int) (*Skill) {
	var skill Skill

	db.Where("id = ?", id).First(&skill)
	db.Model(&skill).Related(&skill.Prop)

	return &skill
}

// update skill
func UpdateSkill(id int, skill *Skill) (*Skill) {
	var Skill Skill
	db.First(&Skill, id)
	db.Model(&Skill).Updates(skill).Related(&Skill.Prop)

	return &Skill
}

// delete skill
func DeleteSkill(id int) (bool) {
	var skill Skill
	db.First(&skill, id)

	if skill.ID == 0 {
		return false
	}

	db.Delete(&skill)
	return true
}


