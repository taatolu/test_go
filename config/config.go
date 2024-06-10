package config

import(
	"log"
	"main/utils"
	"github.com/go-ini/ini"
)


type WebConfigList struct {
	Logfile	string
}

type DbConfigList struct {
	Host	string
	Port	string
	UserName	string
	Password	string
	DbName	string
}

var WebConfig WebConfigList
var DbConfig DbConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(WebConfig.Logfile)
}


func LoadConfig(){
	cfg, err := ini.Load("config.ini")
	
	if err != nil {
		log.Fatalln(err)
	}

	WebConfig = WebConfigList{
		Logfile: cfg.Section("web").Key("logfile").String(),
	}

	DbConfig = DbConfigList{
		Host: cfg.Section("db").Key("host").String(),
		Port: cfg.Section("db").Key("port").String(),
		UserName: cfg.Section("db").Key("userName").String(),
		Password: cfg.Section("db").Key("password").String(),
		DbName: cfg.Section("db").Key("dbName").String(),
	}
}