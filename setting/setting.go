// load conf files
// diff DEV_ENV and PROD_ENV
package setting

import (
	"log"
	"time"
	"github.com/go-ini/ini"
	"os"
)

var (
	Cfg *ini.File

	RunMode string
	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
)

func init()  {
	var err error

	_, fileErr := os.Stat("conf/config-dev.ini")
	if fileErr == nil {
		Cfg, err = ini.Load("conf/config-dev.ini")
		if err != nil {
			log.Fatalf("Faild to parse 'conf/config-dev.ini': %v", err)
		}
	// use prod
	} else{
		Cfg, err = ini.Load("conf/config-prod.ini")
		if err != nil {
			log.Fatalf("Faild to parse 'conf/config-prod.ini': %v", err)
		}
	}
}
func LoadBase()  {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer()  {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("faild to get setion 'server' : %v", err)
	}

	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

