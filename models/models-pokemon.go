package models

import (
	"github.com/jinzhu/gorm"
)

type Pokemon struct {
	gorm.Model
	Name	string	`json:"name"`
	Num		int		`json:"number"`

	Tall	float32	`json:" tall"`
	Weight	float32	`json:"weight"`

	Nature	string 	`json:"nature"` // 性格
	Commentary	string	`json:"commentary"` // 解说
	Comment	string	`json:"comment"` // 评论

	Evolution []Evolution `json:"evolution"` // 进化， one-to-many pokemonId是外键

	Prop	[]Prop	`json:"prop" gorm:"many2many:pokemon_prop;"` //属性 many-to-many pokemon_prop
	Skill	[]Skill `json:"skill" gorm:"many2many:pokemon_skill;"` //技能 many-to-many pokemon_skill
	Type	[]Type `json:"skill" gorm:"many2many:pokemon_type;"` //类型 many-to-many pokemon_type
	Special	[]Special `json:"skill" gorm:"many2many:pokemon_special;"` //特性 many-to-many pokemon_special

}

// add Pokemon
func AddPokemon(Pokemon *Pokemon) (*Pokemon, error) {
	return Pokemon, db.Save(&Pokemon).Error
}

// fetch All Pokemons
func FetchAllPokemon() (*gorm.DB){
	var Pokemons []Pokemon
	return db.Find(&Pokemons)
}

// fetch one
func FetchOnePokemon(id int) (*Pokemon) {
	var Pokemon Pokemon

	db.First(&Pokemon, id)

	return &Pokemon
}

// update Pokemon
func UpdatePokemon(id int, name string) (*Pokemon) {
	var Pokemon Pokemon
	db.First(&Pokemon,  id)
	db.Model(&Pokemon).Update("name", name)

	return &Pokemon
}

// delete Pokemon
func DeletePokemon(id int) (bool) {
	var Pokemon Pokemon
	db.First(&Pokemon, id)

	if Pokemon.ID == 0 {
		return false
	}

	db.Delete(&Pokemon)
	return true
}


