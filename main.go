package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"pokemon/routers"
)


func main() {
	routers.InitRouter()
}


