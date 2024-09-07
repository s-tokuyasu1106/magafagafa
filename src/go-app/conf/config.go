package config

import (
	"log"

	"github.com/magafagafa/go-app/utils"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
	Host      string
	Database  string
	User      string
	Password  string
	Static    string
	DBPort    string
	Sslpem    string
	Sslkey    string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LogginSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		Host:      cfg.Section("db").Key("host").String(),
		Database:  cfg.Section("db").Key("database").String(),
		User:      cfg.Section("db").Key("user").String(),
		Password:  cfg.Section("db").Key("password").String(),
		DBPort:    cfg.Section("db").Key("dbport").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
		Static:    cfg.Section("web").Key("static").String(),
		Sslpem:    cfg.Section("web").Key("sslpem").String(),
		Sslkey:    cfg.Section("web").Key("sslkey").String(),
	}
}
