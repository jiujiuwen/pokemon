package models

import (
	"github.com/jinzhu/gorm"
	"log"
	"pokemon/setting"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)


var db *gorm.DB

func init(){
	var (
		err error
		dbType, dbName, user, passwd, tablePrefix string
		maxIdleConns, maxOpenConns int
	)

	// Load config
	dbConfig, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "faild to get section 'database': %v,", err)
	}

	dbType = dbConfig.Key("TYPE").String()
	dbName = dbConfig.Key("NAME").String()
	user = dbConfig.Key("USER").String()
	passwd = dbConfig.Key("PASSWORD").String()
	//host = dbConfig.Key("HOST").String()
	tablePrefix = dbConfig.Key("TABLE_PREFIX").String()
	maxIdleConns = dbConfig.Key("MAX_IDLE_CONNS").MustInt(10)
	maxOpenConns = dbConfig.Key("MAX_OPEN_CONNS").MustInt(100)

	// Connect db
	db, err = gorm.Open(dbType, user + ":" + passwd + "@/" + dbName + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
		panic("failed to connect database")
	}

	// set default table name
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.DB().SetMaxIdleConns(maxIdleConns)
	db.DB().SetMaxOpenConns(maxOpenConns)
	//Migrate the schema
	db.AutoMigrate(
		&Skill{},
		&Prop{},
		&Type{},
		&Special{})
}


func closeDB()  {
	defer db.Close()
}
